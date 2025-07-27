package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Check the protocol of a response from a server",
	Run: func(cmd *cobra.Command, args []string) {
		checkProtocol(addr, port)
	},
}

func init() {
	protoCmd.Flags().StringVarP(&addr, "addr", "a", "localaddr", "addr to connect to")
	protoCmd.Flags().IntVarP(&port, "port", "p", 1234, "Port to connect to")
	rootCmd.AddCommand(protoCmd)
}

func checkProtocol(addr string, port int) {
	log.Printf("Checking protocol for %s:%d\n", addr, port)

	//  check if connect if a grpc server
	resp, err := grpcCheck(addr, port)
	if err != nil {
		log.Printf("Server is not gRPC: %v", err)
	} else {
		log.Printf("gRPC check response: %s", resp)
		log.Printf("Server is gRPC protocol: %s:%d", addr, port)
		return
	}

	// check if connect if a http server
	resp, err = httpCheck(addr, port)
	if err != nil {
		log.Printf("Server is not HTTP: %v", err)
	} else {
		log.Printf("HTTP check response: %s", resp)
		log.Printf("Server is HTTP protocol: %s:%d", addr, port)
	}
}

func grpcCheck(a string, p int) (string, error) {
	log.Printf("Checking if server is gRPC protocol for %s:%d\n", addr, port)
	var grpcResponse string
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hostStr := fmt.Sprintf("%s:%d", a, p)
	conn, err := grpc.DialContext(ctx, hostStr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	grpcResponse = "✅ Seems to be a gRPC server"
	return grpcResponse, nil
}

func httpCheck(a string, p int) (string, error) {
	log.Printf("Checking if server is HTTP protocol for %s:%d\n", addr, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", a, p))
	if err != nil {
		return "", fmt.Errorf("failed to connect to %s:%d: %v", a, p, err)
	}
	defer conn.Close()
	// Send a simple HTTP request
	request := "GET / HTTP/1.1 Host: " + a + "\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}
	return string(buffer[:n]), nil
}
