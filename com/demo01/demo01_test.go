package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func StringPlus() string  {

	var s string

	s+="nicheng"+"飞絮无情"+"\n"

	s+="heke"

	return s

}

func StringFmt() string  {
	return fmt.Sprint("nicheng","飞絮无情","\n","heke")
	
}

func StringJoin() string{
	s := [] string{"nicheng","飞絮无情","\n","heke"}
	return strings.Join(s, "")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("nicheng")
	b.WriteString("飞絮无情")
	b.WriteString("\n")
	b.WriteString("heke")
	return b.String()
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("nicheng")
	b.WriteString("飞絮无情")
	b.WriteString("\n")
	b.WriteString("heke")
	return b.String()

	
}


func BenchmarkStringPlus(b *testing.B) {
	for i:=0;i<b.N;i++{
		//StringPlus()
		StringFmt()
	}
}

func BenchmarkStringFmt(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringPlus()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringJoin()
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringBuffer()
	}
}


func BenchmarkStringBuilder(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringBuilder()
	}
}
