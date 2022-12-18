package main

import (
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := userpb.NewCreditCardServiceClient(conn)
	u := userpb.NewUserServiceClient(conn)

	//user client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	personName := "Tom Muller"

	x, err := u.GetUserByName(ctx, &userpb.User{
		Name: personName,
	})
	if err != nil {
		log.Fatalf("could not get person data: %v", err)
	}
	//credit card client
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCreditCardByUserName(ctx2, &userpb.CreditCard{
		Name: x.GetName(),
	})
	if err != nil {
		log.Fatalf("could not get person data: %v", err)
	}
	fmt.Println(colorCyan, r)
}
