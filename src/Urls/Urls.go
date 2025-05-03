package Urls


// declare Urls interface
type Urls interface  {
	UrlPool() map[string]Url
}



// make a map of urls
func UrlPool() map[string]Url {
	pool := make(map[string]Url)
	return pool
}


// the url object
type Url struct {
	protocol string
	domain_name string
	domain_ext string
	path string
	header map[string]string
}


// build an url object 
func BuildUrl(protocol string, domain_name string, domain_ext string, path string) Url { 
	header := make(map[string]string)
	url := Url {
		protocol: protocol,
		domain_name: domain_name,
		domain_ext: domain_ext,
		path: path,
		header: header,
	}

	return url
}


// [url method]: add new key/value to header
func (url *Url) AddHeader(key string , value string) {
	url.header[key] = value;
}


// [url method]: generate a complete string from the url object
func (url *Url) UrlToString() string {
	buffUrl := url.protocol
	buffUrl += url.domain_name
	buffUrl += url.domain_ext
	buffUrl += url.path
	return buffUrl
}
