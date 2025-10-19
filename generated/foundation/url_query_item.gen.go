// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLQueryItem] class.
var (
	uRLQueryItemClass     _URLQueryItemClass
	uRLQueryItemClassOnce sync.Once
)

func getURLQueryItemClass() _URLQueryItemClass {
	uRLQueryItemClassOnce.Do(func() {
		uRLQueryItemClass = _URLQueryItemClass{objc.GetClass("NSURLQueryItem")}
	})
	return uRLQueryItemClass
}

type _URLQueryItemClass struct {
	class objc.Class
}

// An interface definition for the [URLQueryItem] class.
type IURLQueryItem interface {
	objectivec.IObject
}

// An object representing a single name/value pair for an item in the query portion of a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem
type URLQueryItem struct {
	objectivec.Object
}

// URLQueryItemFrom constructs a [URLQueryItem] from an unsafe.Pointer.
//
// An object representing a single name/value pair for an item in the query portion of a URL.
func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLQueryItemClass) Alloc() URLQueryItem {
	rv := objc.Send[URLQueryItem](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLQueryItemClass) New() URLQueryItem {
	rv := objc.Send[URLQueryItem](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLQueryItem) Init() URLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLQueryItem) Autorelease() URLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLQueryItem creates a new URLQueryItem instance.
func NewURLQueryItem() URLQueryItem {
	return getURLQueryItemClass().New()
}




