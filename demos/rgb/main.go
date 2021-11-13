// rgb
// rgb bar with GBA's Mode 3 and Back Ground 2 (BG2)

package main

import (
	"runtime/volatile"
	"unsafe"
)

const (
	VRAM       = 0x06000000
	LCD_CTRL   = 0x04000000
	LCD_BG2EN  = 0x04
	LCD_MODE3  = 0x03
	LCD_WIDTH  = 240
	LCD_HEIGHT = 160
)

var displayRegister = (*[2]volatile.Register8)(unsafe.Pointer(uintptr(LCD_CTRL)))

var vram = (*[LCD_HEIGHT * LCD_WIDTH]volatile.Register16)(unsafe.Pointer(uintptr(VRAM)))

func BGR(r uint16, g uint16, b uint16) uint16 {
	return b<<10 + g<<5 + r
}

func main() {

	// init display with Mode3 and BG2 Enabled
	displayRegister[0].Set(LCD_MODE3)
	displayRegister[1].Set(LCD_BG2EN)

	// Draw RGB Bar
	ptr := 0
	for i := 0; i < (LCD_HEIGHT / 3); i++ {
		for j := 0; j < LCD_WIDTH; j++ {
			vram[ptr].Set(BGR(31, 0, 0)) // Red
			ptr++
		}
	}
	for i := 0; i < (LCD_HEIGHT / 3); i++ {
		for j := 0; j < LCD_WIDTH; j++ {
			vram[ptr].Set(BGR(0, 31, 0)) // Green
			ptr++
		}
	}
	for i := 0; i < (LCD_HEIGHT / 3); i++ {
		for j := 0; j < LCD_WIDTH; j++ {
			vram[ptr].Set(BGR(0, 0, 31)) // Blue
			ptr++
		}
	}

	for {
		// forever
	}
}
