package SafeRandom

import "testing"

func TestGenerateRandom(t *testing.T) {

 	gotKeyOne := GenerateRandom(32)
 	gotKeyTwo := GenerateRandom(32)


	if gotKeyOne == gotKeyTwo  {
		t.Errorf("gotKeyOne : %q, gotKeyTwo : %q", gotKeyOne, gotKeyTwo)
	}
}



func TestGenerateRandomByte(t *testing.T) {

	// check return array size
	randSize := 32

	got32Random := GenerateRandomByte(uint(randSize))
		
	if len(got32Random) != randSize  {
		t.Errorf("randSize: %q , got32Random : %q", randSize, len(got32Random))
	}
}




func TestByteToAscii(t *testing.T) {

	randSize := 32

	// is returned strings same  
	got32RandomOne := GenerateRandom(uint(randSize))
	got32RandomTwo := GenerateRandom(uint(randSize))

	if got32RandomOne == got32RandomTwo {
		t.Errorf("got32RandomOne: %q , got32RandomTwo : %q", got32RandomOne, got32RandomTwo)
	}	
	
}
