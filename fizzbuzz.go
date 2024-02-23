package main

import ("fmt"
		"strings"
		"io"
		"os"
)

func Fizzbuzz(a int) string {
	// Test 1 & 2
	/*if a==1 {
		return "1"
	}
	if a==3 {
		return "3 => Fizz"
	}
	return ""*/
	// +Test 3 & 4
	/*fizzbuzz := ""
	if a%3 == 0 && a%5 == 0 {
		fizzbuzz = "FizzBuzz"
	} else if a%3 == 0 {
		fizzbuzz = "Fizz"
	} else if a%5 == 0 {
		fizzbuzz = "Buzz"
	} else {
		fizzbuzz = fmt.Sprintf("%d",a)
	}
	return fizzbuzz*/
	//Refactor Test 4
	/*fizzbuzz := ""
	if a%3 == 0 {
		fizzbuzz += "Fizz"
	} 
	if a%5 == 0 {
		fizzbuzz += "Buzz"
	}
	if len(fizzbuzz) == 0 {
		fizzbuzz = fmt.Sprintf("%d",a)
	}
	return fizzbuzz*/
	// Tests 5 à 10 (tous les cas de figure sont testés)
	/*fizzbuzz := ""
	if a%3 == 0 {
		fizzbuzz += "Fizz"
	} else {
		toString := fmt.Sprintf("%d",a)
		if strings.Contains(toString, "3") {
			fizzbuzz += "Fizz"
		}
	}
	if a%5 == 0 {
		fizzbuzz += "Buzz"
	} else {
		toString := fmt.Sprintf("%d",a)
		if strings.Contains(toString, "5") {
			fizzbuzz += "Buzz"
		}
	}
	if len(fizzbuzz) == 0 {
		fizzbuzz = fmt.Sprintf("%d",a)
	}
	return fizzbuzz*/
	//On refactorise le Fizz et le Buzz chacune dans une fonction séparée pour plus de lisibilité
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