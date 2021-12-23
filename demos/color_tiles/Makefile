GO111MODULE=off

all: main.bin
	mgba -1 main.bin
main.bin: main.go
	tinygo build -opt 2 -target gameboy-advance -o main.bin main.go
clean:
	rm *.bin *.sav
