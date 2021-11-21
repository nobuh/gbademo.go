// tile

package main

import (
	"runtime/volatile"
	"unsafe"
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
	bgColor[2].Set(BGR(0, 0, 31))   // Blue
	bgColor[3].Set(BGR(0, 31, 0))   // Green
	bgColor[4].Set(BGR(0, 31, 31))  // Light teal
	bgColor[5].Set(BGR(31, 0, 0))   // Red
	bgColor[6].Set(BGR(31, 0, 31))  // Purple
	bgColor[7].Set(BGR(31, 31, 0))  // Yellow
	bgColor[8].Set(BGR(31, 31, 31)) // White

	// initialize tiles
	var color, i uint16
	for color = 1; color <= NUM_COLOR; color++ {
		for i = 0; i < LCD_VWIDTH; i++ {
			tileData[i+(color-1)*LCD_VWIDTH].Set(color<<8 + color) // upper and lower bytes use the same color
		}
	}

	// initialze map which is the 16x10 blocks part of the virtual screen
	var x, y, z uint16
	z = 0
	for y = 0; y < 10; y++ {
		for x = 0; x < 15; x++ {
			color = ((z + y + x) % NUM_COLOR) + 1
			mapData[(y*LCD_VWIDTH+x)*2].Set(color)
			mapData[(y*LCD_VWIDTH+x)*2+1].Set(color)
			mapData[(y*LCD_VWIDTH+x)*2+LCD_VWIDTH].Set(color)
			mapData[(y*LCD_VWIDTH+x)*2+LCD_VWIDTH+1].Set(color)
		}
	}
	z++

	// initialize display

	bgRegister[0].Set(0x80) // 256 color and tile block index is 0
	bgRegister[1].Set(0x1C) // map block index index is 28 and 32x32 screen
	displayRegister[0].Set(LCD_MODE0)
	displayRegister[1].Set(LCD_BG0EN)

	// draw the screen
	for {
		for y = 0; y < 10; y++ {
			for x = 0; x < 15; x++ {
				color = ((z + y + x) % NUM_COLOR) + 1
				mapData[(y*LCD_VWIDTH+x)*2].Set(color)
				mapData[(y*LCD_VWIDTH+x)*2+1].Set(color)
				mapData[(y*LCD_VWIDTH+x)*2+LCD_VWIDTH].Set(color)
				mapData[(y*LCD_VWIDTH+x)*2+LCD_VWIDTH+1].Set(color)
			}
		}
		if z < 15 {
			z++
		} else {
			z = 0
		}
	}
}
