package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/GirigiriG/blob_storage/schema"
	"github.com/GirigiriG/blob_storage/utils"
)

var REMOTE_HOST = "localhost:9090"

func main() {
	packet, err := schema.CreatePacket("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var packetLen uint32 = packet.ContentLength
	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, packetLen)

	conn, err := net.Dial("tcp", REMOTE_HOST)
	if err != nil {
		fmt.Println("uanble to connect to tcp server @ " + REMOTE_HOST)
		return
	}

	fmt.Println("[sending] " + packet.Data.FileName)
	conn.Write(buf)
	conn.Write(utils.ParsePayloadToJSON(packet.Data))
}
