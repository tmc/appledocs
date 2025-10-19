// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVComposition] class.
var aVCompositionClass = _AVCompositionClass{objc.GetClass("AVComposition")}

type _AVCompositionClass struct {
	class objc.Class
}

// An interface definition for the [AVComposition] class.
type IAVComposition interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]

type AVComposition struct {
	objectivec.Object
}

// AVCompositionFrom constructs a [AVComposition] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVCompositionFrom(ptr unsafe.Pointer) AVComposition {
	return AVComposition{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCompositionClass) Alloc() AVComposition {
	rv := objc.Send[AVComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCompositionClass) New() AVComposition {
	rv := objc.Send[AVComposition](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVComposition) Init() AVComposition {
	rv := objc.Send[AVComposition](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVComposition) Autorelease() AVComposition {
	rv := objc.Send[AVComposition](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVComposition creates a new AVComposition instance.
func NewAVComposition() AVComposition {
	return aVCompositionClass.New()
}




