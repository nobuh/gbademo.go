// cp437
package main

import (
	"runtime/volatile"
	"unsafe"

	"gbademo.go/demos/cp437/font"
)

const (
	LCD_CTRL     = 0x04000000
	LCD_BG0      = LCD_CTRL + 8
	BG_PALETTE   = 0x05000000
	VRAM         = 0x06000000
	LCD_BG0EN    = 0x01
	LCD_MODE0    = 0x00
	LCD_WIDTH    = 240
	LCD_HEIGHT   = 160
	LCD_VHEIGHT  = 32
	LCD_VWIDTH   = 32
	LCD_SIZE3232 = 0x0000
	LCD_COLOR256 = 0x0080
	NUM_COLOR    = 8
)

func BGR(r uint16, g uint16, b uint16) uint16 {
	return b<<10 + g<<5 + r
}

func VRAM_TILE(blockNum uint32) uint32 {
	// head address of the 16KB tile block.
	return VRAM + blockNum*0x4000
}

func VRAM_MAP(blockNum uint32) uint32 {
	// head address of the 2KB map block.
	return VRAM + blockNum*0x0800
}

var (
	bgColor         = (*[NUM_COLOR + 1]volatile.Register16)(unsafe.Pointer(uintptr(BG_PALETTE))) // +1 for skipping [0]
	tileData        = (*[8092]volatile.Register16)(unsafe.Pointer(uintptr(VRAM_TILE(0))))        // 16KB
	mapData         = (*[1024]volatile.Register16)(unsafe.Pointer(uintptr(VRAM_MAP(28))))        // 2KB
	displayRegister = (*[2]volatile.Register8)(unsafe.Pointer(uintptr(LCD_CTRL)))
	bgRegister      = (*[2]volatile.Register8)(unsafe.Pointer(uintptr(LCD_BG0)))
)

func main() {

	// initialize palette
	bgColor[1].Set(BGR(0, 0, 0))    // Black
	bgColor[2].Set(BGR(31, 31, 31)) // White

	// initialize tiles
	for i := 0; i < 256; i++ {
		for row := 0; row < 8; row++ {
			for col := 7; col >= 0; col-- {
				var val, bit uint16
				if font.Char8x8[i][row]&(1<<col) > 0 {
					bit = 2
				} else {
					bit = 1
				}
				if col%2 > 0 {
					val = uint16(bit)
				} else {
					tileData[(i*LCD_VWIDTH)+(row*4)+((7-col)/2)].Set(val + uint16(bit<<8))
				}
			}
		}
	}

	// initialize display
	bgRegister[0].Set(0x80) // set 256 color and the tile block index is 0
	bgRegister[1].Set(0x1C) // the map block index is 28 and set 32x32 screen
	displayRegister[0].Set(LCD_MODE0)
	displayRegister[1].Set(LCD_BG0EN)

	// draw the screen
	var i, j uint16
	for i = 0; i < 16; i++ {
		for j = 0; j < 16; j++ {
			mapData[(i+2)*LCD_VWIDTH+(j+7)].Set(i*16 + j)
		}
	}

	// for ever
	for {
	}
}