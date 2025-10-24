package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/GirigiriG/blob_storage/schema"
)

func main() {
	lisnter, err := net.Listen("tcp", ":9090")
	fmt.Println("starting on port 9090")
	if err != nil {
		log.Panic(err)
	}

	for {
		connection, err := lisnter.Accept()
		if err != nil {
			fmt.Println("error accepting connection ", err)
			continue
		}
		go handleConnections(connection)
	}
}

func handleConnections(conn net.Conn) {
	lengthBuf := make([]byte, 4)
	io.ReadFull(conn, lengthBuf)

	payloadLen := binary.BigEndian.Uint32(lengthBuf)
	payloadBuff := make([]byte, payloadLen)

	n, err := io.ReadFull(conn, payloadBuff)
	if err != io.EOF {
		fmt.Println("done reading")
	}

	var payloads []schema.Payload

	if err := json.Unmarshal(payloadBuff[:n], &payloads); err != nil {
		fmt.Println("error reading Unmarshaling payload struct", err)
	}

	fmt.Println(payloads[0].FileName)

	newFile := payloads[0]

	if err := os.WriteFile("data/"+newFile.FileName, newFile.Data, 0644); err != nil {
		fmt.Println("error writing file to data", err)
	}
}
