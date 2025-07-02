package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"quanfuxia/internal/model/gens"
	"quanfuxia/pkg/config"
)

var genQueryCmd = &cobra.Command{
	Use:   "gen-query",
	Short: "Generate type-safe query code from existing model structs",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(genConfig)
		db, _ := gorm.Open(mysql.Open(config.Cfg.MySQL.DSN))

		g := gen.NewGenerator(gen.Config{
			OutPath:      "internal/model/query",
			ModelPkgPath: "internal/model/query", // 读取 model
			Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
		})
		g.UseDB(db)
		// ⚠️ 注意：你必须定义 model.User{}，否则不会生成 query
		g.ApplyBasic(gens.City{}, gens.WaUser{}) // 👈 你已有的模型 struct

		g.Execute()
		fmt.Println("✅ 查询代码生成成功")
	},
}

func init() {
	genQueryCmd.Flags().StringVar(&genConfig, "config", "", "配置路径")
	RootCmd.AddCommand(genQueryCmd)
}
