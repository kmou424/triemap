## TrieMap

A Go implementation of a TrieMap, a data structure that maps keys to values using a trie data structure.


### Roadmap

- [x] Implement basic functions of TrieMap
  - [x] Put
  - [x] Get
  - [ ] Delete
  - [ ] More...
- [ ] **Performance improvements**
- [x] Add tests
- [x] Add benchmarks
- [ ] Add more documentation


### Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/kmou424/triemap
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkTrieMap_Put
BenchmarkTrieMap_Put-12    	 1110148	      1256 ns/op
BenchmarkTrieMap_Get
BenchmarkTrieMap_Get-12    	22554152	        53.78 ns/op
BenchmarkSyncMap_Put
BenchmarkSyncMap_Put-12    	 1000000	      1393 ns/op
BenchmarkSyncMap_Get
BenchmarkSyncMap_Get-12    	57608248	        21.65 ns/op
BenchmarkMap_Put
BenchmarkMap_Put-12        	 1829899	       662.5 ns/op
BenchmarkMap_Get
BenchmarkMap_Get-12        	305837685	         3.887 ns/op
PASS
```