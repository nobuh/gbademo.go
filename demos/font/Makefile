export GO111MODULE := off
all: main.gba
	mgba -2 main.gba
main.gba: main.go font/cp437.go
	tinygo build -opt 2 -target gameboy-advance -o main.gba main.go 
font/cp437.go: cp437-8x8.bdf
	awk -f bdf2go.awk cp437-8x8.bdf > font/cp437.go	
clean:
	rm *.bin *.sav *.gba
