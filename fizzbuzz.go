package main

import ("fmt"
		"strings"
		"io"
		"os"
)

func Fizzbuzz(a int) string {
	fizzbuzz := ""
	fizzbuzz += Fizz(a)
	fizzbuzz += Buzz(a)
	if len(fizzbuzz) == 0 {
		fizzbuzz = fmt.Sprintf("%d",a)
	}
	return fizzbuzz
}

func Fizz(a int) string {
	value := ""
	if a%3 == 0 {
		value += "Fizz"
	} else {
		toString := fmt.Sprintf("%d",a)
		if strings.Contains(toString, "3") {
			value += "Fizz"
		}
	}
	return value
}

func Buzz(a int) string {
	value := ""
	if a%5 == 0 {
		value += "Buzz"
	} else {
		toString := fmt.Sprintf("%d",a)
		if strings.Contains(toString, "5") {
			value += "Buzz"
		}
	}
	return value
}

func Compute(out io.Writer) {
	for i := 0; i < 100; i++ {
		fmt.Fprintf(out, "%d => %q\n",i+1,Fizzbuzz(i+1))
	}
}

func main() {
	Compute(os.Stdout)
}