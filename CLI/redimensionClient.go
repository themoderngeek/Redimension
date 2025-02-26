package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var serverAddress string

func main() {
	rootCmd := &cobra.Command{Use: "redimension"}

	setCmd := &cobra.Command{
		Use:   "SET [key] [value]",
		Short: "Set a value in the redimension store",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			value := args[1]
			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			defer conn.Close()
			fmt.Fprintf(conn, "SET %s %s\n", key, value)
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}
			fmt.Print(response)
		},
	}

	getCmd := &cobra.Command{
		Use:   "GET [key]",
		Short: "Get a value from the redimension store",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]
			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			fmt.Fprintf(conn, "GET %s\n", key)
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}
			fmt.Print(response)
		},
	}

	rootCmd.PersistentFlags().StringVar(&serverAddress, "server", "localhost:4000", "Address of the redimension server")
	rootCmd.AddCommand(setCmd, getCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
