package main //goroutines

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("\n Total execution time:%v", time.Since(t0))
	fmt.Println("\n Results are:", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	// m.Lock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result from db is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}
