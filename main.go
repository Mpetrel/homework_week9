package main

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
)

type Header struct {
	PackageLength int32
	HeaderLength int32
	Version int32
	Operation int32
	SequenceId int32
}

func ByteToInt(data []byte) int32 {
	buffer := bytes.NewBuffer(data)
	var x int32
	_ = binary.Read(buffer, binary.BigEndian, &x)
	return x
}

func DepackHeader(buffer []byte) (*Header, error) {
	if len(buffer) < 16 {
		return nil, errors.New("invalid buffer length")
	}
	header := Header{
		PackageLength: ByteToInt(buffer[:4]),
		HeaderLength:  ByteToInt(buffer[4:6]),
		Version:       ByteToInt(buffer[6:8]),
		Operation:     ByteToInt(buffer[8:12]),
		SequenceId:    ByteToInt(buffer[12:16]),
	}
	return &header, nil
}