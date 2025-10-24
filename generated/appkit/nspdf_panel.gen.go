// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPDFPanel */


/* debug [class_header]: Header for NSPDFPanel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFPanel */
// An interface definition for the [PDFPanel] class.
type IPDFPanel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFPanel */
	// properties:
	AccessoryController() IViewController
	SetAccessoryController(value IViewController)
	DefaultFileName() objc.IObject /* cross-framework: NSString */
	SetDefaultFileName(value objc.IObject /* cross-framework: NSString */)
	Options() objectivec.IObject
	SetOptions(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFPanel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFPanel */
// Alloc allocates a new instance without initialization.
func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := objc.Send[PDFPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFPanel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFPanel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFPanel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFPanel */

// A view controller for the accessory view that the panel can present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/accessoryController
func (p_ PDFPanel) AccessoryController() IViewController {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("accessoryController"))
	return rv
}/* debug [instance_properties/getter]: accessoryController */


// A view controller for the accessory view that the panel can present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/accessoryController
func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAccessoryController:"), value)
}/* debug [instance_properties/setter]: accessoryController */


// The initial value for the user-editable filename shown in the name field of the PDF panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/defaultfilename
func (p_ PDFPanel) DefaultFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("defaultFileName"))
	return rv
}/* debug [instance_properties/getter]: defaultFileName */


// The initial value for the user-editable filename shown in the name field of the PDF panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/defaultfilename
func (p_ PDFPanel) SetDefaultFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultFileName:"), value)
}/* debug [instance_properties/setter]: defaultFileName */


// A set of configuration options that determine the accessory views the PDF panel should display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/options-swift.property
func (p_ PDFPanel) Options() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// A set of configuration options that determine the accessory views the PDF panel should display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfpanel/options-swift.property
func (p_ PDFPanel) SetOptions(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPDFPanel */



