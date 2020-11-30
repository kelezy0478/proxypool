package main

import (
	"fmt"
	"os"
	log "github.com/kataras/golog"
	"path/filepath"
)

var dir string

func main() {
	/**/
	log.SetLevel("info")
	log.Info("Set Log Level: Debug.")

	path,_ := os.Executable()
	fmt.Println("当前路径：",path)
	dir := filepath.Dir(path)
	fmt.Println("当前路径：",dir)
	CrawlerRun()
}
