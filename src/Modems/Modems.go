package Modems

// modem object
type Modem struct {
	name string
	username string
	password string
	model string 	
}

// generate a modem object
func GenerateNewModem() Modem {
	newModem := Modem {}
	return newModem
}

// [modem method] : add setting to current exist modem object
func (modem *Modem) addModemSettings(name, username, password, model string ) {
	modem.name = name
	modem.username = username 
	modem.password = password
	modem.model = model
}

// [modem method] : update current modem object info
func (modem *Modem) updateModemSettings(newModem *Modem) {
	modem.name = newModem.name
	modem.username = newModem.username 
	modem.password = newModem.password
	modem.model = newModem.model
}
