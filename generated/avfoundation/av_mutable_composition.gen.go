// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVMutableComposition] class.
var (
	aVMutableCompositionClass     _AVMutableCompositionClass
	aVMutableCompositionClassOnce sync.Once
)

func getAVMutableCompositionClass() _AVMutableCompositionClass {
	aVMutableCompositionClassOnce.Do(func() {
		aVMutableCompositionClass = _AVMutableCompositionClass{objc.GetClass("AVMutableComposition")}
	})
	return aVMutableCompositionClass
}

type _AVMutableCompositionClass struct {
	class objc.Class
}

// An interface definition for the [AVMutableComposition] class.
type IAVMutableComposition interface {
	IAVComposition
}

// An object that you use to create a new composition from existing assets.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition
type AVMutableComposition struct {
	AVComposition
}

// AVMutableCompositionFrom constructs a [AVMutableComposition] from an unsafe.Pointer.
//
// An object that you use to create a new composition from existing assets.
func AVMutableCompositionFrom(ptr unsafe.Pointer) AVMutableComposition {
	return AVMutableComposition{
		AVComposition: AVCompositionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVMutableCompositionClass) Alloc() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVMutableCompositionClass) New() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVMutableComposition) Init() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVMutableComposition) Autorelease() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableComposition creates a new AVMutableComposition instance.
func NewAVMutableComposition() AVMutableComposition {
	return getAVMutableCompositionClass().New()
}




