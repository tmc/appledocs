// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HTTPCookieStorage] class.
var (
	HTTPCookieStorageClass     _HTTPCookieStorageClass
	HTTPCookieStorageClassOnce sync.Once
)

func getHTTPCookieStorageClass() _HTTPCookieStorageClass {
	HTTPCookieStorageClassOnce.Do(func() {
		HTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}
	})
	return HTTPCookieStorageClass
}

type _HTTPCookieStorageClass struct {
	class objc.Class
}

// An interface definition for the [HTTPCookieStorage] class.
type IHTTPCookieStorage interface {
	objectivec.IObject
}

// A container that manages the storage of cookies.
//
// Each stored cookie is represented by an instance of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage
type HTTPCookieStorage struct {
	objectivec.Object
}

// HTTPCookieStorageFrom constructs a [HTTPCookieStorage] from an unsafe.Pointer.
//
// A container that manages the storage of cookies.
func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieStorageClass) Alloc() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HTTPCookieStorageClass) New() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPCookieStorage) Init() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPCookieStorage) Autorelease() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookieStorage creates a new HTTPCookieStorage instance.
func NewHTTPCookieStorage() HTTPCookieStorage {
	return getHTTPCookieStorageClass().New()
}

// The shared cookie storage instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (hc _HTTPCookieStorageClass) SharedHTTPCookieStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("sharedHTTPCookieStorage"))
	return rv
}

// The shared cookie storage instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (h_ HTTPCookieStorage) SharedHTTPCookieStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sharedHTTPCookieStorage"))
	return rv
}
