package fizzbuzz

import "testing"

func TestOne(t *testing.T) {
	num := 1
	expect := "1"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestTwo(t *testing.T) {
	num := 2
	expect := "2"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestThree(t *testing.T) {
	num := 3
	expect := "Fizz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestFour(t *testing.T) {
	num := 4
	expect := "4"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestFive(t *testing.T) {
	num := 5
	expect := "Buzz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestSix(t *testing.T) {
	num := 6
	expect := "Fizz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestSeven(t *testing.T) {
	num := 7
	expect := "7"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestEight(t *testing.T) {
	num := 8
	expect := "8"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestNine(t *testing.T) {
	num := 9
	expect := "Fizz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestTen(t *testing.T) {
	num := 10
	expect := "Buzz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestEleven(t *testing.T) {
	num := 11
	expect := "11"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestTwelve(t *testing.T) {
	num := 12
	expect := "Fizz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestThirteen(t *testing.T) {
	num := 13
	expect := "13"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestFourteen(t *testing.T) {
	num := 14
	expect := "14"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}

func TestFifteen(t *testing.T) {
	num := 15
	expect := "FizzBuzz"
	actual := fizzBuzz(num)
	if expect != actual {
		t.Errorf("given num (%d) expect %q actual %q\n", num, expect, actual)
	}
}
