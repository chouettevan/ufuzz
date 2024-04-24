package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
)
	
type args struct {
	Tls bool	`help:"Add this flag if you wish to connect via https"`
	Host string `arg:"-h,required" help:"target host"`
	Port int `arg:"-p" default:"80" help:"http(s) server port"`
	Config string `arg:"-f,required" help:"path to config file"`
	Wordlists []string `arg:"-w,separate"`
	Threads int `arg:"-t", default:"50"`
}
type Task struct {
	Request string
	Params string
}
func (args) Description() string {
return `ufuzz,unlike most web fuzzers,uses a config file containing an http Request in which it will substitute the placeholders S1,S2,S3.. with input from the wordlists.
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
	var mu sync.Mutex
	channel := make(chan Task)
	arg.MustParse(&args)
	for i := 0;i < args.Threads;i++ {
		var conn io.ReadWriteCloser
		var err error
		if args.Tls {
			conn,err = tls.Dial("tcp",fmt.Sprintf("%s:%d",args.Host,args.Port),nil)
		} else {
			conn,err = net.Dial("tcp",fmt.Sprintf("%s:%d",args.Host,args.Port))
		}
		if err != nil {
			panic(err)
		}
		go fuzzer(conn,&mu,&channel);
	}
}

func fuzzer(connection io.ReadWriteCloser,mu *sync.Mutex,ch *chan Task) {
	defer connection.Close()
	scan := bufio.NewScanner(connection)	
	var status uint16
	var delay int64
	for {
		mu.Lock()
		tsk := <- *ch
		mu.Unlock()
		delay = time.Now().Unix()
		_,err := connection.Write([]byte(tsk.Request))
		delay = time.Now().Unix() - delay
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		scan.Scan() 
		fmt.Sscanf(scan.Text(),"HTTP/1.1 %d",status)
		fmt.Printf("%s  %d  %d",tsk.Params,status,delay);
	}

}
