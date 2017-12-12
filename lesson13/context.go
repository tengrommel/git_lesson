package main

import (
	"context"
	"time"
)

func callRedis(ctx context.Context)  {

}

func callMysql(ctx context.Context)  {

}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	go callRedis(ctx)
	go callMysql(ctx)
	user := make(chan bool)
	select{
	case <-ctx.Done():
	case <-user:
		cancel()
	}
}
