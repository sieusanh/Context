package main

import (
	"context"
	"fmt"
	"time"
)

/*
func main() {
	//context.TODO() is used when we are not sure what context type we will apply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	
	// cancel after 1 second
	// time.AfterFunc(time.Second, func() {
	// 	cancel()
	// })

	defer cancel()

	// cancel after 3 seconds
	select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
	}
}
*/

/*
func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)
	doSomething(ctx)
}
 
func doSomething(ctx context.Context) {

	canceledChannel := make(chan bool)

	// goroutine for running async with time.Sleep() function below
	go func() {
		select {
			case <- ctx.Done():
				fmt.Println(ctx.Err())
				canceledChannel <- true
				return
		}
	}()

	isCanceledChannel := <- canceledChannel
	if isCanceledChannel {
		close(canceledChannel)
		return
	}

	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
*/

/*
func main() {
	ctx := context.WithValue(context.Background(), "number", 10)
	A(ctx)
}

func A(ctx context.Context) {
	if value := ctx.Value("number"); value != nil { // if value exists
		ctx := context.WithValue(ctx, "number1", 10 / 2)
		B(ctx)
	}
}

func B(ctx context.Context) {
	value1 := ctx.Value("number").(int)
	value2 := ctx.Value("number1").(int)
	fmt.Println(value1 + value2)
}
*/

/*
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select{
		case <- ctx.Done():
			fmt.Println("Done roi ne!")
	}
}
*/

// Context tree
func main() {
	parentCtx, cancel := context.WithCancel(context.Background())

	timeoutCtx, _ := context.WithTimeout(parentCtx, time.Second * 10)

	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select {
		case <- timeoutCtx.Done():
			fmt.Println("timeout")
	}
}