GO111MODULE=off

all: rgb.bin 
	VisualBoyAdvance -1 rgb.bin
rgb.bin: main.go 
	tinygo build -target gameboy-advance -o rgb.bin main.go
# build shows the warning "tinygo:ld.lld: warning: lld uses blx instruction, no object with architecture supporting feature detected"
# hoever, it seems no problem. 
clean:
	rm *.bin
