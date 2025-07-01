package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple To-Do List")

	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %s, proceed as given!\n\n", name)

	fmt.Println("Please input the number of tasks you want to store:")
	var list int
	fmt.Scanf("%d", &list)
	reader.ReadString('\n') 

	
	tasks := make([]string, 0, list)

	
	for i := 0; i < list; i++ {
		fmt.Printf("Input the task number %d:\n", i+1)
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		tasks = append(tasks, task)
	}

	
	fmt.Println("\nYour To-Do List:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
