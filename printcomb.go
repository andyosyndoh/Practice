package piscine

import "github.com/01-edu/z01"

func PrintComb() {
  for i := '1'; i <='9'; i++ {
    for j := '1'; j <='9'; i++ {
      for k := '1'; k <= '9'; k++ {
        if i < j && j < k {
          z01.PrintRune(i)
          z01.PrintRune(j)
          z01.PrintRune(k)
          z01.PrintRune('\n')
        }
      }
    }
  }
}
