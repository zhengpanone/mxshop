package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// CallerPath 获取项目根目录
func CallerPath() (path string, err error) {
	path = getCurrentAbPathByExecutable()
	if strings.Contains(path, getTmpDir()) {
		path = getCurrentAbPathByCaller()
	}
	path = strings.Replace(path, "common/utils/filesystem", "", 1)
	return
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	if dir == "" {
		dir = "tmp"
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取最上层的调用者所在目录（跳过工具包的内部调用）（go run）
func getCurrentAbPathByCaller() string {
	pc := make([]uintptr, 10)   // 存储调用栈
	n := runtime.Callers(2, pc) // 跳过 `runtime.Callers` 和 `getCurrentAbPathByCaller`
	frames := runtime.CallersFrames(pc[:n])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if !strings.Contains(frame.File, "/utils/") {
			return filepath.Dir(frame.File)
		}
	}
	log.Fatal("无法找到调用者文件路径")
	return ""
}
