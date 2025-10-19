// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetReaderOutput] class.
var (
	aVAssetReaderOutputClass     _AVAssetReaderOutputClass
	aVAssetReaderOutputClassOnce sync.Once
)

func getAVAssetReaderOutputClass() _AVAssetReaderOutputClass {
	aVAssetReaderOutputClassOnce.Do(func() {
		aVAssetReaderOutputClass = _AVAssetReaderOutputClass{objc.GetClass("AVAssetReaderOutput")}
	})
	return aVAssetReaderOutputClass
}

type _AVAssetReaderOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetReaderOutput] class.
type IAVAssetReaderOutput interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]
type AVAssetReaderOutput struct {
	objectivec.Object
}

// AVAssetReaderOutputFrom constructs a [AVAssetReaderOutput] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVAssetReaderOutputFrom(ptr unsafe.Pointer) AVAssetReaderOutput {
	return AVAssetReaderOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetReaderOutputClass) Alloc() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetReaderOutputClass) New() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetReaderOutput) Init() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetReaderOutput) Autorelease() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderOutput creates a new AVAssetReaderOutput instance.
func NewAVAssetReaderOutput() AVAssetReaderOutput {
	return getAVAssetReaderOutputClass().New()
}




