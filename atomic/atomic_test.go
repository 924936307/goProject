package main

import "testing"

//单元测试

func TestAtomic(t *testing.T) {
	atomic := &Atomic{}
	test(atomic)
}
