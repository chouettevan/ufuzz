package main;

import (
	"fmt"
	"io"
    "bufio"
	"os"
	"sync"
	"time"
    "net/http"
)
func fuzzer(connection io.ReadWriteCloser,mu *sync.Mutex,ch *chan Task,wg *sync.WaitGroup) {
	defer connection.Close()
	var delay int64
	for {
		mu.Lock()
		tsk := <- *ch
		mu.Unlock()
		wg.Add(1)
		delay = time.Now().UnixMilli()
		_,err := connection.Write([]byte(tsk.Request))
		delay = time.Now().UnixMilli() - delay
		if err != nil {
            fmt.Fprintf(os.Stderr,"%s %s \n",tsk.Params,err.Error())
            wg.Done()
            continue
		}
        res,err := http.ReadResponse(bufio.NewReader(connection),nil)
        if err != nil {
            fmt.Fprintf(os.Stderr,"%s %s \n",tsk.Params,err.Error())
            wg.Done()
            continue
        }
        fmt.Printf("%s %d     %d\n",tsk.Params,res.StatusCode,res.ContentLength) 
		wg.Done()
	}
}
