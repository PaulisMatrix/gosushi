package snippets_test

import (
	"practice/snippets"
	"testing"
)

/*
All benchmarks will run at for 2s to find upto 1000

❯ go test -bench=. -test.bench BenchmarkTrialDivision -benchtime=2s
goos: darwin
goarch: arm64
pkg: practice/snippets
BenchmarkTrialDivision-8   	  489745 -> no of iterations to complete benchtime of 2s 	4744 ns/op -> avg time
PASS


❯ go test -bench=. -test.bench BenchmarkSieve -benchtime=2s
goos: darwin
goarch: arm64
pkg: practice/snippets
BenchmarkSieve-8   	  866653	      2747 ns/op
PASS


❯ go test -bench=. -test.bench BenchmarkDjisktraAlgo -benchtime=2s
goos: darwin
goarch: arm64
pkg: practice/snippets
BenchmarkDjisktraAlgo-8   	   16492	    147123 ns/op -> why this??? extra overhead in managing the heapq?
PASS

*/

func BenchmarkTrialDivision(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		snippets.TrialDivision()
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		snippets.SieveOfEratos()
	}
}

func BenchmarkDjisktraAlgo(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		snippets.DjisktraPrimAlgo()
	}
}

/*

In golang:
gosushi on  main [✘!?] via 🐹 v1.21.6 on ☁️  rushikesh@verloop.io(asia-south1)
❯ go run main.go
---Trial Division---
Space: 88 KB
Time: 0.074128 seconds

gosushi on  main [✘!?] via 🐹 v1.21.6 on ☁️  rushikesh@verloop.io(asia-south1)
❯ go run main.go
---Sieve of Eratosthenes---
Space: 1088 KB
Time: 0.004287 seconds

gosushi on  main [✘!?] via 🐹 v1.21.6 on ☁️  rushikesh@verloop.io(asia-south1)
❯ go run main.go
---Dijkstra\'s Approach---
Space: 176 KB
Time: 0.376415 seconds


In python:
❯ python3 prime_finding_algs.py
---Trial Division---
Space: 0.632824 MB
Time: 3.576855792 seconds


~/Downloads via 🐳 desktop-linux via 🐍 v3.9.17 on ☁️  rushikesh@verloop.io(asia-south1) took 3s
❯ python3 prime_finding_algs.py
---Sieve of Eratosthenes---
Space: 8.632888 MB
Time: 0.164537459 seconds


~/Downloads via 🐳 desktop-linux via 🐍 v3.9.17 on ☁️  rushikesh@verloop.io(asia-south1)
❯ python3 prime_finding_algs.py
---Dijkstra's Approach---
Space: 1.265648 MB
Time: 2.750727667 seconds
*/
