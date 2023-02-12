package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"time"
)

func main() {
	testeUid()
	testeCanal()
}

func testeUid() {
	requestId, _ := uuid.NewV4()
	idStr := requestId.String()

	//_, _ := uuid.ParseHex(idStr)

	if requestId.String() == idStr {
		log.Println("sucesso")
	} else {
		log.Println("falha :(")
	}

}

func testeCanal() {
	out := make(chan string)

	go func() {
		time.Sleep(time.Second)
		out <- "Yeste"
	}()

	stringSaida := <-out

	log.Println(stringSaida)

}
