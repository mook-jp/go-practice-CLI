package utils

import (
	"os"
	"path/filepath"
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
		LastSubSection string `toml:"last_subsection"` // 最後に使用されたサブセクションディレクトリ
	} `toml:"context"`
	Generate struct {
		LessonReadme     bool `toml:"lesson_readme` // レッスンごとに README.md を生成するか
		SectionReadme    bool `toml:"lesson_readme` // レッスンごとに README.md を生成するか
		SubSectionReadme bool `toml:"lesson_readme` // レッスンごとに README.md を生成するか
	}
}

func LoadConfig() Config {
	var config Config
	_ = os.MkdirAll()
}
