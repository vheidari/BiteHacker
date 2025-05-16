package Modems

import "testing"

// testing GenerateNewModem
func TestGenerateNewModem(t *testing.T) {

	gotObject := GenerateNewModem() 
	expectObject := Modem {}

	if gotObject != expectObject {
		t.Errorf("gotObject : %q,  expectObject : %q", gotObject , expectObject)
	}

}


// testing addModemSettings method
func TestAddModemSettings(t *testing.T) {

	gotObject := GenerateNewModem() 

	gotObject.addModemSettings("Dlink", "Admin", "Admin", "Adsl-2")
	
	expectObject := Modem {
		name: "Dlink",
		username : "Admin",
		password : "Admin",
		model: "Adsl-2",
	}

	if gotObject != expectObject {
		t.Errorf("gotObject : %q,  expectObject : %q", gotObject , expectObject)
	}

}


// testing updateModemSettings method 
func TestUpdateModemSettings(t *testing.T) {

	gotOldObject := GenerateNewModem()
	gotNewObject := GenerateNewModem();
	gotNewObject.addModemSettings("Dlink", "Admin", "Admin", "Adsl-2")
	
	gotOldObject.updateModemSettings(&gotNewObject)


	if gotOldObject != gotNewObject {
		t.Errorf("gotObject : %q,  expectObject : %q", gotOldObject , gotNewObject)
	}
	
}




