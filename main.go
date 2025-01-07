package main

import (
  "fmt"
  "strings"
  "errors"
)

func cleanInput(text string) ([]string, error) {
  text = strings.Trim(text, " ")
  words := strings.Split(text, " ")

  for i, _  := range words {
    words[i] = strings.ToLower(words[i])
    words[i] = strings.Trim(words[i], " ")

    if len(words[i]) == 0 ||  strings.Contains(words[i], " ") {
      err := errors.New("Empty string")
      return []string{}, err
    }
  }
  return words, nil
}

func main() {
  text := "  Hey  there  "
  cleanedText, err := cleanInput(text);

  if  err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  
  for _, w := range cleanedText {
    fmt.Println(w, len(w))
  }
}
