package main

type BenchCase struct {
	Name string
	N    int
}

type BenchEntry struct {
	Iter int
	Ns   int64
}

type BenchData map[BenchCase][]BenchEntry

func BenchToQPSData(benchData BenchData) QPSData {
	var data QPSData = map[BenchCase]QPS{}
	for c, entries := range benchData {
		var (
			l       = len(entries)
			sumIter int
			sumNs   int64
		)
		if l == 0 {
			continue
		}
		for _, entry := range entries {
			sumIter += entry.Iter
			sumNs += entry.Ns
		}
		avgIter := float64(sumIter) / float64(l)
		avgNs := float64(sumNs) / float64(l)
		qps := (avgIter * float64(c.N)) / (avgNs / 1e9)
		data[c] = QPS(qps)
	}
	return data
}
