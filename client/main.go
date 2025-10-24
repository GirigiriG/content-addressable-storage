package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/GirigiriG/blob_storage/schema"
	"github.com/GirigiriG/blob_storage/utils"
)

func main() {
	packet, err := schema.CreatePacket("vector_db")
	if err != nil {
		log.Fatal(err)
	}
	var packetLen uint32 = packet.HeadAndPayloadLength
	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, packetLen)
	remote := "localhost:9090"
	conn, err := net.Dial("tcp", remote)
	if err != nil {
		fmt.Println("uanble to connect to tcp server @ " + remote)
	}

	fmt.Println("[sending] " + packet.Payload.FileName)
	conn.Write(buf)
	conn.Write(utils.ParsePayloadToJSON(packet.Payload))

}
