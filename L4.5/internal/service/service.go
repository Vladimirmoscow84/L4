package service

import (
	"sort"
	"time"

	"L4.5/internal/model"
)

type Statistic struct{}

func New() *Statistic {
	return &Statistic{}
}

func (s *Statistic) GetStats(nums model.Numbers) *model.Response {
	start := time.Now()

	count := len(nums.Data)

	var sum float64
	for _, v := range nums.Data {
		sum += v
	}

	avg := 0.0
	if count > 0 {
		avg = sum / float64(count)
	}

	sorted := make([]float64, count)
	copy(sorted, nums.Data)
	//До оптимизации
	//sortedNums := quickSort(sorted)
	sort.Float64s(sorted)

	//До оптимизации
	// median := 0.0
	// if count > 0 {
	// 	mid := count / 2
	// 	if count%2 == 0 {
	// 		median = (sortedNums[mid-1] + sortedNums[mid]) / 2
	// 	} else {
	// 		median = sortedNums[mid]
	// 	}
	// }

	median := 0.0
	if count > 0 {
		mid := count / 2
		if count%2 == 0 {
			median = (sorted[mid-1] + sorted[mid]) / 2
		} else {
			median = sorted[mid]
		}
	}

	_ = time.Since(start).Milliseconds()

	return &model.Response{
		Sum:    sum,
		Avg:    avg,
		Median: median,
		//Sorted: sortedNums,  - до потимизации
		Sorted: sorted,
		Count:  count,
	}

}

// quickSort - функция быстрой сортировки слайса до оптимизации
func quickSort(slice []float64) []float64 {
	if len(slice) < 2 {
		return slice
	}
	pivot := slice[0]
	var lower, greater []float64
	for _, v := range slice[1:] {
		if v <= pivot {
			lower = append(lower, v)
		} else {
			greater = append(greater, v)
		}
	}
	answer := append(quickSort(lower), pivot)
	answer = append(answer, quickSort(greater)...)
	return answer
}
