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
