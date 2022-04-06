package main

import (
	"fmt"
	"sync"
)

func main() {
	urls := []string{"https://br.indeed.com/jobs?q&l=Bras%C3%ADlia%2C%20DF&vjk=bd05e6fb411225f9"}
	//internships := make(map[string][]string)
	//messages := make(chan map[string][]string)
	routines := sync.WaitGroup{}

	for _, link := range urls {
		routines.Add(1)
		go get(link, &routines)
	}

	//go func() {
	fmt.Println("WAITING")
	routines.Wait()
	//close(messages)
	//}()

	//for m := range messages {
	//	//internships[m[0]] = m[1:]
	//}
	//fmt.Println(internships)
	fmt.Println("DONE")
}
