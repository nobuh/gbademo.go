package main

import (
	"machine"
	"runtime/volatile"
	"unsafe"
)

const (
	BG_PALETTE     = 0x05000000
	VRAM           = 0x06000000
	LCD_TILE_HWORD = 8 * 4
)

var (
	display = machine.Display
	_bcolor = (*[9]volatile.Register16)(unsafe.Pointer(uintptr(BG_PALETTE)))
	_tile   = (*[0x200]volatile.Register16)(unsafe.Pointer(uintptr(VRAM_TILE(0))))
	_map    = (*[0x200]volatile.Register16)(unsafe.Pointer(uintptr(VRAM_MAP(28))))
)

func VRAM_TILE(n uint8) uint16 {
	return VRAM + n*0x4000
}

func VRAM_MAP(n uint8) uint16 {
	return VRAM + n*0x0800
}

func BGR(r uint8, g uint8, b uint8) unit16 {
	return b<<10 + g<<5 + r
}

func main() {

	_bcolor[1] = BGR(0, 0, 0)    // black
	_bcolor[2] = BGR(0, 0, 31)   // blue
	_bcolor[3] = BGR(0, 31, 0)   // green
	_bcolor[4] = BGR(0, 31, 31)  // light teal
	_bcolor[5] = BGR(31, 0, 0)   // red
	_bcolor[6] = BGR(31, 0, 31)  // purple
	_bcolor[7] = BGR(31, 31, 0)  // yellow
	_bcolor[8] = BGR(31, 31, 31) // white

	// init tiles with colors
	for i := 0; i < 8; i++ {
		for j := 0; j < LCD_TILE_HWORD; j++ {
			tile[i*LCD_TILE_HWORD+j].Set((i+1)<<8 + (i + 1))
		}
	}
}
