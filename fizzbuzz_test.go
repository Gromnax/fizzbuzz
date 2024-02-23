package main

import ("testing"
		"fmt"
		"bytes"
)

func TestFizzBuzz(t *testing.T) {
	//Étape 1
	/*
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
	})*/
	//Refactor des tests une fois l'étape 1 validée ; les tests sont dans l'ordre, le commentaire au dessus de chaque jeu.
	cases := []struct {
		Number int
		FizzBuzzed  string
	}{
		//Teste si un nombre non FizzBuzz est renvoyé tel quel
		{Number: 1, FizzBuzzed: "1"},
		{Number: 2, FizzBuzzed: "2"},
		//Teste si un nombre est Fizz
		{Number: 3, FizzBuzzed: "Fizz"},
		//Teste si un nombre est Buzz
		{Number: 5, FizzBuzzed: "Buzz"},
		//Teste si un nombre est FizzBuzz
		{Number: 15, FizzBuzzed: "FizzBuzz"},
		//Teste si 13 est bien Fizz (présence d'un 3, non divisible par 3)
		{Number: 13, FizzBuzzed: "Fizz"},
		//Teste si 52 est bien Buzz (présence d'un 5, non divisible par 5)
		{Number: 52, FizzBuzzed: "Buzz"},
		//Teste si 51 est bien FizzBuzz (présence d'un 5, divisible par 3, pas par 5)
		{Number: 51, FizzBuzzed: "FizzBuzz"},
		// Teste si 130 est bien FizzBuzz (présence d'un 3, divisible par 5, pas par 3)
		{Number: 130, FizzBuzzed: "FizzBuzz"},
		// Teste si 358 est bien FizzBuzz (présence d'un 3 et d'un 5, divisible ni par 3 ni par 5)
		{Number: 358, FizzBuzzed: "FizzBuzz"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Number, test.FizzBuzzed), func(t *testing.T) {
			got := Fizzbuzz(test.Number)
			assertCorrectMessage(t,got,test.FizzBuzzed)
		})
	}
	//Premier test de l'étape 2, ajouté à la série de tests.
	/*t.Run("5. 13 is now a Fizz number", func(t *testing.T) {
		got := Fizzbuzz(13)
		want := "Fizz"
		assertCorrectMessage(t,got,want)
	})*/
}

//Fonction pour tester l'étape 3
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