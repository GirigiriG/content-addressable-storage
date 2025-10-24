package schema

type Header struct {
	Magic         uint16
	Version       uint8
	ContentLength uint32
}

func NewHeader(payloadLength int) Header {
	return Header{
		Magic:         0xAE0,
		Version:       1,
		ContentLength: uint32(payloadLength),
	}
}
