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
			res.Json(rw, err.Error(), 400)
			return
		}

		exists, _ := es.Exists(body.Email)

		if exists {
			hash, err := es.GetHashByEmail(body.Email)
			if err != nil {
				res.Json(rw, err.Error(), 400)
				return
			}
			localLink := fmt.Sprintf("http://localhost:8087/verify/%s", hash)
			fmt.Println("Local Link: ", localLink)
			err = mymail.SendEmail(body.Email, localLink)
			if err != nil {
				res.Json(rw, err.Error(), 400)
			}
			res.Json(rw, "Success", 200)
			return
		}

		randomHash, err := hash.GenerateRandomHash(10)
		if err != nil {
			res.Json(rw, err.Error(), 400)
			return
		}

		es.Add(body.Email, randomHash)

		localLink := fmt.Sprintf("http://localhost:8081/verify/%s", randomHash)

		err = mymail.SendEmail(body.Email, localLink)
		if err != nil {
			res.Json(rw, err.Error(), 400)
		}
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

		err := es.DeleteByHash(hash)

		if err != nil {
			res.Json(w, false, 400)
		}

		res.Json(w, true, 200)

	}
}
