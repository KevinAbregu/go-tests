package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1700)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, _ := io.ReadAll(res.Body)

	fmt.Printf("download image of size %v\n", len(imageData))

}
