package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect to a server and send a message",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide a message to send")
		}
		message := args[0]
		Client(message)
	},
}

func init() {
	clientCmd.Flags().IntVarP(&port, "port", "p", 1234, "clientPort to connect to")
	rootCmd.AddCommand(clientCmd)
}

func Client(message string) {
	log.Printf("Connecting to server on port %d\n", port)

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Printf("Connected to server at %s", conn.RemoteAddr())

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent: %s", message)

	// Optionally, read the response from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}
	log.Printf("%s\n", string(buffer[:n]))
	os.Exit(0)
	conn.Close()
}
