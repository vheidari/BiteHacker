package main

import "fmt"
import "vheidari/main/src/Modems"
// import "vheidari/main/src/SafeRandom"

import "vheidari/main/src/Urls"




// entry point
func main() {
	fmt.Println("Hello World!")

	Modems.PrintS()


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



	getUrl := Urls.BuildUrl("http://", "192.168.1.1", "")

	Urls.AddHeader("keyOne", "valueOne", &getUrl)
	Urls.AddHeader("keyTwo", "valueTwo", &getUrl)
	Urls.AddHeader("keyThree", "valueThree", &getUrl)
	Urls.AddHeader("set-cookie", "0xid=1?username=vheidari?jobs=programmer?expire=866642131864", &getUrl)
	Urls.AddHeader("set-settion", "soofjsdXojojfojoojojfsdlkjfdsojo", &getUrl)
	Urls.AddHeader("range", "norange", &getUrl)


	urlToString := Urls.UrlToString(&getUrl)

	


	fmt.Println(getUrl)
	fmt.Println(urlToString)

}
