// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MetadataItemValueRequest] class.
var (
	MetadataItemValueRequestClass     _MetadataItemValueRequestClass
	MetadataItemValueRequestClassOnce sync.Once
)

func getMetadataItemValueRequestClass() _MetadataItemValueRequestClass {
	MetadataItemValueRequestClassOnce.Do(func() {
		MetadataItemValueRequestClass = _MetadataItemValueRequestClass{objc.GetClass("AVMetadataItemValueRequest")}
	})
	return MetadataItemValueRequestClass
}

type _MetadataItemValueRequestClass struct {
	class objc.Class
}





// An interface definition for the [MetadataItemValueRequest] class.
type IMetadataItemValueRequest interface {
	objectivec.IObject
	

	// properties:
	MetadataItem() IAVMetadataItem


	

	// methods:
	RespondWithError(error_ foundation.foundation.INSError)
	RespondWithValue(value unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataItemValueRequestClass) Alloc() MetadataItemValueRequest {
	rv := objc.Send[MetadataItemValueRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataItemValueRequestClass) New() MetadataItemValueRequest {
	rv := objc.Send[MetadataItemValueRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataItemValueRequest) Init() MetadataItemValueRequest {
	rv := objc.Send[MetadataItemValueRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataItemValueRequest) Autorelease() MetadataItemValueRequest {
	rv := objc.Send[MetadataItemValueRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataItemValueRequest creates a new MetadataItemValueRequest instance.
func NewMetadataItemValueRequest() MetadataItemValueRequest {
	return getMetadataItemValueRequestClass().New()
}





// An object that responds to a request to load the value of a metadata item.


// An object that responds to a request to load the value of a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest
type MetadataItemValueRequest struct {
	objectivec.Object
}

// MetadataItemValueRequestFrom constructs a [MetadataItemValueRequest] from an unsafe.Pointer.
//
// An object that responds to a request to load the value of a metadata item.
func MetadataItemValueRequestFrom(ptr unsafe.Pointer) MetadataItemValueRequest {
	return MetadataItemValueRequest{objectivec.Object{objc.ID(ptr)}}
}




















// Returns an error when the system fails to load the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/respond(error:)
func (m_ MetadataItemValueRequest) RespondWithError(error_ foundation.foundation.INSError) {
	objc.Send[objc.ID](m_.ID, objc.Sel("respondWithError:"), error_)
}


// Returns the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/respond(value:)
func (m_ MetadataItemValueRequest) RespondWithValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("respondWithValue:"), value)
}







// The metadata item to request a value for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/metadataItem
func (m_ MetadataItemValueRequest) MetadataItem() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("metadataItem"))
	return rv
}








