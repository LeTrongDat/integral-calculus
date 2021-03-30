package main 

import (
	"itmo/invoker"
)


func main() {
	invoker := &invoker.InvokerImpl{}
	invoker.Invoke()
}
