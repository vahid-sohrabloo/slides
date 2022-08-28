package main

import (
	"bufio"
	"bytes"
	"io"
	"net"
)

func Read(conn net.Conn, delim byte) (string, error) {
	//BUFIO START OMIT
	reader := bufio.NewReaderSize(conn, 4096)
	//BUFIO END OMIT
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}
