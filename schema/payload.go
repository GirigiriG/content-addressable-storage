package schema

type Payload struct {
	MimeType string
	FileName string
	Data     []byte
}

func NewPayload(mimeType string, fileName string, data []byte) Payload {
	return Payload{
		MimeType: mimeType,
		FileName: fileName,
		Data:     data,
	}
}
