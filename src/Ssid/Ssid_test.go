package Ssid

import "testing"

func TestNewSSIDAndPassword(t *testing.T) {
	ssidSize := 32
	passwordSize := 15
	visibility := false

	gotSSIDObj := NewSSIDAndPassword(uint(ssidSize), uint(passwordSize), visibility)


	if len(gotSSIDObj.ssid) != ssidSize {
		t.Errorf("gotSSIDObj.ssid size: %q , ssidSize : %q", gotSSIDObj.ssid, ssidSize )		
	}
	
	if len(gotSSIDObj.password) != passwordSize {
		t.Errorf("gotSSIDObj.password size: %q , passwordSize : %q", gotSSIDObj.password,  passwordSize  )
	}
	
	if gotSSIDObj.visibility != visibility {
		isVisible1 := ""
		isVisible2 := ""

		if visibility {
		 	isVisible1 = "false"
		 	isVisible2 = "true"
		} else {
		 	isVisible1 = "true"
		 	isVisible2 = "false"
		}

		t.Errorf("gotSSIDObje.visibility: %q , got32Random : %q", isVisible1, isVisible2  )
		
	}
	
		
}



func TestGenerateWPASettings(t *testing.T) {
	securityMode  := "WPA2"
	wpaEncryption := "TPIK+AES"
	wpaMode := "WPA2-PSK"

	gotObjOne := WPA {
		securityMode  : securityMode,
		wpaEncryption : wpaEncryption,
		wpaMode : wpaMode,
	}	

	gotObjTwo := GenerateWPASettings(securityMode, wpaEncryption, wpaMode)
	
	if gotObjOne != gotObjTwo {
		
		t.Errorf("gotObjOne: %q , gotObjTwo : %q", gotObjOne, gotObjTwo  )
	}
	
}



func TestAddWPASettings(t *testing.T) {

	ssidSize  := 32
	passwordSize := 15
	visibility := false

	gotObjOne := NewSSIDAndPassword(uint(ssidSize), uint(passwordSize), visibility)

	sample := WPA {}

	// is sample and wirelessSeucrity same
	if sample != gotObjOne.wirelessSecurity {
		t.Errorf("sample : %q , gotObjOne : %q", sample, gotObjOne.wirelessSecurity)
	}

	// ----

	securityMode  := "WPA2"
	wpaEncryption := "TPIK+AES"
	wpaMode := "WPA2-PSK"

	gotObjTwo := WPA {
		securityMode  : securityMode,
		wpaEncryption : wpaEncryption,
		wpaMode : wpaMode,
	}	

	// add WPA setting to an exist object 
	gotObjOne.AddWPASettings(&gotObjTwo)
	
	if gotObjOne.wirelessSecurity != gotObjTwo {
		t.Errorf("gotObjTow : %q , gotObjOne.wirelessSecurity : %q", gotObjTwo , gotObjOne.wirelessSecurity)
	}
	
}

