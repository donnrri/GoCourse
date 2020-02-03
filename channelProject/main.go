package main

/*
This section dels with:
concurrency = running multiple process on
one CPU (i.e swapping between processes)

parrelliem = actually running on different cores
you need more than one core to get that.
*/
import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facbook.com",
		"http://facbook.com",
	}

	c := make(chan string) // make a channel (it is typed)

	/*
			Sequentail:
			Each request runs and waits for response and then next
			for _, link := range links {
				checkLink(link)
			}
		OR
			Parralell
			Using Go routines, each request runs indepenently at same time
			using the 'go' keyword
			Basically registers the process with the Go scheduler

			But you have to set up a channel for the routine , to ensure
			it completes before the main go routine finish.

			Add a routine to a channel since channels can nbe litened to
			by main routine and will not complete until the
			channels have all sent back a message
	*/
	for _, link := range links {
		go checkLink(link, c) // new go routine created
		fmt.Println(<-c)      // add a channel to each routine
	}

	/*
		The course uses a seond loop to add routine listeners
		for i:= 0; i < len(links); i++ {
			fmt.Println(<-c)
		}
	*/

}

// have to pass channel in if using channels
// chan sring a channel typed for a string type
func checkLink(link string, c chan string) {
	res, err := http.Get(link)

	if err != nil {
		fmt.Println("The following website might be  down ", link)
		c <- "might be down"
		return
	}
	fmt.Println(link, " is OK status code ", res.StatusCode)
	c <- "It's up"
}
