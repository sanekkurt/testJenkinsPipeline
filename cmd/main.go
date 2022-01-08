package main

import (
	"testJenkinsPipeline/funcs"
	"time"
)

func main() {
	funcs.Sum(5, 5)
	time.Sleep(1 * time.Hour)
}
