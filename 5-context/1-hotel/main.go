package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	//  O select tem quase a mesma função do switch, mas ele espera uma variavel comparavel com os cases chegarem pra começar a operação
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel Booked")
	}
}
