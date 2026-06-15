package function

func WithReturn(x int, y int) int {
	return x + y
}

func WithNamedReturn(x int, y int) (result int) {
	result = x + y
	return
}

func WithMultipleNamedReturn(x int, y int) (result int, msg string) {
	result = x + y
	msg = "Sucesso na operação"
	return
}

func WithMultipleReturn(x int, y int) (int, string) {
	return x + y, "Operação bem sucedida"
}
