package piscine

func StrLen(s string) int {
  count := 0
  for i := 0; i < len(s); i++ {
    if s[i] >= 32 && s[i] <= 126 {
      count++
    }
  }
  return count
}
