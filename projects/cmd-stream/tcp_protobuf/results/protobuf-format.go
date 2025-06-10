package results

import (
	"io"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
	"google.golang.org/protobuf/proto"
)

var EchoResultMUS = echoResultMUS{}

type echoResultMUS struct{}

func (s echoResultMUS) Marshal(v EchoResult, w muss.Writer) (n int, err error) {
	bs, err := proto.Marshal(v.ProtoData)
	if err != nil {
		return
	}
	n, err = varint.PositiveInt.Marshal(len(bs), w)
	if err != nil {
		return
	}
	n, err = w.Write(bs)
	return
}

func (s echoResultMUS) Unmarshal(r muss.Reader) (v EchoResult, n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		bs = make([]byte, l)
		n1 int
	)
	n1, err = io.ReadFull(r, bs)
	n += n1
	if err != nil {
		return
	}
	data := common.ProtoData{}
	err = proto.Unmarshal(bs, &data)
	if err != nil {
		return
	}
	v = EchoResult{ProtoData: &data}
	return
}

func (s echoResultMUS) Size(result EchoResult) (size int) {
	panic("not implemented")
}

func (s echoResultMUS) Skip(r muss.Reader) (n int, err error) {
	panic("not implemented")
}
