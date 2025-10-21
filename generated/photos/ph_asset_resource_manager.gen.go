// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHAssetResourceManager] class.
var (
	PHAssetResourceManagerClass     _PHAssetResourceManagerClass
	PHAssetResourceManagerClassOnce sync.Once
)

func getPHAssetResourceManagerClass() _PHAssetResourceManagerClass {
	PHAssetResourceManagerClassOnce.Do(func() {
		PHAssetResourceManagerClass = _PHAssetResourceManagerClass{objc.GetClass("PHAssetResourceManager")}
	})
	return PHAssetResourceManagerClass
}

type _PHAssetResourceManagerClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResourceManager] class.
type IPHAssetResourceManager interface {
	objectivec.IObject
	CancelDataRequest(requestID unsafe.Pointer)
	RequestDataForAssetResourceOptionsDataReceivedHandlerCompletionHandler(resource unsafe.Pointer, options unsafe.Pointer, handler unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer
	WriteDataForAssetResourceToFileOptionsCompletionHandler(resource unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
}

// A resource manager for the data storage underlying a Photos asset.
//
// An asset can have multiple underlying data resources—for example, both original and edited versions—each of which is represented by a object. Unlike the class, which provides and caches the primary representations of assets as thumbnails, image objects, or video objects, the asset resource manager provides direct access to these underlying data resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceManager
type PHAssetResourceManager struct {
	objectivec.Object
}

// PHAssetResourceManagerFrom constructs a [PHAssetResourceManager] from an unsafe.Pointer.
//
// A resource manager for the data storage underlying a Photos asset.
func PHAssetResourceManagerFrom(ptr unsafe.Pointer) PHAssetResourceManager {
	return PHAssetResourceManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceManagerClass) Alloc() PHAssetResourceManager {
	rv := objc.Send[PHAssetResourceManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceManagerClass) New() PHAssetResourceManager {
	rv := objc.Send[PHAssetResourceManager](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResourceManager) Init() PHAssetResourceManager {
	rv := objc.Send[PHAssetResourceManager](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResourceManager) Autorelease() PHAssetResourceManager {
	rv := objc.Send[PHAssetResourceManager](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResourceManager creates a new PHAssetResourceManager instance.
func NewPHAssetResourceManager() PHAssetResourceManager {
	return getPHAssetResourceManagerClass().New()
}


// Returns the shared asset resource manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceManager/default()
func (pc _PHAssetResourceManagerClass) DefaultManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("defaultManager"))
	return rv
}

// Cancels an asynchronous request.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceManager/cancelDataRequest(_:)
func (p_ PHAssetResourceManager) CancelDataRequest(requestID unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelDataRequest:"), requestID)
}

// Requests the underlying data for the specified asset resource, to be delivered asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceManager/requestData(for:options:dataReceivedHandler:completionHandler:)
func (p_ PHAssetResourceManager) RequestDataForAssetResourceOptionsDataReceivedHandlerCompletionHandler(resource unsafe.Pointer, options unsafe.Pointer, handler unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("requestDataForAssetResource:options:dataReceivedHandler:completionHandler:"), resource, options, handler, completionHandler)
	return rv
}

// Requests the underlying data for the specified asset resource, to be asynchronously written to a local file.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceManager/writeData(for:toFile:options:completionHandler:)
func (p_ PHAssetResourceManager) WriteDataForAssetResourceToFileOptionsCompletionHandler(resource unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("writeDataForAssetResource:toFile:options:completionHandler:"), resource, fileURL, options, completionHandler)
}



