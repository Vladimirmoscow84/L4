.PHONY: run build clean

run:
	@echo "🚀 Запуск GC Exporter..."
	go run main.go

build:
	@echo "🔧 Сборка бинарника..."
	go build -o gc-exporter main.go

clean:
	@echo "🧹 Очистка..."
	rm -f gc-exporter
