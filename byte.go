package cast

import (
	"encoding/binary"
	"fmt"
	"math"
)

func ToBytes(i any, b binary.ByteOrder) []byte {
	v, _ := ToBytesE(i, b)
	return v
}

func ToBytesE(i any, b binary.ByteOrder) ([]byte, error) {
	i, _ = indirect(i)

	if i == nil {
		return []byte{}, fmt.Errorf("unable to cast %#v of type %T to []byte", i, i)
	}

	switch v := i.(type) {
	case []byte:
		return v, nil
	case byte:
		return []byte{v}, nil
	case int8:
		if v < 0 {
			return []byte{}, errNegativeNotAllowed
		}
		return []byte{byte(v)}, nil
	case int16:
		if v < 0 {
			return []byte{}, errNegativeNotAllowed
		}
		a := make([]byte, 2)
		b.PutUint16(a, uint16(v))
		return a, nil
	case int32:
		if v < 0 {
			return []byte{}, errNegativeNotAllowed
		}
		a := make([]byte, 4)
		b.PutUint32(a, uint32(v))
		return a, nil
	case int64:
		if v < 0 {
			return []byte{}, errNegativeNotAllowed
		}
		a := make([]byte, 8)
		b.PutUint64(a, uint64(v))
		return a, nil
	case uint16:
		a := make([]byte, 2)
		b.PutUint16(a, v)
		return a, nil
	case uint32:
		a := make([]byte, 4)
		b.PutUint32(a, v)
		return a, nil
	case uint64:
		a := make([]byte, 8)
		b.PutUint64(a, v)
		return a, nil
	case int:
		switch true {
		case v < 0:
			return []byte{}, errNegativeNotAllowed
		case v <= math.MaxUint32:
			a := make([]byte, 4)
			b.PutUint32(a, uint32(v))
			return a, nil
		default:
			a := make([]byte, 8)
			b.PutUint64(a, uint64(v))
			return a, nil
		}
	case uint:
		switch true {
		case v < 0:
			return []byte{}, errNegativeNotAllowed
		case v <= math.MaxUint32:
			a := make([]byte, 4)
			b.PutUint32(a, uint32(v))
			return a, nil
		case v <= math.MaxUint64:
			a := make([]byte, 8)
			b.PutUint64(a, uint64(v))
			return a, nil
		default:
			return []byte{}, fmt.Errorf("unable to cast %#v of type %T to []byte", i, i)
		}
	case string:
		return []byte(v), nil
	default:
		return []byte{}, fmt.Errorf("unable to cast %#v of type %T to []byte", i, i)
	}
}
