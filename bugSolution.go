package main

import (

        "fmt"

        "os"

)

func main() {

        file, err := os.Open("my_file.txt")

        if err != nil {

                fmt.Println("Error opening file:", err)

                return

        }

        defer file.Close()

        // Process the file here... 

        fmt.Println("File processed successfully.")

} 