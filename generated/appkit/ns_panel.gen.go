// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPanel */


/* debug [class_header]: Header for NSPanel */
// The class instance for the [Panel] class.
var (
	PanelClass     _PanelClass
	PanelClassOnce sync.Once
)

func getPanelClass() _PanelClass {
	PanelClassOnce.Do(func() {
		PanelClass = _PanelClass{objc.GetClass("NSPanel")}
	})
	return PanelClass
}

type _PanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Panel */
// An interface definition for the [Panel] class.
type IPanel interface {
	IWindow
	
/* debug [class_interface_properties]: Properties for Panel */
	// properties:
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
	FloatingPanel() bool
	SetFloatingPanel(value bool)
	WorksWhenModal() bool
	SetWorksWhenModal(value bool)
	IsFloatingPanel() bool
	SetIsFloatingPanel(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Panel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Panel */
// Alloc allocates a new instance without initialization.
func (pc _PanelClass) Alloc() Panel {
	rv := objc.Send[Panel](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PanelClass) New() Panel {
	rv := objc.Send[Panel](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Panel) Init() Panel {
	rv := objc.Send[Panel](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Panel) Autorelease() Panel {
	rv := objc.Send[Panel](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPanel creates a new Panel instance.
func NewPanel() Panel {
	return getPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Panel */
// A special kind of window that typically performs a function that is auxiliary to the main window.
//
// For details about how panels work (especially to find out how their behavior differs from window behavior), see .


// A special kind of window that typically performs a function that is auxiliary to the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel
type Panel struct {
	Window
}

// PanelFrom constructs a [Panel] from an unsafe.Pointer.
//
// A special kind of window that typically performs a function that is auxiliary to the main window.
func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: WindowFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Panel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Panel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Panel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Panel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Panel */

// A Boolean value that indicates whether the receiver becomes the key window only when needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("becomesKeyOnlyIfNeeded"))
	return rv
}/* debug [instance_properties/getter]: becomesKeyOnlyIfNeeded */


// A Boolean value that indicates whether the receiver becomes the key window only when needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBecomesKeyOnlyIfNeeded:"), value)
}/* debug [instance_properties/setter]: becomesKeyOnlyIfNeeded */


// A Boolean value that indicates whether the receiver is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) FloatingPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("floatingPanel"))
	return rv
}/* debug [instance_properties/getter]: floatingPanel */


// A Boolean value that indicates whether the receiver is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/isFloatingPanel
func (p_ Panel) SetFloatingPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFloatingPanel:"), value)
}/* debug [instance_properties/setter]: floatingPanel */


// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) WorksWhenModal() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("worksWhenModal"))
	return rv
}/* debug [instance_properties/getter]: worksWhenModal */


// A Boolean value that indicates whether the panel receives keyboard and mouse events even when some other window is being run modally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanel/worksWhenModal
func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWorksWhenModal:"), value)
}/* debug [instance_properties/setter]: worksWhenModal */


// A Boolean value that indicates whether the receiver is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/isfloatingpanel
func (p_ Panel) IsFloatingPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFloatingPanel"))
	return rv
}/* debug [instance_properties/getter]: isFloatingPanel */


// A Boolean value that indicates whether the receiver is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspanel/isfloatingpanel
func (p_ Panel) SetIsFloatingPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFloatingPanel:"), value)
}/* debug [instance_properties/setter]: isFloatingPanel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPanel */



