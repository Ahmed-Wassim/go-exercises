package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func load(i int) []byte {
	var base = fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)

	resp, err := http.Get(base)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func main() {

	var (
		output     io.WriteCloser = os.Stdout
		err        error
		cnt, fails int
		data       []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails < 2; i++ {
		data = load(i)
		if data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",") // OB1
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))
		if err != nil {
			panic(err)
		}

		fails = 0
		cnt++
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
