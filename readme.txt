Semaphore Concurrency Control in Go
This is a simple Go program that demonstrates how to use a semaphore to limit the number of concurrent goroutines. It processes a set of tasks concurrently, but ensures that only a limited number of tasks are processed at any given time.

Key Features:
Semaphore Control: Limits the number of concurrent goroutines to avoid overwhelming the system.
WaitGroup: Ensures that the main function waits for all tasks to complete before exiting.
Simulated Task Processing: Tasks are simulated with a sleep function to represent time-consuming work.
Requirements:
Go 1.18 or higher
Installation:
Clone the repository to your local machine:

bash
Copy
Edit
git clone https://github.com/yourusername/semaphore-concurrency-example.git
cd semaphore-concurrency-example
Build and run the program:

bash
Copy
Edit
go run main.go
Program Breakdown:
Semaphore (semaphore := make(chan struct{}, 3)): A channel is used as a semaphore to limit the number of concurrent goroutines. In this example, the capacity of the semaphore is set to 3, which means a maximum of 3 goroutines can run concurrently.

WaitGroup (var wg sync.WaitGroup): A sync.WaitGroup is used to ensure that the main goroutine waits for all other goroutines to complete before it exits. The wg.Add(1) increments the counter for each goroutine, and wg.Done() decrements the counter once the goroutine completes.

Goroutines: For each task, a goroutine is launched, and the semaphore is used to acquire and release slots. This ensures that only 3 tasks can run at once. If there are more than 3 tasks, the other tasks will wait until a slot is freed up.

Task Processing: The tasks are simulated using time.Sleep(2 * time.Second) to mimic time-consuming operations (such as database queries or file processing).

Example Output:
css
Copy
Edit
Processing task 1
Processing task 2
Processing task 3
Completed task 1
Processing task 4
Completed task 2
Processing task 5
Completed task 3
Completed task 4
Processing task 6
Completed task 5
Processing task 7
Completed task 6
Processing task 8
Completed task 7
Processing task 9
Completed task 8
Processing task 10
Completed task 9
Completed task 10
All tasks completed.
How It Works:
Start Goroutines: A loop starts 10 tasks, each in its own goroutine.
Semaphore Controls Concurrency: The semaphore ensures that only 3 tasks are processed at once. If more than 3 tasks are launched, the remaining tasks wait until one of the 3 active goroutines finishes.
Wait for Completion: The sync.WaitGroup ensures that the main goroutine does not exit until all tasks are finished.
Customize the Program:
Concurrency Limit: Change the 3 in make(chan struct{}, 3) to any other number to adjust the concurrency limit.
Task Simulation: Modify the time.Sleep(2 * time.Second) line to simulate different lengths of task processing time.
Task Count: Change the loop in main() to process more or fewer tasks.


