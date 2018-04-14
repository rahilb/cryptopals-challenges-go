package cryptopals_challenges_go

import (
	"reflect"
	"fmt"
	"testing"
	"encoding/base64"
	"encoding/hex"
	"bufio"
	"os"
	"log"
)

func TestHexToBase64(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}
	for _, c := range cases {
		got := hexToBase64(c.in)
		if got != c.want {
			t.Errorf("HexToBase64(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestEncodeBase64(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"Hello World"},
	}
	for _, c := range cases {
		bytearray := []byte(c.in)
		got := encodeBase64(bytearray)
		want := base64.StdEncoding.EncodeToString(bytearray)
		if got != want {
			t.Errorf("HexToBase64(%q) == %q, want %q", c.in, got, want)
		}
	}
}

func TestDecodeHex(t *testing.T) {
	cases := []struct {
		hexString string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"},
	}
	for _, in := range cases {
		got := decodeHex(in.hexString)
		expected, _ := hex.DecodeString(in.hexString)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("decodeHex(%q) == %q, want %q", in.hexString, got, expected)
		} else {
			fmt.Println(string(got))
		}
	}
}

func TestFixedXor(t *testing.T) {
	cases := []struct {
		input, xorInput, expected string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}
	for _, c := range cases {
		got := fixedXor(decodeHex(c.input), decodeHex(c.xorInput))
		expected, _ := hex.DecodeString(c.expected)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("fixedXor(%q, %q) == %q, want %q", c.input, c.xorInput, got, expected)
		} else {
			fmt.Println(string(got))
		}
	}
}

func TestFindSingleByteXor(t *testing.T) {
	plaintext := findSingleByteXor(decodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	fmt.Println(string(plaintext))
}

func TestChallenge4(t *testing.T) {
	file, err := os.Open("set1_challenge4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxScore := 0
	bestPlaintext := ""

	for scanner.Scan() {
		plaintext := findSingleByteXor(decodeHex(scanner.Text()))
		score := plaintextScore(string(plaintext))
		if score > maxScore {
			maxScore = score
			bestPlaintext = string(plaintext)
		}
	}

	fmt.Printf("Best Score: %v, Best Plaintext: %v", maxScore, bestPlaintext)

}