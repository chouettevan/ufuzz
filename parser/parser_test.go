package parser

import (
	"testing"
	"strings"
	"io"
)

type testMap[K any,T any] []struct{
	name string
	input K
	output T
}

func TestGetStatus(t *testing.T) {
	var statusTests  = testMap[string,int]{
		{"404 Not found response","HTTP/1.1 404 Not Found",404},
		{"200 Ok response","HTTP/2.0 200 Ok",200},
		{"302 Found response test","HTTP/1.1 302 Found",302},
	};
	for _,test := range statusTests {
		t.Run(test.name,func(t *testing.T) {
			if getStatus(test.input) != test.output {
				t.Errorf("function detected status %d in response %s",getStatus(test.input),test.input)
			}
		})
	}
}

const charset string  = "\n ";
func TestHttpParse(t *testing.T) {
	var parseTests = testMap[io.Reader,string]{
		{"200 Ok with keep-alive",strings.NewReader("HTTP/1.1 200 Ok\nConnection:keep-alive"),"200    37"},
	}
	for _,test := range parseTests {
		t.Run(test.name,func(t *testing.T) {
			out := strings.Trim(HttpParse(test.input),charset)
			if out != strings.Trim(test.output,charset) {
				t.Errorf("function generated summary %s instead of %s",out,strings.Trim(test.output,charset))
			}
		})
	}
}
