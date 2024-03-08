package piscine

func LastRune(s string) rune {
  size := len(s)
  slice :=[]rune(s)
  return slice[size -1]
}