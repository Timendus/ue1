package main

func assert(assertion bool, message string) {
	if !assertion {
		fail(message)
	}
}

func fail(message string) {
	panic(message)
}
