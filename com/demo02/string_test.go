package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func StringPlus(p []string) string{
var s string
l:=len(p)
for i:=0;i<l;i++{
s+=p[i]
}
return s
}

func StringFmt(p []interface{}) string{
return fmt.Sprint(p...)
}

func StringJoin(p []string) string{
return strings.Join(p,"")
}

func StringBuffer(p []string) string {
var b bytes.Buffer
l:=len(p)
for i:=0;i<l;i++{
b.WriteString(p[i])
}
return b.String()
}

func StringBuilder(p []string) string {
var b strings.Builder
l:=len(p)
for i:=0;i<l;i++{
b.WriteString(p[i])
}
return b.String()
}

const BLOG  = "http://www.flysnow.org/"

func initStrings(N int) []string{
s:=make([]string,N)
for i:=0;i<N;i++{
s[i]=BLOG
}
return s;
}

func initStringi(N int) []interface{}{
s:=make([]interface{},N)
for i:=0;i<N;i++{
s[i]=BLOG
}
return s;
}

func BenchmarkStringPlus10(b *testing.B) {
p:= initStrings(10)
b.ResetTimer()
for i:=0;i<b.N;i++{
StringPlus(p)
}
}

func BenchmarkStringFmt10(b *testing.B) {
p:= initStringi(10)
b.ResetTimer()
for i:=0;i<b.N;i++{
StringFmt(p)
}
}

func BenchmarkStringJoin10(b *testing.B) {
p:= initStrings(10)
b.ResetTimer()
for i:=0;i<b.N;i++{
StringJoin(p)
}
}

func BenchmarkStringBuffer10(b *testing.B) {
p:= initStrings(10)
b.ResetTimer()
for i:=0;i<b.N;i++{
StringBuffer(p)
}
}

func BenchmarkStringBuilder10(b *testing.B) {
p:= initStrings(10)
b.ResetTimer()
for i:=0;i<b.N;i++{
StringBuilder(p)
}
}