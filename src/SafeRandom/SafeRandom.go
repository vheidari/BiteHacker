package SafeRandom

import(
	"crypto/rand"
	"log"
	"math"
)

func byteToAscii(keys *[]byte) string {

	// template string holds a string sample
	tempString := "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789!@#$%"

	// get keys and temp string length
	keysLength := len(*keys)
	tempStringLength := len(tempString)


	// calculation Quotient 
	calcQuotient := float64(256) / float64(tempStringLength)


	var bufferString string = ""
	
	// extract keys and generate a string buffer
	for item := range keysLength {

		// get key item
		getKey := float64((*keys)[item])
		
		// calculate alphabet position 
		calcAlphaBetPosition := uint32(math.Round( getKey / calcQuotient ))

		if calcAlphaBetPosition >= uint32(tempStringLength) {
			calcAlphaBetPosition--
		}

		// make a buffer of the string
		bufferString += string(tempString[calcAlphaBetPosition])
	}

	return bufferString
}


// generate a random number
func GenerateRandom(size uint) string  {

	// create an empty array
	keys := GenerateRandomByte(size)

	_, err := rand.Read(keys)

	if nil != err {
		log.Fatal("Problem to Create a Random number")
			
	}

	// convert all keys to ascii
	getBuffer := byteToAscii(&keys)

	return getBuffer
}



func GenerateRandomByte(size uint) []byte {
	randomByte := make([]byte, size)

	return randomByte
}
