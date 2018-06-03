// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package marshaling

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/mitchellh/mapstructure"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack"
	"go.thethings.network/lorawan-stack/pkg/errors"
)

var (
	msgpUnmarshalerType = reflect.TypeOf((*msgp.Unmarshaler)(nil)).Elem()
)

// Unflattened unflattens m and returns the result
func Unflattened(m map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(m))
	for k, v := range m {
		skeys := strings.Split(k, Separator)
		parent := out
		for _, sk := range skeys[:len(skeys)-1] {
			sm, ok := parent[sk]
			if !ok {
				sm = make(map[string]interface{})
				parent[sk] = sm
			}
			parent = sm.(map[string]interface{})
		}
		parent[skeys[len(skeys)-1]] = v
	}
	return out
}

// IsNillableKind reports whether k represents a nillable reflect.Kind.
func IsNillableKind(k reflect.Kind) bool {
	return k == reflect.Ptr ||
		k == reflect.Map ||
		k == reflect.Interface ||
		k == reflect.Chan ||
		k == reflect.Func ||
		k == reflect.Slice
}

// IsNillableType reports whether t represents a nillable reflect.Type.
func IsNillableType(t reflect.Type) bool {
	return IsNillableKind(t.Kind())
}

// MapUnmarshaler is the interface implemented by an object that can
// unmarshal a map[string]interface{} representation of itself.
//
// UnmarshalMap must be able to decode the form generated by MarshalMap.
// UnmarshalMap must deep copy the data if it wishes to retain the data after returning.
type MapUnmarshaler interface {
	UnmarshalMap(map[string]interface{}) error
}

// UnmarshalMap parses the map-encoded data and stores the result
// in the value pointed to by v.
//
// UnmarshalMap uses the inverse of the encodings that
// Marshal uses.
func UnmarshalMap(m map[string]interface{}, v interface{}, hooks ...mapstructure.DecodeHookFunc) error {
	if mu, ok := v.(MapUnmarshaler); ok {
		return mu.UnmarshalMap(m)
	}

	if len(m) == 0 {
		return nil
	}

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr && rv.IsValid() {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return ErrInvalidData.NewWithCause(nil, errors.Errorf("Indirected value is nil"))
	}

	for mk, mv := range m {
		skeys := strings.Split(mk, Separator)

		fv := rv
		for _, sk := range skeys {
			for fv.Kind() == reflect.Ptr {
				if fv.IsNil() {
					fv.Set(reflect.New(fv.Type().Elem()))
				}
				fv = fv.Elem()
			}
			if fv = fv.FieldByName(sk); !fv.IsValid() {
				return errors.Errorf("Field `%s` specified, but does not exist on structs of type `%s`", sk, fv.Type())
			}
		}

		rmv := reflect.ValueOf(mv)
		switch vk := rmv.Kind(); vk {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
			if vk != fv.Kind() {
				return ErrInvalidData.NewWithCause(nil, errors.Errorf("Field value `%s` has kind `%s`, when `%s` is expected", mk, fv.Kind(), vk))
			}
			fv.Set(rmv)

		case reflect.Invalid:
			continue

		case reflect.Slice:
			if rmv.Type().Elem().Kind() == reflect.Uint8 {
				iv, err := BytesToType(rmv.Bytes(), fv.Type())
				if err != nil {
					return err
				}
				if iv != nil {
					fv.Set(reflect.ValueOf(iv))
				}
				continue
			}
			fallthrough

		default:
			return ErrInvalidData.NewWithCause(nil, errors.Errorf("Unmatched map value kind: `%s`", vk))
		}
	}
	return nil
}

// BytesToType decodes []byte value in b into a new value of type typ.
// It assumes that nil values are represented using an empty b(nil or 0-length).
// Otherwise, it expects the first byte in b to represent the encoding method
// used to encode the value and attempts to decode accordingly.
func BytesToType(b []byte, typ reflect.Type) (interface{}, error) {
	if len(b) == 0 {
		if !IsNillableType(typ) {
			return nil, ErrInvalidData.NewWithCause(nil, errors.Errorf("Type `%s` is not nullable, but zero-length byte slice specified", typ))
		}
		return nil, nil
	}

	enc := Encoding(b[0])
	b = b[1:]

	pv := reflect.New(typ)
	ev := pv.Elem()
	iv := ev
	for iv.Kind() == reflect.Ptr {
		iv.Set(reflect.New(iv.Type().Elem()))
		iv = iv.Elem()
	}

	switch iv.Kind() {
	case reflect.Slice:
		iv.Set(reflect.MakeSlice(iv.Type(), 0, 0))
	case reflect.Map:
		iv.Set(reflect.MakeMap(iv.Type()))
	case reflect.Chan:
		iv.Set(reflect.MakeChan(iv.Type(), 0))
	case reflect.Func:
		iv.Set(reflect.MakeFunc(iv.Type(), func([]reflect.Value) []reflect.Value { return nil }))
	}

	switch enc {
	case JSONEncoding:
		err := json.Unmarshal(b, pv.Interface())
		return pv.Elem().Interface(), err

	case ProtoEncoding:
		rv := ev
		if !ev.Type().Implements(protoMessageType) {
			if !pv.Type().Implements(protoMessageType) {
				return nil, errors.Errorf("Expected %s or %s to implement %s", ev.Type(), pv.Type(), protoMessageType)
			}
			rv = pv
		}

		err := proto.Unmarshal(b, rv.Interface().(proto.Message))
		return pv.Elem().Interface(), err

	case MsgPackEncoding:
		rv := ev
		if !ev.Type().Implements(msgpUnmarshalerType) {
			if !pv.Type().Implements(msgpUnmarshalerType) {
				err := msgpack.Unmarshal(b, pv.Interface())
				return pv.Elem().Interface(), err
			}
			rv = pv
		}

		b, err := rv.Interface().(msgp.Unmarshaler).UnmarshalMsg(b)
		if len(b) > 0 {
			return nil, ErrInvalidData.NewWithCause(nil, errors.New("Unmarshaled data left in buffer"))
		}
		return pv.Elem().Interface(), err

	case GobEncoding:
		err := gob.NewDecoder(bytes.NewReader(b)).DecodeValue(pv.Elem())
		return pv.Elem().Interface(), err
	default:
		return nil, ErrInvalidData.NewWithCause(nil, errors.Errorf("Unmatched encoding"))
	}
}

// ByteMapUnmarshaler is the interface implemented by an object that can
// unmarshal a map[string][]byte representation of itself.
//
// UnmarshalByteMap must be able to decode the form generated by MarshalByteMap.
// UnmarshalByteMap must deep copy the data if it wishes to retain the data after returning.
type ByteMapUnmarshaler interface {
	UnmarshalByteMap(map[string][]byte) error
}

// UnmarshalByteMap parses the byte map-encoded data and stores the result in the value pointed to by v.
// UnmarshalByteMap uses the inverse of the encodings that MarshalByteMap uses.
func UnmarshalByteMap(bm map[string][]byte, v interface{}, hooks ...mapstructure.DecodeHookFunc) error {
	if mm, ok := v.(ByteMapUnmarshaler); ok {
		return mm.UnmarshalByteMap(bm)
	}

	if len(bm) == 0 {
		return nil
	}

	im := make(map[string]interface{}, len(bm))
	for k, v := range bm {
		im[k] = v
	}

	return UnmarshalMap(im, v)
}