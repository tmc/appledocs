// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureSynchronizedData] class.
var (
	aVCaptureSynchronizedDataClass     _AVCaptureSynchronizedDataClass
	aVCaptureSynchronizedDataClassOnce sync.Once
)

func getAVCaptureSynchronizedDataClass() _AVCaptureSynchronizedDataClass {
	aVCaptureSynchronizedDataClassOnce.Do(func() {
		aVCaptureSynchronizedDataClass = _AVCaptureSynchronizedDataClass{objc.GetClass("AVCaptureSynchronizedData")}
	})
	return aVCaptureSynchronizedDataClass
}

type _AVCaptureSynchronizedDataClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureSynchronizedData] class.
type IAVCaptureSynchronizedData interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]
type AVCaptureSynchronizedData struct {
	objectivec.Object
}

// AVCaptureSynchronizedDataFrom constructs a [AVCaptureSynchronizedData] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVCaptureSynchronizedDataFrom(ptr unsafe.Pointer) AVCaptureSynchronizedData {
	return AVCaptureSynchronizedData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureSynchronizedDataClass) Alloc() AVCaptureSynchronizedData {
	rv := objc.Send[AVCaptureSynchronizedData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureSynchronizedDataClass) New() AVCaptureSynchronizedData {
	rv := objc.Send[AVCaptureSynchronizedData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureSynchronizedData) Init() AVCaptureSynchronizedData {
	rv := objc.Send[AVCaptureSynchronizedData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureSynchronizedData) Autorelease() AVCaptureSynchronizedData {
	rv := objc.Send[AVCaptureSynchronizedData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSynchronizedData creates a new AVCaptureSynchronizedData instance.
func NewAVCaptureSynchronizedData() AVCaptureSynchronizedData {
	return getAVCaptureSynchronizedDataClass().New()
}




