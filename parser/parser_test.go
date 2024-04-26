package parser

import (
	"fmt"
	"testing"
	// "testing/iotest"
)

func TestGetStatus(t *testing.T) {
	var statusTests  = map[string]int{
		"HTTP/1.1 404 Not Found":404,
		"HTTP/2.0 200 Ok":200,
		"HTTP/1.1 302 Found":302,
	};
	for key,value := range statusTests {
		t.Run(fmt.Sprintf("%d Status test",value),func(t *testing.T) {
			if getStatus(key) != value {
				t.Errorf("function detected status %d in response %s",getStatus(key),key)
			}
		})
	}
}
