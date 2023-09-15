package header

import (
	"time"

	"github.com/gnzlabs/keyring"
)

const ByteCount int = 82

type Header struct {
	PrecedingBlockDigest []byte
	IssuerFingerprint    []byte
	CreationTime         time.Time
	BlockType            uint16
	ContentLength        uint32
	SignatureLength      uint32
}

func New(precedingBlockId []byte, blockType uint16, contentLength int, issuer any) *Header {
	Header := Header{
		PrecedingBlockDigest: precedingBlockId,
		IssuerFingerprint:    issuer.(keyring.PublicKeyRing).Identifier(),
		CreationTime:         time.Now(),
		BlockType:            blockType,
		ContentLength:        uint32(contentLength),
	}
	return &Header
}
