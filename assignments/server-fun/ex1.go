package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	port := ":8080"
	li, e := net.Listen("tcp", port)
	fmt.Println("Listening on port ", port)
	if e != nil {
		log.Fatalln(e)
	}
	defer li.Close()

	for {
		cx, e := li.Accept()
		fmt.Println("> Connected")
		if e != nil {
			log.Fatalln(e)
		}

		io.WriteString(cx, "Connection made!\n\rHELLO!  And GOODBYE!!")
		cx.Close()
		fmt.Println("> Closed")
	}

}
