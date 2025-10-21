// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHImageRequestOptions] class.
var (
	PHImageRequestOptionsClass     _PHImageRequestOptionsClass
	PHImageRequestOptionsClassOnce sync.Once
)

func getPHImageRequestOptionsClass() _PHImageRequestOptionsClass {
	PHImageRequestOptionsClassOnce.Do(func() {
		PHImageRequestOptionsClass = _PHImageRequestOptionsClass{objc.GetClass("PHImageRequestOptions")}
	})
	return PHImageRequestOptionsClass
}

type _PHImageRequestOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHImageRequestOptions] class.
type IPHImageRequestOptions interface {
	objectivec.IObject
}

// A set of options affecting the delivery of still image representations of Photos assets you request from an image manager.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions
type PHImageRequestOptions struct {
	objectivec.Object
}

// PHImageRequestOptionsFrom constructs a [PHImageRequestOptions] from an unsafe.Pointer.
//
// A set of options affecting the delivery of still image representations of Photos assets you request from an image manager.
func PHImageRequestOptionsFrom(ptr unsafe.Pointer) PHImageRequestOptions {
	return PHImageRequestOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHImageRequestOptionsClass) Alloc() PHImageRequestOptions {
	rv := objc.Send[PHImageRequestOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHImageRequestOptionsClass) New() PHImageRequestOptions {
	rv := objc.Send[PHImageRequestOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHImageRequestOptions) Init() PHImageRequestOptions {
	rv := objc.Send[PHImageRequestOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHImageRequestOptions) Autorelease() PHImageRequestOptions {
	rv := objc.Send[PHImageRequestOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHImageRequestOptions creates a new PHImageRequestOptions instance.
func NewPHImageRequestOptions() PHImageRequestOptions {
	return getPHImageRequestOptionsClass().New()
}


// A Boolean value that determines whether Photos processes the image request synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/issynchronous
func (p_ PHImageRequestOptions) IsSynchronous() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSynchronous"))
	return rv
}


// SetIsSynchronous sets the value of the isSynchronous property.
// A Boolean value that determines whether Photos processes the image request synchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/issynchronous
func (p_ PHImageRequestOptions) SetIsSynchronous(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSynchronous:"), value)
}

// A Boolean value that specifies whether Photos can download the requested image from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/isnetworkaccessallowed
func (p_ PHImageRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}


// SetIsNetworkAccessAllowed sets the value of the isNetworkAccessAllowed property.
// A Boolean value that specifies whether Photos can download the requested image from iCloud.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/isnetworkaccessallowed
func (p_ PHImageRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/allowSecondaryDegradedImage
func (p_ PHImageRequestOptions) AllowSecondaryDegradedImage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowSecondaryDegradedImage"))
	return rv
}


// SetAllowSecondaryDegradedImage sets the value of the allowSecondaryDegradedImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/allowSecondaryDegradedImage
func (p_ PHImageRequestOptions) SetAllowSecondaryDegradedImage(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowSecondaryDegradedImage:"), value)
}

// The requested image quality and delivery priority.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/deliveryMode
func (p_ PHImageRequestOptions) DeliveryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deliveryMode"))
	return rv
}


// SetDeliveryMode sets the value of the deliveryMode property.
// The requested image quality and delivery priority.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/deliveryMode
func (p_ PHImageRequestOptions) SetDeliveryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeliveryMode:"), value)
}

// A Boolean value that specifies whether Photos can download the requested image from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/isNetworkAccessAllowed
func (p_ PHImageRequestOptions) NetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("networkAccessAllowed"))
	return rv
}


// SetNetworkAccessAllowed sets the value of the networkAccessAllowed property.
// A Boolean value that specifies whether Photos can download the requested image from iCloud.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/isNetworkAccessAllowed
func (p_ PHImageRequestOptions) SetNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNetworkAccessAllowed:"), value)
}

// A Boolean value that determines whether Photos processes the image request synchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/isSynchronous
func (p_ PHImageRequestOptions) Synchronous() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("synchronous"))
	return rv
}


// SetSynchronous sets the value of the synchronous property.
// A Boolean value that determines whether Photos processes the image request synchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/isSynchronous
func (p_ PHImageRequestOptions) SetSynchronous(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSynchronous:"), value)
}

// A rectangle for requesting a cropped version of the original image.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/normalizedCropRect
func (p_ PHImageRequestOptions) NormalizedCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("normalizedCropRect"))
	return rv
}


// SetNormalizedCropRect sets the value of the normalizedCropRect property.
// A rectangle for requesting a cropped version of the original image.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/normalizedCropRect
func (p_ PHImageRequestOptions) SetNormalizedCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalizedCropRect:"), value)
}

// A block that Photos calls periodically while downloading the image.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/progressHandler
func (p_ PHImageRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}


// SetProgressHandler sets the value of the progressHandler property.
// A block that Photos calls periodically while downloading the image.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/progressHandler
func (p_ PHImageRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}

// A mode that specifies how to resize the requested image.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/resizeMode
func (p_ PHImageRequestOptions) ResizeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resizeMode"))
	return rv
}


// SetResizeMode sets the value of the resizeMode property.
// A mode that specifies how to resize the requested image.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/resizeMode
func (p_ PHImageRequestOptions) SetResizeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResizeMode:"), value)
}

// The version of the image to be requested.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/version
func (p_ PHImageRequestOptions) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The version of the image to be requested.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptions/version
func (p_ PHImageRequestOptions) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersion:"), value)
}



