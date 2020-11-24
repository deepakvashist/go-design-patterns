package main

import (
	"fmt"

	"github.com/deepakvashist/go-design-patterns/builder/clients"
)

func main() {
	builder := &clients.AwsBuilder{}
	director := &clients.Director{}

	director.SetBuilder(builder)

	awsClient := director.BuildClient()

	fmt.Println(fmt.Sprintf("Client version is: %s", awsClient.GetVersion()))
	fmt.Println(fmt.Sprintf("Client base URL is: %s", awsClient.GetBaseURL()))
	fmt.Println(fmt.Sprintf("Client max retries is: %d", awsClient.GetMaxRetries()))
	fmt.Println(fmt.Sprintf("Client allowed status codes is: %v", awsClient.GetAllowedStatusCodes()))
}
