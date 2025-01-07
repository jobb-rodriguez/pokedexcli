package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input     string
    expected  []string
  }{
    {
      input:    "  hello world  ",
      expected: []string{"hello", "world"},
    },
    {
      input:    " pokedex pikachu CHARMANDER",
      expected: []string{"pokedex", "pikachu", "charmander"},
    },
    {
      input:    "ash  ketchum  pallet  town",
      expected: []string{"ash", "ketchum", "pallet", "town"},
    },
  }
  
  for _, c := range cases {
    actual, err := cleanInput(c.input)

    if err == nil && len(actual) != len(c.expected) {
      t.Errorf("Expected: %v; Actual: %v", len(actual), len(c.expected))
      t.Fail()
    }

    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]

      if err == nil && word != expectedWord {
        t.Errorf("Expected: %v; Actual: %v", expectedWord, word)
        t.Fail()
      }
    } 
  }
}
