// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFPanel] class.
var (
	PDFPanelClass     _PDFPanelClass
	PDFPanelClassOnce sync.Once
)

func getPDFPanelClass() _PDFPanelClass {
	PDFPanelClassOnce.Do(func() {
		PDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}
	})
	return PDFPanelClass
}

type _PDFPanelClass struct {
	class objc.Class
}

// An interface definition for the [PDFPanel] class.
type IPDFPanel interface {
	objectivec.IObject
	// properties:
	AccessoryController() IViewController
	SetAccessoryController(value IViewController)
	DefaultFileName() objc.IObject /* cross-framework: NSString */
	SetDefaultFileName(value objc.IObject /* cross-framework: NSString */)
	Options() unsafe.Pointer
	SetOptions(value unsafe.Pointer)
	// methods:
}

// A Save or Export as PDF panel that’s consistent with the macOS user interface.
//
// A PDF panel has a variety of built-in customization controls, such as page orientation, paper size, and tags. It also supports the use of a custom accessory view controller that allows an app to specify how a PDF file should be created.


// A Save or Export as PDF panel that’s consistent with the macOS user interface.
//
// [Full Topic]
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



// A view controller for the accessory view that the panel can present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/accessoryController
func (p_ PDFPanel) AccessoryController() IViewController {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("accessoryController"))
	return rv
}


// A view controller for the accessory view that the panel can present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/accessoryController
func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAccessoryController:"), value)
}


// The initial value for the user-editable filename shown in the name field of the PDF panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/defaultfilename
func (p_ PDFPanel) DefaultFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("defaultFileName"))
	return rv
}


// The initial value for the user-editable filename shown in the name field of the PDF panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/defaultfilename
func (p_ PDFPanel) SetDefaultFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultFileName:"), value)
}


// A set of configuration options that determine the accessory views the PDF panel should display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/options-swift.property
func (p_ PDFPanel) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("options"))
	return rv
}


// A set of configuration options that determine the accessory views the PDF panel should display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/options-swift.property
func (p_ PDFPanel) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}



