all:	build-mod run	





build-mod:	go.mod
	go mod tidy
run:	main.go 
	go run main.go