package common

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 定义命令行参数
var (
	Port         = flag.Int("port", 3000, "the listening port") // 监听端口号
	PrintVersion = flag.Bool("version", false, "print version and exit") // 打印版本号并退出程序
	PrintHelp    = flag.Bool("help", false, "print help and exit") // 打印帮助信息并退出程序
	LogDir       = flag.String("log-dir", "", "specify the log directory") // 指定日志目录
)

// 打印帮助信息
func printHelp() {
	fmt.Println("One API " + Version + " - All in one API service for OpenAI API.")
	fmt.Println("Copyright (C) 2023 JustSong. All rights reserved.")
	fmt.Println("GitHub: https://github.com/songquanpeng/one-api")
	fmt.Println("Usage: one-api [--port <port>] [--log-dir <log directory>] [--version] [--help]")
}

// 初始化函数
func init() {
	flag.Parse()

	// 打印版本号并退出
	if *PrintVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	// 打印帮助信息并退出
	if *PrintHelp {
		printHelp()
		os.Exit(0)
	}

	// 配置环境变量
	if os.Getenv("SESSION_SECRET") != "" {
		SessionSecret = os.Getenv("SESSION_SECRET")
	}
	if os.Getenv("SQLITE_PATH") != "" {
		SQLitePath = os.Getenv("SQLITE_PATH")
	}

	// 设置日志目录
	if *LogDir != "" {
		var err error
		*LogDir, err = filepath.Abs(*LogDir)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(*LogDir); os.IsNotExist(err) {
			err = os.Mkdir(*LogDir, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
