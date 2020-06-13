package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// i++ は式ではなく文である
	for i := 1; i < len(os.Args); i++ {
		// P6. 2次のオーダーの処理:
		// 結合するたびにコピーが走っている (コピー量がnの2乗になる)
		// 文字列は普遍なので一度作られると書き換えたりすることができないため
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
