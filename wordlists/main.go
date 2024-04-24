package wordlists

import (
	"os"
	"bufio"
	"github.com/chouettevan/ufuzz/"
	"strings"
)

func SendTasks(ch* chan main.Task,wordlists []string,config string,num uint64,params string) error {
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
				strings.ReplaceAll(config,"S"+string(num),scanner.Text()),
				num + 1,
				params + "    " + scanner.Text())
		} else {
			var tsk main.Task		
			tsk.Request = strings.ReplaceAll(config,"S"+string(num),scanner.Text())
			tsk.Params = params + "     " + scanner.Text()
		}
	}
	return nil
}
