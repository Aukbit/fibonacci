#Fibonacci [![Circle CI](https://circleci.com/bb/_paulo/fibonacci.svg?style=svg)](https://circleci.com/bb/_paulo/fibonacci)

Test and benchmark several implementations of the fibonacci numbers in Go

```
$ go test . -bench=.
BenchmarkFiboRecursive-8         3000000               485 ns/op
BenchmarkFiboClosure-8          20000000                85.4 ns/op
BenchmarkFiboConcurrent-8        1000000              1251 ns/op
```
