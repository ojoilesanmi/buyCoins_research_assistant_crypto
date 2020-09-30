package main

import "fmt"

func main() {
	//return false
	fmt.Println(isProthPrime(25))
	fmt.Println(isProthPrime(73))
	fmt.Println(isProthPrime(19))

	//return true
	fmt.Println(isProthPrime(3))
	fmt.Println(isProthPrime(5))
	fmt.Println(isProthPrime(13))
	fmt.Println(isProthPrime(17))
	fmt.Println(isProthPrime(41))
	fmt.Println(isProthPrime(97))
	fmt.Println(isProthPrime(113))
	fmt.Println(isProthPrime(193))
	fmt.Println(isProthPrime(241))
	fmt.Println(isProthPrime(257))
	fmt.Println(isProthPrime(353))
	fmt.Println(isProthPrime(449))
	fmt.Println(isProthPrime(577))
	fmt.Println(isProthPrime(641))
	fmt.Println(isProthPrime(673))
	fmt.Println(isProthPrime(769))
	fmt.Println(isProthPrime(929))
	fmt.Println(isProthPrime(1153))
	fmt.Println(isProthPrime(1217))
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPowerOfTwo(n int) bool {

	return n != 0 && ((n & (n - 1)) == 0)

}

func isProthNumber(n int) bool {

	if n <= 1 {
		return false
	}

	n = n - 1

	k := 1

	for k < (n / k) {
		if n%k == 0 {
			if isPowerOfTwo(n / k) {
				return true
			}
		}

		k = k + 2
	}

	return false

}

func isProthPrime(n int) bool {

	if isPrime(n) && isProthNumber(n) {
		return true
	}
	return false
}
