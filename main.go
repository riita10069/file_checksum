package main

type (
	handler struct {
		HashService IHashService
	}
	IHandler interface {
		handler() error
	}
)

func newHandler(hashService IHashService) *handler {
	return &handler{HashService: hashService}
}

func main() {
	err := newHandler(NewHashService()).handler()
	if err != nil {
		panic(err.Error())
	}
}

func (h *handler) handler() error {
	fileReader := NewFileAdapter("read.txt")
	fileWriter := NewFileAdapter("write.txt")

	err := fileReader.ReadByLine()
	if err != nil {
		return err
	}

	fileWriter.lines = h.HashService.SHA256(fileReader.lines)
	err = fileWriter.WriteByLine()
	if err != nil {
		return err
	}

	return nil
}
