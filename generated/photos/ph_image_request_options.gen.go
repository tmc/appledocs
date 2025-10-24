// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	AllowSecondaryDegradedImage() bool
	SetAllowSecondaryDegradedImage(value bool)
	DeliveryMode() unsafe.Pointer
	SetDeliveryMode(value unsafe.Pointer)
	IsNetworkAccessAllowed() bool
	SetIsNetworkAccessAllowed(value bool)
	IsSynchronous() bool
	SetIsSynchronous(value bool)
	NormalizedCropRect() objc.IObject /* cross-framework: Rect */
	SetNormalizedCropRect(value objc.IObject /* cross-framework: Rect */)
	ProgressHandler() unsafe.Pointer
	SetProgressHandler(value unsafe.Pointer)
	ResizeMode() unsafe.Pointer
	SetResizeMode(value unsafe.Pointer)
	Version() unsafe.Pointer
	SetVersion(value unsafe.Pointer)
	// methods:
}

// A set of options affecting the delivery of still image representations of Photos assets you request from an image manager.

// A set of options affecting the delivery of still image representations of Photos assets you request from an image manager.
//
// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/allowsecondarydegradedimage
func (p_ PHImageRequestOptions) AllowSecondaryDegradedImage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowSecondaryDegradedImage"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/allowsecondarydegradedimage
func (p_ PHImageRequestOptions) SetAllowSecondaryDegradedImage(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowSecondaryDegradedImage:"), value)
}

// The requested image quality and delivery priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/deliverymode
func (p_ PHImageRequestOptions) DeliveryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deliveryMode"))
	return rv
}

// The requested image quality and delivery priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/deliverymode
func (p_ PHImageRequestOptions) SetDeliveryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeliveryMode:"), value)
}

// A Boolean value that specifies whether Photos can download the requested image from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/isnetworkaccessallowed
func (p_ PHImageRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}

// A Boolean value that specifies whether Photos can download the requested image from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/isnetworkaccessallowed
func (p_ PHImageRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

// A Boolean value that determines whether Photos processes the image request synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/issynchronous
func (p_ PHImageRequestOptions) IsSynchronous() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSynchronous"))
	return rv
}

// A Boolean value that determines whether Photos processes the image request synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/issynchronous
func (p_ PHImageRequestOptions) SetIsSynchronous(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSynchronous:"), value)
}

// A rectangle for requesting a cropped version of the original image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/normalizedcroprect
func (p_ PHImageRequestOptions) NormalizedCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("normalizedCropRect"))
	return rv
}

// A rectangle for requesting a cropped version of the original image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/normalizedcroprect
func (p_ PHImageRequestOptions) SetNormalizedCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalizedCropRect:"), value)
}

// A block that Photos calls periodically while downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/progresshandler
func (p_ PHImageRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}

// A block that Photos calls periodically while downloading the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/progresshandler
func (p_ PHImageRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}

// A mode that specifies how to resize the requested image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/resizemode
func (p_ PHImageRequestOptions) ResizeMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resizeMode"))
	return rv
}

// A mode that specifies how to resize the requested image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/resizemode
func (p_ PHImageRequestOptions) SetResizeMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResizeMode:"), value)
}

// The version of the image to be requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/version
func (p_ PHImageRequestOptions) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("version"))
	return rv
}

// The version of the image to be requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagerequestoptions/version
func (p_ PHImageRequestOptions) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersion:"), value)
}
