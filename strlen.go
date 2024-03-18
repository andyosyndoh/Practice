package piscine

func StrLen(s string) int {
  count := 0
  for i := 0; i < len(s); i++ {
    if i >= 'A' || i <= 'z' {
      count++
    }
  }
  return count
}
