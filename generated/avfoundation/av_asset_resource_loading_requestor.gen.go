// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetResourceLoadingRequestor] class.
var (
	AssetResourceLoadingRequestorClass     _AssetResourceLoadingRequestorClass
	AssetResourceLoadingRequestorClassOnce sync.Once
)

func getAssetResourceLoadingRequestorClass() _AssetResourceLoadingRequestorClass {
	AssetResourceLoadingRequestorClassOnce.Do(func() {
		AssetResourceLoadingRequestorClass = _AssetResourceLoadingRequestorClass{objc.GetClass("AVAssetResourceLoadingRequestor")}
	})
	return AssetResourceLoadingRequestorClass
}

type _AssetResourceLoadingRequestorClass struct {
	class objc.Class
}





// An interface definition for the [AssetResourceLoadingRequestor] class.
type IAssetResourceLoadingRequestor interface {
	objectivec.IObject
	

	// properties:
	ProvidesExpiredSessionReports() bool


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoadingRequestorClass) Alloc() AssetResourceLoadingRequestor {
	rv := objc.Send[AssetResourceLoadingRequestor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetResourceLoadingRequestorClass) New() AssetResourceLoadingRequestor {
	rv := objc.Send[AssetResourceLoadingRequestor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceLoadingRequestor) Init() AssetResourceLoadingRequestor {
	rv := objc.Send[AssetResourceLoadingRequestor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceLoadingRequestor) Autorelease() AssetResourceLoadingRequestor {
	rv := objc.Send[AssetResourceLoadingRequestor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceLoadingRequestor creates a new AssetResourceLoadingRequestor instance.
func NewAssetResourceLoadingRequestor() AssetResourceLoadingRequestor {
	return getAssetResourceLoadingRequestorClass().New()
}





// An object that contains information about the originator of a resource-loading request.


// An object that contains information about the originator of a resource-loading request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestor
type AssetResourceLoadingRequestor struct {
	objectivec.Object
}

// AssetResourceLoadingRequestorFrom constructs a [AssetResourceLoadingRequestor] from an unsafe.Pointer.
//
// An object that contains information about the originator of a resource-loading request.
func AssetResourceLoadingRequestorFrom(ptr unsafe.Pointer) AssetResourceLoadingRequestor {
	return AssetResourceLoadingRequestor{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether the requestor provides expired session reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestor/providesExpiredSessionReports
func (a_ AssetResourceLoadingRequestor) ProvidesExpiredSessionReports() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("providesExpiredSessionReports"))
	return rv
}








