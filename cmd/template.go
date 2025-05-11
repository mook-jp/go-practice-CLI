/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	templatePathVar string
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template <name> [templateString]",
	Short: "テンプレートファイルを作成します。",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var templateString string
		var templateName string
		if len(args) == 0 {
			return fmt.Errorf("テンプレート名称は必須です。名称を指定してください。")
		} else if len(args) == 1 {
			if templatePathVar == "" {
				return fmt.Errorf(
					"templateString および テンプレートファイルが未指定です。\n" +
						"テンプレートファイルのパスを指定するか、templateString を指定してください。\n" +
						"例: gopractice template <name> <templateString>",
				)
			}
			path := strings.TrimSpace(templatePathVar)
			info, err := os.Stat(path)
			if err != nil {
				return fmt.Errorf("ファイルが存在しないか、読み取り権限がありません: %s\n%w", path, err)
			} else if info.IsDir() {
				return fmt.Errorf("指定されたパスはディレクトリです: %s", path)
			}

			buf, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("ファイルの読み込みに失敗しました: %s\n%w", path, err)
			}
			templateString = string(buf)
		} else if len(args) == 2 {
			templateString = args[1]
		}

		templateName = args[0]

		// ここでテンプレートファイルの作成処理を実行
		fmt.Printf("テンプレートファイルの作成が完了しました\nName:%s\ntemplate:\n%s\n", templateName, templateString)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringVarP(&templatePathVar, "file", "f", "", "テンプレートファイルのパスを指定します。")
}
