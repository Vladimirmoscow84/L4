package or

import (
	"fmt"
	"time"
)

func OrExample() {
	signal := func(task string, duration time.Duration) <-chan any {
		doneCh := make(chan any)
		go func() {
			defer close(doneCh)
			time.Sleep(duration)
			fmt.Println(task, "is done")
		}()
		return doneCh
	}

	begin := time.Now()
	<-Or(signal("First Task", 1*time.Second), signal("Second Task", 2*time.Second))

	fmt.Println(time.Since(begin))
}
