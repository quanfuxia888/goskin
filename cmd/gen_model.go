package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"quanfuxia/pkg/config"
	"strings"
)

var (
	genConfig      string
	genModelTables string
)
var genModelCmd = &cobra.Command{
	Use:   "gen-model",
	Short: "Generate model struct from database table",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(genConfig)
		db, _ := gorm.Open(mysql.Open(config.Cfg.MySQL.DSN))

		g := gen.NewGenerator(gen.Config{
			OutPath:      "internal/model/gens",
			ModelPkgPath: "internal/model/gens", // 读取 model
		})
		g.UseDB(db)

		if genModelTables != "" {
			tables := parseTables(genModelTables)
			for _, t := range tables {
				g.GenerateModel(t) // 只生成 struct，不生成 query
			}
		} else {
			g.GenerateAllTable()
		}

		g.Execute()
		fmt.Println("✅ 模型生成成功")
	},
}

func init() {
	genModelCmd.Flags().StringVar(&genModelTables, "tables", "", "表名，逗号分隔")
	genModelCmd.Flags().StringVar(&genConfig, "config", "", "配置路径")
	RootCmd.AddCommand(genModelCmd)
}

func parseTables(s string) []string {
	parts := strings.Split(s, ",")
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
