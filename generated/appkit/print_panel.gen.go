// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintPanel] class.
var (
	printPanelClass     _PrintPanelClass
	printPanelClassOnce sync.Once
)

func getPrintPanelClass() _PrintPanelClass {
	printPanelClassOnce.Do(func() {
		printPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}
	})
	return printPanelClass
}

type _PrintPanelClass struct {
	class objc.Class
}

// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	objectivec.IObject
}

// The Print panel that queries the user for information about a print job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel
type PrintPanel struct {
	objectivec.Object
}

// PrintPanelFrom constructs a [PrintPanel] from an unsafe.Pointer.
//
// The Print panel that queries the user for information about a print job.
func PrintPanelFrom(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.Send[PrintPanel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrintPanel) Autorelease() PrintPanel {
	rv := objc.Send[PrintPanel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrintPanel creates a new PrintPanel instance.
func NewPrintPanel() PrintPanel {
	return getPrintPanelClass().New()
}




