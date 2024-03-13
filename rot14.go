package piscine

func Rot14(s string) string {
  new := make([]rune, len(s))
  for i, char := range s {
    if (char >= 'a' && char <= 'z') {
      if char + 14 > 'z' {
        new[i] = char + 14 - 26
      } else {
        new[i] = char + 14
      }
    } else {
      new[i] = char
    }
    if (char >= 'A' && char <= 'Z') {
      if char + 14 > 'Z' {
        new[i] = char + 14 -26
      } else {
        new[i] = char + 14
      }
    }
  }
  return string(new)
} 