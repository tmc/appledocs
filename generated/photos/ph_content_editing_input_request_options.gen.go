// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHContentEditingInputRequestOptions] class.
var (
	PHContentEditingInputRequestOptionsClass     _PHContentEditingInputRequestOptionsClass
	PHContentEditingInputRequestOptionsClassOnce sync.Once
)

func getPHContentEditingInputRequestOptionsClass() _PHContentEditingInputRequestOptionsClass {
	PHContentEditingInputRequestOptionsClassOnce.Do(func() {
		PHContentEditingInputRequestOptionsClass = _PHContentEditingInputRequestOptionsClass{objc.GetClass("PHContentEditingInputRequestOptions")}
	})
	return PHContentEditingInputRequestOptionsClass
}

type _PHContentEditingInputRequestOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHContentEditingInputRequestOptions] class.
type IPHContentEditingInputRequestOptions interface {
	objectivec.IObject
	// properties:
	CanHandleAdjustmentData() bool
	SetCanHandleAdjustmentData(value bool)
	IsNetworkAccessAllowed() bool
	SetIsNetworkAccessAllowed(value bool)
	ProgressHandler() unsafe.Pointer
	SetProgressHandler(value unsafe.Pointer)
	// methods:
}

// A set of options affecting the delivery of image or video data when you request to edit the content of a Photos asset.
//
// You use the class with the method for editing the contents of a object. This class doesn’t affect photo editing extensions.

// A set of options affecting the delivery of image or video data when you request to edit the content of a Photos asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInputRequestOptions
type PHContentEditingInputRequestOptions struct {
	objectivec.Object
}

// PHContentEditingInputRequestOptionsFrom constructs a [PHContentEditingInputRequestOptions] from an unsafe.Pointer.
//
// A set of options affecting the delivery of image or video data when you request to edit the content of a Photos asset.
func PHContentEditingInputRequestOptionsFrom(ptr unsafe.Pointer) PHContentEditingInputRequestOptions {
	return PHContentEditingInputRequestOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHContentEditingInputRequestOptionsClass) Alloc() PHContentEditingInputRequestOptions {
	rv := objc.Send[PHContentEditingInputRequestOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHContentEditingInputRequestOptionsClass) New() PHContentEditingInputRequestOptions {
	rv := objc.Send[PHContentEditingInputRequestOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHContentEditingInputRequestOptions) Init() PHContentEditingInputRequestOptions {
	rv := objc.Send[PHContentEditingInputRequestOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHContentEditingInputRequestOptions) Autorelease() PHContentEditingInputRequestOptions {
	rv := objc.Send[PHContentEditingInputRequestOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHContentEditingInputRequestOptions creates a new PHContentEditingInputRequestOptions instance.
func NewPHContentEditingInputRequestOptions() PHContentEditingInputRequestOptions {
	return getPHContentEditingInputRequestOptionsClass().New()
}

// A block to be called when Photos needs to determine whether your app can continue previous edits made to an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/canhandleadjustmentdata
func (p_ PHContentEditingInputRequestOptions) CanHandleAdjustmentData() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canHandleAdjustmentData"))
	return rv
}

// A block to be called when Photos needs to determine whether your app can continue previous edits made to an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/canhandleadjustmentdata
func (p_ PHContentEditingInputRequestOptions) SetCanHandleAdjustmentData(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanHandleAdjustmentData:"), value)
}

// A Boolean value that specifies whether Photos can download the asset from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/isnetworkaccessallowed
func (p_ PHContentEditingInputRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}

// A Boolean value that specifies whether Photos can download the asset from iCloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/isnetworkaccessallowed
func (p_ PHContentEditingInputRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

// A block Photos calls periodically while downloading the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/progresshandler
func (p_ PHContentEditingInputRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}

// A block Photos calls periodically while downloading the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginputrequestoptions/progresshandler
func (p_ PHContentEditingInputRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}
