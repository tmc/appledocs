// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKFilterBrowserPanel */


/* debug [class_header]: Header for IKFilterBrowserPanel */
// The class instance for the [IKFilterBrowserPanel] class.
var (
	IKFilterBrowserPanelClass     _IKFilterBrowserPanelClass
	IKFilterBrowserPanelClassOnce sync.Once
)

func getIKFilterBrowserPanelClass() _IKFilterBrowserPanelClass {
	IKFilterBrowserPanelClassOnce.Do(func() {
		IKFilterBrowserPanelClass = _IKFilterBrowserPanelClass{objc.GetClass("IKFilterBrowserPanel")}
	})
	return IKFilterBrowserPanelClass
}

type _IKFilterBrowserPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKFilterBrowserPanel */
// An interface definition for the [IKFilterBrowserPanel] class.
type IIKFilterBrowserPanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for IKFilterBrowserPanel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKFilterBrowserPanel */
	// methods:
	BeginWithOptionsModelessDelegateDidEndSelectorContextInfo(inOptions objc.IObject /* cross-framework: NSDictionary */, modelessDelegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject)
	BeginSheetWithOptionsModalForWindowModalDelegateDidEndSelectorContextInfo(inOptions objc.IObject /* cross-framework: NSDictionary */, docWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject)
	FilterBrowserViewWithOptions(inOptions objc.IObject /* cross-framework: NSDictionary */) IKFilterBrowserView
	FilterName() foundation.String
	Finish(sender objc.IObject)
	RunModalWithOptions(inOptions objc.IObject /* cross-framework: NSDictionary */) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKFilterBrowserPanel */
// Alloc allocates a new instance without initialization.
func (ic _IKFilterBrowserPanelClass) Alloc() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKFilterBrowserPanelClass) New() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKFilterBrowserPanel) Init() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKFilterBrowserPanel) Autorelease() IKFilterBrowserPanel {
	rv := objc.Send[IKFilterBrowserPanel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKFilterBrowserPanel creates a new IKFilterBrowserPanel instance.
func NewIKFilterBrowserPanel() IKFilterBrowserPanel {
	return getIKFilterBrowserPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKFilterBrowserPanel */
// Presents a user interface for browsing filters.
//
// The class provides a user interface that allows users to browse Core Image filters ( ), to preview a filter, and to get additional information about the filter, such as its description. An object can be displayed as: a separate panel, that is, a utility window that floats on top of document windows a modal dialog a sheet, that is, a dialog that is attached to its parent window and must be dismissed by the user a view that an application can insert into a custom user interface An object can be configured through a style mask to use either the default or brushed metal look for windows. The size and number of visible controls are specified through an options dictionary. An object communicates selection changes through notifications. The class allows the user to create filter collections that are stored with the key in the property list located in .


// Presents a user interface for browsing filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel
type IKFilterBrowserPanel struct {
	appkit.Panel
}

// IKFilterBrowserPanelFrom constructs a [IKFilterBrowserPanel] from an unsafe.Pointer.
//
// Presents a user interface for browsing filters.
func IKFilterBrowserPanelFrom(ptr unsafe.Pointer) IKFilterBrowserPanel {
	return IKFilterBrowserPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKFilterBrowserPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKFilterBrowserPanel */

// Creates a shared instance of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/filterBrowserPanel(withStyleMask:)
func (ic _IKFilterBrowserPanelClass) FilterBrowserPanelWithStyleMask(styleMask objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("filterBrowserPanelWithStyleMask:"), styleMask)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FilterBrowserPanelWithStyleMask) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKFilterBrowserPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKFilterBrowserPanel */

// Displays the filter browser in a new utility window, unless the filter browser is already open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/begin(options:modelessDelegate:didEnd:contextInfo:)
func (i_ IKFilterBrowserPanel) BeginWithOptionsModelessDelegateDidEndSelectorContextInfo(inOptions objc.IObject /* cross-framework: NSDictionary */, modelessDelegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("beginWithOptions:modelessDelegate:didEndSelector:contextInfo:"), inOptions, modelessDelegate, didEndSelector, contextInfo)
}/* debug [instance_methods/method]: BeginWithOptionsModelessDelegateDidEndSelectorContextInfo */


// Displays the filter browser in a sheet—that is, a dialog that is attached to its parent window and must be dismissed by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/beginSheet(options:modalFor:modalDelegate:didEnd:contextInfo:)
func (i_ IKFilterBrowserPanel) BeginSheetWithOptionsModalForWindowModalDelegateDidEndSelectorContextInfo(inOptions objc.IObject /* cross-framework: NSDictionary */, docWindow appkit.Window, modalDelegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("beginSheetWithOptions:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), inOptions, docWindow, modalDelegate, didEndSelector, contextInfo)
}/* debug [instance_methods/method]: BeginSheetWithOptionsModalForWindowModalDelegateDidEndSelectorContextInfo */


// Returns a view that contains a filter browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/filterBrowserView(options:)
func (i_ IKFilterBrowserPanel) FilterBrowserViewWithOptions(inOptions objc.IObject /* cross-framework: NSDictionary */) IKFilterBrowserView {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("filterBrowserViewWithOptions:"), inOptions)
	return rv
}/* debug [instance_methods/method]: FilterBrowserViewWithOptions */


// Returns the name of the filter that is currently selected in the filter browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/filterName()
func (i_ IKFilterBrowserPanel) FilterName() foundation.String {
	rv := objc.Send[foundation.String](i_.ID, objc.Sel("filterName"))
	return rv
}/* debug [instance_methods/method]: FilterName */


// Closes a filter browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/finish(_:)
func (i_ IKFilterBrowserPanel) Finish(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("finish:"), sender)
}/* debug [instance_methods/method]: Finish */


// Displays the filter browser in a modal dialog that must be dismissed by the user but that is not attached to a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKFilterBrowserPanel/runModal(options:)
func (i_ IKFilterBrowserPanel) RunModalWithOptions(inOptions objc.IObject /* cross-framework: NSDictionary */) int {
	rv := objc.Send[int](i_.ID, objc.Sel("runModalWithOptions:"), inOptions)
	return rv
}/* debug [instance_methods/method]: RunModalWithOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKFilterBrowserPanel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKFilterBrowserPanel */



