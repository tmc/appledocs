// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFPanel] class.
var (
	pDFPanelClass     _PDFPanelClass
	pDFPanelClassOnce sync.Once
)

func getPDFPanelClass() _PDFPanelClass {
	pDFPanelClassOnce.Do(func() {
		pDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}
	})
	return pDFPanelClass
}

type _PDFPanelClass struct {
	class objc.Class
}

// An interface definition for the [PDFPanel] class.
type IPDFPanel interface {
	objectivec.IObject
}

// A Save or Export as PDF panel that’s consistent with the macOS user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel

type PDFPanel struct {
	objectivec.Object
}

// PDFPanelFrom constructs a [PDFPanel] from an unsafe.Pointer.
//
// A Save or Export as PDF panel that’s consistent with the macOS user interface.
func PDFPanelFrom(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := objc.Send[PDFPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFPanelClass) New() PDFPanel {
	rv := objc.Send[PDFPanel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFPanel) Init() PDFPanel {
	rv := objc.Send[PDFPanel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFPanel) Autorelease() PDFPanel {
	rv := objc.Send[PDFPanel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFPanel creates a new PDFPanel instance.
func NewPDFPanel() PDFPanel {
	return getPDFPanelClass().New()
}




