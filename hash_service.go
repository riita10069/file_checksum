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
	var allWg = sync.WaitGroup{}
	var smallWg = sync.WaitGroup{}

	allWg.Add(len(lines))
	hashedLines := make([]string, len(lines))
	for i, v := range lines {
		smallWg.Add(1)
		go func() {
			num := i
			text := v
			smallWg.Done()
			hashedLines[num] = NewHashDomain(text).HexDumpBySHA256()
			allWg.Done()
		}()
		smallWg.Wait()
	}
	allWg.Wait()
	return hashedLines
}
