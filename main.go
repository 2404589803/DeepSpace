package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	DeepSpace.AddCommand(
		startCommand(),
		listCommand(),
		inspectCommand(),
		cleanupCommand(),
		exportCommand(),
	)
}

var (
	DeepSpace = &cobra.Command{
		Use:           "deepspace",
		Version:       "v0.0.1",
		Short:         "DeepSpace is a command-line tool for debugging the DeepSeek AI HTTP API",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func main() {
	// 如果没有命令行参数，启动 GUI 模式
	if len(os.Args) == 1 {
		StartGUI()
		return
	}

	// 否则使用命令行模式
	if err := DeepSpace.Execute(); err != nil {
		logFatal(err)
	}
}
