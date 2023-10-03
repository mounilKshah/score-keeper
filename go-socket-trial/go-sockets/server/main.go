package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

// ANSI escape codes for text colors
var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
}

var clients = make(map[string]net.Conn) // Map to store connections using UUID
var mutex = &sync.Mutex{}

func main() {
	fmt.Println("Starting", connType, "server on", connHost+":"+connPort)

	/**
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error connecting: %v", err)
		}
		go handleConnection(c)
	}
	*/

	ticker := time.NewTicker(2 * time.Second)

	// Creating channel using make
	tickerChan := make(chan bool)

	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer l.Close()

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatalf("Error connecting: %v", err)
			}
			go handleConnection(c)
		}
		// for {
		// 	select {
		// 	case <-tickerChan:
		// 		return
		// 		// interval task
		// 	case tm := <-ticker.C:
		// 		fmt.Println("The Current time is: ", tm)
		// 		// go intervalTask()

		// 		for {
		// 			c, err := l.Accept()
		// 			if err != nil {
		// 				log.Fatalf("Error connecting: %v", err)
		// 			}
		// 			go handleConnection(c)
		// 		}
		// 	}
		// }
	}()

	// Calling Sleep() method
	time.Sleep(10 * time.Second)

	// Calling Stop() method
	ticker.Stop()

	// Setting the value of channel
	tickerChan <- true

	// Printed when the ticker is turned off
	fmt.Println("Ticker is turned off!")
}

func intervalTask() {
	fmt.Println("This is interval task")
}

func handleConnection(conn net.Conn) {
	// Generate a UUID for the client
	clientUUID := uuid.New().String()
	fmt.Println("handle connection called")

	// Pick a random color and send it to the client
	rand.Seed(time.Now().UnixNano())
	chosenColor := colors[rand.Intn(len(colors))]
	conn.Write([]byte(chosenColor + "\n"))

	fmt.Println("New client connected:", clientUUID)

	mutex.Lock()
	clients[clientUUID] = conn
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, clientUUID)
		mutex.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client", clientUUID, "left.")
			return
		}
		// message = "temp data"
		// Broadcast the message to all clients except the sender
		for id, c := range clients {
			if id != clientUUID {
				c.Write([]byte("[" + clientUUID + "]: " + message))
			}
		}
	}
}
