package ydbase

type Student struct {
	XH, Name string
	Age      int
	Height   float32
}

var StudentTable = []Student{}

func init() {
	StudentTable = make([]Student, 0, 100)
}
