package tinycompiler

import (
	"fmt"
	"os"
)

type EmitterIface interface {
	//produce the equivalent C words
	Emit(string)
	//produce the whole C code line
	EmitLine(string)
	//produce the headers to the C code e.g for adding necesseary header files
	HeaderLine(string)
	//write to the output file
	WriteToFile()
}

var _ EmitterIface = (*Emitter)(nil)

type Emitter struct {
	fullPath string //path to the generate C code file
	header   string //prepending header info such as #include to the C code
	code     string //equivalent C code
}

func (e *Emitter) Emit(code string) {
	e.code += code
}

func (e *Emitter) EmitLine(code string) {
	e.code += code + "\n"
}

func (e *Emitter) HeaderLine(code string) {
	e.header += code + "\n"
}

func (e *Emitter) WriteToFile() {
	fd, _ := os.OpenFile(e.fullPath, os.O_CREATE|os.O_WRONLY, 0644)
	defer fd.Close()

	_, err := fd.WriteString(e.header + e.code)
	if err != nil {
		panic(fmt.Errorf("error in writing to the file %w", err))
	}
}
