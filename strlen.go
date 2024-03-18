package piscine

func StrLen(s string) int {
  count := 0
  for i := 0; i < len(s); i++ {
    if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 32 && s[i] <= 33 ){
      count++
    }
  }
  return count
}
