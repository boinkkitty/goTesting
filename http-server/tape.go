package main

import "io"

type tape struct {
	file io.ReadWriteSeeker
}
