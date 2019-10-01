package handlers

import (
	"fmt"
	"io"
	"log"
)

// ResponseError writes error response to writer
func ResponseError(w io.Writer, status int, message string) {
	_, err := fmt.Fprintf(w, `
{
	"code": %d,
	"message": "%s"
}`, status, message)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
