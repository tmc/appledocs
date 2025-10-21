// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ContextMenuElementInfo] class.
type IContextMenuElementInfo interface {
	objectivec.IObject
}

// An object that contains information about a link the user clicked in a webpage, and which you use to configure a context menu for that link.
//
// A object contains the URL of a link in the web view’s content. You don’t create instances of this class directly. Instead, the web view creates them and passes them to the methods of its associated object when the user interacts with the link. In your delegate method implementations, use the URL in this object to determine how to configure the contextual menu.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ContextMenuElementInfoClass) Alloc() ContextMenuElementInfo {
	rv := objc.Send[ContextMenuElementInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The URL of the link that the user clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContextMenuElementInfo/linkURL
func (c_ ContextMenuElementInfo) LinkURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("linkURL"))
	return rv
}



