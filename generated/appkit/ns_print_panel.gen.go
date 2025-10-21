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
	PrintPanelClass     _PrintPanelClass
	PrintPanelClassOnce sync.Once
)

func getPrintPanelClass() _PrintPanelClass {
	PrintPanelClassOnce.Do(func() {
		PrintPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}
	})
	return PrintPanelClass
}

type _PrintPanelClass struct {
	class objc.Class
}

// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	objectivec.IObject
}

// The Print panel that queries the user for information about a print job.
//
// A Print panel may let the user select the range of pages to print and the number of copies before executing the Print command. Print panels can display a simplified interface when printing certain types of data. For example, the panel can display a list of print-setting presets, which lets the user enable print settings in groups as opposed to individually. Assigning an appropriate string to the property activates the simplified interface and identifies which presets to display. For design guidance, see .
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


// The type of settings that the print panel displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) JobStyleHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("jobStyleHint"))
	return rv
}


// SetJobStyleHint sets the value of the jobStyleHint property.
// The type of settings that the print panel displays.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) SetJobStyleHint(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobStyleHint:"), value)
}

// The information associated with the running Print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printInfo
func (p_ PrintPanel) PrintInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("printInfo"))
	return rv
}



