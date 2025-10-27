// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CaptureDeferredPhotoProxy] class.
var (
	CaptureDeferredPhotoProxyClass     _CaptureDeferredPhotoProxyClass
	CaptureDeferredPhotoProxyClassOnce sync.Once
)

func getCaptureDeferredPhotoProxyClass() _CaptureDeferredPhotoProxyClass {
	CaptureDeferredPhotoProxyClassOnce.Do(func() {
		CaptureDeferredPhotoProxyClass = _CaptureDeferredPhotoProxyClass{objc.GetClass("AVCaptureDeferredPhotoProxy")}
	})
	return CaptureDeferredPhotoProxyClass
}

type _CaptureDeferredPhotoProxyClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDeferredPhotoProxy] class.
type ICaptureDeferredPhotoProxy interface {
	ICapturePhoto
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDeferredPhotoProxyClass) Alloc() CaptureDeferredPhotoProxy {
	rv := objc.Send[CaptureDeferredPhotoProxy](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeferredPhotoProxyClass) New() CaptureDeferredPhotoProxy {
	rv := objc.Send[CaptureDeferredPhotoProxy](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeferredPhotoProxy) Init() CaptureDeferredPhotoProxy {
	rv := objc.Send[CaptureDeferredPhotoProxy](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeferredPhotoProxy) Autorelease() CaptureDeferredPhotoProxy {
	rv := objc.Send[CaptureDeferredPhotoProxy](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeferredPhotoProxy creates a new CaptureDeferredPhotoProxy instance.
func NewCaptureDeferredPhotoProxy() CaptureDeferredPhotoProxy {
	return getCaptureDeferredPhotoProxyClass().New()
}





// A lightly-processed photo with data that the system may use to process and fetch a higher-resolution asset at a later time.
//
// A photo proxy behaves like a normal , and approximates the look of the final rendered image. This object represents intermediate data that the system can render into a final image and ingested into the user’s photo library using the framework. The intermediate data aren’t accessible by the calling process.


// A lightly-processed photo with data that the system may use to process and fetch a higher-resolution asset at a later time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeferredPhotoProxy
type CaptureDeferredPhotoProxy struct {
	CapturePhoto
}

// CaptureDeferredPhotoProxyFrom constructs a [CaptureDeferredPhotoProxy] from an unsafe.Pointer.
//
// A lightly-processed photo with data that the system may use to process and fetch a higher-resolution asset at a later time.
func CaptureDeferredPhotoProxyFrom(ptr unsafe.Pointer) CaptureDeferredPhotoProxy {
	return CaptureDeferredPhotoProxy{
		CapturePhoto: CapturePhotoFrom(ptr),
	}
}































