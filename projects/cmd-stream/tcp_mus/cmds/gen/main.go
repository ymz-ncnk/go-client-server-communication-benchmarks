package main

import (
	"os"
	"reflect"

	"github.com/cmd-stream/core-go"
	musgen "github.com/mus-format/musgen-go/mus"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	assert "github.com/ymz-ncnk/assert/panic"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream/tcp_mus/cmds"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream/tcp_protobuf/receiver"
)

func init() {
	assert.On = true
}

func main() {
	g, err := musgen.NewCodeGenerator(
		genops.WithPkgPath("github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream/tcp_mus/cmds"),
		genops.WithStream(),
	)
	assert.EqualError(err, nil)

	echoCmdType := reflect.TypeFor[cmds.EchoCmd]()
	err = g.AddStruct(echoCmdType)
	assert.EqualError(err, nil)

	err = g.AddDTS(echoCmdType)
	assert.EqualError(err, nil)

	err = g.AddInterface(reflect.TypeFor[core.Cmd[receiver.Receiver]](),
		introps.WithImpl(echoCmdType),
		introps.WithMarshaller(),
	)
	assert.EqualError(err, nil)

	// Generate
	bs, err := g.Generate()
	assert.EqualError(err, nil)
	err = os.WriteFile("./mus-format.gen.go", bs, 0755)
	assert.EqualError(err, nil)
}
