package or

import (
	"testing"
	"time"
)

func doSignal(duration time.Duration) <-chan any {
	resChan := make(chan any)
	go func() {
		defer close(resChan)
		time.Sleep(duration)
	}()
	return resChan
}

func TestOr(t *testing.T) {
	begin := time.Now()

	<-Or(
		doSignal(2*time.Second),
		doSignal(1*time.Second),
		doSignal(900*time.Millisecond),
		doSignal(700*time.Millisecond),
		doSignal(4*time.Minute),
		doSignal(3*time.Second),
	)
	workTime := time.Since(begin)
	if workTime > 800*time.Millisecond {
		t.Error("Test not complited")
	}
}
