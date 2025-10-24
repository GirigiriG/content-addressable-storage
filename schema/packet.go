package schema

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/GirigiriG/blob_storage/utils"
)

type Packet struct {
	HeadAndPayloadLength uint32
	Header               Header
	Payload              Payload
}

func CreatePacket(fileName string) (*Packet, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("unable to open file", err)
		return nil, err
	}
	mimeType := ""

	fileNameAndExtension := strings.Split(fileName, ".")
	if len(fileNameAndExtension) == 2 {
		mimeType = fileNameAndExtension[1]
	} else {
		fmt.Println("mimetype not included in filename. attempting to extra it from the file data")
		ext := strings.Split(http.DetectContentType(file[:512]), "/")[1]
		mimeType = ext
	}

	fileNameWithExt := fmt.Sprintf("%v.%v", fileNameAndExtension[0], mimeType)
	payload := NewPayload(mimeType, fileNameWithExt, file)
	header := NewHeader(utils.GetStructLength(payload))

	packet := Packet{
		HeadAndPayloadLength: uint32(utils.GetStructLength(header, payload)),
		Header:               header,
		Payload:              payload,
	}

	return &packet, nil
}
