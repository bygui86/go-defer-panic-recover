package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	goodErr := good()
	if goodErr != nil {
		fmt.Println("good:", goodErr.Error())
	}

	betterErr := better()
	if betterErr != nil {
		fmt.Println("better:", betterErr.Error())
	}

	bestErr := best()
	if bestErr != nil {
		fmt.Println("best:", bestErr.Error())
	}
}

func good() error {
	f, err := os.Open("./defer-error/book.txt")
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			// do something like log
		}
	}()

	// some other code

	return nil
}

func better() (err error) {
	f, err := os.Open("./defer-error/book.txt")
	if err != nil {
		return err
	}
	defer func() {
		ferr := f.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	// some other code

	return nil
}

func best() (err error) {
	f, err := os.Open("./defer-error/book.txt")
	if err != nil {
		return err
	}
	defer func() {
		ferr := f.Close()
		if ferr != nil {
			err = errors.Wrap(err, ferr.Error())
		}
	}()

	// some other code

	return nil
}
