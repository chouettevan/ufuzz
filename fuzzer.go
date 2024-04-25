package main;

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)
func fuzzer(connection io.ReadWriteCloser,mu *sync.Mutex,ch *chan Task,wg *sync.WaitGroup) {
	defer connection.Close()
	wg.Add(1)
	defer wg.Done()
	scan := bufio.NewScanner(connection)	
	var delay int64
	for {
		mu.Lock()
		tsk := <- *ch
		mu.Unlock()
		delay = time.Now().Unix()
		_,err := connection.Write([]byte(tsk.Request))
		delay = time.Now().Unix() - delay
		if err != nil {
			fmt.Fprintln(os.Stderr,err.Error())
			continue
		}
		scan.Scan() 
		fmt.Printf("%s  %s  %d \n",tsk.Params,scan.Text(),delay);
		for scan.Scan() {}
	}

}
