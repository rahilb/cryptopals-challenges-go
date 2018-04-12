package cryptopals_challenges_go

import "strings"

var hexChars = [16]string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}

func decodeHex(hexString string) []byte {
	if (len(hexString) % 2) > 0 {
		panic("Uneven Hex String")
	}
	var hexMap = make(map[string]uint8)
	for i := 0; i < 16; i = i+1 {
		hexMap[hexChars[i]] = uint8(i)
	}
	var bytes = make([]byte, len(hexString) / 2)
	var stringChars = strings.Split(hexString,"")
	for i, j := 0, 0; i < len(stringChars) - 1; i, j = i+2, j+1 {
		firstHex := hexMap[strings.ToUpper(stringChars[i])]
		secondHex := hexMap[strings.ToUpper(stringChars[i + 1])]
		bytes[j] = uint8(firstHex << 4 | secondHex)
	}
	return bytes
}

func encodeBase64(bytes []byte) string {
	return "Unimplemented"
}

func hexToBase64(hexString string) string {
	return encodeBase64(decodeHex(hexString))
}