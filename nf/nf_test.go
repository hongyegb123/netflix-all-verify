package nf

import (
	"fmt"
	"testing"
)

func TestNF(t *testing.T) {
	proxyUrl := "localhost:1081"
	ok, out := NF("http://" + proxyUrl)
	fmt.Println(ok, out)
}
