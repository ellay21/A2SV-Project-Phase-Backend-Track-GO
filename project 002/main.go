package main
import "fmt"
func main(){

	//custom test for frequency count
	ans := freq_count("hello world hello")
	for k, v := range ans {
		fmt.Printf("%s: %d\n", k, v)
	}
	//custom test for palindrome to show the ignore cases and some punctuation i have used this test
	palindromeTest := "A man, a plan, a canal: Panama"
	fmt.Printf("%q is palindrome: %v\n", palindromeTest, isPalindrome(palindromeTest))
}