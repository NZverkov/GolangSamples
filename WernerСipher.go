package main

import (
	"fmt"
	"strings"
)

func main() {

	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	CompleteString := stringcreator(len(cipherText), len(keyword), keyword)
	fmt.Println(cipherText)
	fmt.Println(len(cipherText))
	fmt.Println(CompleteString)
	finisharray := decoder(CompleteString, cipherText)
	fmt.Print(finisharray)
}
func stringcreator(cipherTextlength int, keywordlength int, keyword string) string {
	bothstr := strings.Repeat(keyword, (cipherTextlength/keywordlength)+1)[:cipherTextlength]
	return bothstr
}
func decoder(CompleteString string, cipherText string) []string {

	dictionary := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25,
	}
	length := len(cipherText)
	var (
		finisharray    = make([]int, length)
		numArrayComp   = make([]int, length)
		numArrayCip    = make([]int, length)
		stringarrayfin = make([]string, length)
	)
	CompleteStringarr := strings.Split(CompleteString, "")

	for i := 0; i < len(CompleteStringarr); i++ {
		numArrayComp[i] = dictionary[CompleteStringarr[i]]
	}
	cipherTextarr := strings.Split(cipherText, "")
	for i := 0; i < len(cipherTextarr); i++ {
		numArrayCip[i] = dictionary[cipherTextarr[i]]
	}
	for i := 0; i < len(cipherTextarr); i++ {
		if (numArrayCip[i] - numArrayComp[i]) < 0 {
			finisharray[i] = numArrayCip[i] + len(dictionary) - numArrayComp[i]
		} else {
			finisharray[i] = (numArrayCip[i] - numArrayComp[i])
		}
	}
	for i := 0; i < len(cipherTextarr); i++ {
		stringarrayfin[i] = mapkey(dictionary, finisharray[i])

	}
	return stringarrayfin
}
func mapkey(m map[string]int, value int) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return key
		}
	}
	return
}
