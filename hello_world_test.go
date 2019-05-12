package main

import (
	"testing"
)

// 单元测试
func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}

// 性能测试
func BenchmarkHelloWorld(b *testing.B) {
	a1 := 100
	b1 := 2
	c1 := a1 / b1
	b.Log("hello world", c1)
}

// 覆盖率测试
func TestHelloWorld_PrintHello(t *testing.T) {
	h := &HelloWorld{}
	h.PrintHello()
}
