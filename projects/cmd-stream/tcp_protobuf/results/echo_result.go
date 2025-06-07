package results

import "github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"

type EchoResult struct {
	*common.ProtoData
}

func (r EchoResult) LastOne() bool {
	return true
}
