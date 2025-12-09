package service

import (
	"testing"

	"L4.5/internal/model"
)

func BenchGetStatsSmall(b *testing.B) {
	nums := model.Numbers{Data: []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}}
	stat := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stat.GetStats(nums)
	}
}

func BenchGetStatsBig(b *testing.B) {
	data := make([]float64, 250000)
	for i := range data {
		data[i] = float64(i%100) + float64(i/23)
	}
	nums := model.Numbers{Data: data}
	stat := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stat.GetStats(nums)
	}

}
