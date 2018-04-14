package cryptopals_challenges_go

import (
	"strings"
	"bytes"
)

var hexChars = [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func decodeHex(hexString string) []byte {
	if (len(hexString) % 2) > 0 {
		panic("Uneven Hex String")
	}
	var hexMap = make(map[string]uint8)
	for i := 0; i < 16; i = i + 1 {
		hexMap[hexChars[i]] = uint8(i)
	}
	var bytearray = make([]byte, len(hexString)/2)
	var stringChars = strings.Split(hexString, "")
	for i, j := 0, 0; i < len(stringChars)-1; i, j = i+2, j+1 {
		firstHex, secondHex := hexMap[strings.ToUpper(stringChars[i])], hexMap[strings.ToUpper(stringChars[i+1])]
		bytearray[j] = uint8(firstHex<<4 | secondHex)
	}
	return bytearray
}

var base64Chars = map[int]string{
	0:  "A", 16: "Q", 32: "g", 48: "w",
	1:  "B", 17: "R", 33: "h", 49: "x",
	2:  "C", 18: "S", 34: "i", 50: "y",
	3:  "D", 19: "T", 35: "j", 51: "z",
	4:  "E", 20: "U", 36: "k", 52: "0",
	5:  "F", 21: "V", 37: "l", 53: "1",
	6:  "G", 22: "W", 38: "m", 54: "2",
	7:  "H", 23: "X", 39: "n", 55: "3",
	8:  "I", 24: "Y", 40: "o", 56: "4",
	9:  "J", 25: "Z", 41: "p", 57: "5",
	10: "K", 26: "a", 42: "q", 58: "6",
	11: "L", 27: "b", 43: "r", 59: "7",
	12: "M", 28: "c", 44: "s", 60: "8",
	13: "N", 29: "d", 45: "t", 61: "9",
	14: "O", 30: "e", 46: "u", 62: "+",
	15: "P", 31: "f", 47: "v", 63: "/",
}

func encodeBase64(bytearray []byte) string {
	var buffer bytes.Buffer
	groupsOfThreeBytes := (len(bytearray)/3) * 3
	for i := 0; i < groupsOfThreeBytes; i = i + 3 {
		b1, b2, b3 := bytearray[i], bytearray[i+1], bytearray[i+2]
		chunk := uint32(uint32(b1)<<16 | uint32(b2)<<8 | uint32(b3))
		firstChar, secondChar, thirdChar, fourthChar := chunk >> 18, (chunk >> 12) & 0x3f, (chunk >> 6) & 0x3f, chunk & 0x3f
		buffer.WriteString(base64Chars[int(firstChar)])
		buffer.WriteString(base64Chars[int(secondChar)])
		buffer.WriteString(base64Chars[int(thirdChar)])
		buffer.WriteString(base64Chars[int(fourthChar)])
	}

	remainingBytes := len(bytearray) - groupsOfThreeBytes

	if remainingBytes == 0 {
		return buffer.String()
	}

	finalChunk := uint32(bytearray[groupsOfThreeBytes]) << 16

 	if remainingBytes == 2 {
		finalChunk |= uint32(bytearray[groupsOfThreeBytes + 1]) << 8
	}
	buffer.WriteString(base64Chars[int(finalChunk >> 18)])
	buffer.WriteString(base64Chars[int((finalChunk >> 12) & 0x3f)])

	if remainingBytes == 1 {
		buffer.WriteString("==")
		return buffer.String()
	} else if remainingBytes == 2 {
		buffer.WriteString(base64Chars[int(finalChunk >> 6 & 0x3f)])
		buffer.WriteString("=")
		return buffer.String()
	}
	return buffer.String()
}

func hexToBase64(hexString string) string {
	return encodeBase64(decodeHex(hexString))
}

func fixedXor(left []byte, right []byte) []byte {
	if len(left) != len(right) {
		panic("Uneven input arrays")
	} else if len(left) == 0 {
		return make([]byte, 0)
	}
	dst := make([]byte, len(left))
	for i := 0; i < len(left); i++ {
		dst[i] = left[i] ^ right[i]
	}
	return dst
}