package ipfs

import (
	"bytes"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
)

func AddImageFile() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	file_path := "datalayer01.webp"
	data, err1 := os.ReadFile(file_path)
	if err1 != nil {
		fmt.Println(err1)
	}
	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
}
