package main

func main() {

	/*
		First execution run by Bob. Bob has generated a public/private RSA key so that Alice can
		send him a private DES key. This way Bob will be able to communitcate with her secretly. He
		stores his private to his local disk, and then release his public key into a shared drive
		where the world can access
	*/

	// secret, public := generateRSAKeys()
	// codedPubKey := publicKeyToString(public)
	// codedPrivateKey := privateKeyToString(secret)
	// releasePublicRSAKey(codedPubKey)
	// storePrivateRSAKey(codedPrivateKey)

	/*
		Second execution run by Alice. Alice has been notified that Bob has release his public RSA key.
		She will retrieve it from the public drive and use it to encrypt a new DES private key that
		she has generated. She will release the encrypted DES key back into the shared drive for everyone
		to see and save the unencrypted version locally to her private hard drive. She will also generate
		an IV to use for their aggreed upon mode of operation, Output Feedback Mode. She will also publicly
		release the IV
	*/
	// codedPubKey := getPublicRSAKeyFromFile()
	// public, _ := decodePubKey(codedPubKey)
	// desKey := generateDESKey()
	// IV := generateIV()
	// saveDESKeyToFile(desKey)
	// cypherDESKey := encryptRSA(public, desKey)
	// releaseEncryptedDESKey(cypherDESKey)
	// releaseIV(IV)

	/*
		Third execution run by Bob. He now has the DES Key that he needs to start send Alice a message stream.
		He will encode his stream using the agreed upon OFB method, using the IV she released publicly. since she
		has sent him a 256 bit key, he will use DES 4 times, or 64 bits of the key on each iteration. Then he will
		release his encrypted message
	*/
	// codedSecretKey := getPrivateRSAKeyFromFile()
	// secret, _ := decodeSecretKey(codedSecretKey)
	// cypher := getEncryptedDESKeyFromFile()
	// key := decrptRSA(secret, cypher)
	// saveDESKeyToFile(key)
	// keys := splitKeyintoKeys(strconv.FormatUint(key, 10))
	// IV := getIVFromFile()
	// formatedIV := formatIV(IV)
	// message := "Hey Alice, I'm glad that we're able to communicate now in this super secret channel, just wanted you to know the eagle has landed"
	// cyphertext:= RunDesFourTimes(keys,formatIV, message)
	// releaseCypher(cyphertext)

	/*
		Final execution run by Alice. Bob has release the encrypted message he wanted to deliver to Alice. She
		will reach out and grab it from a public drive and then decode the message, where she can finally read it.
	*/

	// key := getDesKeyfromFile()
	// keys := splitKeyintoKeys(strconv.FormatUint(key, 10))
	// IV := getIVFromFile()
	// formatedIV := formatIV(IV)
	// cyphertext := retrieveCypherTextFromFile()
	// message:= RunDesFourTimes(keys,formatIV, cyphertest)
	// println(message)

}

/*































	// var cText cypherText
	// var mText cypherText
// s := ""
// chunks := stringToEightByteArray(message)
// for i, chunk := range chunks {
// 	var tempString []string
// 	// cypherText = append(cypherText, desEnc(chunk, keys[0]))
// 	if i == 0 {
// 		tempString = desEnc(formatedIV, keys[0])
// 		for _, piece := range tempString {
// 			s = ""
// 			s = s + piece
// 		}
// 	} else {
// 		tempString = desEnc(s, keys[0])
// 		for _, piece := range tempString {
// 			s = ""
// 			s = s + piece
// 		}
// 	}
// 	c := xor(tempString, stringToBinaryArray(chunk))
// 	cText = append(cText, c)
// }

// s = ""
// cout := stringToEightByteArray(cText.formatToString())
// for i, chunk := range cout {
// 	var tempString []string
// 	// cypherText = append(cypherText, desEnc(chunk, keys[0]))
// 	if i == 0 {
// 		tempString = desEnc(formatedIV, keys[0])
// 		for _, piece := range tempString {
// 			s = ""
// 			s = s + piece
// 		}
// 	} else {
// 		tempString = desEnc(s, keys[0])
// 		for _, piece := range tempString {
// 			s = ""
// 			s = s + piece
// 		}
// 	}
// 	c := xor(tempString, stringToBinaryArray(chunk))
// 	mText = append(mText, c)
// }
// // println(mText.formatToString())

*/
