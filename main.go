package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"9fans.net/go/acme"
)

func addr() (int, int, error) {
	winid := os.Getenv("winid")
	id, err := strconv.Atoi(winid)
	if err != nil {
		return 0, 0, errors.New("cannot find Acme window")
	}
	win, err := acme.Open(id, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("cannot open window with id %v: %v\n", id, err)
	}
	defer win.CloseFiles()

	// read to open addr file:
	if _, _, err := win.ReadAddr(); err != nil {
		return 0, 0, fmt.Errorf("cannot open addr file: ", err)
	}
	if err := win.Ctl("addr=dot"); err != nil {
		return 0, 0, fmt.Errorf("invalid ctl write: ", err)
	}
	q0, q1, err := win.ReadAddr()
	if err != nil {
		return 0, 0, fmt.Errorf("invalid addr read: ", err)
	}
	return q0, q1, err
}

func main() {
	q0, q1, err := addr()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("#%d,#%d\n", q0, q1)
	os.Exit(0)
}