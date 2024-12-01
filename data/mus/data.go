package mus

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/raw"
	"github.com/mus-format/mus-stream-go/unsafe"
	"github.com/mus-format/mus-stream-go/varint"
)

type Data struct {
	Bool    bool
	Int64   int64
	String  string
	Float64 float64
}

func MarshalDataMUS(d Data, w muss.Writer) (n int, err error) {
	_, err = ord.MarshalBool(d.Bool, w)
	if err != nil {
		return
	}
	_, err = varint.MarshalInt64(d.Int64, w)
	if err != nil {
		return
	}
	_, err = unsafe.MarshalString(d.String, nil, w)
	if err != nil {
		return
	}
	_, err = raw.MarshalFloat64(d.Float64, w)
	return
}

func UnmarshalDataMUS(r muss.Reader) (d Data, n int, err error) {
	d.Bool, _, err = ord.UnmarshalBool(r)
	if err != nil {
		return
	}
	d.Int64, _, err = varint.UnmarshalInt64(r)
	if err != nil {
		return
	}
	d.String, _, err = unsafe.UnmarshalString(nil, r)
	if err != nil {
		return
	}
	d.Float64, _, err = raw.UnmarshalFloat64(r)
	return
}

func SizeDataMUS(d Data) (size int) {
	size += ord.SizeBool(d.Bool)
	size += varint.SizeInt64(d.Int64)
	size += unsafe.SizeString(d.String, nil)
	return size + raw.SizeFloat64(d.Float64)
}
