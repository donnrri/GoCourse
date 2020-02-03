package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
Part of interfaces course

*/

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Errpr : ", err)
		os.Exit(1)
	}

	/*
	 old way of declaing a byteslice bs := []byte{}
	 instead use the make function ,
	 here you can specify the size also.
	*/
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
