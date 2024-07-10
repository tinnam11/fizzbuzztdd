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
