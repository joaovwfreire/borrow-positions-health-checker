package main

import (
	"sync"
	//scripts "main/borrowersGenerator"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(4)
	go scripts(11362579, 11396312, &wg)
	go scripts(13500000, 13504000, &wg)
	go scripts(15000000, 15000400, &wg)
	go scripts(16500000, 16500400, &wg)

	wg.Wait()
}
