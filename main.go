package main

import (
	"github.com/alexflint/go-arg"
)
	
type args struct {
	Tls bool	`help:"Add this flag if you wish to connect via https"`
	Host string `arg:"-h,required" help:"target host"`
	Port int `arg:"-p" default:"80" help:"http(s) server port"`
	Config string `arg:"-f,required" help:"path to config file"`
	Wordlists []string `arg:"-w,separate"`
}
func (args) Description() string {
return `ufuzz,unlike most web fuzzers,uses a config file containing an http request in which it will substitute the placeholders S1,S2,S3.. with input from the wordlists.
Example:
	ufuzz -f ufuzz.conf -h 1.1.1.1 -p 443 -tls -w /path/to/wordlist | grep -v 404
	with ufuzz.conf containing:
		GET /S1 HTTP/1.1
		Host:1.1.1.1
Will fuzz for directories
Placeholders can be used anywhere within the config file.
`
}

func main() {
	var args args
	arg.MustParse(&args)
}
