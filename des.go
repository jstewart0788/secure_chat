package main

func desEnc(message string, key string) []string {
	//Permutation Initialization
	initialPermutation := getInitPerm()
	initialPermutationInv := getInitPermInv()
	pc1 := getPC1()
	pc2 := getPC2()

	//generate  keys
	k0 := pc1.runPermutation(stringToBinaryArray(key), 56)
	c0, d0 := k0[:28], k0[28:]
	c1 := rotate(c0, 1, LEFT)
	d1 := rotate(d0, 1, LEFT)
	c2 := rotate(c1, 1, LEFT)
	d2 := rotate(d1, 1, LEFT)
	c3 := rotate(c2, 2, LEFT)
	d3 := rotate(d2, 2, LEFT)
	c4 := rotate(c3, 2, LEFT)
	d4 := rotate(d3, 2, LEFT)
	c5 := rotate(c4, 2, LEFT)
	d5 := rotate(d4, 2, LEFT)
	c6 := rotate(c5, 2, LEFT)
	d6 := rotate(d5, 2, LEFT)
	c7 := rotate(c6, 2, LEFT)
	d7 := rotate(d6, 2, LEFT)
	c8 := rotate(c7, 2, LEFT)
	d8 := rotate(d7, 2, LEFT)
	c9 := rotate(c8, 1, LEFT)
	d9 := rotate(d8, 1, LEFT)
	c10 := rotate(c9, 2, LEFT)
	d10 := rotate(d9, 2, LEFT)
	c11 := rotate(c10, 2, LEFT)
	d11 := rotate(d10, 2, LEFT)
	c12 := rotate(c11, 2, LEFT)
	d12 := rotate(d11, 2, LEFT)
	c13 := rotate(c12, 2, LEFT)
	d13 := rotate(d12, 2, LEFT)
	c14 := rotate(c13, 2, LEFT)
	d14 := rotate(d13, 2, LEFT)
	c15 := rotate(c14, 2, LEFT)
	d15 := rotate(d14, 2, LEFT)
	c16 := rotate(c15, 1, LEFT)
	d16 := rotate(d15, 1, LEFT)

	k1_56 := append(c1, d1...)
	k2_56 := append(c2, d2...)
	k3_56 := append(c3, d3...)
	k4_56 := append(c4, d4...)
	k5_56 := append(c5, d5...)
	k6_56 := append(c6, d6...)
	k7_56 := append(c7, d7...)
	k8_56 := append(c8, d8...)
	k9_56 := append(c9, d9...)
	k10_56 := append(c10, d10...)
	k11_56 := append(c11, d11...)
	k12_56 := append(c12, d12...)
	k13_56 := append(c13, d13...)
	k14_56 := append(c14, d14...)
	k15_56 := append(c15, d15...)
	k16_56 := append(c16, d16...)

	k1 := pc2.runPermutation(k1_56, 48)
	k2 := pc2.runPermutation(k2_56, 48)
	k3 := pc2.runPermutation(k3_56, 48)
	k4 := pc2.runPermutation(k4_56, 48)
	k5 := pc2.runPermutation(k5_56, 48)
	k6 := pc2.runPermutation(k6_56, 48)
	k7 := pc2.runPermutation(k7_56, 48)
	k8 := pc2.runPermutation(k8_56, 48)
	k9 := pc2.runPermutation(k9_56, 48)
	k10 := pc2.runPermutation(k10_56, 48)
	k11 := pc2.runPermutation(k11_56, 48)
	k12 := pc2.runPermutation(k12_56, 48)
	k13 := pc2.runPermutation(k13_56, 48)
	k14 := pc2.runPermutation(k14_56, 48)
	k15 := pc2.runPermutation(k15_56, 48)
	k16 := pc2.runPermutation(k16_56, 48)

	//generate first message
	m0 := initialPermutation.runPermutation(stringToBinaryArray(message), 64)
	l0, r0 := m0[:32], m0[32:]
	//round 1
	r1 := xor(l0, function(r0, k1))
	l1 := r0
	//round2
	r2 := xor(l1, function(r1, k2))
	l2 := r1
	//round 3
	r3 := xor(l2, function(r2, k3))
	l3 := r2
	//round4
	r4 := xor(l3, function(r3, k4))
	l4 := r3
	//round 5
	r5 := xor(l4, function(r4, k5))
	l5 := r4
	//round6
	r6 := xor(l5, function(r5, k6))
	l6 := r5
	//round 7
	r7 := xor(l6, function(r6, k7))
	l7 := r6
	//round8
	r8 := xor(l7, function(r7, k8))
	l8 := r7
	//round 9
	r9 := xor(l8, function(r8, k9))
	l9 := r8
	//round10
	r10 := xor(l9, function(r9, k10))
	l10 := r9
	//round 11
	r11 := xor(l10, function(r10, k11))
	l11 := r10
	//round12
	r12 := xor(l11, function(r11, k12))
	l12 := r11
	//round 13
	r13 := xor(l12, function(r12, k13))
	l13 := r12
	//round14
	r14 := xor(l13, function(r13, k14))
	l14 := r13
	//round 15
	r15 := xor(l14, function(r14, k15))
	l15 := r14
	//round16
	r16 := xor(l15, function(r15, k16))
	l16 := r15
	//combine L and R
	mf := append(r16, l16...)
	cypherText := initialPermutationInv.runPermutation(mf, 64)
	return cypherText
}

func desDec(cypherText []string, key string) string {
	//Permutation Initialization
	initialPermutation := getInitPerm()
	initialPermutationInv := getInitPermInv()
	pc1 := getPC1()
	pc2 := getPC2()

	//generate  keys
	k0 := pc1.runPermutation(stringToBinaryArray(key), 56)
	// round 1
	c16, d16 := k0[:28], k0[28:]
	// round 2
	c15 := rotate(c16, 1, RIGHT)
	d15 := rotate(d16, 1, RIGHT)
	// round 3
	c14 := rotate(c15, 2, RIGHT)
	d14 := rotate(d15, 2, RIGHT)
	// round 4
	c13 := rotate(c14, 2, RIGHT)
	d13 := rotate(d14, 2, RIGHT)
	// round 5
	c12 := rotate(c13, 2, RIGHT)
	d12 := rotate(d13, 2, RIGHT)
	// round 6
	c11 := rotate(c12, 2, RIGHT)
	d11 := rotate(d12, 2, RIGHT)
	// round 7
	c10 := rotate(c11, 2, RIGHT)
	d10 := rotate(d11, 2, RIGHT)
	// round 8
	c9 := rotate(c10, 2, RIGHT)
	d9 := rotate(d10, 2, RIGHT)
	// round 9
	c8 := rotate(c9, 1, RIGHT)
	d8 := rotate(d9, 1, RIGHT)
	// round 10
	c7 := rotate(c8, 2, RIGHT)
	d7 := rotate(d8, 2, RIGHT)
	// round 11
	c6 := rotate(c7, 2, RIGHT)
	d6 := rotate(d7, 2, RIGHT)
	// round 12
	c5 := rotate(c6, 2, RIGHT)
	d5 := rotate(d6, 2, RIGHT)
	// round 13
	c4 := rotate(c5, 2, RIGHT)
	d4 := rotate(d5, 2, RIGHT)
	// round 14
	c3 := rotate(c4, 2, RIGHT)
	d3 := rotate(d4, 2, RIGHT)
	// round 15
	c2 := rotate(c3, 2, RIGHT)
	d2 := rotate(d3, 2, RIGHT)
	// round 16
	c1 := rotate(c2, 1, RIGHT)
	d1 := rotate(d2, 1, RIGHT)

	k1_56 := append(c1, d1...)
	k2_56 := append(c2, d2...)
	k3_56 := append(c3, d3...)
	k4_56 := append(c4, d4...)
	k5_56 := append(c5, d5...)
	k6_56 := append(c6, d6...)
	k7_56 := append(c7, d7...)
	k8_56 := append(c8, d8...)
	k9_56 := append(c9, d9...)
	k10_56 := append(c10, d10...)
	k11_56 := append(c11, d11...)
	k12_56 := append(c12, d12...)
	k13_56 := append(c13, d13...)
	k14_56 := append(c14, d14...)
	k15_56 := append(c15, d15...)
	k16_56 := append(c16, d16...)

	k1 := pc2.runPermutation(k1_56, 48)
	k2 := pc2.runPermutation(k2_56, 48)
	k3 := pc2.runPermutation(k3_56, 48)
	k4 := pc2.runPermutation(k4_56, 48)
	k5 := pc2.runPermutation(k5_56, 48)
	k6 := pc2.runPermutation(k6_56, 48)
	k7 := pc2.runPermutation(k7_56, 48)
	k8 := pc2.runPermutation(k8_56, 48)
	k9 := pc2.runPermutation(k9_56, 48)
	k10 := pc2.runPermutation(k10_56, 48)
	k11 := pc2.runPermutation(k11_56, 48)
	k12 := pc2.runPermutation(k12_56, 48)
	k13 := pc2.runPermutation(k13_56, 48)
	k14 := pc2.runPermutation(k14_56, 48)
	k15 := pc2.runPermutation(k15_56, 48)
	k16 := pc2.runPermutation(k16_56, 48)

	//generate first message
	m0 := initialPermutation.runPermutation(cypherText, 64)
	l0, r0 := m0[:32], m0[32:]
	//round 1
	r1 := xor(l0, function(r0, k16))
	l1 := r0
	//round2
	r2 := xor(l1, function(r1, k15))
	l2 := r1
	//round 3
	r3 := xor(l2, function(r2, k14))
	l3 := r2
	//round4
	r4 := xor(l3, function(r3, k13))
	l4 := r3
	//round 5
	r5 := xor(l4, function(r4, k12))
	l5 := r4
	//round6
	r6 := xor(l5, function(r5, k11))
	l6 := r5
	//round 7
	r7 := xor(l6, function(r6, k10))
	l7 := r6
	//round8
	r8 := xor(l7, function(r7, k9))
	l8 := r7
	//round 9
	r9 := xor(l8, function(r8, k8))
	l9 := r8
	//round10
	r10 := xor(l9, function(r9, k7))
	l10 := r9
	//round 11
	r11 := xor(l10, function(r10, k6))
	l11 := r10
	//round12
	r12 := xor(l11, function(r11, k5))
	l12 := r11
	//round 13
	r13 := xor(l12, function(r12, k4))
	l13 := r12
	//round14
	r14 := xor(l13, function(r13, k3))
	l14 := r13
	//round 15
	r15 := xor(l14, function(r14, k2))
	l15 := r14
	//round16
	r16 := xor(l15, function(r15, k1))
	l16 := r15
	//combine L and R
	mf := append(r16, l16...)
	message := initialPermutationInv.runPermutation(mf, 64)
	return string(binaryArrayToByteSlice(message))
}
