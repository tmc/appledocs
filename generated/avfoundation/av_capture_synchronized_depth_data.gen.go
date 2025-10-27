// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CaptureSynchronizedDepthData] class.
var (
	CaptureSynchronizedDepthDataClass     _CaptureSynchronizedDepthDataClass
	CaptureSynchronizedDepthDataClassOnce sync.Once
)

func getCaptureSynchronizedDepthDataClass() _CaptureSynchronizedDepthDataClass {
	CaptureSynchronizedDepthDataClassOnce.Do(func() {
		CaptureSynchronizedDepthDataClass = _CaptureSynchronizedDepthDataClass{objc.GetClass("AVCaptureSynchronizedDepthData")}
	})
	return CaptureSynchronizedDepthDataClass
}

type _CaptureSynchronizedDepthDataClass struct {
	class objc.Class
}





// An interface definition for the [CaptureSynchronizedDepthData] class.
type ICaptureSynchronizedDepthData interface {
	ICaptureSynchronizedData
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedDepthDataClass) Alloc() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedDepthDataClass) New() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedDepthData) Init() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedDepthData) Autorelease() CaptureSynchronizedDepthData {
	rv := objc.Send[CaptureSynchronizedDepthData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedDepthData creates a new CaptureSynchronizedDepthData instance.
func NewCaptureSynchronizedDepthData() CaptureSynchronizedDepthData {
	return getCaptureSynchronizedDepthDataClass().New()
}





// A container for scene depth information collected using synchronized capture.


// A container for scene depth information collected using synchronized capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDepthData
type CaptureSynchronizedDepthData struct {
	CaptureSynchronizedData
}

// CaptureSynchronizedDepthDataFrom constructs a [CaptureSynchronizedDepthData] from an unsafe.Pointer.
//
// A container for scene depth information collected using synchronized capture.
func CaptureSynchronizedDepthDataFrom(ptr unsafe.Pointer) CaptureSynchronizedDepthData {
	return CaptureSynchronizedDepthData{
		CaptureSynchronizedData: CaptureSynchronizedDataFrom(ptr),
	}
}






























