package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	} else if limit == 2 {
		return []int{2}
	}

	currentNum := 2
	numberField := make([]bool, limit+1)
	var primeNums []int

	for i := 0; i <= limit; i++ {
		// pos 0 and 1 do not actually matter
		numberField[i] = true
	}

	for {
		// if past the limit, return what you have
		if currentNum > limit {
			return primeNums
		}

		// if the current number is prime
		if numberField[currentNum] {
			// flag the current one as prime
			primeNums = append(primeNums, currentNum)

			// flag all multiples as not-prime
			for i := currentNum; i <= limit; i += currentNum {
				numberField[i] = false
			}
		}
		currentNum++
	}
}
