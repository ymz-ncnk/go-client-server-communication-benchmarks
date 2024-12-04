package mus

// func GenerateDataFrom(g [][]data.Data) (d [][]Data) {
// 	d = make([][]Data, len(g))
// 	for i := 0; i < len(g); i++ {
// 		d[i] = make([]Data, len(g[i]))
// 		for j := 0; j < len(g[i]); j++ {
// 			d[i][j] = Data{
// 				Bool:    g[i][j].Bool,
// 				Int64:   g[i][j].Int64,
// 				String:  g[i][j].String,
// 				Float64: g[i][j].Float64,
// 			}
// 		}
// 	}
// 	return
// }

type DataSetMUS [][]Data
