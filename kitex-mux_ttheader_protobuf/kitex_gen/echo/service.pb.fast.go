// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package echo

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *KitexData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_KitexData[number], err)
}

func (x *KitexData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Bool, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *KitexData) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Int64, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *KitexData) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.String_, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *KitexData) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Float64, offset, err = fastpb.ReadDouble(buf, _type)
	return offset, err
}

func (x *KitexData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *KitexData) fastWriteField1(buf []byte) (offset int) {
	if !x.Bool {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetBool())
	return offset
}

func (x *KitexData) fastWriteField2(buf []byte) (offset int) {
	if x.Int64 == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetInt64())
	return offset
}

func (x *KitexData) fastWriteField3(buf []byte) (offset int) {
	if x.String_ == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetString_())
	return offset
}

func (x *KitexData) fastWriteField4(buf []byte) (offset int) {
	if x.Float64 == 0 {
		return offset
	}
	offset += fastpb.WriteDouble(buf[offset:], 4, x.GetFloat64())
	return offset
}

func (x *KitexData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *KitexData) sizeField1() (n int) {
	if !x.Bool {
		return n
	}
	n += fastpb.SizeBool(1, x.GetBool())
	return n
}

func (x *KitexData) sizeField2() (n int) {
	if x.Int64 == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetInt64())
	return n
}

func (x *KitexData) sizeField3() (n int) {
	if x.String_ == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetString_())
	return n
}

func (x *KitexData) sizeField4() (n int) {
	if x.Float64 == 0 {
		return n
	}
	n += fastpb.SizeDouble(4, x.GetFloat64())
	return n
}

var fieldIDToName_KitexData = map[int32]string{
	1: "Bool",
	2: "Int64",
	3: "String_",
	4: "Float64",
}
