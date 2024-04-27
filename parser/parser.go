package parser

import (
	"bufio"
	"fmt"
	"io"
)

func HttpParse(conn io.Reader) string {
	scan := bufio.NewScanner(conn)
	var size uint64
	scan.Scan()
	status := getStatus(scan.Text())
	size += uint64(len(scan.Bytes()))
	for scan.Scan() {
		size += uint64(len(scan.Bytes()))
		// add 1 byte for the \n character
		size++
	}
	return fmt.Sprintf("%d    %d\n",status,size)
}

func getStatus(line string ) int {
	var status int
	var version float64
	fmt.Sscanf(line,"HTTP/%f %d ",&version,&status)
	return status+1
}
