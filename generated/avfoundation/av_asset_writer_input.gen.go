// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetWriterInput] class.
var (
	aVAssetWriterInputClass     _AVAssetWriterInputClass
	aVAssetWriterInputClassOnce sync.Once
)

func getAVAssetWriterInputClass() _AVAssetWriterInputClass {
	aVAssetWriterInputClassOnce.Do(func() {
		aVAssetWriterInputClass = _AVAssetWriterInputClass{objc.GetClass("AVAssetWriterInput")}
	})
	return aVAssetWriterInputClass
}

type _AVAssetWriterInputClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetWriterInput] class.
type IAVAssetWriterInput interface {
	objectivec.IObject
}

// An object that appends media samples to a track in an asset writer’s output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput
type AVAssetWriterInput struct {
	objectivec.Object
}

// AVAssetWriterInputFrom constructs a [AVAssetWriterInput] from an unsafe.Pointer.
//
// An object that appends media samples to a track in an asset writer’s output file.
func AVAssetWriterInputFrom(ptr unsafe.Pointer) AVAssetWriterInput {
	return AVAssetWriterInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetWriterInputClass) Alloc() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetWriterInputClass) New() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetWriterInput) Init() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetWriterInput) Autorelease() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriterInput creates a new AVAssetWriterInput instance.
func NewAVAssetWriterInput() AVAssetWriterInput {
	return getAVAssetWriterInputClass().New()
}




