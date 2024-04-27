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

var statusTests  = testMap[string,int]{
	{"404 Not found response","HTTP/1.1 404 Not Found",404},
	{"200 Ok response","HTTP/2.0 200 Ok",200},
	{"302 Found response test","HTTP/1.1 302 Found",302},
};

func TestGetStatus(t *testing.T) {
	for _,test := range statusTests {
		t.Run(test.name,func(t *testing.T) {
			if getStatus(test.input) != test.output {
				t.Errorf("function detected status %d in response %s",getStatus(test.input),test.input)
			}
		})
	}
}

func BenchmarkGetStatus(b *testing.B) {
	for _,test := range statusTests {
		b.Run(test.name,func(b *testing.B) {
			for i := 0;i < b.N;i++ {
				getStatus(test.input)
			}
		})
	}
}

const charset string  = "\n ";
var parseTests = testMap[io.Reader,string]{
	{"200 Ok with keep-alive",strings.NewReader("HTTP/1.1 200 Ok\nConnection:keep-alive"),"200    37"},
}
func TestHttpParse(t *testing.T) {
	for _,test := range parseTests {
		t.Run(test.name,func(t *testing.T) {
			out := strings.Trim(HttpParse(test.input),charset)
			if out != strings.Trim(test.output,charset) {
				t.Errorf("function generated summary %s instead of %s",out,strings.Trim(test.output,charset))
			}
		})
	}
}
func BenchmarkHttpParse(b *testing.B) {
	for _,test := range parseTests {
		b.Run(test.name,func(b *testing.B) {
			for i := 0;i < b.N;i++ {
				HttpParse(test.input)
			}
		})
	}

}

