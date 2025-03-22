package utils

import (
	"fmt"
	"os"
	"syscall"
)

func redln(args ...any) {
	args = append([]any{"\033[31m"}, args...)
	fmt.Println(args...)
}

func greln(args ...any) {
	args = append([]any{"\033[32m"}, args...)
	fmt.Println(args...)
}

func redf(format string, args ...any) {
	fmt.Printf("\033[31m"+format, args...)
}

func gref(format string, args ...any) {
	fmt.Printf("\033[32m"+format, args...)
}

func Mkdir() {
	fileInfo, err := os.Lstat("./")
	if err != nil {
		fmt.Println(err)
	}
	fileMode := fileInfo.Mode()
	unixPerms := fileMode & os.ModePerm // 0777 と ./のパーミッションのアンドを取得
	if err := os.Mkdir("xx/", unixPerms); err != nil {
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
	}
}
