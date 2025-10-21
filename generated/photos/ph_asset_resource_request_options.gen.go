// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHAssetResourceRequestOptions] class.
var (
	PHAssetResourceRequestOptionsClass     _PHAssetResourceRequestOptionsClass
	PHAssetResourceRequestOptionsClassOnce sync.Once
)

func getPHAssetResourceRequestOptionsClass() _PHAssetResourceRequestOptionsClass {
	PHAssetResourceRequestOptionsClassOnce.Do(func() {
		PHAssetResourceRequestOptionsClass = _PHAssetResourceRequestOptionsClass{objc.GetClass("PHAssetResourceRequestOptions")}
	})
	return PHAssetResourceRequestOptionsClass
}

type _PHAssetResourceRequestOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResourceRequestOptions] class.
type IPHAssetResourceRequestOptions interface {
	objectivec.IObject
}

// A set of options affecting the delivery of underlying asset data that you request from the asset resource manager.
//
// You use this class when requesting the underlying data for photo, video, and Live Photo asset resources from a object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceRequestOptions
type PHAssetResourceRequestOptions struct {
	objectivec.Object
}

// PHAssetResourceRequestOptionsFrom constructs a [PHAssetResourceRequestOptions] from an unsafe.Pointer.
//
// A set of options affecting the delivery of underlying asset data that you request from the asset resource manager.
func PHAssetResourceRequestOptionsFrom(ptr unsafe.Pointer) PHAssetResourceRequestOptions {
	return PHAssetResourceRequestOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceRequestOptionsClass) Alloc() PHAssetResourceRequestOptions {
	rv := objc.Send[PHAssetResourceRequestOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceRequestOptionsClass) New() PHAssetResourceRequestOptions {
	rv := objc.Send[PHAssetResourceRequestOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResourceRequestOptions) Init() PHAssetResourceRequestOptions {
	rv := objc.Send[PHAssetResourceRequestOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResourceRequestOptions) Autorelease() PHAssetResourceRequestOptions {
	rv := objc.Send[PHAssetResourceRequestOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResourceRequestOptions creates a new PHAssetResourceRequestOptions instance.
func NewPHAssetResourceRequestOptions() PHAssetResourceRequestOptions {
	return getPHAssetResourceRequestOptionsClass().New()
}


// A Boolean value that specifies whether Photos can download the requested asset resource data from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresourcerequestoptions/isnetworkaccessallowed
func (p_ PHAssetResourceRequestOptions) IsNetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isNetworkAccessAllowed"))
	return rv
}


// SetIsNetworkAccessAllowed sets the value of the isNetworkAccessAllowed property.
// A Boolean value that specifies whether Photos can download the requested asset resource data from iCloud.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresourcerequestoptions/isnetworkaccessallowed
func (p_ PHAssetResourceRequestOptions) SetIsNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsNetworkAccessAllowed:"), value)
}

// A Boolean value that specifies whether Photos can download the requested asset resource data from iCloud.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceRequestOptions/isNetworkAccessAllowed
func (p_ PHAssetResourceRequestOptions) NetworkAccessAllowed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("networkAccessAllowed"))
	return rv
}


// SetNetworkAccessAllowed sets the value of the networkAccessAllowed property.
// A Boolean value that specifies whether Photos can download the requested asset resource data from iCloud.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceRequestOptions/isNetworkAccessAllowed
func (p_ PHAssetResourceRequestOptions) SetNetworkAccessAllowed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNetworkAccessAllowed:"), value)
}

// A block that Photos calls periodically while downloading the asset resource data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceRequestOptions/progressHandler
func (p_ PHAssetResourceRequestOptions) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("progressHandler"))
	return rv
}


// SetProgressHandler sets the value of the progressHandler property.
// A block that Photos calls periodically while downloading the asset resource data.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceRequestOptions/progressHandler
func (p_ PHAssetResourceRequestOptions) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProgressHandler:"), value)
}



