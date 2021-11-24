package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	/*pid, _, _ := syscall.Syscall(39,0,0,0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24,0,0,0)
	fmt.Println("My uid is", uid)
	message := []byte{'h','e','l','l','o'}
	fd := 1
	syscall.Write(syscall.Handle(fd), message)*/
	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}