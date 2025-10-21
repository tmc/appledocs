// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHLivePhotoRequestOptions] class.
var (
	PHLivePhotoRequestOptionsClass     _PHLivePhotoRequestOptionsClass
	PHLivePhotoRequestOptionsClassOnce sync.Once
)

func getPHLivePhotoRequestOptionsClass() _PHLivePhotoRequestOptionsClass {
	PHLivePhotoRequestOptionsClassOnce.Do(func() {
		PHLivePhotoRequestOptionsClass = _PHLivePhotoRequestOptionsClass{objc.GetClass("PHLivePhotoRequestOptions")}
	})
	return PHLivePhotoRequestOptionsClass
}

type _PHLivePhotoRequestOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHLivePhotoRequestOptions] class.
type IPHLivePhotoRequestOptions interface {
	objectivec.IObject
}

// A set of options affecting the delivery of Live Photo assets you request from an image manager.
//
// A Live Photo is a picture that includes movement and sound from the moments just before and after its capture.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions
type PHLivePhotoRequestOptions struct {
	objectivec.Object
}

// PHLivePhotoRequestOptionsFrom constructs a [PHLivePhotoRequestOptions] from an unsafe.Pointer.
//
// A set of options affecting the delivery of Live Photo assets you request from an image manager.
func PHLivePhotoRequestOptionsFrom(ptr unsafe.Pointer) PHLivePhotoRequestOptions {
	return PHLivePhotoRequestOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHLivePhotoRequestOptionsClass) Alloc() PHLivePhotoRequestOptions {
	rv := objc.Send[PHLivePhotoRequestOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHLivePhotoRequestOptionsClass) New() PHLivePhotoRequestOptions {
	rv := objc.Send[PHLivePhotoRequestOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHLivePhotoRequestOptions) Init() PHLivePhotoRequestOptions {
	rv := objc.Send[PHLivePhotoRequestOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHLivePhotoRequestOptions) Autorelease() PHLivePhotoRequestOptions {
	rv := objc.Send[PHLivePhotoRequestOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHLivePhotoRequestOptions creates a new PHLivePhotoRequestOptions instance.
func NewPHLivePhotoRequestOptions() PHLivePhotoRequestOptions {
	return getPHLivePhotoRequestOptionsClass().New()
}


// The requested Live Photo quality and delivery priority.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/deliveryMode
func (p_ PHLivePhotoRequestOptions) DeliveryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deliveryMode"))
	return rv
}


// SetDeliveryMode sets the value of the deliveryMode property.
// The requested Live Photo quality and delivery priority.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/deliveryMode
func (p_ PHLivePhotoRequestOptions) SetDeliveryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeliveryMode:"), value)
}
// A Boolean value that specifies whether Photos can download the requested Live Photo data from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/isNetworkAccessAllowed
func (p_ PHLivePhotoRequestOptions) NetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("networkAccessAllowed"))
	return rv
}


// SetNetworkAccessAllowed sets the value of the networkAccessAllowed property.
// A Boolean value that specifies whether Photos can download the requested Live Photo data from iCloud.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/isNetworkAccessAllowed
func (p_ PHLivePhotoRequestOptions) SetNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNetworkAccessAllowed:"), value)
}
// A block that Photos calls periodically while downloading the Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/progressHandler
func (p_ PHLivePhotoRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}


// SetProgressHandler sets the value of the progressHandler property.
// A block that Photos calls periodically while downloading the Live Photo.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/progressHandler
func (p_ PHLivePhotoRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}
// The version of the Live Photo to be requested.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/version
func (p_ PHLivePhotoRequestOptions) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The version of the Live Photo to be requested.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoRequestOptions/version
func (p_ PHLivePhotoRequestOptions) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersion:"), value)
}


