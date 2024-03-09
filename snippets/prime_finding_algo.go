package snippets

import (
	"container/heap"
	"fmt"
	"time"
)

const LIMIT int = 1000

type PrimeMultiples struct {
	Prime    int
	Multiple int
}

type PrimesPool []*PrimeMultiples

// time complexity O(NLogN)
// space complexity O(N)
func TrialDivision() {
	start := time.Now()
	primes := make([]int, 0)
	var isPrime bool

	for i := 2; i <= LIMIT; i++ {
		isPrime = true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				// its not a prime
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}

	}
	fmt.Println("---Trial Division---")
	fmt.Printf("Space: %d KB\n", cap(primes)/1e3)
	fmt.Printf("Time: %f seconds\n", time.Since(start).Seconds())
}

// time complexity O(NLogLogN)
// space complexity O(N)
// https://iq.opengenus.org/sieve-of-eratosthenes-analysis/
func SieveOfEratos() {
	primes := make([]int, 0)
	sieve := make([]bool, LIMIT+1)
	start := time.Now()

	// assume all nums are prime
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false

	for i := 2; i*i <= LIMIT; i++ {
		if sieve[i] {
			for j := i * i; j <= LIMIT; j += i {
				// mark all multiples of i as false
				if sieve[j] {
					sieve[j] = false
				}
			}
		}
	}

	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	fmt.Println("---Sieve of Eratosthenes---")
	fmt.Printf("Space: %d KB\n", cap(primes)/1e3+cap(sieve)/1e3)
	fmt.Printf("Time: %f seconds\n", time.Since(start).Seconds())
}

// time complexity
// space complexity

func (pm PrimesPool) Len() int { return len(pm) }

func (pm PrimesPool) Less(i, j int) bool {
	// give the lowest current multiple available
	return pm[i].Multiple < pm[j].Multiple
}

func (pm PrimesPool) Swap(i, j int) {
	pm[i], pm[j] = pm[j], pm[i]
}

func (pm *PrimesPool) Push(x any) {
	item, ok := x.(*PrimeMultiples)
	if !ok {
		panic("invalid concrete type!")
	}
	*pm = append(*pm, item)
}

func (pm *PrimesPool) Pop() any {
	old := *pm
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pm = old[0 : n-1]
	return item
}

func DjisktraPrimAlgo() {
	start := time.Now()

	initPrim := &PrimeMultiples{
		Prime:    2,
		Multiple: 4,
	}
	primes := []int{2}
	primeHeap := make(PrimesPool, 0)
	primeHeap = append(primeHeap, initPrim)

	heap.Init(&primeHeap)

	// you always check the smallest multiple of the top
	// if it matches then obv its not a prime, we update the multiple of this number
	// if its less than the top then its a prime
	// if its greater then we need to update top to match any new multiple

	for i := 3; i <= LIMIT; i++ {
		// top of the heap
		top := (primeHeap)[0].Multiple

		if top > i {
			// its a prime, add it and its square to the pool(initially)
			primes = append(primes, i)
			heap.Push(&primeHeap, &PrimeMultiples{
				Prime:    i,
				Multiple: i * i,
			})
			continue
		}

		// i is greater than smallest multiple available
		for top <= i {
			item := heap.Pop(&primeHeap).(*PrimeMultiples)
			item.Multiple += item.Prime
			heap.Push(&primeHeap, item)
			top = (primeHeap)[0].Multiple
		}

	}

	fmt.Println("---Dijkstra\\'s Approach---")
	fmt.Printf("Space: %d KB\n", cap(primes)/1e3+cap(primeHeap)/1e3)
	fmt.Printf("Time: %f seconds\n", time.Since(start).Seconds())
}

func Prims() {
	// finding prime upto LIMIT
	// TrialDivision()
	// SieveOfEratos()
	DjisktraPrimAlgo()
}
