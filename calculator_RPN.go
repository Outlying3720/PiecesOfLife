package main

import "os"
import "bufio"
import "fmt"
import "strconv"
import "stack/stack"

func main() {
	calc1 := stack.MyStack{}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the RPN formula(enter 'q' to exit): ")
	
	for {
		token, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Something wrong.")
			os.Exit(1)
		}
		token = token[:len(token) - 1] // remove "\r\n"
		// fmt.Printf("--%s--\n",token)  // debug statement
		switch {
			case token == "q": // stop als invoer = "q"
				fmt.Println("Calculator stopped")
				return
			case token >= "0" && token <= "999999":
				i, _ := strconv.Atoi(token)
				calc1.Push(i)
			case token == "+":
				q, _ := calc1.Pop()
				p, _ := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
			case token == "-":
				q, _ := calc1.Pop()
				p, _ := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))
	
			case token == "*":
				q, _ := calc1.Pop()
				p, _ := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))
	
			case token == "/":
				q, _ := calc1.Pop()
				p, _ := calc1.Pop()
				fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
			default:
				fmt.Println("No valid input")
		}
	}
}