// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariExtensionHandler] class.
var (
	SFSafariExtensionHandlerClass     _SFSafariExtensionHandlerClass
	SFSafariExtensionHandlerClassOnce sync.Once
)

func getSFSafariExtensionHandlerClass() _SFSafariExtensionHandlerClass {
	SFSafariExtensionHandlerClassOnce.Do(func() {
		SFSafariExtensionHandlerClass = _SFSafariExtensionHandlerClass{objc.GetClass("SFSafariExtensionHandler")}
	})
	return SFSafariExtensionHandlerClass
}

type _SFSafariExtensionHandlerClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariExtensionHandler] class.
type ISFSafariExtensionHandler interface {
	objectivec.IObject
	// properties:
	SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A base class that you subclass to handle events in your Safari app extension.


// A base class that you subclass to handle events in your Safari app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionHandler
type SFSafariExtensionHandler struct {
	objectivec.Object
}

// SFSafariExtensionHandlerFrom constructs a [SFSafariExtensionHandler] from an unsafe.Pointer.
//
// A base class that you subclass to handle events in your Safari app extension.
func SFSafariExtensionHandlerFrom(ptr unsafe.Pointer) SFSafariExtensionHandler {
	return SFSafariExtensionHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionHandlerClass) Alloc() SFSafariExtensionHandler {
	rv := objc.Send[SFSafariExtensionHandler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariExtensionHandlerClass) New() SFSafariExtensionHandler {
	rv := objc.Send[SFSafariExtensionHandler](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionHandler) Init() SFSafariExtensionHandler {
	rv := objc.Send[SFSafariExtensionHandler](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionHandler) Autorelease() SFSafariExtensionHandler {
	rv := objc.Send[SFSafariExtensionHandler](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionHandler creates a new SFSafariExtensionHandler instance.
func NewSFSafariExtensionHandler() SFSafariExtensionHandler {
	return getSFSafariExtensionHandlerClass().New()
}



// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariExtensionHandler) SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}



