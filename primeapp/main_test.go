package main

import (
    "bufio"
    "io"
    "os"
    "strings"
    "testing"
)

func Test_IsPrime(t *testing.T) {
	primeTests := []struct {
		name       string
		testNumber int
		expected   bool
		msg        string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative", -1, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNumber)
		if e.expected && !result {
			t.Errorf("%s: expected true, got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false, got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s expected %s, got %s", e.name, e.msg, msg)
		}
	}
}

func Test_Prompt(t *testing.T) {
    // save a copy of os.Stdout
    oldOut := os.Stdout

    // create a read and write pipe
    r, w, _ := os.Pipe()

    // set os.Stdout to our write pipe
    os.Stdout = w

    prompt()

    // close our writer
    _ = w.Close()

    // reset os.Stdout to what it was before
    os.Stdout = oldOut

    // read the output of our prompt() fun from our read pipe
    out, _ := io.ReadAll(r)

    if string(out) != "-> " {
        t.Errorf("incorrect prompt: expected -> but got %s", string(out))
    }
}

func Test_Intro(t *testing.T) {
    // save a copy of os.Stdout
    oldOut := os.Stdout

    // create a read and write pipe
    r, w, _ := os.Pipe()

    // set os.Stdout to our write pipe
    os.Stdout = w

    intro()

    // close our writer
    _ = w.Close()

    // reset os.Stdout to what it was before
    os.Stdout = oldOut

    // read the output of our prompt() fun from our read pipe
    out, _ := io.ReadAll(r)

    if !strings.Contains(string(out), "Enter a number to see if it is prime.") {
        t.Errorf("intro text not correct; got %s", string(out))
    }
}

type checkNumbersTT struct {
    name string
    input string
    expected string
}

func Test_checkNumbers(t *testing.T) {
    tests := []checkNumbersTT{
        {name: "empty", input: "", expected: "Please enter a number or 'q' to quit the program"},
        {name: "typed", },
        {name: "zero", input: "0", expected: "0 is not prime, by definition!"},
        {name: "one", input: "1", expected: "1 is not prime, by definition!"},
        {name: "two", input: "2", expected: "2 is a prime number"},
        {name: "three", input: "3", expected: "3 is a prime number"},
        {name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
    }

    for _, tt := range tests {
        input := strings.NewReader(tt.input)
        reader := bufio.NewScanner(input)
        res, _ := checkNumbers(reader)

        if !strings.EqualFold(res, tt.expected) {
            t.Errorf("%s::expected::%s, but got %s", tt.name, tt.expected, res)
        }
    }
}
