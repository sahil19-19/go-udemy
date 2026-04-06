package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string) // we make a channel of type string

	for _, link := range sites {
		go checkLink(link, c)
	}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(<-c)
	// }

	// an infinite loop
	// when a link is checked once, we want to check it again
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternate way to write the above code
	// wait for the channel to return smthing, when it does assign it to `link`
	// a lot more clear code
	// time.Second > 1 sec
	for link := range c {
		// time.Sleep(5 * time.Second)
		// placing sleep here would cause the main function to sleep for 5 seconds
		// we dont want that as that would stop the main function
		// go checkLink(link, c)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(link) // we pass the link to the function literal so that even if the value of `link` in the outer scope changes
		// nothing happens to the value that we pass
	}

}

// giving access of channel c to checkLink function
// we can use it to communicate with any go routine using this function to main function
func checkLink(link string, c chan string) {
	// time.Sleep(time.Second * 5) // writing sleep here is not correct as well
	// as we dont want to wait for the result when sm1 runs checkLink()
	_, err := http.Get(link) // this is a blocking call, it takes some time to complete
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
