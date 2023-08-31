# bench-go-slices-libs
benchmarking for slice helper

```text
Go Funk - Contains

Starting...  1
goos: darwin
goarch: arm64
pkg: github.com/vitordm/bench-go-slices-libs
BenchmarkContainsNumberWithFunk-2   	Starting...  100
Starting...  10000
Starting...  1000000
Starting...  17072313
17072313	       333.8 ns/op	     144 B/op	      12 allocs/op
BenchmarkContainsNumberWithFunk-2   	Starting...  1
Starting...  100
Starting...  10000
Starting...  1000000
Starting...  17961639
17961639	       333.8 ns/op	     144 B/op	      12 allocs/op
PASS
ok  	github.com/vitordm/bench-go-slices-libs	12.973s
```

```text
slice - Contains

Starting...  1
goos: darwin
goarch: arm64
pkg: github.com/vitordm/bench-go-slices-libs
BenchmarkContainsNumberWithSlice-2   	Starting...  100
Starting...  10000
Starting...  1000000
Starting...  100000000
Starting...  1000000000
1000000000	         4.035 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsNumberWithSlice-2   	Starting...  1
Starting...  100
Starting...  10000
Starting...  1000000
Starting...  100000000
Starting...  1000000000
1000000000	         4.037 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/vitordm/bench-go-slices-libs	9.502s
```
