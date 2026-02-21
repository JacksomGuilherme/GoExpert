package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Teste")
}

// por convenção toda variável de context sempre deve ir no inicio dos parâmetros
func bookHotel(ctx context.Context, hotelName string) {
	token := ctx.Value("token")
	fmt.Println(token)
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
	case <-time.After(5 * time.Second):
		fmt.Println(fmt.Sprintf("%s - Hotel booked.", hotelName))
	}
}
