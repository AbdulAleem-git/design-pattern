package main

import (
	"fmt"
	"sync"
)

type single struct{}
var singleinstance *single
var lock = &sync.Mutex{}

func getInstance ()*single{
	 if singleinstance == nil {
		 lock.Lock()
		 defer lock.Unlock()
		 if singleinstance == nil {
			 fmt.Println("Creating Single instance now")
			 singleinstance = &single{}
		 }else{
			 fmt.Println("Single instance already created")
		 }
		 return singleinstance
	 }else {
        fmt.Println("Single instance already created.")
    }

    return singleinstance
}

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}