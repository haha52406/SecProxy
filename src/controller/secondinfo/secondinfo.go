package secondinfo

import "net/http"

type SecondInfo struct{}

func (si *SecondInfo) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

func (si *SecondInfo) PreHandle(res http.ResponseWriter, req *http.Request) {

}
