package main

import "context"

import "fmt"

import "time"

func main() {
	ctx := context.Background()

	sonctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	go work(sonctx, "haha")
	time.Sleep(7 * time.Second)
	cancel()
	for {
	}
}

func work(ctx context.Context, name string) error {
	var i int
	for {
		i++
		fmt.Printf("for[%d]:\n", i)
		select {
		case <-ctx.Done():
			fmt.Println("\tcontext done:", ctx.Err().Error())
			return ctx.Err()
		default:
			ticker := time.NewTicker(3 * time.Second)
			time.Sleep(1 * time.Second)
			go func(ctx context.Context, num int) {
				select {
				case <-ticker.C:
					fmt.Printf("\t\tticker[%d]:%v\n", num, time.Now())
				case <-ctx.Done():
					fmt.Printf("\t\tdone[%d]\n", num)
				}
			}(ctx, i)
		}
	}
}
