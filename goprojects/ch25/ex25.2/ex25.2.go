package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파일 열어서 라인 읽기
func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("The file could not be found.", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
