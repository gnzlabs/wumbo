package wumbo

import (
	"crypto/rand"

	"github.com/gnzlabs/keyring"
	"github.com/gnzlabs/wumbo/header"
)

func New(precedingBlockId []byte, contentType uint16, data any, issuer any) (*Block, error) {
	blockContent, err := MarshalContent(data)
	if err != nil {
		return nil, err
	}
	newBlock := Block{
		Header:  header.New(precedingBlockId, contentType, len(blockContent), issuer),
		Content: blockContent,
	}
	if newBlock.Signature, err = issuer.(keyring.PrivateKeyRing).Sign(rand.Reader, newBlock.Digest(), nil); err != nil {
		return nil, err
	}
	return &newBlock, nil
}
