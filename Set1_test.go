package cryptopals_challenges_go

import "testing"
import "encoding/hex"
import "encoding/base64"
import (
	"reflect"
	"fmt"
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