// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HTTPCookie] class.
var (
	hTTPCookieClass     _HTTPCookieClass
	hTTPCookieClassOnce sync.Once
)

func getHTTPCookieClass() _HTTPCookieClass {
	hTTPCookieClassOnce.Do(func() {
		hTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}
	})
	return hTTPCookieClass
}

type _HTTPCookieClass struct {
	class objc.Class
}

// An interface definition for the [HTTPCookie] class.
type IHTTPCookie interface {
	objectivec.IObject
}

// A representation of an HTTP cookie.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie
type HTTPCookie struct {
	objectivec.Object
}

// HTTPCookieFrom constructs a [HTTPCookie] from an unsafe.Pointer.
//
// A representation of an HTTP cookie.
func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.Send[HTTPCookie](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPCookie) Autorelease() HTTPCookie {
	rv := objc.Send[HTTPCookie](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookie creates a new HTTPCookie instance.
func NewHTTPCookie() HTTPCookie {
	return getHTTPCookieClass().New()
}




