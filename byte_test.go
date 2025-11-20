package cast

import (
	"encoding/binary"
	"math"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestToBytesE(t *testing.T) {
	c := qt.New(t)

	//Overkill?
	expected8 := []byte{math.MaxInt8}
	expectedLittle16 := append([]byte{math.MaxUint8}, math.MaxInt8)
	expectedBig16 := append(expected8, math.MaxUint8)
	expectedLittle32 := append([]byte{math.MaxUint8, math.MaxUint8}, expectedLittle16...)
	expectedBig32 := append(expectedBig16, math.MaxUint8, math.MaxUint8)
	expectedLittle64 := append([]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8}, expectedLittle32...)
	expectedBig64 := append(expectedBig32, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8)

	expectedU8 := []byte{math.MaxUint8}
	expectedLittleU16 := append(expectedU8, expectedU8...)
	expectedLittleU32 := append(expectedLittleU16, expectedLittleU16...)
	expectedLittleU64 := append(expectedLittleU32, expectedLittleU32...)
	expectedBigU16 := append(expected8, expectedU8...)
	expectedBigU32 := append(expectedBigU16, expectedLittleU16...)
	expectedBigU64 := append(expectedBigU32, expectedLittleU32...)

	tests := []struct {
		input     any
		byteOrder binary.ByteOrder
		expect    []byte
		iserr     bool
	}{
		//LittleEndian
		{byte(math.MaxUint8), binary.LittleEndian, expectedU8, false},
		{math.MaxInt32, binary.LittleEndian, expectedLittle32, false},
		{math.MaxInt64, binary.LittleEndian, expectedLittle64, false},
		{uint(math.MaxUint32), binary.LittleEndian, expectedLittleU32, false},
		{uint(math.MaxUint64), binary.LittleEndian, expectedLittleU64, false},
		{int8(math.MaxInt8), binary.LittleEndian, expected8, false},
		{int16(math.MaxInt16), binary.LittleEndian, expectedLittle16, false},
		{int32(math.MaxInt32), binary.LittleEndian, expectedLittle32, false},
		{int64(math.MaxInt64), binary.LittleEndian, expectedLittle64, false},
		{uint8(math.MaxUint8), binary.LittleEndian, expectedU8, false},
		{uint16(math.MaxUint16), binary.LittleEndian, expectedLittleU16, false},
		{uint32(math.MaxUint32), binary.LittleEndian, expectedLittleU32, false},
		{uint64(math.MaxUint64), binary.LittleEndian, expectedLittleU64, false},
		{[]byte("one time"), binary.LittleEndian, []byte("one time"), false},
		{"one more time", binary.LittleEndian, []byte("one more time"), false},
		//BigEndian
		{byte(math.MaxUint8), binary.BigEndian, expectedU8, false},
		{math.MaxInt32, binary.BigEndian, expectedBig32, false},
		{math.MaxInt64, binary.BigEndian, expectedBig64, false},
		{uint(math.MaxInt32), binary.BigEndian, expectedBigU32, false},
		{uint(math.MaxInt64), binary.BigEndian, expectedBigU64, false},
		{int8(math.MaxInt8), binary.BigEndian, expected8, false},
		{int16(math.MaxInt16), binary.BigEndian, expectedBig16, false},
		{int32(math.MaxInt32), binary.BigEndian, expectedBig32, false},
		{int64(math.MaxInt64), binary.BigEndian, expectedBig64, false},
		{uint8(math.MaxUint8), binary.BigEndian, expectedU8, false},
		{uint16(math.MaxInt16), binary.BigEndian, expectedBigU16, false},
		{uint32(math.MaxInt32), binary.BigEndian, expectedBigU32, false},
		{uint64(math.MaxInt64), binary.BigEndian, expectedBigU64, false},
		{[]byte("one time"), binary.BigEndian, []byte("one time"), false},
		{"one more time", binary.BigEndian, []byte("one more time"), false},
		// errors
		{testing.T{}, binary.LittleEndian, []byte{}, true},
		{int8(-8), binary.LittleEndian, []byte{}, true},
		{int16(-136), binary.LittleEndian, []byte{}, true},
		{int32(-2184), binary.LittleEndian, []byte{}, true},
		{int64(-34952), binary.LittleEndian, []byte{}, true},
		{true, binary.LittleEndian, []byte{}, true},
		{false, binary.LittleEndian, []byte{}, true},
		{nil, binary.LittleEndian, []byte{}, true},
	}

	for i, test := range tests {
		errmsg := qt.Commentf("i = %d", i) // assert helper message

		v, err := ToBytesE(test.input, test.byteOrder)
		if test.iserr {
			c.Assert(err, qt.IsNotNil, errmsg)
			continue
		}

		c.Assert(err, qt.IsNil, errmsg)
		c.Assert(v, qt.DeepEquals, test.expect, errmsg)

		// Non-E test
		v = ToBytes(test.input, test.byteOrder)
		c.Assert(v, qt.DeepEquals, test.expect, errmsg)
	}
}
