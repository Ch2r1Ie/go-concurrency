# ***Semaphore Concurrency Control in Go***

This is a simple Go program that demonstrates how to use a semaphore to limit the number of concurrent goroutines. It processes a set of tasks concurrently while ensuring that only a limited number of tasks are processed at any given time.

***

## **🔹 Key Features**

✅ **Semaphore Control:** Limits the number of concurrent goroutines to avoid overwhelming the system.  
✅ **WaitGroup:** Ensures that the main function waits for all tasks to complete before exiting.  
✅ **Simulated Task Processing:** Tasks are simulated with a sleep function to represent time-consuming work.  

***

## **🛠 Requirements**

- Go 1.18 or higher  

***

## **📥 Installation**

### **1️⃣ Clone the repository:**  
```bash
git clone https://github.com/yourusername/semaphore-concurrency-example.git
cd semaphore-concurrency-example
```
### **2️⃣ Build and run the program:**  
```bash
go run main.go
```

***

## **📌 Program Breakdown**

### 🔹 **Semaphore (`semaphore := make(chan struct{}, 3)`)**  
A channel is used as a semaphore to limit the number of concurrent goroutines. In this example, the capacity of the semaphore is set to **3**, meaning a maximum of **3 goroutines** can run concurrently.  

### 🔹 **WaitGroup (`var wg sync.WaitGroup`)**  
A `sync.WaitGroup` ensures that the main goroutine waits for all other goroutines to complete before it exits.  

- `wg.Add(1)`: Increments the counter for each goroutine.  
- `wg.Done()`: Decrements the counter once the goroutine completes.  
- `wg.Wait()`: Blocks execution until all goroutines finish.  

### 🔹 **Goroutines**  
For each task, a goroutine is launched, and the semaphore is used to **acquire and release** slots.  

🚀 If more than 3 tasks are running, the extra tasks **wait** until a slot is freed up.  

### 🔹 **Task Processing**  
The tasks are simulated using `time.Sleep(2 * time.Second)` to mimic time-consuming operations (such as database queries or file processing).  

***

## **📊 Example Output**  

```plaintext
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
```
