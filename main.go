package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var (
		limit      = flag.Int("l", 128, "depth limit")
		showHidden = flag.Bool("a", false, "hidden files")
	)
	flag.Parse()
	fArgs := flag.Args()
	root := "."
	if len(fArgs) > 0 {
		root = fArgs[0]
	}
	err := walk(root, root, 0, *limit, *showHidden)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func walk(root string, name string, level int, limit int, showHidden bool) error {
	if level >= limit {
		return nil
	}
	for i := 0; i < level; i++ {
		for n := 0; n < 2; n++ {
			fmt.Print(" ")
		}
	}
	fmt.Println(name)
	// read sub entries
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := file.Name()
		if filename == "." || filename == ".." {
			continue
		}
		if filename[0] != '.' || showHidden {
			walk(filepath.Join(root, filename), filename, level+1, limit, showHidden)
		}
	}
	return nil
}
