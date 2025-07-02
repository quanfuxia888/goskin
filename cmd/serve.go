package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"quanfuxia/internal/common"
	"quanfuxia/internal/model/query"
	"quanfuxia/internal/route"
	"quanfuxia/pkg/config"
	"quanfuxia/pkg/logger"
	"quanfuxia/pkg/mq"
	"quanfuxia/pkg/mysql"
	"quanfuxia/pkg/redis"
)

var configPath string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(configPath) // 加载配置
		// 初始化 i18n + 校验翻译器
		if err := common.InitI18n(); err != nil {
			log.Fatalf("初始化 i18n 失败: %v", err)
		}
		fmt.Println("服务名称：", config.Cfg.App.Name)
		fmt.Println("服务端口：", config.Cfg.App.Port)
		// 后续初始化日志、DB、Gin 等
		logger.Init()
		mysql.Init()
		logger.L().Info("服务启动成功")
		//ctx := context.Background()
		query.SetDefault(mysql.DB)
		//city, err := query.City.WithContext(ctx).Where(query.City.ID.Eq(113)).First()
		//
		//if err != nil {
		//	fmt.Println(err.Error())
		//} else {
		//	fmt.Println(city.Name)
		//}
		// 在 main.go 或 serve.go
		redis.Init()
		redis.InitLocker()
		mq.Init()
		r := route.InitRouter()
		port := config.Cfg.App.Port
		addr := fmt.Sprintf(":%s", port)
		fmt.Println("✅ 服务启动在 " + addr)
		r.Run(addr)
	},
}

func init() {
	serveCmd.Flags().StringVar(&configPath, "config", "", "配置文件路径，默认使用 configs/config.yaml")
	RootCmd.AddCommand(serveCmd)
}
