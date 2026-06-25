package main

import (
	"go.uber.org/zap"
)

func main() {
	// добавляем предустановленный логер NewDevelopment
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic("cannot initialize zap")
	}
	// это нужно добавить, если логер буферизован
	// в данном случае не буферизован, но привычка хорошая
	defer logger.Sync()

	// для примера берём простой URL
	const url = "http://example.com"

	// делаем логер SugaredLogger
	sugar := logger.Sugar()

	// выводим сообщение уровня Info с парой "url": url в виде JSON, это SugaredLogger
	sugar.Infow(
		"Failed to fetch URL",
		"url", url,
	)

	// выводим сообщение уровня Info, но со строкой URL, это тоже SugaredLogger
	sugar.Infof("Failed to fetch URL: %s", url)
	// выводим сообщение уровня Error со строкой URL, и это SugaredLogger
	sugar.Errorf("Failed to fetch URL: %s", url)

	// переводим в обычный Logger
	plain := sugar.Desugar()

	// выводим сообщение уровня Info обычного регистратора (не SugaredLogger)
	plain.Info("Hello, Go!")
	// также уровня Warn (не SugaredLogger)
	plain.Warn("Simple warning")
	// и уровня Error, но добавляем строго типизированное поле "url" (не SugaredLogger)
	plain.Error("Failed to fetch URL", zap.String("url", url))
}
