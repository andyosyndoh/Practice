package piscine

func StrLen(s string) int {
  count := 0
  for i := 1; i < len(s); i++ {
    count++
  }
  return count
}
