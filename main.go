package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// first i take in the arguments from the command line and only focus on the indexes 1 continuing
	args := os.Args[1:]

	// then i read the contents of the fisrt file i.e index 0
	text, _ := os.ReadFile(args[0])

	// splits the string text in to strings and store them in n array of strings called words
	words := strings.Split(string(text), " ")

	// iterate throung the whole array of strings
	for i, ch := range words {
		//if there is a string containing (hex) in the array...
		if ch == "(hex)" {
			//...a variable of int called new will be created where and the string before the "(hex)" will be converted to an int from base(16)(hexadecimal) to base(64) and stored in ne
			new, _ := strconv.ParseInt(words[i-1], 16, 64)
			// new will be converted form int to string and stored in the string at index before "(hex)"
			words[i-1] = fmt.Sprint(new)
			// the strings at the index before "(hex)" and "(hex)" itself will be replaces by out new string from above 
			words = append(words[:i], words[i+1:]...)
		//if there is a string containing (bin) in the array...
		} else if ch == "(bin)" {
			//...a variable of int called new will be created where and the string before the "(bin)" will be converted to an int from binary to base(64) and stored in ne
			new, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = fmt.Sprint(new)
			words = append(words[:i], words[i+1:]...)
		}
	}
	// call the functions while updating our array
	output := up(words)
	output1 := cap(output)
	output2 := low(output1)
	output3 := vow(output2)
	output4 := punc(output3)

	// storing my output in the file in the argument at index 2 in the command line
	// first join the array to one string using spaces 
	final := strings.Join(output4, " ")
	// make byte slice of final and store in new
	new := []byte(final)
	//use write file to store new in the arg at inex 1 with permision 644 
	os.WriteFile(args[1], new, 0o644)
}

// this is the function that capitalizes the first letter of each word before the letter (cap)
func cap(s []string) []string {
	for i := 0; i < len(s); i++ {
		// this if checks if a cap is appearring anywhere in the array
		if strings.Contains(s[i], "(cap") {
			// this if checks if the cap has a comma preceding it and that means it has an integer
			if strings.Contains(s[i], "(cap,") {
				// since the next index contains the number which we want and a bracket, the bracket must be removed and the integer stored in variable number
				number, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				// an iteration is done from th point which 'number' is subtracted from i(which is the index of our "(cap," ) upto the index before i 
				for j := i - number; j < i; j++ {
					// the first letter of all strings included in the iteration are converted to capital letters
					s[j] = strings.Title(s[j])
				}
				// remove the two strings I.e "(cap," and "6)"
				s = append(s[:i], s[i+2:]...)
			// thid part will work if it finds "(cap)""
			} else {
				// it will change the first letter of the previous index to capital letter
				s[i-1] = strings.Title(s[i-1])
				//this will remove "(cap)"
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	// a modified array of strings is retuned
	return s
}

// this func changes the word before "low" to a lower case using the same logic in function cap
func low(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(low") {
			if strings.Contains(s[i], "(low,") {
				number, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for j := i - number; j < i; j++ {
					s[j] = strings.ToLower(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToLower(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

// this func changes the word before "up" to upper case using the same logic in function cap
func up(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(up") {
			if strings.Contains(s[i], "(up,") {
				number, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for j := i - number; j < i; j++ {
					s[j] = strings.ToUpper(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

// this function changes "a" before a vowel and h to "an" 
func vow(s []string) []string {
	vowels := []byte{'a', 'e', 'i', 'u', 'o', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}

	for i, word := range s {
		for _, letter := range vowels {
			// checks if a word is a(small letter) and if the first letter of the preceeding word is a vowel
			if word == "a" && s[i+1][0] == letter {
				// changes the a to an
				s[i] = "an"
			// checks if a word is A(capital letter) and if the first letter of the preceeding word is a vowel
			} else if word == "A" && s[i+1][0] == letter {
				// changes the A to An
				s[i] = "An"
			}
		}
	}
	return s
}

// this functino handles punctuations correctly
func punc(s []string) []string {
	// define the punctuations and store then in the variable puns  as a string
	puns := []string{",",".","!","?",":",";"}

	// this pat checks if at the begiing of a string in an array is a puctuation and moves the punctuation to the previous string
	for i, word := range s {
		for _, puncs := range puns {
			// this checks if the beggining of a string is a punctuation and the end of it is not a punctuiation
			if string(word[0]) == puncs && string(word[len(word)-1]) != puncs {
				// this part adds the punctuation to the end of the previous string
				s[i-1] = s[i-1] + puncs
				// this part removes the punctuation at the beggining of the current word
				s[i] = word[1:]
			}
		}
	}
	
	// corrects punctuations at the end of the string
	for i, word := range s {
		for _, puncs := range puns {
			// checks if the first index if the current string is a punctuation and if the current string is at the end of the array
			if (string(word[0]) == puncs) && (s[len(s)-1] == s[i]) {
				s[i-1] = s[i-1] + word
				// removes the last part of the array
				s = s[:len(s)-1]
			}
		}
	}

	// checks if the first index of the string is a punctuation and the last index of the string is a punctuation and the string is not at the end of the string
	for i, word := range s {
		for _, puncs := range puns {
			if string(word[0]) == puncs && string(word[len(word)-1]) == puncs && s[i] != s[len(s)-1]  {
				s[i-1] = s[i-1] + word
				// removes the current index
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	// deals with the first apostrophe
	count := 0
	for i, word := range s {
		// checks if the current string in the array is an apostrophe and if the count is still at 0
		if word == "'" && count == 0 {
			// increase the count to that this is no repeated again
			count = count + 1
			// add the apostrophe to the next word(string) and store it in the next string
			s[i+1] = word + s[i+1]
			// remove the apostrophe
			s = append(s[:i], s[i+1:]...)
		}
	}

	// handles the second apostophe
	for i, word := range s {
		// checks if the current position is an apostrophe
		if word == "'" {
			// add the apostrophe to the previous word and store it at the previous index
			s[i-1] = s[i-1] + word
			//removes the apostrophe
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
