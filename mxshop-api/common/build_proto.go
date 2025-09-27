// build_proto.go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Config struct {
	ProtoDir        string   // proto 文件目录
	OutputDir       string   // 输出目录
	Recursive       bool     // 是否递归搜索
	ExcludePatterns []string // 排除模式
	SpecificFiles   []string // 指定文件
}

var PROTO_CONFIG = Config{
	ProtoDir:        "proto",
	OutputDir:       "proto/pb",
	Recursive:       true,
	ExcludePatterns: []string{".*_test\\.proto$", "^test_.*\\.proto$"},
	SpecificFiles:   []string{},
}

func build() bool {
	fmt.Println("🚀 自动搜索并编译 Proto 文件...")
	fmt.Printf("📁 源目录: %s\n", PROTO_CONFIG.ProtoDir)
	fmt.Printf("📁 输出目录: %s\n", PROTO_CONFIG.OutputDir)
	fmt.Printf("🔍 递归搜索: %v\n", PROTO_CONFIG.Recursive)

	// 创建输出目录
	_ = os.MkdirAll(PROTO_CONFIG.OutputDir, os.ModePerm)

	// 获取要编译的 proto 文件
	protoFiles := discoverProtoFiles()
	if len(protoFiles) == 0 {
		fmt.Println("❌ 未找到 proto 文件")
		return false
	}

	fmt.Printf("\n📄 找到 %d 个 proto 文件:\n", len(protoFiles))
	for i, f := range protoFiles {
		fmt.Printf("  %d. %s\n", i+1, f)
	}
	fmt.Println()

	// 执行 protoc
	return runProtoc(protoFiles)
}

func discoverProtoFiles() []string {
	if len(PROTO_CONFIG.SpecificFiles) > 0 {
		fmt.Println("📋 使用指定的文件列表")
		return PROTO_CONFIG.SpecificFiles
	}

	var files []string
	err := filepath.Walk(PROTO_CONFIG.ProtoDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if !PROTO_CONFIG.Recursive && path != PROTO_CONFIG.ProtoDir {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(info.Name(), ".proto") {
			if shouldExclude(path) {
				fmt.Printf("⏭️  跳过: %s\n", path)
				return nil
			}
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("❌ 扫描文件失败: %v\n", err)
		return nil
	}

	sort.Strings(files)
	return files
}

func shouldExclude(path string) bool {
	for _, pattern := range PROTO_CONFIG.ExcludePatterns {
		if matched, _ := regexp.MatchString(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

func runProtoc(files []string) bool {
	args := []string{
		"-I", PROTO_CONFIG.ProtoDir,
		"--go_out=" + PROTO_CONFIG.OutputDir,
		"--go_opt=paths=source_relative",
		"--go-grpc_out=" + PROTO_CONFIG.OutputDir,
		"--go-grpc_opt=paths=source_relative",
	}
	args = append(args, files...)

	fmt.Println("⚙️  执行编译命令...")
	fmt.Printf("   protoc %s\n", strings.Join(args, " "))

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ 编译失败: %v\n", err)
		return false
	}
	return true
}

func clean() {
	fmt.Println("🧹 清理生成的文件...")
	outputDir := PROTO_CONFIG.OutputDir
	patterns := []string{"*_pb.go"}

	removed := 0
	for _, p := range patterns {
		matches, _ := filepath.Glob(filepath.Join(outputDir, p))
		for _, file := range matches {
			_ = os.Remove(file)
			fmt.Printf("  - 删除 %s\n", file)
			removed++
		}
	}
	if removed == 0 {
		fmt.Println("📄 没有找到需要清理的文件")
	}
}

func listFiles() {
	fmt.Println("📋 扫描 proto 文件...")
	files := discoverProtoFiles()
	if len(files) == 0 {
		fmt.Println("❌ 未找到 proto 文件")
		return
	}
	fmt.Printf("\n📄 找到 %d 个文件:\n", len(files))
	for i, f := range files {
		info, _ := os.Stat(f)
		fmt.Printf("  %2d. %s (%d bytes)\n", i+1, f, info.Size())
	}
}

func printHelp() {
	fmt.Println(`
🛠️  Proto 自动编译工具 (Go 版)

用法:
  go run build_proto.go           # 自动编译所有 proto 文件
  go run build_proto.go clean     # 清理生成的文件
  go run build_proto.go list      # 列出会被编译的文件
  go run build_proto.go help      # 显示此帮助

配置修改:
  修改 PROTO_CONFIG 变量即可调整:
  - ProtoDir: proto 文件目录
  - OutputDir: 输出目录
  - Recursive: 是否递归搜索
  - ExcludePatterns: 排除模式
  - SpecificFiles: 指定文件 (优先级最高)
`)
}

func main() {
	args := os.Args
	command := "build"
	if len(args) > 1 {
		command = args[1]
	}

	switch command {
	case "clean":
		clean()
	case "list":
		listFiles()
	case "help":
		printHelp()
	default:
		if !build() {
			os.Exit(1)
		}
	}
}
