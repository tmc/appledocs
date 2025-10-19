// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLComponents] class.
var (
	uRLComponentsClass     _URLComponentsClass
	uRLComponentsClassOnce sync.Once
)

func getURLComponentsClass() _URLComponentsClass {
	uRLComponentsClassOnce.Do(func() {
		uRLComponentsClass = _URLComponentsClass{objc.GetClass("NSURLComponents")}
	})
	return uRLComponentsClass
}

type _URLComponentsClass struct {
	class objc.Class
}

// An interface definition for the [URLComponents] class.
type IURLComponents interface {
	objectivec.IObject
}

// An object that parses URLs into and constructs URLs from their constituent parts.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents
type URLComponents struct {
	objectivec.Object
}

// URLComponentsFrom constructs a [URLComponents] from an unsafe.Pointer.
//
// An object that parses URLs into and constructs URLs from their constituent parts.
func URLComponentsFrom(ptr unsafe.Pointer) URLComponents {
	return URLComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLComponentsClass) Alloc() URLComponents {
	rv := objc.Send[URLComponents](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLComponentsClass) New() URLComponents {
	rv := objc.Send[URLComponents](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLComponents) Init() URLComponents {
	rv := objc.Send[URLComponents](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLComponents) Autorelease() URLComponents {
	rv := objc.Send[URLComponents](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLComponents creates a new URLComponents instance.
func NewURLComponents() URLComponents {
	return getURLComponentsClass().New()
}





