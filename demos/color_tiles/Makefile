GO111MODULE=off

all: demo.bin 
	VisualBoyAdvance -1 demo.bin
demo.bin: main.go 
	tinygo build -target gameboy-advance -o demo.bin main.go
# build shows the warning "tinygo:ld.lld: warning: lld uses blx instruction, no object with architecture supporting feature detected"
# hoever, it seems no problem. 
clean:
	rm *.bin
