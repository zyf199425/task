package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var num int = 10
	add(&num)
	fmt.Println(num)

	nums := []int{1, 2, 3, 4, 5}
	mutiply(&nums)
	fmt.Println(nums)

	printEvenAndOdd()
	tasks := []func(){
		func() {
			// 统计1 - 100 的奇数个数
			var count int
			for i := 1; i <= 1000; i++ {
				if i%2 != 0 {
					count++
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			}
			fmt.Println("奇数个数：", count)
		},
		func() {
			// 统计1 - 100 的偶数个数
			var count int
			for i := 1; i <= 1000; i++ {
				if i%2 == 0 {
					count++
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			}
			fmt.Println("偶数个数：", count)
		},
	}
	taskScheduler(tasks)

	r := Rectangle{}
	shap(&r)
	c := Circle{}
	shap(&c)

	e := Employee{Person: Person{Name: "张三", Age: 20}, EmployeeID: 1001}
	e.PrintInfo()

	chanTest()
	chanTest2()

	lockTest1()
	atomicTest()
}

// 指针题目1 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值
func add(num *int) {
	*num += 10
}

// 指针题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func mutiply(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// Goroutine 题目1 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printEvenAndOdd() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("printOdd", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("printEven", i)
			}
		}
	}()
	wg.Wait()
}

// Goroutine题目2 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func taskScheduler(tasks []func()) {
	if len(tasks) == 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i, task := range tasks {
		go func() {
			defer wg.Done()
			startTime := time.Now()
			task()
			elapsed := time.Since(startTime)
			fmt.Println("task:", i, "耗费时间：", elapsed)
		}()
	}
	wg.Wait()
}

// 面向对象 题目1 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle Area()")
}
func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter()")
}

type Circle struct {
}

func (c *Circle) Area() {
	fmt.Println("Circle Area()")
}
func (c *Circle) Perimeter() {
	fmt.Println("Circle Perimeter()")
}

func shap(s Shape) {
	s.Area()
	s.Perimeter()
}

// 面向对象 题目2 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}

// Channel 题目1 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func chanTest() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendOnly(ch, 10)
	}()
	go func() {
		defer wg.Done()
		receiveOnly(ch)
	}()
	wg.Wait()
}

func sendOnly(ch chan<- int, end int) {
	defer close(ch)
	for i := 1; i <= end; i++ {
		ch <- i
		fmt.Println("发送数据：", i)
	}
}
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println("接收数据：", v)
	}
}

// Channel 题目2 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func chanTest2() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendOnly(ch, 100)
	}()
	go func() {
		defer wg.Done()
		receiveOnly(ch)
	}()
	wg.Wait()
}

// 锁机制 题目1 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func lockTest1() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.add()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.getCounter())
}

type Counter struct {
	Mut   sync.Mutex
	Count int
}

func (c *Counter) add() {
	c.Mut.Lock()
	c.Count++
	c.Mut.Unlock()
}

func (c *Counter) getCounter() int {
	return c.Count
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atomicTest() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", count)
}

type AtomicCounter struct {
	Count atomic.Int64
}
