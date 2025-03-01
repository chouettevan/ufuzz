package main;

import (
	"fmt"
	"io"
    "bufio"
	"os"
	"sync"
	"time"
    "net/http"
	"net"
	"crypto/tls"
)
func fuzzer(args *args,mu *sync.Mutex,ch *chan Task,wg *sync.WaitGroup) {
	var delay int64
	for {
		mu.Lock()
		tsk := <- *ch
		mu.Unlock()
		wg.Add(1)
		delay = time.Now().UnixMilli()
		var conn io.ReadWriteCloser
		var err error
		if args.Tls {
			conn,err = tls.Dial("tcp",fmt.Sprintf("%s:%d",args.Host,args.Port),nil)
		} else {
			conn,err = net.Dial("tcp",fmt.Sprintf("%s:%d",args.Host,args.Port))
		}
		if err != nil {
			fmt.Fprintln(os.Stderr,err.Error())
		}
		_,err = conn.Write([]byte(tsk.Request))
		if err != nil {
            fmt.Fprintf(os.Stderr,"%s %s \n",tsk.Params,err.Error())
            wg.Done()
			conn.Close()
			continue
		}
        res,err := http.ReadResponse(bufio.NewReader(conn),nil)
		delay = time.Now().UnixMilli() - delay
        if err != nil {
            fmt.Fprintf(os.Stderr,"%s %s \n",tsk.Params,err.Error())
            wg.Done()
			conn.Close()
            continue
        }
        fmt.Printf("%s %d              %d          %d\n",tsk.Params,res.StatusCode,res.ContentLength,delay) 
		wg.Done()
		conn.Close()
	}
}
