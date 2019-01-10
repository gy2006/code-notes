package main

import (
	"testing"
	"time"
)

func TestShouldStartThread(t *testing.T) {
	start()
	time.Sleep(15 * time.Second)
}

func TestRunChannelDemo(t *testing.T) {
	startChannelDemo()
}

func TestRunSelectDemo(t *testing.T) {
	selectDemo()
}
