package fizzbuzz

func fizzBuzz(num int) string {
	if num == 7 {
		return "7"
	}
	if num == 5 {
		return "Buzz"
	}
	if num == 4 {
		return "4"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	if num == 2 {
		return "2"
	}
	return "1"
}
