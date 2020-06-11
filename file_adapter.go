package main

import (
	"bufio"
	"os"
)

type FileAdapter struct {
	filename string
	lines    []string
}

func NewFileAdapter(filename string) *FileAdapter {
	return &FileAdapter{filename: filename}
}

func (a *FileAdapter) fileScan(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a.lines = append(a.lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (a *FileAdapter) ReadByLine() error {
	file, err := os.Open(a.filename)
	defer file.Close()
	if err != nil {
		return err
	}
	if err = a.fileScan(file); err != nil {
		return err
	}
	return nil
}

func (a *FileAdapter) writelines(file *os.File) error {
	for _, line := range a.lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *FileAdapter) WriteByLine() error {
	file, err := os.Create(a.filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = a.writelines(file); err != nil {
		return err
	}
	return nil
}
