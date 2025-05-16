package main

import "fmt"
import _ "vheidari/main/pkg/SafeRandom"
import "vheidari/main/pkg/Urls"
import "vheidari/main/pkg/Ssid"
import "vheidari/main/pkg/HttpRequest"



// entry point
func main() {
	fmt.Println("Hello World!")



	// for item := range 10000 {

	// fmt.Println(SafeRandom.GenerateRandom(64))
	// fmt.Println(SafeRandom.GenerateRandom(32))
	// fmt.Println(SafeRandom.GenerateRandom(128))
	// fmt.Println(SafeRandom.GenerateRandom(256))
	// fmt.Println(SafeRandom.GenerateRandom(512))
	// 
	// fmt.Println(item)
	// fmt.Println("-----------------")
	// 	
	
	// }



	getUrl := Urls.BuildUrl("http://", "192.168.1.1", "", "/page/target")

	getUrl.AddHeader("keyOne", "valueOne")
	getUrl.AddHeader("keyTwo", "valueTwo")
	getUrl.AddHeader("keyThree", "valueThree")
	getUrl.AddHeader("set-cookie", "0xid=1?username=vheidari?jobs=programmer?expire=866642131864")
	getUrl.AddHeader("set-settion", "soofjsdXojojfojoojojfsdlkjfdsojo")
	getUrl.AddHeader("range", "norange")


	urlToString := getUrl.UrlToString()

	


	fmt.Println(getUrl)
	fmt.Println(urlToString)


	urlsPool := Urls.UrlPool()

	urlsPool["UpdateSSID"] = getUrl

	fmt.Println(urlsPool)



	for k,v := range urlsPool {
		
		fmt.Println(k)
		fmt.Println(v)
	}



	//  new ssid 
// 	NewSsid := Ssid.NewSSIDAndPassword(16, 32, false)
// 
// 	wpaSettings := Ssid.GenerateWPASettings("WPA2", "TKIP+AES", "WPA2-PSK");
// 	
// 	NewSsid.AddWPASettings(&wpaSettings)
// 
// 	fmt.Println(NewSsid)



	 // httpRequest
	sampelUrl := Urls.BuilderUrl("http://", "192.168.1.1", "", "/page/target")
	

	




}
