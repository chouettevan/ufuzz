package main

import (
	"github.com/alexflint/go-arg"
	"crypto/tls"
	"io"
	"os"
)
	
var args struct {
	Tls bool
	Host string `arg:"-h,required" help:"target host"`
	Port int `arg:"-p" default:"80" help:"http(s) server port"`
	Wordlists []string `arg:"-w,separate"`
}
func main() {
	arg.MustParse(&args)
	conn,err := tls.Dial("tcp","mail.google.com:443",nil)
	if (err != nil ) {
		panic(err);
	}
	conn.Write([]byte("GET / HTTP/1.1\nHost:mail.google.com\n\n"))
	io.Copy(os.Stdout,conn);
	
}
