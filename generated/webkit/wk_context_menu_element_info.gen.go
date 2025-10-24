// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKContextMenuElementInfo */

/* debug [class_header]: Header for WKContextMenuElementInfo */
// The class instance for the [ContextMenuElementInfo] class.
var (
	ContextMenuElementInfoClass     _ContextMenuElementInfoClass
	ContextMenuElementInfoClassOnce sync.Once
)

func getContextMenuElementInfoClass() _ContextMenuElementInfoClass {
	ContextMenuElementInfoClassOnce.Do(func() {
		ContextMenuElementInfoClass = _ContextMenuElementInfoClass{objc.GetClass("WKContextMenuElementInfo")}
	})
	return ContextMenuElementInfoClass
}

type _ContextMenuElementInfoClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for ContextMenuElementInfo */
// An interface definition for the [ContextMenuElementInfo] class.
type IContextMenuElementInfo interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for ContextMenuElementInfo */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for ContextMenuElementInfo */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for ContextMenuElementInfo */
// Alloc allocates a new instance without initialization.
func (cc _ContextMenuElementInfoClass) Alloc() ContextMenuElementInfo {
	rv := objc.Send[ContextMenuElementInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContextMenuElementInfoClass) New() ContextMenuElementInfo {
	rv := objc.Send[ContextMenuElementInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContextMenuElementInfo) Init() ContextMenuElementInfo {
	rv := objc.Send[ContextMenuElementInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContextMenuElementInfo) Autorelease() ContextMenuElementInfo {
	rv := objc.Send[ContextMenuElementInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContextMenuElementInfo creates a new ContextMenuElementInfo instance.
func NewContextMenuElementInfo() ContextMenuElementInfo {
	return getContextMenuElementInfoClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for ContextMenuElementInfo */
// An object that contains information about a link the user clicked in a webpage, and which you use to configure a context menu for that link.
//
// A object contains the URL of a link in the web view’s content. You don’t create instances of this class directly. Instead, the web view creates them and passes them to the methods of its associated object when the user interacts with the link. In your delegate method implementations, use the URL in this object to determine how to configure the contextual menu.

// An object that contains information about a link the user clicked in a webpage, and which you use to configure a context menu for that link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContextMenuElementInfo
type ContextMenuElementInfo struct {
	objectivec.Object
}

// ContextMenuElementInfoFrom constructs a [ContextMenuElementInfo] from an unsafe.Pointer.
//
// An object that contains information about a link the user clicked in a webpage, and which you use to configure a context menu for that link.
func ContextMenuElementInfoFrom(ptr unsafe.Pointer) ContextMenuElementInfo {
	return ContextMenuElementInfo{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for ContextMenuElementInfo */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for ContextMenuElementInfo */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for ContextMenuElementInfo */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for ContextMenuElementInfo */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for ContextMenuElementInfo */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKContextMenuElementInfo */
