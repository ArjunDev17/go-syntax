package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// TODO :create pool of bytes.Buffers which can be reused
var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, debug string) {
	// var bf bytes.Buffer
	bf := bufPool.Get().(*bytes.Buffer)
	bf.WriteString(time.Now().Format("15:03:02"))
	bf.WriteString(" : ")
	bf.WriteString(debug)
	bf.WriteString("\n")

	w.Write(bf.Bytes())
	bufPool.Put(bf)
}
func main() {
	log(os.Stdout, "deb-str_1")
	log(os.Stdout, "deb-str_2")
}
