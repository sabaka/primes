package eratosthenes

var (
	FRAME_SIZE = 1000
)

func PrimesUnder(limit int) []int {
	numbers := filledSlice(0, limit)
	primes := make([]int, 0)
	for i := 2; i < limit; i++ {
		if numbers[i] != 0 {
			for j := 2; i*j < limit; j++ {
				numbers[i*j] = 0
			}
			primes = append(primes, numbers[i])
		}
	}
	return primes
}

func Prime(index int) int {
	return Primes(index)[index-1]
}

func Primes(size int) []int {
	start, end := 0, FRAME_SIZE-1
	sieve := filledSlice(start, FRAME_SIZE)
	sieve[1] = 0
	primes := make([]int, 0)
	for {

		if len(primes) > 0 {
			for _, vp := range primes {
				for i, vs := range sieve {
					if vs%vp == 0 {
						sieve[i] = 0
					}
				}
			}
		}

		for i := 0; i < FRAME_SIZE; i++ {
			if sieve[i] > 0 {
				for j := sieve[i] * sieve[i]; j < FRAME_SIZE; j += sieve[i] {
					sieve[j] = 0
				}
				primes = append(primes, sieve[i])
				if len(primes) == size {
					return primes
				}
			}
		}
		start += FRAME_SIZE
		end += FRAME_SIZE
		sieve = filledSlice(start, end)
	}
	return primes
}

func filledSlice(firstNumber, size int) []int {
	res := make([]int, 0)
	for i := 0; i < size; i++ {
		res = append(res, firstNumber+i)
	}
	return res
}
