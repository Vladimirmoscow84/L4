/*Утилита анализа GC и памяти (runtime, профилирование)
Необходимо разработать программу на Go, которая показывает через HTTP-endpoint в формате Prometheus текущую информацию о памяти и сборщике мусора.

Используйте runtime.ReadMemStats, debug.SetGCPercent, профилирование (pprof).

Примеры метрик:

количество аллокаций

количество сборок мусора

используемая память

последнее время GC

другие — по вашему желанию

Результат: директория с кодом сервера, инструкцией по запуску (README), примерами запросов.*/

package main

import (
	"net/http"
	"runtime"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

//Определение метрик для сбора информации

var (
	//heapAlloc - метрика текущего объема памяти в куче
	heapAlloc = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "memory_alloc_bytes_in_heap",
			Help: "текущий объем памяти в куче(байт)",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.HeapAlloc)
		},
	)

	//totalAlloc - метрика общего объема памяти всех аллокаций
	totalAlloc = prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "total_memory_alloc",
			Help: "общий объем всех аллокаций(байт)",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.TotalAlloc)
		},
	)

	//heapSys - метрика полученной памяти из системы
	heapSys = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "memory_sys_bytes_in_heap",
			Help: "полученная память из системы(байт)",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.HeapSys)
		},
	)

	//malloc - метрика общего колличества аллокаций
	mallocs = prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "memory_total_mallocs",
			Help: "общее количество аллокаций",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.Mallocs)
		},
	)

	//frees - метрика количества освобождений памяти
	frees = prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "total_memory_frees",
			Help: "количество освобождений памяти",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.Frees)
		},
	)

	// numGC - метрика количества циклов работы сборщика мусора
	numGC = prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "num_gc",
			Help: "количество циклов сборщика мусора",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.NumGC)
		},
	)

	//lastGC - метрика времени последней сборки мусора
	lastGCTime = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "last_gc_time",
			Help: "время последней сборки мусора(секунды)",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.LastGC) / 1e9
		},
	)

	// pauseTotal - метрика общего времени всех пауз сборщика мусора
	pauseTotal = prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "pause_total",
			Help: "общее время пауз GC (наносекунды)",
		},
		func() float64 {
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			return float64(ms.PauseTotalNs)
		},
	)

	//goroutines - метрика количества активных горутин
	numGoroutines = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "num_goroutines",
			Help: "количество активных горутин",
		},
		func() float64 {
			return float64(runtime.NumGoroutine())
		},
	)
)

// регистрация метрик в prometheos
func init() {
	prometheus.MustRegister(
		heapAlloc,
		heapSys,
		totalAlloc,
		mallocs,
		frees,
		numGC,
		lastGCTime,
		pauseTotal,
		numGoroutines,
	)
}

// gcHandler - ручка для управления GC
func gcHandler(c *gin.Context) {
	set := c.Query("set")
	if set == "" {
		current := debug.SetGCPercent(-1)
		_ = debug.SetGCPercent(current)

		c.string(http.StatusOK, "Значение CG в текущий момент = %d", current)
		return
	}

	val, err := strconv.Atoi(set)
	if err != nil {
		c.string(http.StatusBadRequest, "Некоректное значение параметра 'set'")
		return
	}

	previos := debug.SetGCPercent(val)
	c.string(http.StatusOK, "Установлено новое значение GOGC = %d\n Предыдущее значение GOGC = %d", val, previos)
}

func main() {}
