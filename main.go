package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	port int
	path string
)

func init() {
	flag.IntVar(&port, "port", 80, "")
	flag.StringVar(&path, "path", "./", "")
}

func main() {
	flag.Parse()

	e := gin.Default()
	e.Static("/", path)
	e.Run(fmt.Sprintf("0.0.0.0:%d", port))

}
