package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func SendTasks(ch* chan Task,wordlists []string,config string,num uint64,params string) error {
	file,err := os.Open(wordlists[0])
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(wordlists) != 1 {
			SendTasks(
				ch,
				wordlists[1:],
				strings.ReplaceAll(config,"S"+fmt.Sprint(num),scanner.Text()),
				num + 1,
				params + scanner.Text() + "    ")
		} else {
			var tsk Task		
			tsk.Request = strings.ReplaceAll(config,"S"+fmt.Sprint(num),scanner.Text())
			tsk.Params = params + scanner.Text() + "     " 
			*ch <- tsk
		}
	}
	return nil
}

