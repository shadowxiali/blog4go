// Copyright (c) 2015, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package blog4go

import (
	"fmt"
	"testing"
)

func BenchmarkSocketWriter(b *testing.B) {
	b.StopTimer()
	writer, err := NewSocketWriter("udp", "127.0.0.1:12124")
	defer writer.Close()
	if nil != err {
		fmt.Println(err.Error())
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		writer.Debugf("haha %s. en\\en, always %d and %f", "eddie", 18, 3.1415)
	}
}