package main

import "sync"

type (
	HashService struct {}
	IHashService interface {
		SHA256([]string) []string
	}
)

func NewHashService() *HashService {
	return &HashService{}
}

func (*HashService) SHA256(lines []string) []string {
	var wg = sync.WaitGroup{}
	hashedLines := make([]string, len(lines))
	for i, v := range lines {
		wg.Add(1)
		go func() {
			hashedLines[i] = NewHashDomain(v).HexDumpBySHA256()
			wg.Done()
		}()
	}
	wg.Wait()
	return hashedLines
}

