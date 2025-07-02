package common

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"os"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/yaml.v2"
)

var (
	UniTrans    *ut.UniversalTranslator
	Translators = make(map[string]ut.Translator)     // zh、en translator 实例
	langMap     = make(map[string]map[string]string) // 业务自定义翻译 map
)

func InitI18n() error {
	zhT := zh.New()
	enT := en.New()
	UniTrans = ut.New(enT, zhT, enT)

	langs := []string{"zh", "en"}
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil
	}

	for _, lang := range langs {
		trans, found := UniTrans.GetTranslator(lang)
		if !found {
			continue
		}
		Translators[lang] = trans

		// 注册校验器翻译
		switch lang {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(validate, trans)
		case "en":
			_ = en_translations.RegisterDefaultTranslations(validate, trans)
		}

		// 加载业务翻译文件
		loadManualLang(lang)
	}

	return nil
}

func loadManualLang(lang string) {
	path := fmt.Sprintf("configs/i18n/%s.yaml", lang)
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	messages := make(map[string]string)
	_ = yaml.Unmarshal(data, &messages)
	langMap[lang] = messages
}

func Translate(lang, key string) string {
	// 先从业务自定义翻译中取
	if trans, ok := langMap[lang][key]; ok {
		return trans
	}
	return key
}
