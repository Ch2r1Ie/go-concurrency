package gosemaphore

import (
	"fmt"
	"sync"
)

func semaphore() {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5)

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(taskID int) {

			semaphore <- struct{}{}

			defer func() { <-semaphore }()

			{
				fmt.Printf("Processing task %d\n", taskID)

				fmt.Printf("Completed task %d\n", taskID)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}

func Init() {

	{
		semaphore()
	}

}
