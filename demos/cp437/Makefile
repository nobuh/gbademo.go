export GO111MODULE := off
all: main.go 
	tinygo run main.go > cp437.txt
font.go: cp437-8x8.bdf
	awk -f bdf2go.awk cp437-8x8.bdf > font/cp437.go	
clean:
	rm *.bin *.sav *.txt 
