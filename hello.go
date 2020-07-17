package main

import (
	"fmt"

	"github.com/hello-jenkins-larkintuckerllc/hello-jenkins/sum"
)

func main() {
	s := sum.Sum(2, 3)
	fmt.Println(s)
}
