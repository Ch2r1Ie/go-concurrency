##1. Why Use sync.WaitGroup?
The sync.WaitGroup is used to wait for a collection of goroutines to finish executing before the main function exits. Without it, the main function might terminate before all goroutines have completed their tasks, potentially causing incomplete execution.

###Explanation:

A goroutine is a lightweight thread of execution. When you launch multiple goroutines, the main function could finish executing before all goroutines complete, especially if they are running concurrently.
sync.WaitGroup helps track the number of goroutines that need to finish:
Add(n): Increments the counter by n for the number of goroutines being launched.
Done(): Decrements the counter for each goroutine when it completes.
Wait(): Blocks the main function (or any function) until the counter reaches zero, ensuring all tasks are finished.
In your example, wg.Add(1) increments the counter for each goroutine. After the task completes, wg.Done() is called to decrement the counter. The wg.Wait() at the end makes sure the main function waits until all goroutines are done before exiting.

##2. Why Use Concurrency Control (Semaphore)?
Concurrency control (using a semaphore in this case) is essential when you want to limit the number of concurrent goroutines accessing shared resources or performing tasks that might be too expensive or risky to execute simultaneously.

###Explanation:

In Go, when you launch a goroutine, it runs concurrently with other goroutines. While Go is designed to manage thousands (or even millions) of goroutines efficiently, there are times when you don’t want all goroutines to run at once, especially if they involve external resources like APIs, databases, or file systems.
A semaphore is a concurrency pattern used to limit the number of concurrent operations. In your code, the semaphore is implemented using a channel (make(chan struct{}, 3)), which can hold up to 3 items at a time.
Each time a goroutine starts, it tries to "acquire" a spot in the semaphore by sending an empty struct (semaphore <- struct{}{}). If the semaphore is full (i.e., 3 goroutines are already running), the goroutine will block until a spot is available.
Once the task is done, the goroutine "releases" the semaphore by removing an item (<-semaphore), making room for another goroutine to start.
Why Semaphore is Useful:

It limits the number of concurrent operations, which is crucial when interacting with systems that have limitations, such as APIs with rate limits or databases with connection constraints.
It prevents overwhelming external resources or the system itself by controlling load and ensuring efficient resource usage.
##3. Combined Effect of WaitGroup and Semaphore:
The WaitGroup ensures that the main function doesn't exit prematurely, waiting for all the goroutines to finish.
The semaphore controls how many goroutines can run concurrently, ensuring that only a limited number of tasks are processed at the same time.
Together, these mechanisms allow you to safely and efficiently handle concurrent tasks in Go, making sure that:

No tasks are left unfinished when the program ends.
External systems aren’t overloaded by too many concurrent requests.
