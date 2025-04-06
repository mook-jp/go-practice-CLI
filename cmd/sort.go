// The MIT License (MIT)

// Copyright © 2025 mook-jp <mook24.jp@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort Directory",
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "./"
		var numSlice []int

		// dir内のエントリを取得する
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return
		}

		// フォルダーのみ表示
		for _, entry := range entries {
			if entry.IsDir() && entry.Name()[0:2] == "gp" {
				i, err := strconv.Atoi(entry.Name()[2:4])
				if err != nil {
					fmt.Println(err)
					return
				}
				numSlice = append(numSlice, i)
			}
		}

		// 逆順に並び替え
		// sort.Slice(numSlice, func(i, j int) bool {
		// 	return numSlice[i] > numSlice[j]
		// })
		// xx := len(numSlice)

		for num, dirNumber := range numSlice {
			oldName := "./gp" + fmt.Sprintf("%02d", dirNumber)
			newName := "./gp" + fmt.Sprintf("%02d", num+1)
			// 逆順に並び替え
			// newName := "./gp" + fmt.Sprintf("%02d", xx-num)

			err := os.Rename(oldName, newName)
			// エラーの出し方を検討する必要あり。
			// すでにソート済みのフォルダについてはエラーを表示しない
			// 00から始まっている場合に、01から始めるようなサブコマンドを作成
			if err != nil {
				fmt.Println("フォルダ名の変更に失敗: ", err)
			} else {
				fmt.Printf("フォルダ名を変更(old: %s -> new: %s)\n", oldName, newName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
}
