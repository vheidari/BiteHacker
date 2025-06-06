package Ssid 

import "vheidari/main/pkg/SafeRandom"


type WPA struct {
	securityMode string
	wpaEncryption string
	wpaMode string
}


// ssid struct hold ssid
type SSID struct {
 	ssid string
 	password string
 	visibility bool
 	wirelessSecurity WPA
 	
}


func NewSSIDAndPassword(ssidSize uint, passwordSize uint, visibility bool ) SSID {
	ssid := SafeRandom.GenerateRandom(ssidSize)
	password := SafeRandom.GenerateRandom(passwordSize)
	visibility = false

	wirelessSecurity := WPA  {}
	
	newSsid := SSID {
		ssid : ssid,
		password : password,
		visibility: visibility,
		wirelessSecurity: wirelessSecurity,
	}

	return newSsid
}


func GenerateWPASettings (securityMode string , wpaEncryption string , wpaMode string ) WPA {
	return WPA {
		securityMode : securityMode,
		wpaEncryption: wpaEncryption,
		wpaMode : wpaMode,
	}
}


func (ssid *SSID) AddWPASettings(wpaSettings *WPA) {
		ssid.wirelessSecurity =  *wpaSettings
}

