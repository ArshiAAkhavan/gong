package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	names := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter name: ")

		scanner.Scan()

		text := scanner.Text()

		if len(text) != 0 {

			fmt.Println(text)
			names = append(names, text)
		} else {
			break
		}
	}

	fmt.Println(names)
}

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"syscall"
// )

// var (
// 	dll            = syscall.MustLoadDLL("kernel32")
// 	setConsoleMode = dll.MustFindProc("SetConsoleMode")
// )

// func SetInputConsoleMode(h syscall.Handle, m uint32) error {
// 	r, _, err := setConsoleMode.Call(uintptr(h), uintptr(m))
// 	if r == 0 {
// 		return err
// 	}
// 	return nil
// }

// func main() {
// 	syscall.Must
// 	h := syscall.Handle(os.Stdin.Fd())
// 	var m uint32
// 	if err := syscall.GetConsoleMode(h, &m); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := SetInputConsoleMode(h, 0); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer SetInputConsoleMode(h, m)

// 	fmt.Printf("press any key to exit ...")

// 	b := make([]byte, 10)
// 	if _, err := os.Stdin.Read(b); err != nil {
// 		log.Fatal(err)
// 	}
// }

/*func  main() {

	// char, _, err := keyboard.GetSingleKey()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("You pressed: %q\r\n", char)
	//--------------------------------------------------------------------------------------------------------------------
	// // disable input buffering
	// exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// // // do not display entered characters on the screen
	// // exec.Command("stty", "-F", "/dev/tty", "-echo").Rtun()

	// var b []byte = make([]byte, 1)
	// for {
	// 	os.Stdin.Read(b)
	// 	fmt.Println("I got the byte", b, "("+string(b)+")")
	// }
}
*/
