// Package vigenere enables encryption and decrytion of strings according to Vigenere system.
package vigenere

import (
	"fmt"
)

// Encrypt crypts strings according to Vigenere system.
func Encrypt(text, key string) (string, error) {

	tS, kS := []rune(text), []rune(key)

	// makes key as long as text
	var i int
	for len(tS) > len(kS) {
		kS = append(kS, kS[i])
		i++
	}

	var encrypted []rune
	for n := range tS {
		if tS[n] > 'z' || tS[n] < 'a' {
			return "", fmt.Errorf("runa non valida: %v", tS[n])
		}
		encrypted = append(encrypted, rune((tS[n]-97+kS[n]-97)%26)+97)
	}
	return string(encrypted), nil
}

// Decrypt decrypts strings according to Vigenere system.
func Decrypt(text, key string) (string, error) {

	tS, kS := []rune(text), []rune(key)

	// makes key as long as text
	var i int
	for len(tS) > len(kS) {
		kS = append(kS, kS[i])
		i++
	}

	var decrypted []rune
	for n := range tS {
		if tS[n] > 'z' || tS[n] < 'a' {
			return "", fmt.Errorf("runa non valida: %v", tS[n])
		}
		decrypted = append(decrypted, rune((tS[n]-kS[n]+26)%26)+97)
	}

	return string(decrypted), nil
}
