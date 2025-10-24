// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHVideoRequestOptions] class.
var (
	PHVideoRequestOptionsClass     _PHVideoRequestOptionsClass
	PHVideoRequestOptionsClassOnce sync.Once
)

func getPHVideoRequestOptionsClass() _PHVideoRequestOptionsClass {
	PHVideoRequestOptionsClassOnce.Do(func() {
		PHVideoRequestOptionsClass = _PHVideoRequestOptionsClass{objc.GetClass("PHVideoRequestOptions")}
	})
	return PHVideoRequestOptionsClass
}

type _PHVideoRequestOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHVideoRequestOptions] class.
type IPHVideoRequestOptions interface {
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

// A set of options affecting the delivery of video asset data that you request from an image manager.

// A set of options affecting the delivery of video asset data that you request from an image manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptions
type PHVideoRequestOptions struct {
	objectivec.Object
}

// PHVideoRequestOptionsFrom constructs a [PHVideoRequestOptions] from an unsafe.Pointer.
//
// A set of options affecting the delivery of video asset data that you request from an image manager.
func PHVideoRequestOptionsFrom(ptr unsafe.Pointer) PHVideoRequestOptions {
	return PHVideoRequestOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHVideoRequestOptionsClass) Alloc() PHVideoRequestOptions {
	rv := objc.Send[PHVideoRequestOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHVideoRequestOptionsClass) New() PHVideoRequestOptions {
	rv := objc.Send[PHVideoRequestOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHVideoRequestOptions) Init() PHVideoRequestOptions {
	rv := objc.Send[PHVideoRequestOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHVideoRequestOptions) Autorelease() PHVideoRequestOptions {
	rv := objc.Send[PHVideoRequestOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHVideoRequestOptions creates a new PHVideoRequestOptions instance.
func NewPHVideoRequestOptions() PHVideoRequestOptions {
	return getPHVideoRequestOptionsClass().New()
}

// A mode specifying the requested video quality and delivery priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/deliverymode
func (p_ PHVideoRequestOptions) DeliveryMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deliveryMode"))
	return rv
}

// A mode specifying the requested video quality and delivery priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/deliverymode
func (p_ PHVideoRequestOptions) SetDeliveryMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeliveryMode:"), value)
}

// A Boolean value that specifies whether Photos can download the requested video from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/isnetworkaccessallowed
func (p_ PHVideoRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}

// A Boolean value that specifies whether Photos can download the requested video from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/isnetworkaccessallowed
func (p_ PHVideoRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

// A block Photos calls periodically while downloading the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/progresshandler
func (p_ PHVideoRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}

// A block Photos calls periodically while downloading the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/progresshandler
func (p_ PHVideoRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}

// The version of the video to request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/version
func (p_ PHVideoRequestOptions) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("version"))
	return rv
}

// The version of the video to request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phvideorequestoptions/version
func (p_ PHVideoRequestOptions) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVersion:"), value)
}
