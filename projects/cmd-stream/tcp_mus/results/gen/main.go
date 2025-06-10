package main

import (
	"os"
	"reflect"

	"github.com/cmd-stream/core-go"
	musgen "github.com/mus-format/musgen-go/mus"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	assert "github.com/ymz-ncnk/assert/panic"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/projects/cmd-stream/tcp_mus/results"
)

func init() {
	assert.On = true
}

func main() {
	g, err := musgen.NewCodeGenerator(
		genops.WithPkgPath("github.com/ymz-ncnk/go-client-server-communication-benchmarks/projects/cmd-stream/tcp_mus/results"),
		genops.WithStream(),
	)
	assert.EqualError(err, nil)

	echoResultType := reflect.TypeFor[results.EchoResult]()
	err = g.AddStruct(echoResultType)
	assert.EqualError(err, nil)

	err = g.AddDTS(echoResultType)
	assert.EqualError(err, nil)

	err = g.AddInterface(reflect.TypeFor[core.Result](),
		introps.WithImpl(echoResultType),
		introps.WithMarshaller(),
	)
	assert.EqualError(err, nil)

	bs, err := g.Generate()
	assert.EqualError(err, nil)
	err = os.WriteFile("./mus-format.gen.go", bs, 0755)
	assert.EqualError(err, nil)
}
