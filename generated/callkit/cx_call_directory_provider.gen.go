// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallDirectoryProvider] class.
var (
	CXCallDirectoryProviderClass     _CXCallDirectoryProviderClass
	CXCallDirectoryProviderClassOnce sync.Once
)

func getCXCallDirectoryProviderClass() _CXCallDirectoryProviderClass {
	CXCallDirectoryProviderClassOnce.Do(func() {
		CXCallDirectoryProviderClass = _CXCallDirectoryProviderClass{objc.GetClass("CXCallDirectoryProvider")}
	})
	return CXCallDirectoryProviderClass
}

type _CXCallDirectoryProviderClass struct {
	class objc.Class
}

// An interface definition for the [CXCallDirectoryProvider] class.
type ICXCallDirectoryProvider interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The principal object for a Call Directory app extension for a host app.


// The principal object for a Call Directory app extension for a host app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryProvider
type CXCallDirectoryProvider struct {
	objectivec.Object
}

// CXCallDirectoryProviderFrom constructs a [CXCallDirectoryProvider] from an unsafe.Pointer.
//
// The principal object for a Call Directory app extension for a host app.
func CXCallDirectoryProviderFrom(ptr unsafe.Pointer) CXCallDirectoryProvider {
	return CXCallDirectoryProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryProviderClass) Alloc() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallDirectoryProviderClass) New() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryProvider) Init() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryProvider) Autorelease() CXCallDirectoryProvider {
	rv := objc.Send[CXCallDirectoryProvider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryProvider creates a new CXCallDirectoryProvider instance.
func NewCXCallDirectoryProvider() CXCallDirectoryProvider {
	return getCXCallDirectoryProviderClass().New()
}



