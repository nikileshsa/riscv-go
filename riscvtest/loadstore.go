package main

func main() {
	// check that basic loads and stores work;
	// if they don't, we'll fault before we make it to the exit.
	a := new(int)
	*a = 0
	riscvexit(0)
}
