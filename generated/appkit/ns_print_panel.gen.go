// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPrintPanel */


/* debug [class_header]: Header for NSPrintPanel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PrintPanel */
// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PrintPanel */
	// properties:
	AccessoryControllers() []ViewController
	HelpAnchor() HelpAnchorName /* typedef */
	SetHelpAnchor(value HelpAnchorName /* typedef */)
	JobStyleHint() PrintPanelJobStyleHint /* typedef */
	SetJobStyleHint(value PrintPanelJobStyleHint /* typedef */)
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	PrintInfo() IPrintInfo
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PrintPanel */
	// methods:
	AddAccessoryController(accessoryController unsafe.Pointer)
	BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer)
	DefaultButtonTitle() foundation.String
	RemoveAccessoryController(accessoryController unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	SetDefaultButtonTitle(defaultButtonTitle objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PrintPanel */
// Alloc allocates a new instance without initialization.
func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PrintPanel */
// The Print panel that queries the user for information about a print job.
//
// A Print panel may let the user select the range of pages to print and the number of copies before executing the Print command. Print panels can display a simplified interface when printing certain types of data. For example, the panel can display a list of print-setting presets, which lets the user enable print settings in groups as opposed to individually. Assigning an appropriate string to the property activates the simplified interface and identifies which presets to display. For design guidance, see .


// The Print panel that queries the user for information about a print job.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PrintPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PrintPanel */

// Returns a new print panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printPanel
func (pc _PrintPanelClass) PrintPanel() IPrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.class), objc.Sel("printPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrintPanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PrintPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PrintPanel */

// Adds a custom controller to the Print panel to manage an accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/addAccessoryController(_:)
func (p_ PrintPanel) AddAccessoryController(accessoryController unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAccessoryController:"), accessoryController)
}/* debug [instance_methods/method]: AddAccessoryController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/beginSheet(using:on:completionHandler:)
func (p_ PrintPanel) BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginSheetUsingPrintInfo:onWindow:completionHandler:"), printInfo, parentWindow, handler)
}/* debug [instance_methods/method]: BeginSheetUsingPrintInfoOnWindowCompletionHandler */


// Returns the title of the Print panel’s default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/defaultButtonTitle()
func (p_ PrintPanel) DefaultButtonTitle() foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("defaultButtonTitle"))
	return rv
}/* debug [instance_methods/method]: DefaultButtonTitle */


// Removes the specified controller and accessory view from the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/removeAccessoryController(_:)
func (p_ PrintPanel) RemoveAccessoryController(accessoryController unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAccessoryController:"), accessoryController)
}/* debug [instance_methods/method]: RemoveAccessoryController */


// Displays the Print panel and begins the modal loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal()
func (p_ PrintPanel) RunModal() int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Displays the Print panel and runs the modal loop using the specified printing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal(with:)
func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}/* debug [instance_methods/method]: RunModalWithPrintInfo */


// Sets the title of the Print panel’s default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/setDefaultButtonTitle(_:)
func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultButtonTitle:"), defaultButtonTitle)
}/* debug [instance_methods/method]: SetDefaultButtonTitle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PrintPanel */

// The array of controller objects that manage the Print panel’s accessory views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/accessoryControllers
func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.Send[[]ViewController](p_.ID, objc.Sel("accessoryControllers"))
	return rv
}/* debug [instance_properties/getter]: accessoryControllers */


// The HTML help anchor associated with the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/helpAnchor
func (p_ PrintPanel) HelpAnchor() HelpAnchorName /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("helpAnchor"))
	return rv
}/* debug [instance_properties/getter]: helpAnchor */


// The HTML help anchor associated with the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/helpAnchor
func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHelpAnchor:"), value)
}/* debug [instance_properties/setter]: helpAnchor */


// The type of settings that the print panel displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("jobStyleHint"))
	return rv
}/* debug [instance_properties/getter]: jobStyleHint */


// The type of settings that the print panel displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobStyleHint:"), value)
}/* debug [instance_properties/setter]: jobStyleHint */


// The current configuration options for the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/options-swift.property
func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.Send[PrintPanelOptions](p_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The current configuration options for the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/options-swift.property
func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// The information associated with the running Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printInfo
func (p_ PrintPanel) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}/* debug [instance_properties/getter]: printInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPrintPanel */



