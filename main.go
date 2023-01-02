package main 

import "fmt"

func main() {
    var conferenceName string = "GopherCon"
    const conferenceTickets = 50
    var remainingTickets = conferenceTickets

    fmt.Println("Welcome to " + conferenceName)
    fmt.Println("There are ", remainingTickets, " tickets remaining")
    fmt.Println("Get your tickets here to attend")

     
    var userName string

    fmt.Println("Please enter your name")
    fmt.Scanln(&userName)
}

