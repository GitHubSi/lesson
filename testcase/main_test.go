package testcase

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testing.Init()
	os.Exit(m.Run())
}

func TestArgs(t *testing.T) {
	str := flag.String("c", "haha", "config")
	flag.Parse()
	fmt.Println(*str, "==")
	//fmt.Println("==")
}
