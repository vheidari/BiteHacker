package Urls

type Url struct {
	protocol string
	domain_name string
	domain_ext string
	header map[string]string
}


// Build An Url 
func BuildUrl(protocol string, domain_name string, domain_ext string) Url { 
	header := make(map[string]string)
	url := Url {
		protocol: protocol,
		domain_name: domain_name,
		domain_ext: domain_ext,
		header: header,
	}

	return url
}


// Add New Key Value Header
func AddHeader(key string , value string, url *Url) {
	url.header[key] = value;
}


// Return Whole Url as string
func UrlToString(url *Url) string {
	buffUrl := url.protocol
	buffUrl += url.domain_name
	buffUrl += url.domain_ext
	return buffUrl
}

