package pkg

import (
	"fmt"
	"io"
	"os"
)

//Println line with \n
func Println(val interface{}) {

	io.WriteString(os.Stdout, fmt.Sprintf("%v\n", val))
}

//Print line
func Print(val interface{}) {

	io.WriteString(os.Stdout, fmt.Sprintf("%v", val))
}
