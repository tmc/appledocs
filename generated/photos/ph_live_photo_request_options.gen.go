// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	DeliveryMode() unsafe.Pointer
	SetDeliveryMode(value unsafe.Pointer)
	IsNetworkAccessAllowed() bool
	SetIsNetworkAccessAllowed(value bool)
	ProgressHandler() unsafe.Pointer
	SetProgressHandler(value unsafe.Pointer)
	Version() unsafe.Pointer
	SetVersion(value unsafe.Pointer)
	// methods:
}

// A set of options affecting the delivery of Live Photo assets you request from an image manager.
//
// A Live Photo is a picture that includes movement and sound from the moments just before and after its capture.

// A set of options affecting the delivery of Live Photo assets you request from an image manager.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/deliverymode
func (p_ PHLivePhotoRequestOptions) DeliveryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deliveryMode"))
	return rv
}

// The requested Live Photo quality and delivery priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/deliverymode
func (p_ PHLivePhotoRequestOptions) SetDeliveryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeliveryMode:"), value)
}

// A Boolean value that specifies whether Photos can download the requested Live Photo data from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/isnetworkaccessallowed
func (p_ PHLivePhotoRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}

// A Boolean value that specifies whether Photos can download the requested Live Photo data from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/isnetworkaccessallowed
func (p_ PHLivePhotoRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

// A block that Photos calls periodically while downloading the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/progresshandler
func (p_ PHLivePhotoRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}

// A block that Photos calls periodically while downloading the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/progresshandler
func (p_ PHLivePhotoRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}

// The version of the Live Photo to be requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/version
func (p_ PHLivePhotoRequestOptions) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("version"))
	return rv
}

// The version of the Live Photo to be requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotorequestoptions/version
func (p_ PHLivePhotoRequestOptions) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersion:"), value)
}
