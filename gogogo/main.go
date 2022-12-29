package main

import (
	"fmt"
	"net"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		c.Close()
		return true
	}
	return false
}

type Node struct {
	Data      int
	NextPoint *Node
	PrePoint  *Node
}
type LinkedList struct {
	Head    *Node
	Current *Node
	Tail    *Node
}

func CreateLinkedList() {
	data := []int{1, 21, 31, 51, 62, 2, 3, 43, 55, 23, 12}
	link := LinkedList{}
	var currentNode *Node
	for i := 0; i < len(data); i++ {
		currentNode = new(Node)
		currentNode.Data = data[i]
		InsertNode(&link, currentNode)
	}
	ShowLinkedList(link)
}

func InsertNode(link *LinkedList, node *Node) {
	if link.Head == nil {
		link.Head = node
		link.Tail = node
		link.Current = node
	} else {
		link.Tail.NextPoint = node
		node.PrePoint = link.Tail
		link.Tail = node
	}
}

func ShowLinkedList(link LinkedList) {
	var currentNode *Node
	currentNode = link.Head
	for {
		fmt.Println("Node: ", currentNode.Data)
		if currentNode.NextPoint == nil {
			break
		} else {
			currentNode = currentNode.NextPoint
		}
	}
}

func main() {
	// ports := []int{}

	// wg := &sync.WaitGroup{}
	// lock := &sync.Mutex{}
	// timeout := time.Millisecond * 200
	// for port := 1; port < 100; port++ {
	// 	wg.Add(1)
	// 	go func (p int)  {
	// 		opened := isOpen("www.taobao.com",p ,timeout)
	// 		if opened {
	// 			lock.Lock()
	// 			ports = append(ports, p)
	// 			lock.Unlock()
	// 		}
	// 		wg.Done()
	// 	}(port)
	// }
	// wg.Wait()
	// fmt.Println(ports)
	CreateLinkedList()
}
