package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var (
	sectionName      string
	subSectionName   string
	sectionNumber    string
	subSectionNumber string
	configPath       = filepath.Join("tour", ".gopractice.toml")
)

// Config は CLI の設定ファイル (.gopractice.toml) の構造です。
type Config struct {
	Context struct {
		LastSection    string `toml:"last_section"`    // 最後に使用されたセクションディレクトリ
		LastSubsection string `toml:"last_subsection"` // 最後に使用されたサブセクションディレクトリ
	} `toml:"context"`
	Generate struct {
		LessonReadme     bool `toml:"lesson_readme"`     // レッスンごとに README.md を生成するか
		SectionReadme    bool `toml:"section_readme"`    // セクションごとに README.md を生成するか
		SubsectionReadme bool `toml:"subsection_readme"` // サブセクションごとに README.md を生成するか
	} `toml:"generate"`
}

// loadConfig は tour/.gopractice.toml を読み込み、設定を返します。
func loadConfig() Config {
	var config Config
	_ = os.MkdirAll("tour", os.ModePerm)
	_, err := os.Stat(configPath)
	if err == nil {
		_, err = toml.DecodeFile(configPath, &config)
		if err != nil {
			fmt.Println("Failed to read config:", err)
		}
	}
	return config
}

// saveConfig は設定を tour/.gopractice.toml に保存します。
func saveConfig(config Config) {
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Failed to save config:", err)
		return
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		fmt.Println("Failed to encode config:", err)
	}
}

// newCmd は gopractice new コマンドで新しいレッスンを作成します。
var newCmd = &cobra.Command{
	Use:   "new [lesson-name]",
	Short: "Create a new Go lesson with automatic folder structure",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		base := "tour"
		lessonRaw := args[0]
		_ = os.MkdirAll(base, os.ModePerm)

		config := loadConfig()
		finalPath := base

		// セクションの指定と作成／再利用
		if sectionNumber != "" {
			secDir := getDirByNumber(base, sectionNumber)
			if secDir == "" {
				fmt.Printf("No section found with number %s\n", sectionNumber)
				return
			}
			finalPath = filepath.Join(base, secDir)
			config.Context.LastSection = secDir
		} else if sectionName != "" {
			secDir := getOrCreateNumberedDir(base, sectionName)
			finalPath = filepath.Join(base, secDir)
			config.Context.LastSection = secDir
			if config.Generate.SectionReadme {
				createReadme(filepath.Join(base, secDir), sectionName)
			}
		} else if config.Context.LastSection != "" {
			finalPath = filepath.Join(base, config.Context.LastSection)
		}

		// サブセクションの指定と作成／再利用
		if subSectionNumber != "" && finalPath != base {
			subDir := getDirByNumber(finalPath, subSectionNumber)
			if subDir == "" {
				fmt.Printf("No subsection found with number %s\n", subSectionNumber)
				return
			}
			finalPath = filepath.Join(finalPath, subDir)
			config.Context.LastSubsection = subDir
		} else if subSectionName != "" && finalPath != base {
			subDir := getOrCreateNumberedDir(finalPath, subSectionName)
			finalPath = filepath.Join(finalPath, subDir)
			config.Context.LastSubsection = subDir
			if config.Generate.SubsectionReadme {
				createReadme(finalPath, subSectionName)
			}
		} else if config.Context.LastSubsection != "" && finalPath != base {
			finalPath = filepath.Join(finalPath, config.Context.LastSubsection)
		}

		// レッスンディレクトリの作成
		lessonDir := getNextNumberedDir(finalPath, lessonRaw)
		lessonPath := filepath.Join(finalPath, lessonDir)
		_ = os.MkdirAll(lessonPath, os.ModePerm)

		// main.go の作成
		mainFile := filepath.Join(lessonPath, "main.go")
		mainContent := `package main

func main() {

}
`
		_ = os.WriteFile(mainFile, []byte(mainContent), 0644)

		// lesson の README.md を作成（設定次第）
		if config.Generate.LessonReadme {
			createReadme(lessonPath, lessonRaw)
		}

		saveConfig(config)

		fmt.Printf("Lesson created at: %s\n", lessonPath)
	},
}

// createReadme は与えられたパスに README.md を作成します（既存ならスキップ）
func createReadme(path, title string) {
	readme := filepath.Join(path, "README.md")
	if _, err := os.Stat(readme); err == nil {
		return // すでに存在 → 何もしない
	}
	content := fmt.Sprintf("# %s\n\nこのフォルダは Go 学習用のセクション/レッスンです。\n", title)
	_ = os.WriteFile(readme, []byte(content), 0644)
}

// getOrCreateNumberedDir は指定名に対応するディレクトリを取得または作成します（番号付き）
func getOrCreateNumberedDir(base, name string) string {
	existing := getAllNumberedDirs(base)
	for _, dir := range existing {
		if strings.HasSuffix(dir, "_"+name) {
			return dir
		}
	}
	newIndex := len(existing) + 1
	dirName := fmt.Sprintf("%02d_%s", newIndex, name)
	_ = os.MkdirAll(filepath.Join(base, dirName), os.ModePerm)
	return dirName
}

// getNextNumberedDir は base 配下に新しい番号を振ったディレクトリ名を返します
func getNextNumberedDir(base, name string) string {
	existing := getAllNumberedDirs(base)
	newIndex := len(existing) + 1
	return fmt.Sprintf("%02d_%s", newIndex, name)
}

// getDirByNumber は番号（例：01）に一致するディレクトリを返します
func getDirByNumber(base, num string) string {
	prefix := fmt.Sprintf("%02s_", num)
	entries, err := os.ReadDir(base)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), prefix) {
			return e.Name()
		}
	}
	return ""
}

// getAllNumberedDirs は base 配下の "01_xxx" のようなディレクトリを一覧で返します
func getAllNumberedDirs(base string) []string {
	entries, err := os.ReadDir(base)
	if err != nil {
		return nil
	}
	var dirs []string
	for _, e := range entries {
		if e.IsDir() && isNumberedDir(e.Name()) {
			dirs = append(dirs, e.Name())
		}
	}
	sort.Strings(dirs)
	return dirs
}

// isNumberedDir は "01_xxx" 形式の名前かどうかを判定します
func isNumberedDir(name string) bool {
	parts := strings.SplitN(name, "_", 2)
	if len(parts) < 2 {
		return false
	}
	_, err := strconv.Atoi(parts[0])
	return err == nil
}

// init は new コマンドのフラグ初期化を行います
func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVar(&sectionName, "sec", "", "Section name")
	newCmd.Flags().StringVar(&subSectionName, "subsec", "", "Subsection name")
	newCmd.Flags().StringVar(&sectionNumber, "secno", "", "Section number (e.g. 01)")
	newCmd.Flags().StringVar(&subSectionNumber, "subsecno", "", "Subsection number (e.g. 01)")
}
