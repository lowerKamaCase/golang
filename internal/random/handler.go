package random

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func Random(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("Got request ", *request)
	randomFrom1To6 := rand.IntN(6) + 1
	result := fmt.Sprintf("%d", randomFrom1To6)

	rw.Write([]byte(result))
}
