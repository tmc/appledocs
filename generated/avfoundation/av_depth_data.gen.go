// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVDepthData] class.
var (
	aVDepthDataClass     _AVDepthDataClass
	aVDepthDataClassOnce sync.Once
)

func getAVDepthDataClass() _AVDepthDataClass {
	aVDepthDataClassOnce.Do(func() {
		aVDepthDataClass = _AVDepthDataClass{objc.GetClass("AVDepthData")}
	})
	return aVDepthDataClass
}

type _AVDepthDataClass struct {
	class objc.Class
}

// An interface definition for the [AVDepthData] class.
type IAVDepthData interface {
	objectivec.IObject
}

// A container for per-pixel distance or disparity information captured by compatible camera devices.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData
type AVDepthData struct {
	objectivec.Object
}

// AVDepthDataFrom constructs a [AVDepthData] from an unsafe.Pointer.
//
// A container for per-pixel distance or disparity information captured by compatible camera devices.
func AVDepthDataFrom(ptr unsafe.Pointer) AVDepthData {
	return AVDepthData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVDepthDataClass) Alloc() AVDepthData {
	rv := objc.Send[AVDepthData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVDepthDataClass) New() AVDepthData {
	rv := objc.Send[AVDepthData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVDepthData) Init() AVDepthData {
	rv := objc.Send[AVDepthData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVDepthData) Autorelease() AVDepthData {
	rv := objc.Send[AVDepthData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDepthData creates a new AVDepthData instance.
func NewAVDepthData() AVDepthData {
	return getAVDepthDataClass().New()
}




