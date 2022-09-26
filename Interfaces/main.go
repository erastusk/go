// What happens if you wanted a receiver function to take in any type but do the same things within the logic
// You wouldn't write a new function with different receiver type, Interfaces attempt to address this.
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// Problem
// func main() {
// 	eng := englishBot{}
// 	sp := spanishBot{}
// 	eng.printGreeting()
// 	sp.printGreeting()

// }

// func (e englishBot) getGreeting() string {
// 	return "how are You"

// }
// func (s spanishBot) getGreeting() string {
// 	return "Hola"
// }

// func printGreeting(e englishBot) {
// 	fmt.Println(e)
// }

// func printGreeting(s spanishBot) {
// 	fmt.Println(s)
// }

type bot interface {
	//Any type associated to getGreeting, that type is now also associated to any function associated to bot. Essentiall those "other" types
	//become members of type bot and hence have access to any function associated to bot.
	//Note: only types calling getGreeting AND returning string, could be more.
	getGreeting() string
}
func main() {
	eng := englishBot{}
	sp := spanishBot{}
	
	printGreeting(eng)
	printGreeting(sp)

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (e englishBot) getGreeting() string {
	return "how are You"

}
func (s spanishBot) getGreeting() string {
	return "Hola"
}
