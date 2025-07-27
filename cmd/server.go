package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server to listen for incoming connections and messages",
	Run: func(cmd *cobra.Command, args []string) {
		Server()
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 1234, "Port to listen on")
	rootCmd.AddCommand(serverCmd)
}

func Server() {
	log.Printf("Server is running on port %d\n", port)

	connection, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer connection.Close()

	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		msgReader(conn)
		conn.Close()
	}
}

func msgReader(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}
	if n == 0 {
		log.Printf("Connection closed by client: %s", conn.RemoteAddr())
		return
	}
	log.Printf("Received message: %s", string(buffer[:n]))
	// send a response back to the client
	response := fmt.Sprintf("Server connection successful! Received message: %s", string(buffer[:n]))
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
