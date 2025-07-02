package cmd

import (
	"fmt"
	"os"
	"quanfuxia/internal/common"

	"github.com/spf13/cobra"
)

var output string

var genErrDocCmd = &cobra.Command{
	Use:   "gen-errdoc",
	Short: "生成错误码 Markdown 文档",
	Run: func(cmd *cobra.Command, args []string) {
		all := common.GetAllCodes()
		common.InitI18n()
		outputText := "# 错误码文档\n\n| Code | Key | 中文含义 |\n|------|-----|----------|\n"
		for code, key := range all {
			zh := common.Translate("zh", key)
			outputText += fmt.Sprintf("| %d | %s | %s |\n", code, key, zh)
		}

		if err := os.WriteFile(output, []byte(outputText), 0644); err != nil {
			fmt.Println("❌ 写入失败:", err)
			return
		}
		fmt.Println("✅ 错误码文档生成成功:", output)
	},
}

func init() {
	genErrDocCmd.Flags().StringVarP(&output, "output", "o", "docs/error_codes.md", "输出文件路径")
	RootCmd.AddCommand(genErrDocCmd)
}
