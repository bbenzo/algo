DIR=./...

all: test bench

test:
	go test $(DIR)

bench:
	go test -bench=. -benchmem $(DIR)
