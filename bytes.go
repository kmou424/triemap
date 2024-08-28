package triemap

import (
	"fmt"
	"math"
	"reflect"
)

func toBytes(v any) []byte {
	val := reflect.ValueOf(v)
	return deepEncode(val)
}

func deepEncode(val reflect.Value) []byte {
	t := val.Type()
	switch t.Kind() {
	case reflect.String:
		return []byte(val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intToBytes(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintToBytes(val.Uint())
	case reflect.Float32, reflect.Float64:
		return floatToBytes(val.Float())
	case reflect.Bool:
		if val.Bool() {
			return []byte{0x01}
		}
		return []byte{0x00}
	case reflect.Slice, reflect.Array:
		length := val.Len()
		data := make([]byte, length*int(val.Index(0).Type().Size()))
		for i := 0; i < length; i++ {
			subSliceData := deepEncode(val.Index(i))
			copy(data[i*len(subSliceData):i*len(subSliceData)+len(subSliceData)], subSliceData)
		}
		return data
	case reflect.Struct:
		var b []byte
		for i := 0; i < t.NumField(); i++ {
			b = append(b, deepEncode(val.Field(i))...)
		}
		return b
	default:
		panic(fmt.Sprintf("Unsupported type: %s", t))
	}
}

func intToBytes(x int64) []byte {
	return []byte{byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24)}
}

func uintToBytes(x uint64) []byte {
	return []byte{byte(x), byte(x >> 8), byte(x >> 16), byte(x >> 24)}
}

func floatToBytes(x float64) []byte {
	var buf [8]byte
	n := math.Float64bits(x)
	buf[0] = byte(n >> 56)
	buf[1] = byte(n >> 48)
	buf[2] = byte(n >> 40)
	buf[3] = byte(n >> 32)
	buf[4] = byte(n >> 24)
	buf[5] = byte(n >> 16)
	buf[6] = byte(n >> 8)
	buf[7] = byte(n)
	return buf[:]
}
