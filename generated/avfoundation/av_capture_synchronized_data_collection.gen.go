// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureSynchronizedDataCollection] class.
var (
	CaptureSynchronizedDataCollectionClass     _CaptureSynchronizedDataCollectionClass
	CaptureSynchronizedDataCollectionClassOnce sync.Once
)

func getCaptureSynchronizedDataCollectionClass() _CaptureSynchronizedDataCollectionClass {
	CaptureSynchronizedDataCollectionClassOnce.Do(func() {
		CaptureSynchronizedDataCollectionClass = _CaptureSynchronizedDataCollectionClass{objc.GetClass("AVCaptureSynchronizedDataCollection")}
	})
	return CaptureSynchronizedDataCollectionClass
}

type _CaptureSynchronizedDataCollectionClass struct {
	class objc.Class
}





// An interface definition for the [CaptureSynchronizedDataCollection] class.
type ICaptureSynchronizedDataCollection interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSynchronizedDataCollectionClass) Alloc() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSynchronizedDataCollectionClass) New() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSynchronizedDataCollection) Init() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSynchronizedDataCollection) Autorelease() CaptureSynchronizedDataCollection {
	rv := objc.Send[CaptureSynchronizedDataCollection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSynchronizedDataCollection creates a new CaptureSynchronizedDataCollection instance.
func NewCaptureSynchronizedDataCollection() CaptureSynchronizedDataCollection {
	return getCaptureSynchronizedDataCollectionClass().New()
}





// A set of data samples collected simultaneously from multiple capture outputs.


// A set of data samples collected simultaneously from multiple capture outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDataCollection
type CaptureSynchronizedDataCollection struct {
	objectivec.Object
}

// CaptureSynchronizedDataCollectionFrom constructs a [CaptureSynchronizedDataCollection] from an unsafe.Pointer.
//
// A set of data samples collected simultaneously from multiple capture outputs.
func CaptureSynchronizedDataCollectionFrom(ptr unsafe.Pointer) CaptureSynchronizedDataCollection {
	return CaptureSynchronizedDataCollection{objectivec.Object{objc.ID(ptr)}}
}






























