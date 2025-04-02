package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"syscall"
)

func CreateFile(dirName string) error {
	mainString := `package main

func main(){

}`

	file, err := os.Create(dirName + "/" + "main.go")
	if err != nil {
		redln("ファイル作成エラー: ", err)
		return err
	}
	defer file.Close()

	// ファイルにデータを書き込む
	_, err = file.WriteString(mainString)
	if err != nil {
		redln("書き込みエラー: ", err)
		return err
	}

	return nil
}

func max(nums []int) int {
	maxVal := math.MinInt
	for _, val := range nums {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}

func getNewDirName() string {
	dirPath := "./"
	var numSlice []int

	// dir内のエントリを取得
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("エラー: ", err)
		return ""
	}

	// フォルダーのみを表示
	isExistsDir := false
	for _, entry := range entries {
		if entry.IsDir() && entry.Name()[0:2] == "gp" {
			isExistsDir = true
			i, err := strconv.Atoi(entry.Name()[2:4])
			if err != nil {
				fmt.Println("エラー: 予想しないフォルダ名(例: gtXXにて数字以外が入力されています)")
				return ""
			}
			numSlice = append(numSlice, i)
		}
	}

	newDirName := "gp01"
	if isExistsDir {
		newDirName = "gp" + fmt.Sprintf("%02d", max(numSlice)+1)
	}
	return newDirName
}

func CreateNewPackage() {
	newDirName := getNewDirName()
	if newDirName == "" {
		return
	}
	fileInfo, err := os.Lstat("./")
	if err != nil {
		fmt.Println(err)
	}
	fileMode := fileInfo.Mode()
	unixPerms := fileMode & os.ModePerm // 0777 と ./のパーミッションのアンドを取得
	if err := os.Mkdir(newDirName, unixPerms); err != nil {
		// 説明
		// https://chatgpt.com/c/67cbf18e-e258-8006-afd3-b658e456531f
		// `*os.PathError` にキャスト
		if pathErr, ok := err.(*os.PathError); ok {
			if errno, ok := pathErr.Err.(syscall.Errno); ok {
				// `syscall.Errno` を取得
				switch errno {
				case 1:
					redln("MKDIR failed: 操作許可なし(例: /root/への作成など)")
				case 2:
					redln("MKDIR failed: ファイルが存在しない (例: 親ディレクトリが存在しない等)")
				case 13:
					redln("MKDIR failed: アクセス拒否(例：権限がないディレクトリです)")
				case 17:
					redln("MKDIR failed: ファイルがすでに存在(例: 既にディレクトリが存在)")
				case 28:
					redln("MKDIR failed: ディスク容量不足(例: 空き容量がない)")
				default:
					redf("MKDIR failed: %v (errno: %d)\n", err, errno)
				}
				return
			}
		}
		redln("Other error<other os.PathError>: ", err)
	} else {
		greln("MKDIR successfully")
		err := CreateFile(newDirName)
		if err != nil {
			redln(err)
		}
	}
}
