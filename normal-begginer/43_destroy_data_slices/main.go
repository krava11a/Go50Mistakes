package main

import (
	"bytes"
	"fmt"
)

func main() {

	//BAD
	//path := []byte("AAAA/BBBBBBBBB")
	//sepIndex := bytes.IndexByte(path,'/')
	//dir1 := path[:sepIndex]
	//dir2 := path[sepIndex+1:]
	//fmt.Println("dir1 =>",string(dir1)) // выводит: dir1 => AAAA
	//fmt.Println("dir2 =>",string(dir2)) // выводит: dir2 => BBBBBBBBB
	//
	//dir1 = append(dir1,"suffix"...)
	//path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	//
	//fmt.Println("dir1 =>",string(dir1)) // выводит: dir1 => AAAAsuffix
	//fmt.Println("dir2 =>",string(dir2)) // выводит: dir2 => uffixBBBB (not ok)
	//
	//fmt.Println("new path =>",string(path))

	//BETTER
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex:sepIndex] // полное выражение слайса
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) // выводит: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) // выводит: dir2 => BBBBBBBBB

	dir1 = append(dir1,"suffix"...)
	path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

	fmt.Println("dir1 =>",string(dir1)) // выводит: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) // выводит: dir2 => BBBBBBBBB (ok now)

	fmt.Println("new path =>",string(path))

}
