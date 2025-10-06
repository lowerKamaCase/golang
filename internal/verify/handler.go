package verify

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/pkg/hash"
	"lowerkamacase/golang/pkg/mymail"
	"lowerkamacase/golang/pkg/req"
	"lowerkamacase/golang/pkg/res"
	"lowerkamacase/golang/pkg/storage"
	"net/http"
	"sync"
)

type VerifierDeps struct {
	*configs.Config
}

type Verifier struct {
	*configs.Config
}

var mu sync.RWMutex

func NewVerifierHandler(router *http.ServeMux, deps VerifierDeps) {
	handler := &Verifier{
		Config: deps.Config,
	}
	es := storage.NewEmailStorage("emails.json", &mu)
	router.HandleFunc("POST /send", handler.Send(es))
	router.HandleFunc("GET /verify/{hash}", handler.Verify(es))

}

func (verifier *Verifier) Send(es *storage.EmailStorage) http.HandlerFunc {

	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[SendRequest](&rw, request)
		if err != nil {
			return
		}
		fmt.Println("Send: ", body)

		exists, _ := es.Exists(body.Email)

		if exists {
			hash, err := es.GetHashByEmail(body.Email)
			if err != nil {
				return
			}
			localLink := fmt.Sprintf("http://localhost:8081/verify/%s", hash)
			fmt.Println("Local Link: ", localLink)
			mymail.SendEmail(body.Email, localLink)

			return
		}

		randomHash, err := hash.GenerateRandomHash(10)
		if err != nil {
			return
		}

		fmt.Println(randomHash)

		es.Add(body.Email, randomHash)

		localLink := fmt.Sprintf("http://localhost:8081/verify/%s", randomHash)
		fmt.Println("Local Link: ", localLink)
		mymail.SendEmail(body.Email, localLink)
		res.Json(rw, "Success", 200)

	}
}

func (verifier *Verifier) Verify(es *storage.EmailStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		if hash == "" {
			http.Error(w, "Hash parameter is required", http.StatusBadRequest)
			return
		}

		es.DeleteByHash(hash)

		res.Json(w, "Success Verification", 200)

	}
}
