UNAME := $(shell uname -s)
BINNAME := spg

ifeq ("$(UNAME)", "windows32")
	BINNAME := $(BINNAME).exe
endif

.PHONY: run build install clean

run:
	go run main.go

build:
	go build main.go -o $(BINNAME)

install:
	go install .

clean:
	rm spg.exe
