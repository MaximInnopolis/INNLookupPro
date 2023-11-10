package main

import (
	"INNLookupPro/cmd/logger"
	"INNLookupPro/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"strings"
)

func isINN(s string) bool {
	if len(s) == 10 || len(s) == 12 {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	return false
}

func getInput() (string, error) {
	var inn string
	for {
		fmt.Printf("\nEnter INN: ")
		_, err := fmt.Scanln(&inn)

		if err != nil {
			return "", err
		}

		inn = strings.TrimSpace(inn)
		if !isINN(inn) {
			logger.Println("Error: Invalid INN. Please try again.")
			continue
		}

		return inn, nil
	}
}

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		logger.Printf("Client failed to connect to gRPC server: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	client := protos.NewCompanyInfoServiceClient(conn)

	for {
		inn, err := getInput()
		logger.Println("Client entered proper INN")
		if err != nil {
			continue
		}

		response, err := client.GetCompanyInfo(context.Background(), &protos.CompanyInfoRequest{Inn: inn})
		if err != nil {
			log.Fatal(err)
		}

		logger.Println("Client received entered INN:", response.Inn)
		logger.Println("Client received Kpp:", response.Kpp)
		logger.Println("Client received CompanyName:", response.CompanyName)
		logger.Println("Client received DirectorName:", response.DirectorName)
	}

}
