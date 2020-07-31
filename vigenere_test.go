// Package vigenere enables encryption and decrytion of strings according to Vigenere system.
package vigenere_test

import (
	"fmt"
	"log"
	"testing"
	"vigenere"
)

func ExampleEncrypt() {
	s, err := vigenere.Encrypt("primocontatto", "pluto")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	// Output:
	// eccfcrzhmoiei
}

func ExampleDecrypt() {
	s, err := vigenere.Decrypt("eccfcrzhmoiei", "pluto")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	// Output:
	// primocontatto
}

func TestEncrypt(t *testing.T) {
	type args struct {
		text string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Number", args{key: "pluto", text: "contatto1"}, "", true},
		{"Uppercase", args{key: "pluto", text: "Contatto"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vigenere.Encrypt(tt.args.text, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		text string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Number", args{key: "pluto", text: "contatto1"}, "", true},
		{"Uppercase", args{key: "pluto", text: "Contatto"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vigenere.Decrypt(tt.args.text, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
