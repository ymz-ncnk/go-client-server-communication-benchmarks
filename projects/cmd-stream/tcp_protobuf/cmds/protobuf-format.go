package cmds

import (
	"io"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
	"google.golang.org/protobuf/proto"
)

var EchoCmdMUS = echoCmdMUS{}

type echoCmdMUS struct{}

func (s echoCmdMUS) Marshal(v EchoCmd, w muss.Writer) (n int, err error) {
	bs, err := proto.Marshal(v.ProtoData)
	if err != nil {
		return
	}
	n, err = varint.PositiveInt.Marshal(len(bs), w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = w.Write(bs)
	n += n1
	return
}

func (s echoCmdMUS) Unmarshal(r muss.Reader) (v EchoCmd, n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		n1 int
		bs = make([]byte, l)
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
	v = EchoCmd{ProtoData: &data}
	return
}

func (s echoCmdMUS) Size(v EchoCmd) (size int) {
	panic("not implemented")
}

func (s echoCmdMUS) Skip(r muss.Reader) (n int, err error) {
	panic("not implemented")
}
