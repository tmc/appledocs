// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AccessoryView() View
	AddAccessoryController(accessoryController unsafe.Pointer)
	BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer)
	BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	DefaultButtonTitle() foundation.String
	FinalWritePrintInfo()
	RemoveAccessoryController(accessoryController unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	SetAccessoryView(accessoryView IView)
	SetDefaultButtonTitle(defaultButtonTitle string)
	UpdateFromPrintInfo()
	AccessoryControllers() []ViewController
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value IHelpAnchorName)
	JobStyleHint() PrintPanelJobStyleHint
	SetJobStyleHint(value IPrintPanelJobStyleHint)
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	PrintInfo() NSPrintInfo
}

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



// Returns a new print panel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printPanel
func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.class), objc.Sel("printPanel"))
	return rv
}


// Returns the accessory view of the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/accessoryView
func (p_ PrintPanel) AccessoryView() View {
	rv := objc.Send[View](p_.ID, objc.Sel("accessoryView"))
	return rv
}


// Adds a custom controller to the Print panel to manage an accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/addAccessoryController(_:)
func (p_ PrintPanel) AddAccessoryController(accessoryController unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAccessoryController:"), accessoryController)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/beginSheet(using:on:completionHandler:)
func (p_ PrintPanel) BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginSheetUsingPrintInfo:onWindow:completionHandler:"), printInfo, parentWindow, handler)
}


// Displays a Print panel sheet and runs it modally for the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/beginSheet(with:modalFor:delegate:didEnd:contextInfo:)
func (p_ PrintPanel) BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:"), printInfo, docWindow, delegate, didEndSelector, contextInfo)
}


// Returns the title of the Print panel’s default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/defaultButtonTitle()
func (p_ PrintPanel) DefaultButtonTitle() foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("defaultButtonTitle"))
	return rv
}


// Writes the Print panel’s printing attributes to the current print operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/finalWritePrintInfo
func (p_ PrintPanel) FinalWritePrintInfo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("finalWritePrintInfo"))
}


// Removes the specified controller and accessory view from the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/removeAccessoryController(_:)
func (p_ PrintPanel) RemoveAccessoryController(accessoryController unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAccessoryController:"), accessoryController)
}


// Displays the Print panel and begins the modal loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal()
func (p_ PrintPanel) RunModal() int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModal"))
	return rv
}


// Displays the Print panel and runs the modal loop using the specified printing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal(with:)
func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}


// Sets the accessory view for the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/setAccessoryView:
func (p_ PrintPanel) SetAccessoryView(accessoryView IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAccessoryView:"), accessoryView)
}


// Sets the title of the Print panel’s default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/setDefaultButtonTitle(_:)
func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultButtonTitle:"), objc.String(defaultButtonTitle))
}


// Updates the Print panel with information from the current print operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/updateFromPrintInfo
func (p_ PrintPanel) UpdateFromPrintInfo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateFromPrintInfo"))
}


// The array of controller objects that manage the Print panel’s accessory views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/accessoryControllers
func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.Send[[]ViewController](p_.ID, objc.Sel("accessoryControllers"))
	return rv
}


// The HTML help anchor associated with the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/helpAnchor
func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := objc.Send[HelpAnchorName](p_.ID, objc.Sel("helpAnchor"))
	return rv
}


// The HTML help anchor associated with the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/helpAnchor
func (p_ PrintPanel) SetHelpAnchor(value IHelpAnchorName) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHelpAnchor:"), value)
}


// The type of settings that the print panel displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := objc.Send[PrintPanelJobStyleHint](p_.ID, objc.Sel("jobStyleHint"))
	return rv
}


// The type of settings that the print panel displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p_ PrintPanel) SetJobStyleHint(value IPrintPanelJobStyleHint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobStyleHint:"), value)
}


// The current configuration options for the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/options-swift.property
func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.Send[PrintPanelOptions](p_.ID, objc.Sel("options"))
	return rv
}


// The current configuration options for the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/options-swift.property
func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}


// The information associated with the running Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printInfo
func (p_ PrintPanel) PrintInfo() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}



