package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var (
	query          = "test"
	matches        int
	workerCount    = 0
	maxWorkerCount = 32
	searchRequest  = make(chan string)
	workerDone     = make(chan bool)
	foundMatch     = make(chan bool)
)

func main() {
	start := time.Now()
	workerCount = 1
	go search("/Users/max/", true)
	waitForWorkers()
	fmt.Println("matches", matches)
	fmt.Println(time.Since(start))
}
func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

// 改进
func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- true
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}

// func search(path string) {
// 	files, err := ioutil.ReadDir(path)
// 	if err == nil {
// 		for _, file := range files {
// 			name := file.Name()
// 			if name == query {
// 				matches++
// 			}
// 			if file.IsDir() {
// 				search(path + name + "/")
// 			}
// 		}
// 	}
// }
