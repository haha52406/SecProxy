package secondkill

import (
	"net/http"
)

type SecondKill struct{}

func (sk *SecondKill) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// RETURN:
	// 	res.Write()
}

// func (sk *SecondKill) PreHandle(res http.ResponseWriter, req *http.Request) error {
// 	// requestBody, err := ioutil.ReadAll(req.Body)
// 	// if err != nil {
// 	// 	return errors.New("read request body is err")
// 	// }
// }
