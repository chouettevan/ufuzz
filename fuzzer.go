package main;

import (
	"github.com/chouettevan/ufuzz/parser"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
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
			fmt.Fprintln(os.Stderr,err.Error())
			continue
		}
		fmt.Printf("%s    %d    %s",tsk.Params,delay,parser.HttpParse(connection))
		wg.Done()
	}

}
