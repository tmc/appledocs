// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class QLPreviewPanel */


/* debug [class_header]: Header for QLPreviewPanel */
// The class instance for the [PreviewPanel] class.
var (
	PreviewPanelClass     _PreviewPanelClass
	PreviewPanelClassOnce sync.Once
)

func getPreviewPanelClass() _PreviewPanelClass {
	PreviewPanelClassOnce.Do(func() {
		PreviewPanelClass = _PreviewPanelClass{objc.GetClass("QLPreviewPanel")}
	})
	return PreviewPanelClass
}

type _PreviewPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewPanel */
// An interface definition for the [PreviewPanel] class.
type IPreviewPanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for PreviewPanel */
	// properties:
	CurrentController() objc.ID
	CurrentPreviewItem() unsafe.Pointer
	CurrentPreviewItemIndex() int
	SetCurrentPreviewItemIndex(value int)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DisplayState() objc.ID
	SetDisplayState(value objc.ID)
	InFullScreenMode() bool
	IsInFullScreenMode() bool
	SetIsInFullScreenMode(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewPanel */
	// methods:
	EnterFullScreenModeWithOptions(screen appkit.Screen, options objc.IObject /* cross-framework: NSDictionary */) bool
	ExitFullScreenModeWithOptions(options objc.IObject /* cross-framework: NSDictionary */)
	RefreshCurrentPreviewItem()
	ReloadData()
	UpdateController()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewPanel */
// Alloc allocates a new instance without initialization.
func (pc _PreviewPanelClass) Alloc() PreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewPanelClass) New() PreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewPanel) Init() PreviewPanel {
	rv := objc.Send[PreviewPanel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewPanel) Autorelease() PreviewPanel {
	rv := objc.Send[PreviewPanel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewPanel creates a new PreviewPanel instance.
func NewPreviewPanel() PreviewPanel {
	return getPreviewPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewPanel */
// A class that implements the Quick Look preview panel to display a preview of a list of items.
//
// Every application has a single shared instance of accessible through . The preview panel follows the responder chain and adapts to the first responder willing to control it. A preview panel controller provides the content through methods defined in the protocol. You can’t subclass ; you can, however, customize its behavior using a . See the protocol for the methods to customize a preview panel’s behavior.


// A class that implements the Quick Look preview panel to display a preview of a list of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel
type PreviewPanel struct {
	appkit.Panel
}

// PreviewPanelFrom constructs a [PreviewPanel] from an unsafe.Pointer.
//
// A class that implements the Quick Look preview panel to display a preview of a list of items.
func PreviewPanelFrom(ptr unsafe.Pointer) PreviewPanel {
	return PreviewPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewPanel */

// Returns the shared Quick Look preview panel instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/shared()
func (pc _PreviewPanelClass) SharedPreviewPanel() IPreviewPanel {
	rv := objc.Send[PreviewPanel](objc.ID(pc.class), objc.Sel("sharedPreviewPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedPreviewPanel) */


// Returns a Boolean value that indicates whether the system has created a shared Quick Look preview panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/sharedPreviewPanelExists()
func (pc _PreviewPanelClass) SharedPreviewPanelExists() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("sharedPreviewPanelExists"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedPreviewPanelExists) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewPanel */

// Instructs the panel to enter full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/enterFullScreenMode(_:withOptions:)
func (p_ PreviewPanel) EnterFullScreenModeWithOptions(screen appkit.Screen, options objc.IObject /* cross-framework: NSDictionary */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}/* debug [instance_methods/method]: EnterFullScreenModeWithOptions */


// Instructs the panel to exit full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/exitFullScreenMode(options:)
func (p_ PreviewPanel) ExitFullScreenModeWithOptions(options objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("exitFullScreenModeWithOptions:"), options)
}/* debug [instance_methods/method]: ExitFullScreenModeWithOptions */


// Asks the preview panel to recompute the preview of the current preview item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/refreshCurrentPreviewItem()
func (p_ PreviewPanel) RefreshCurrentPreviewItem() {
	objc.Send[objc.ID](p_.ID, objc.Sel("refreshCurrentPreviewItem"))
}/* debug [instance_methods/method]: RefreshCurrentPreviewItem */


// Asks the preview panel to reload its data from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/reloadData()
func (p_ PreviewPanel) ReloadData() {
	objc.Send[objc.ID](p_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Asks the preview panel to update its current controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/updateController()
func (p_ PreviewPanel) UpdateController() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateController"))
}/* debug [instance_methods/method]: UpdateController */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewPanel */

// The current first responder accepting to control the preview panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/currentController
func (p_ PreviewPanel) CurrentController() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("currentController"))
	return rv
}/* debug [instance_properties/getter]: currentController */


// The currently previewed item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/currentPreviewItem
func (p_ PreviewPanel) CurrentPreviewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentPreviewItem"))
	return rv
}/* debug [instance_properties/getter]: currentPreviewItem */


// The index of the current preview item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/currentPreviewItemIndex
func (p_ PreviewPanel) CurrentPreviewItemIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPreviewItemIndex"))
	return rv
}/* debug [instance_properties/getter]: currentPreviewItemIndex */


// The index of the current preview item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/currentPreviewItemIndex
func (p_ PreviewPanel) SetCurrentPreviewItemIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPreviewItemIndex:"), value)
}/* debug [instance_properties/setter]: currentPreviewItemIndex */


// The preview panel data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/dataSource
func (p_ PreviewPanel) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The preview panel data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/dataSource
func (p_ PreviewPanel) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// The delegate object that controls the preview panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/delegate
func (p_ PreviewPanel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that controls the preview panel’s behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/delegate
func (p_ PreviewPanel) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The preview panel’s display state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/displayState
func (p_ PreviewPanel) DisplayState() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("displayState"))
	return rv
}/* debug [instance_properties/getter]: displayState */


// The preview panel’s display state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/displayState
func (p_ PreviewPanel) SetDisplayState(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayState:"), value)
}/* debug [instance_properties/setter]: displayState */


// The property that indicates whether the panel is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel/isInFullScreenMode
func (p_ PreviewPanel) InFullScreenMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("inFullScreenMode"))
	return rv
}/* debug [instance_properties/getter]: inFullScreenMode */


// The property that indicates whether the panel is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/isinfullscreenmode
func (p_ PreviewPanel) IsInFullScreenMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInFullScreenMode"))
	return rv
}/* debug [instance_properties/getter]: isInFullScreenMode */


// The property that indicates whether the panel is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklookui/qlpreviewpanel/isinfullscreenmode
func (p_ PreviewPanel) SetIsInFullScreenMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInFullScreenMode:"), value)
}/* debug [instance_properties/setter]: isInFullScreenMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewPanel */



