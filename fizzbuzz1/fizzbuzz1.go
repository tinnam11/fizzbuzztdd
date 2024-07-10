package fizzbuzz

import "strconv"

func fizzBuzz(num int) string {
	if num%5 == 0 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(num)
}
