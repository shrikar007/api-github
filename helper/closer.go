package helper

import (
	"io"
	"log"
)

func Close(close io.Closer) {
	err := close.Close()
	if err != nil {
		log.Println("error when trying to close resource:", err.Error())
	}
}
