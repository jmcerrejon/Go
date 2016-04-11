/*
 * Program: coordinates splitter
 * Author: Jose Manuel Cerrejon Gonzalez (ulysess _at_ gmail.com)
 * Version: 0.1 (4/11/16)
 * Description: Get a "x1 y1 x2 y2" one-line coordinate system in a file and split into:
 * x1 y1
 * x2 y2
 * ...
 *
 * TODO:     V Arguments
 * 	     · Check with RegEx the right input format
 * 	     V Split logic on process depending on cores (using Go Routines)
 * 	     · Generate multiplatform binaries
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// Run before main()
func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	if len(os.Args) < 2 {
		// No arguments
		fmt.Println("You must to specify a text file.")
		os.Exit(0)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	fileArgs := os.Args[1]

	// Check type
	// fmt.Println(reflect.TypeOf(os.Args))

	// Take the file name without extension
	var extension = filepath.Ext(fileArgs)
	var fileName = fileArgs[0 : len(fileArgs)-len(extension)]

	fopen, err := os.Open(fileArgs)
	if err != nil {
		// handle the error here
		return
	}
	defer fopen.Close()

	// get the fopen size
	stat, err := fopen.Stat()
	if err != nil {
		return
	}
	// read the fopen
	bs := make([]byte, stat.Size())
	_, err = fopen.Read(bs)
	if err != nil {
		return
	}

	var newContent string
	str := strings.Split(string(bs), " ")

	// GoRoutine
	go func() {

		for i := 0; i < len(str); i = i + 2 {
			newContent += str[i] + " " + str[i+1] + "\n"
		}

		wg.Done()
	}()

	// Wait to end the code above
	wg.Wait()

	// Make & write the file
	fwrite, err := os.Create(fileName + "_converted.txt")
	if err != nil {
		return
	}
	defer fwrite.Close()

	fwrite.WriteString(newContent)
}
