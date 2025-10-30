
# GO Утилита анализа GC и памяти (runtime, профилирование)

![Go Version](https://img.shields.io/badge/Go-1.23+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Prometheus](https://img.shields.io/badge/Prometheus-exporter-orange)
![Gin](https://img.shields.io/badge/Gin--framework-Enabled-purple)

Экспортер метрик **памяти и сборщика мусора (GC)** для приложений на Go.  
Программа показывает через HTTP-endpoint данные о состоянии памяти, количестве сборок мусора, паузах GC и других характеристиках рантайма.

Также поддерживает:
- управление параметром **GOGC** через HTTP
- **профилирование pprof**
- **экспорт в Prometheus**
- периодическое логирование состояния памяти

---

## Возможности

- 📊 Метрики памяти (`runtime.ReadMemStats`)
- ⚙️ Управление параметром GC (`debug.SetGCPercent`)
- 🧩 Профилирование через `pprof`
- 🧵 Подсчёт активных горутин
- 🧠 Экспорт в формате **Prometheus**
- 💬 REST API на **Gin**

---

## Конфигурация

Создайте файл `.env` в корне проекта:

```env
ADDR=:9100
GOGC_DEFAULT=100

Установка и запуск
1. Клонировать репозиторий
git clone https://github.com/<your-username>/go-gc-exporter.git
cd go-gc-exporter

2. Установить зависимости
go mod tidy

3. Запуск
go run main.go


Эндпоинты
| Endpoint         | Описание                                       |
| ---------------- | ---------------------------------------------- |
| `/metrics`       | Метрики в формате Prometheus                   |
| `/gc`            | Управление и просмотр значения GOGC            |
| `/debug/pprof/*` | Профилирование (`heap`, `cpu`, `trace` и т.д.) |


Примеры запросов
curl http://localhost:9100/metrics

Проверить текущее значение GOGC
curl http://localhost:9100/gc

Установить новое значение GOGC
curl "http://localhost:9100/gc?set=200"

Пример ответа:
Установлено новое значение GOGC = 200
Предыдущее значение GOGC = 100

Пример метрик Prometheus
# HELP memory_alloc_bytes_in_heap текущий объем памяти в куче(байт)
# TYPE memory_alloc_bytes_in_heap gauge
memory_alloc_bytes_in_heap 512000

# HELP num_gc количество циклов сборщика мусора
# TYPE num_gc counter
num_gc 17

# HELP num_goroutines количество активных горутин
# TYPE num_goroutines gauge
num_goroutines 9

Профилирование (pprof)
Эндпоинты доступны по адресу:
| Тип профиля     | URL                                                                          |
| --------------- | ---------------------------------------------------------------------------- |
| Список профилей | [/debug/pprof/](http://localhost:9100/debug/pprof/)                          |
| CPU-профиль     | [/debug/pprof/profile](http://localhost:9100/debug/pprof/profile?seconds=30) |
| Память (heap)   | [/debug/pprof/heap](http://localhost:9100/debug/pprof/heap)                  |

Пример команды:
go tool pprof http://localhost:9100/debug/pprof/profile?seconds=30

Пример логов
[main] Утилита анализа GC запущена, слушаем на :9100
heap=5200 KB, sys=10240 KB, num_gc=15, goroutines=9
heap=5230 KB, sys=10400 KB, num_gc=16, goroutines=9

Go GC & Memory Exporter
Автор: <Vladimirmoscow84>
📧 Контакт: <ccr1@yandex.ru.  https://github.com/Vladimirmoscow84 >
🌐 GitHub: https://github.com/Vladimirmoscow84/L4/tree/main/L4.4
