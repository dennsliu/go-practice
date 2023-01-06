package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "http://redseanet.com/", nil)
	if err != nil {
		fmt.Printf("http.NewRequest err: %+v", err)
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do err: %+v", err)
		return
	}
	defer resp.Body.Close()
}
