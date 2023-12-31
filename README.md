# Benchmark Results

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/egtann/strip-wildcard-prefix
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkStripWildcardPrefix-12          7089686               168.8 ns/op
BenchmarkHttpStripPrefix-12             13492386                88.49 ns/op
PASS
ok      github.com/egtann/strip-wildcard-prefix 2.657s
```
