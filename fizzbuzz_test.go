package main

import ("testing"
		"fmt"
		"bytes"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("1. Entering a number not fizz nor buzz", func(t *testing.T) {
		got := Fizzbuzz(1)
		want := "1"
		assertCorrectMessage(t,got,want)
	})
	t.Run("2. Entering a fizz number", func(t *testing.T) {
		got := Fizzbuzz(3)
		want := "Fizz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("3. Entering a buzz number", func(t *testing.T) {
		got := Fizzbuzz(5)
		want := "Buzz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("4. Entering a fizzbuzz number", func(t *testing.T) {
		got := Fizzbuzz(15)
		want := "FizzBuzz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("5. Entering a fizz number containing a 3, but not divisible by 3", func(t *testing.T) {
		got := Fizzbuzz(13)
		want := "Fizz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("6. Entering a buzz number containing a 5, but not divisible by 5", func(t *testing.T) {
		got := Fizzbuzz(52)
		want := "Buzz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("7. Entering a fizzbuzz number containing a 5, divisible by 3, but not divisible by 5", func(t *testing.T) {
		got := Fizzbuzz(51)
		want := "FizzBuzz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("8. Entering a fizzbuzz number containing a 3, divisible byb 5, but not divisible by 3", func(t *testing.T) {
		got := Fizzbuzz(130)
		want := "FizzBuzz"
		assertCorrectMessage(t,got,want)
	})
	t.Run("9. Entering a fizzbuzz number containing a 3 and a 5 but divisible by neither", func(t *testing.T) {
		got := Fizzbuzz(358)
		want := "FizzBuzz"
		assertCorrectMessage(t,got,want)
	})
}

func TestCompute(t *testing.T){
	buffer := bytes.Buffer{}
	Compute(&buffer)
	got := buffer.String()
	want := ""
	for i := 0; i < 100; i++ {
		want += fmt.Sprintf("%d => %q\n",i+1,Fizzbuzz(i+1))
	}
	assertCorrectMessage(t,got,want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}