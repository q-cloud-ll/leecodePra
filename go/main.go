package main

import "fmt"

func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//ch1 := make(chan bool)
	//ch2 := make(chan bool)
	//
	//go func() {
	//	for i := 1; i <= 100; i += 2 {
	//		<-ch1
	//		fmt.Println(i)
	//		ch2 <- true
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for i := 2; i <= 100; i += 2 {
	//		<-ch2
	//		fmt.Println(i)
	//		ch1 <- true
	//	}
	//	wg.Done()
	//}()
	//
	//ch1 <- true
	//
	//wg.Wait()
	// 创建一个父Context
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//// 启动一个goroutine，并将ctx作为参数传递
	//go func(ctx context.Context) {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			// 检测到父Context已取消，退出goroutine
	//			fmt.Println("child goroutine exit")
	//			return
	//		default:
	//			// 执行一些任务
	//			fmt.Println("child goroutine running")
	//			time.Sleep(time.Second)
	//		}
	//	}
	//}(ctx)
	//
	//// 等待一段时间后取消父Context
	//time.Sleep(5 * time.Second)
	//cancel()
	//
	//// 等待一段时间后退出程序
	//time.Sleep(2 * time.Second)
	//fmt.Println("main exit")
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- i
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}
