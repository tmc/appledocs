// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSModuleIdentity] class.
var (
	FSModuleIdentityClass     _FSModuleIdentityClass
	FSModuleIdentityClassOnce sync.Once
)

func getFSModuleIdentityClass() _FSModuleIdentityClass {
	FSModuleIdentityClassOnce.Do(func() {
		FSModuleIdentityClass = _FSModuleIdentityClass{objc.GetClass("FSModuleIdentity")}
	})
	return FSModuleIdentityClass
}

type _FSModuleIdentityClass struct {
	class objc.Class
}

// An interface definition for the [FSModuleIdentity] class.
type IFSModuleIdentity interface {
	objectivec.IObject
}

// An installed file system module.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity
type FSModuleIdentity struct {
	objectivec.Object
}

// FSModuleIdentityFrom constructs a [FSModuleIdentity] from an unsafe.Pointer.
//
// An installed file system module.
func FSModuleIdentityFrom(ptr unsafe.Pointer) FSModuleIdentity {
	return FSModuleIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSModuleIdentityClass) Alloc() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSModuleIdentityClass) New() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSModuleIdentity) Init() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSModuleIdentity) Autorelease() FSModuleIdentity {
	rv := objc.Send[FSModuleIdentity](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSModuleIdentity creates a new FSModuleIdentity instance.
func NewFSModuleIdentity() FSModuleIdentity {
	return getFSModuleIdentityClass().New()
}


// A Boolean value that indicates if the module is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsmoduleidentity/isenabled
func (f_ FSModuleIdentity) IsEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates if the module is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsmoduleidentity/isenabled
func (f_ FSModuleIdentity) SetIsEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsEnabled:"), value)
}

// The module’s bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/bundleIdentifier
func (f_ FSModuleIdentity) BundleIdentifier() string {
	rv := objc.Send[string](f_.ID, objc.Sel("bundleIdentifier"))
	return rv
}

// A Boolean value that indicates if the module is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/isEnabled
func (f_ FSModuleIdentity) Enabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("enabled"))
	return rv
}

// The module’s URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSModuleIdentity/url
func (f_ FSModuleIdentity) Url() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("url"))
	return rv
}



