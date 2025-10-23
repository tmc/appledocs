// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IMKServer] class.
var (
	IMKServerClass     _IMKServerClass
	IMKServerClassOnce sync.Once
)

func getIMKServerClass() _IMKServerClass {
	IMKServerClassOnce.Do(func() {
		IMKServerClass = _IMKServerClass{objc.GetClass("IMKServer")}
	})
	return IMKServerClass
}

type _IMKServerClass struct {
	class objc.Class
}

// An interface definition for the [IMKServer] class.
type IIMKServer interface {
	objectivec.IObject
	Bundle() foundation.Bundle
	LastKeyEventWasDeadKey() bool
	PaletteWillTerminate() bool
}

// The class manages client connections to your input method. When you write the main function for your input method, you create an object. You should never need to override this class.


// The class manages client connections to your input method. When you write the main function for your input method, you create an object. You should never need to override this class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer
type IMKServer struct {
	objectivec.Object
}

// IMKServerFrom constructs a [IMKServer] from an unsafe.Pointer.
//
// The class manages client connections to your input method. When you write the main function for your input method, you create an object. You should never need to override this class.
func IMKServerFrom(ptr unsafe.Pointer) IMKServer {
	return IMKServer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IMKServerClass) Alloc() IMKServer {
	rv := objc.Send[IMKServer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IMKServerClass) New() IMKServer {
	rv := objc.Send[IMKServer](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IMKServer) Init() IMKServer {
	rv := objc.Send[IMKServer](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IMKServer) Autorelease() IMKServer {
	rv := objc.Send[IMKServer](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKServer creates a new IMKServer instance.
func NewIMKServer() IMKServer {
	return getIMKServerClass().New()
}



// Creates and returns a server object from property list information contained in the provided bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:bundleIdentifier:)
func NewIMKServerWithNameBundleIdentifier(name string, bundleIdentifier string) IMKServer {
	instance := getIMKServerClass().Alloc()
	rv := objc.Send[IMKServer](instance.ID, objc.Sel("initWithName:bundleIdentifier:"), objc.String(name), objc.String(bundleIdentifier))
	rv.Autorelease()
	return rv
}


// Creates and returns a server object initialized with the provided parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:controllerClass:delegateClass:)
func NewIMKServerWithNameControllerClassDelegateClass(name string, controllerClassID objc.Class, delegateClassID objc.Class) IMKServer {
	instance := getIMKServerClass().Alloc()
	rv := objc.Send[IMKServer](instance.ID, objc.Sel("initWithName:controllerClass:delegateClass:"), objc.String(name), controllerClassID, delegateClassID)
	rv.Autorelease()
	return rv
}



// Returns an object for the input method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer/bundle()
func (i_ IMKServer) Bundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](i_.ID, objc.Sel("bundle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer/lastKeyEventWasDeadKey()
func (i_ IMKServer) LastKeyEventWasDeadKey() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("lastKeyEventWasDeadKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKServer/paletteWillTerminate()
func (i_ IMKServer) PaletteWillTerminate() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("paletteWillTerminate"))
	return rv
}


