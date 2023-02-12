package util

import "net/http"

func CopyHeaders(headerOrigem http.Header, headersDestino http.Header) {
	for key, values := range headerOrigem {
		for _, value := range values {
			headersDestino.Add(key, value)
		}
	}
}
