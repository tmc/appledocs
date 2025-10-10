// main.go - Call Swift CoreGraphics wrapper from Go using purego

package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
)

// CGContext wraps the Swift-managed CoreGraphics context
type CGContext struct {
	ptr uintptr
}

// SwiftCG provides access to Swift CoreGraphics functions
type SwiftCG struct {
	createBitmapContext func(width, height int32) uintptr
	releaseContext      func(ctx uintptr)
	setFillColor        func(ctx uintptr, r, g, b, a float64)
	setStrokeColor      func(ctx uintptr, r, g, b, a float64)
	setLineWidth        func(ctx uintptr, width float64)
	fillRect            func(ctx uintptr, x, y, width, height float64)
	strokeRect          func(ctx uintptr, x, y, width, height float64)
	fillEllipse         func(ctx uintptr, x, y, width, height float64)
	strokeEllipse       func(ctx uintptr, x, y, width, height float64)
	beginPath           func(ctx uintptr)
	moveTo              func(ctx uintptr, x, y float64)
	addLineTo           func(ctx uintptr, x, y float64)
	closePath           func(ctx uintptr)
	strokePath          func(ctx uintptr)
	fillPath            func(ctx uintptr)
	createPNGData       func(ctx uintptr, outData *unsafe.Pointer, outLen *int) int32
	freePNGData         func(data unsafe.Pointer)
}

func NewSwiftCG() (*SwiftCG, error) {
	lib, err := purego.Dlopen("./libCoreGraphicsSwift.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return nil, fmt.Errorf("failed to load library: %w", err)
	}

	cg := &SwiftCG{}
	purego.RegisterLibFunc(&cg.createBitmapContext, lib, "cg_create_bitmap_context")
	purego.RegisterLibFunc(&cg.releaseContext, lib, "cg_release_context")
	purego.RegisterLibFunc(&cg.setFillColor, lib, "cg_set_fill_color")
	purego.RegisterLibFunc(&cg.setStrokeColor, lib, "cg_set_stroke_color")
	purego.RegisterLibFunc(&cg.setLineWidth, lib, "cg_set_line_width")
	purego.RegisterLibFunc(&cg.fillRect, lib, "cg_fill_rect")
	purego.RegisterLibFunc(&cg.strokeRect, lib, "cg_stroke_rect")
	purego.RegisterLibFunc(&cg.fillEllipse, lib, "cg_fill_ellipse")
	purego.RegisterLibFunc(&cg.strokeEllipse, lib, "cg_stroke_ellipse")
	purego.RegisterLibFunc(&cg.beginPath, lib, "cg_begin_path")
	purego.RegisterLibFunc(&cg.moveTo, lib, "cg_move_to")
	purego.RegisterLibFunc(&cg.addLineTo, lib, "cg_add_line_to")
	purego.RegisterLibFunc(&cg.closePath, lib, "cg_close_path")
	purego.RegisterLibFunc(&cg.strokePath, lib, "cg_stroke_path")
	purego.RegisterLibFunc(&cg.fillPath, lib, "cg_fill_path")
	purego.RegisterLibFunc(&cg.createPNGData, lib, "cg_create_png_data")
	purego.RegisterLibFunc(&cg.freePNGData, lib, "cg_free_png_data")

	return cg, nil
}

func (cg *SwiftCG) CreateContext(width, height int32) *CGContext {
	ptr := cg.createBitmapContext(width, height)
	if ptr == 0 {
		return nil
	}
	return &CGContext{ptr: ptr}
}

func (ctx *CGContext) Release(cg *SwiftCG) {
	if ctx.ptr != 0 {
		cg.releaseContext(ctx.ptr)
		ctx.ptr = 0
	}
}

func (ctx *CGContext) SetFillColor(cg *SwiftCG, r, g, b, a float64) {
	cg.setFillColor(ctx.ptr, r, g, b, a)
}

func (ctx *CGContext) SetStrokeColor(cg *SwiftCG, r, g, b, a float64) {
	cg.setStrokeColor(ctx.ptr, r, g, b, a)
}

func (ctx *CGContext) SetLineWidth(cg *SwiftCG, width float64) {
	cg.setLineWidth(ctx.ptr, width)
}

func (ctx *CGContext) FillRect(cg *SwiftCG, x, y, width, height float64) {
	cg.fillRect(ctx.ptr, x, y, width, height)
}

func (ctx *CGContext) StrokeRect(cg *SwiftCG, x, y, width, height float64) {
	cg.strokeRect(ctx.ptr, x, y, width, height)
}

func (ctx *CGContext) FillEllipse(cg *SwiftCG, x, y, width, height float64) {
	cg.fillEllipse(ctx.ptr, x, y, width, height)
}

func (ctx *CGContext) StrokeEllipse(cg *SwiftCG, x, y, width, height float64) {
	cg.strokeEllipse(ctx.ptr, x, y, width, height)
}

func (ctx *CGContext) SavePNG(cg *SwiftCG, filename string) error {
	var dataPtr unsafe.Pointer
	var length int

	result := cg.createPNGData(ctx.ptr, &dataPtr, &length)
	if result < 0 {
		return fmt.Errorf("failed to create PNG data")
	}
	defer cg.freePNGData(dataPtr)

	data := unsafe.Slice((*byte)(dataPtr), length)
	return os.WriteFile(filename, data, 0644)
}

func main() {
	cg, err := NewSwiftCG()
	if err != nil {
		log.Fatalf("Failed to initialize Swift CG: %v", err)
	}

	// Create a 400x400 context
	ctx := cg.CreateContext(400, 400)
	if ctx == nil {
		log.Fatal("Failed to create context")
	}
	defer ctx.Release(cg)

	// Draw a white background
	ctx.SetFillColor(cg, 1.0, 1.0, 1.0, 1.0)
	ctx.FillRect(cg, 0, 0, 400, 400)

	// Draw a blue rectangle
	ctx.SetFillColor(cg, 0.2, 0.4, 0.8, 1.0)
	ctx.FillRect(cg, 50, 50, 100, 100)

	// Draw a red circle with stroke
	ctx.SetFillColor(cg, 0.8, 0.2, 0.2, 1.0)
	ctx.FillEllipse(cg, 200, 50, 150, 150)

	ctx.SetStrokeColor(cg, 0.0, 0.0, 0.0, 1.0)
	ctx.SetLineWidth(cg, 3.0)
	ctx.StrokeEllipse(cg, 200, 50, 150, 150)

	// Draw a green triangle using path
	cg.beginPath(ctx.ptr)
	cg.moveTo(ctx.ptr, 100, 300)
	cg.addLineTo(ctx.ptr, 200, 250)
	cg.addLineTo(ctx.ptr, 150, 350)
	cg.closePath(ctx.ptr)

	ctx.SetFillColor(cg, 0.2, 0.8, 0.2, 1.0)
	cg.fillPath(ctx.ptr)

	// Save to PNG
	if err := ctx.SavePNG(cg, "output.png"); err != nil {
		log.Fatalf("Failed to save PNG: %v", err)
	}

	fmt.Println("Successfully created output.png using Swift+CoreGraphics via purego!")
	fmt.Println("No cgo, no Objective-C runtime - pure Swift → Go interop")
}
