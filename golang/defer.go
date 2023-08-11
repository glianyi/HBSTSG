package golang

func Call() {
	println("kaishi")
	defer println("print")
	panic(nil)
	// defer println("not print")
}

func LIFO() {
	defer println("B")
	defer println("A")
}

func ParamDefer() {
	count := 10
	defer println("B:", count)
	count += 1
	defer println("A:", count)

}
