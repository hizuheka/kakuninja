package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ファイルの内容に文字列が含まれているかチェックする関数
func containsStringInFile(content string, searchString string) bool {
	return strings.Contains(content, searchString)
}

// 指定されたファイルに対して文字列検索を行う
func searchInFile(filePath string, searchString string) bool {
	// ファイルを一気に読み込む
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ファイルを開けません:", err)
		os.Exit(1)
	}

	// ファイルの内容に検索する文字列が含まれているかチェック
	return containsStringInFile(string(content), searchString)
}

func main() {
	// サブコマンド `contains` の設定
	if len(os.Args) < 2 || os.Args[1] != "contains" {
		fmt.Println("使用法: contains -f <対象ファイル> -t <検索する文字列>")
		os.Exit(1)
	}

	// 引数の解析
	file := flag.String("f", "", "検索対象のファイル")
	searchString := flag.String("t", "", "検索する文字列")
	flag.Parse()

	// ファイルと検索文字列の指定を確認
	if *file == "" || *searchString == "" {
		fmt.Println("ファイルと検索文字列の両方を指定してください")
		os.Exit(1)
	}

	// ファイル内で文字列を検索
	if searchInFile(*file, *searchString) {
		// 文字列が見つかった場合、終了コード0で終了
		os.Exit(0)
	}

	// 文字列が見つからなかった場合、終了コード31で終了
	os.Exit(31)
}
