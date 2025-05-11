/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "GOPプロジェクトを初期化します",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("引数なしの場合、デフォルトのプロジェクト名を使用します")
		} else {
			fmt.Println("引数が指定されました:", args[0])
		}
		// ここでプロジェクトの初期化処理を実行
		fmt.Println("プロジェクトの初期化が完了しました")
		// 例: ディレクトリの作成、初期設定ファイルの生成など
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
