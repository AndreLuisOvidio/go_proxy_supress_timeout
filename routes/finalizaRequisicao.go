package routes

import (
	"net/http"
	"strings"
)

func FinalizaHandler(w http.ResponseWriter, r *http.Request) {
	requestId := r.Header.Get("requestId")
	if requestId == "" {
		requestId = strings.TrimPrefix(r.URL.Path, "/finalizaRequisicao/")
	}
	if requestId == "" {
		w.WriteHeader(400)
		w.Write([]byte("requestId não informado"))
		return
	}
	//c7044da3-feaa-4aa8-7f6d-484d11f91cc9
	requestOut, ok := MapRequest[requestId]
	if ok {
		delete(MapRequest, requestId)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("request já finalizada"))
		return
	}

	AguardaRequestCompletar(w, requestOut, requestId)
}
