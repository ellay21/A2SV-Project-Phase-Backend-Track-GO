package main

import (
	"strings"
	"unicode"
)
func freq_count(st string) map[string]int{
	freq:=make(map[string]int)
	slice_st := strings.Split(st, " ")
	for _,word:=range(slice_st){
		freq[word]++
	}
	return freq
}
func isPalindrome(st string) bool{
	nwst:=[]rune{}
	for i:=range st{
		if unicode.IsLetter(rune(st[i])) {
			nwst = append(nwst, unicode.ToLower(rune(st[i])))
		} else if unicode.IsDigit(rune(st[i])) {
			nwst = append(nwst, rune(st[i]))
		}}
	l,r := 0,len(nwst)-1
	for l<r { 
		if nwst[l] != nwst[r] {
			return false
		}
		l++
		r--
	}
	return true
}