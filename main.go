package main

import (
	"fmt"

	"github.com/lukeoleson/tourofgo/mutex"
	"github.com/lukeoleson/tourofgo/channels"

	"golang.org/x/tour/tree"
)


func main() {

	// channels
	fmt.Println(channels.Same(tree.New(1), tree.New(1)))

	// mutex
	mutex.Crawler()

}
