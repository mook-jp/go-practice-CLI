/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
