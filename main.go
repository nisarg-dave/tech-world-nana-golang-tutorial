// Main package is special because it indicates that the files belongs to an executable program
package main

// Import the fmt package, which contains functions for formatting and printing text. fmt is part of the standard library.
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// This is a slice of strings. It is a collection of values of the same type. Slice is a dynamic array.
var taskItems = []string{"Buy groceries", "Walk the dog", "Make dinner"}
// This is an array of strings. It is a collection of values of the same type. Array is a static array.
// taskArray := [3]string{"Buy groceries", "Walk the dog", "Make dinner"}
// main is the entry point for the program
// to run the program, use the command "go run main.go"
func main() {
	// This is the short hand for declaring a variable. Long hand is var listHeading string = "List of my Todos" or var listHeading = "List of my Todos"
	listHeading := "List of my Todos"

	// This is a function that handles the root route ("/") and calls the greetUser function which is the handler for the root route
	http.HandleFunc("/", greetUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)



	http.ListenAndServe(":8080", nil)
	fmt.Println(listHeading)
	// Don't need to use pointers here as slices have internal pointers to update the original slice
	// printTasks(taskItems)

}

func greetUser(w http.ResponseWriter, r *http.Request) {
	// Println is a function that prints a string to the console
	fmt.Fprintln(w, "Welcome to our Todolist App!")
}

func showTasks(w http.ResponseWriter, r *http.Request) {

	// this is a for loop that iterates over the taskItems slice and prints each item to the console
	for index, item := range taskItems {
		// %d is a placeholder for a decimal number
		// %s is a placeholder for a string
		// fmt.Printf("%d. %s\n", index+1, item)
		fmt.Fprintf(w, "%d. %s\n", index+1, item)

	}
}


func addTask(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var requestBody struct {
        Task string `json:"task"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    taskItems = append(taskItems, requestBody.Task)
}
