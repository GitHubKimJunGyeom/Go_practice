package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 해법
// 1. 찾으려는 단어와 파일 경로를 입력
// 2. 경로에 해당하는 파일을 찾음. 특정 파일 하나만 나타내거나 여러파일을 나타낼 수도 있음
// 3. 파일을 읽고 각 라인에서 해당 단어가 나오는 지 확인
// 4. 특정 단어가 등장하는 라인을 취합하여 마지막으로 결과 출력

// 와일드 카드
// 사용 문자는 *와 ?
// * : 0개 이상의 아무 문자를 나타냄
// ? : 문자 하나를 나타냄

// os.Args변수, 실행 인수를 가져온다
// 첫번쨰 인수에 실행 명령이 들어가고 두번째 인수부터 입력한 인수가 저장된다.

// 파일 핸들링
// 순서 : 파일 열기-> 파일 목록 가져오기 -> 파일 내용 읽기
// 단어 포함 여부 검사

// 실행 인수 읽고 파일 목록 가져오기
func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least two execution arguments are required.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("the word you're looking for", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("The file path is invalid.")
			return
		}
		fmt.Println("List of files you want to find")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
