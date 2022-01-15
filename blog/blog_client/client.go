package main

import (
	"context"
	"fmt"
	"github.com/udodinho/grpc/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer connection.Close()
	c := blogpb.NewBlogServiceClient(connection)

	// create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Udodinho",
		Title:    "My first blog",
		Content:  "Content of my first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5f1f8f8b-f8f8-4f8f-8f8f-f8f8f8f8f8f8"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		log.Fatalf("Error happened while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)
}
