package cryptopals_challenges_go

func hammingDistanceByte(left byte, right byte) int {
	distance := 0
	xoredInput := left ^ right
	for xoredInput > 0 {
		xoredInput &= xoredInput - 1;
		distance += 1
	}
 	return distance
}

func hammingDistance(left []byte, right []byte) int {
	if len(left) != len(right) {
		panic("Can not calculate hamming distance between unequal arrays")
	}
	distance := 0
	for i := 0; i < len(left); i += 1 {
		distance += hammingDistanceByte(left[i], right[i])
	}
	return distance
}

func calculateKeySizes(input []byte, maxKeySize int) []float32 {
	if maxKeySize > (len(input) / 2) {
		panic("Max key size may not be greater than input length")
	}
	results := make([]float32, maxKeySize)
    for i := 1; i <= maxKeySize; i += 1 {
    	dist1 := float32(hammingDistance(input[0:i], input[i: i * 2])) / float32(i + 1)
    	dist2 := float32(hammingDistance(input[i * 2: i * 3], input[i * 3: i * 4])) / float32(i + 1)
		results[i - 1] = (dist1 + dist2) / 2
	}
	return results
}

func transposeBlocks(input []byte, keySize int) [][]byte {
	dst := make([][]byte, keySize)
	for keyPosition := 0; keyPosition < keySize; keyPosition = keyPosition + 1 {
		if keyPosition < (len(input) % keySize) {
			dst[keyPosition] = make([]byte, (len(input) / keySize) + 1)
		} else {
			dst[keyPosition] = make([]byte, len(input) / keySize)
		}
	}
	for i, j := 0, 1; i < len(input); i, j = i + 1, j + 1 {
		dst[i % keySize][int(i / keySize)] = input[i]
	}
	return dst
}