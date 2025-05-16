package HttpRequest

import _ "net/http"
import  _ "encoding/json"
import "vheidari/main/pkg/Urls"


type MethodId uint
// type HttpMethod  struct 

const (
	Get MethodId = iota
	Post
	Put	
	Delete
	Patch
)

// 
var methodName = map[MethodId] string {
	Get : "Get",
	Post: "Post",
	Put: "Put",
	Delete: "Delete",
	Patch : "Patch",
}




type HttpReq struct {
	url Urls.Url
	method MethodId
}


//
func (httpReq *HttpReq) ReturnMethodName() string {
	return  methodName[httpReq.method]
}


// 
func GenerateRequest(url *Urls.Url , methodId MethodId) HttpReq {

	 newRequest := HttpReq {
		url: *url,
		method : methodId,
	}

	return newRequest
}

// 
// passing function to go
// Go(func(a string ) {
// 	fmt.Println(a);
// } )
// 
func (httpRequest *HttpReq) Go(result  func(string)) {

	methodId := httpRequest.methodId
	getMethodName := ReturnMethodName()

	result("this is next")
	
}
