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