package main

import "sync"

type (
	HashService  struct{}
	IHashService interface {
		SHA256([]string) []string
	}
)

func NewHashService() *HashService {
	return &HashService{}
}

func (*HashService) SHA256(lines []string) []string {
	var wg = sync.WaitGroup{}

	wg.Add(len(lines))
	hashedLines := make([]string, len(lines))
	for i, v := range lines {
		go func(i int, v string) {
			hashedLines[i] = NewHashDomain(v).HexDumpBySHA256()
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	return hashedLines
}
