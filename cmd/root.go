// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "quanfuxia",
	Short: "quanfuxia backend CLI tool",
	Long:  "quanfuxia 是一个使用 Gin 开发的 Go 后端项目，支持 HTTP 服务、定时任务、消息消费等子命令",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "命令执行失败: %v\n", err)
		os.Exit(1)
	}
}
