## Getting Started

An out of box example. [![playground][play-simple-img]][play-simple]
```go
package main

import (
	"github.com/phuslu/goid"
)

func main() {
	println(goid.Goid())
}

// Output:
//   1
```

### Performance

```go
// go test -v -cpu=1 -run=none -bench=. -benchtime=10s -benchmem bench_test.go
package main

import (
	"testing"

	choleraehyq "github.com/choleraehyq/pid"
	petermattis "github.com/petermattis/goid"
	phuslu "github.com/phuslu/goid"
)

func BenchmarkCholeraehyq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		choleraehyq.Get()
	}
}

func BenchmarkPetermattis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		petermattis.Get()
	}
}

func BenchmarkPhuslu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		phuslu.Goid()
	}
}
```

A Performance result as below, for daily benchmark results see [github actions][benchmark]
```
goos: linux
goarch: amd64
pkg: bench
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkCholeraehyq
BenchmarkCholeraehyq 	1000000000	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkPetermattis
BenchmarkPetermattis 	1000000000	         1.858 ns/op	       0 B/op	       0 allocs/op
BenchmarkPhuslu
BenchmarkPhuslu      	1000000000	         1.858 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	bench	6.151s
```

[play-simple-img]: https://img.shields.io/badge/playground-VNOBbotvDd6-29BEB0?style=flat&logo=go
[play-simple]: https://go.dev/play/p/VNOBbotvDd6
[benchmark]: https://github.com/phuslu/goid/actions?query=workflow%3Abenchmark
