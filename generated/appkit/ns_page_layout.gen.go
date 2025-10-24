// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPageLayout */


/* debug [class_header]: Header for NSPageLayout */
// The class instance for the [PageLayout] class.
var (
	PageLayoutClass     _PageLayoutClass
	PageLayoutClassOnce sync.Once
)

func getPageLayoutClass() _PageLayoutClass {
	PageLayoutClassOnce.Do(func() {
		PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}
	})
	return PageLayoutClass
}

type _PageLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PageLayout */
// An interface definition for the [PageLayout] class.
type IPageLayout interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PageLayout */
	// properties:
	AccessoryControllers() []ViewController
	PrintInfo() IPrintInfo
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PageLayout */
	// methods:
	AddAccessoryController(accessoryController IViewController)
	BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer)
	RemoveAccessoryController(accessoryController IViewController)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PageLayout */
// Alloc allocates a new instance without initialization.
func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PageLayout) Init() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PageLayout) Autorelease() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPageLayout creates a new PageLayout instance.
func NewPageLayout() PageLayout {
	return getPageLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PageLayout */
// A panel that queries the user for information such as paper type and orientation.
//
// A page layout panel is typically displayed in response to the user selecting the Page Setup menu item. You obtain an instance with the class method. The pane can then be run as a sheet using or modally using or . For design guidance, see .


// A panel that queries the user for information such as paper type and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout
type PageLayout struct {
	objectivec.Object
}

// PageLayoutFrom constructs a [PageLayout] from an unsafe.Pointer.
//
// A panel that queries the user for information such as paper type and orientation.
func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PageLayout *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PageLayout */

// Returns a newly created page layout object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/pageLayout
func (pc _PageLayoutClass) PageLayout() IPageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("pageLayout"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PageLayout) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PageLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PageLayout */

// Adds the specified controller of an accessory view to be presented in the page setup panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/addAccessoryController(_:)
func (p_ PageLayout) AddAccessoryController(accessoryController IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAccessoryController:"), accessoryController)
}/* debug [instance_methods/method]: AddAccessoryController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/beginSheet(using:on:completionHandler:)
func (p_ PageLayout) BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo IPrintInfo, parentWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginSheetUsingPrintInfo:onWindow:completionHandler:"), printInfo, parentWindow, handler)
}/* debug [instance_methods/method]: BeginSheetUsingPrintInfoOnWindowCompletionHandler */


// Removes the specified controller of an accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/removeAccessoryController(_:)
func (p_ PageLayout) RemoveAccessoryController(accessoryController IViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAccessoryController:"), accessoryController)
}/* debug [instance_methods/method]: RemoveAccessoryController */


// Displays the page layout panel and begins the modal loop using the shared print info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/runModal()
func (p_ PageLayout) RunModal() int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Displays the page layout panel and begins the modal loop using the specified print info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/runModal(with:)
func (p_ PageLayout) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.Send[int](p_.ID, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}/* debug [instance_methods/method]: RunModalWithPrintInfo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PageLayout */

// An array of accessory view controllers belonging to the page layout panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/accessoryControllers
func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := objc.Send[[]ViewController](p_.ID, objc.Sel("accessoryControllers"))
	return rv
}/* debug [instance_properties/getter]: accessoryControllers */


// The printing information object used when the page layout panel is run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/printInfo
func (p_ PageLayout) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}/* debug [instance_properties/getter]: printInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPageLayout */



