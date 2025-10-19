// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BundleResourceRequest] class.
var (
	bundleResourceRequestClass     _BundleResourceRequestClass
	bundleResourceRequestClassOnce sync.Once
)

func getBundleResourceRequestClass() _BundleResourceRequestClass {
	bundleResourceRequestClassOnce.Do(func() {
		bundleResourceRequestClass = _BundleResourceRequestClass{objc.GetClass("NSBundleResourceRequest")}
	})
	return bundleResourceRequestClass
}

type _BundleResourceRequestClass struct {
	class objc.Class
}

// An interface definition for the [BundleResourceRequest] class.
type IBundleResourceRequest interface {
	objectivec.IObject
}

// A resource manager you use to download content hosted on the App Store at the time your app needs it.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest
type BundleResourceRequest struct {
	objectivec.Object
}

// BundleResourceRequestFrom constructs a [BundleResourceRequest] from an unsafe.Pointer.
//
// A resource manager you use to download content hosted on the App Store at the time your app needs it.
func BundleResourceRequestFrom(ptr unsafe.Pointer) BundleResourceRequest {
	return BundleResourceRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BundleResourceRequestClass) Alloc() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BundleResourceRequestClass) New() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BundleResourceRequest) Init() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BundleResourceRequest) Autorelease() BundleResourceRequest {
	rv := objc.Send[BundleResourceRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBundleResourceRequest creates a new BundleResourceRequest instance.
func NewBundleResourceRequest() BundleResourceRequest {
	return getBundleResourceRequestClass().New()
}




