package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	port := ":8080"
	li, e := net.Listen("tcp", port)
	log.Println("> Listening on port ", port)
	if e != nil {
		log.Fatalln(e)
	}
	defer li.Close()

	for {
		cx, e := li.Accept()
		log.Println("> Connected")
		if e != nil {
			log.Fatalln(e)
		}

		//io.WriteString(cx, "Connection made!\n\rHELLO!  And GOODBYE!!")
		handleConnection(cx)

		cx.Close()
		log.Println("> Closed")
	}

}

func handleConnection(c net.Conn) {
	rDat, e := ioutil.ReadAll(c)
	if e != nil {
		log.Println("ERROR: Cannot read from connection: ", e)
		return
	}

	log.Println("Received from connection:")
	fmt.Println(string(rDat))
}
