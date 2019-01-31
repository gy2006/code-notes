package main

import (
	"testing"
)

func TestShouldStartThread(t *testing.T) {
	start()
	// time.Sleep(15 * time.Second)
}

func TestRunChannelDemo(t *testing.T) {
	startChannelDemo()
}

func TestRunSelectDemo(t *testing.T) {
	selectDemo()
}
