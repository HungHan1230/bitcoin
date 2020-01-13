package main

import "C"

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

//export AddObjectToIPFS
func AddObjectToIPFS(str string) *C.char {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	// fmt.Printf("added %s", cid)
	return C.CString(cid)
	// C.CString()
}

//export GetObjectFromIPFS
func GetObjectFromIPFS(str *C.char) *C.char {
	sh := shell.NewShell("localhost:5001")

	io, err := sh.Cat(C.GoString(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(io)
	newStr := buf.String()
	// fmt.Println(newStr)
	return C.CString(newStr)
}

func main() {}
