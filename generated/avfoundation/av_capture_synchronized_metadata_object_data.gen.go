// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CaptureSynchronizedMetadataObjectData] class.
var (
	CaptureSynchronizedMetadataObjectDataClass     _CaptureSynchronizedMetadataObjectDataClass
	CaptureSynchronizedMetadataObjectDataClassOnce sync.Once
)

func getCaptureSynchronizedMetadataObjectDataClass() _CaptureSynchronizedMetadataObjectDataClass {
	CaptureSynchronizedMetadataObjectDataClassOnce.Do(func() {
		CaptureSynchronizedMetadataObjectDataClass = _CaptureSynchronizedMetadataObjectDataClass{objc.GetClass("AVCaptureSynchronizedMetadataObjectData")}
	})
	return CaptureSynchronizedMetadataObjectDataClass
}

type _CaptureSynchronizedMetadataObjectDataClass struct {
	class objc.Class
}





// An interface definition for the [CaptureSynchronizedMetadataObjectData] class.
type ICaptureSynchronizedMetadataObjectData interface {
	ICaptureSynchronizedData
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedMetadataObjectDataClass) Alloc() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedMetadataObjectDataClass) New() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedMetadataObjectData) Init() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedMetadataObjectData) Autorelease() CaptureSynchronizedMetadataObjectData {
	rv := objc.Send[CaptureSynchronizedMetadataObjectData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedMetadataObjectData creates a new CaptureSynchronizedMetadataObjectData instance.
func NewCaptureSynchronizedMetadataObjectData() CaptureSynchronizedMetadataObjectData {
	return getCaptureSynchronizedMetadataObjectDataClass().New()
}





// A container for metadata objects collected using synchronized capture.


// A container for metadata objects collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedMetadataObjectData
type CaptureSynchronizedMetadataObjectData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedMetadataObjectDataFrom constructs a [CaptureSynchronizedMetadataObjectData] from an unsafe.Pointer.
//
// A container for metadata objects collected using synchronized capture.
func CaptureSynchronizedMetadataObjectDataFrom(ptr unsafe.Pointer) CaptureSynchronizedMetadataObjectData {
	return CaptureSynchronizedMetadataObjectData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}






























