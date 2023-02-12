package routes

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"go-proxy/util"
	"io"
	"log"
	"net/http"
	"time"
)

var MapRequest = map[string]chan http.Response{}

const timeoutTime = 2

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println("mapRequest: ", MapRequest)

	requestId, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return
	}

	responseOut := make(chan http.Response)

	go sendRequest(responseOut, r)

	AguardaRequestCompletar(w, responseOut, requestId.String())
}

func sendRequest(out chan http.Response, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, "http://localhost:8080"+r.URL.String(), r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	util.CopyHeaders(r.Header, req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	out <- *resp
}

func requestTimeout(w http.ResponseWriter, responseOut chan http.Response, requestId string) {
	w.Header().Add("requestId", requestId)
	w.WriteHeader(http.StatusProcessing)
	w.Write([]byte("Processando: " + requestId))
	fmt.Println("Requisição ainda em andamento, retornando status Processing")

	MapRequest[requestId] = responseOut
}

func finalizaRequest(res http.Response, w http.ResponseWriter) {
	util.CopyHeaders(res.Header, w.Header())
	w.WriteHeader(res.StatusCode)
	defer res.Body.Close()

	byteRes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(byteRes)
}

func AguardaRequestCompletar(w http.ResponseWriter, responseOut chan http.Response, requestId string) {
	select {
	case resp := <-responseOut:
		finalizaRequest(resp, w)
	case <-time.After(timeoutTime * time.Second):
		requestTimeout(w, responseOut, requestId)
	}
}
