// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVMediaSelection] class.
var (
	aVMediaSelectionClass     _AVMediaSelectionClass
	aVMediaSelectionClassOnce sync.Once
)

func getAVMediaSelectionClass() _AVMediaSelectionClass {
	aVMediaSelectionClassOnce.Do(func() {
		aVMediaSelectionClass = _AVMediaSelectionClass{objc.GetClass("AVMediaSelection")}
	})
	return aVMediaSelectionClass
}

type _AVMediaSelectionClass struct {
	class objc.Class
}

// An interface definition for the [AVMediaSelection] class.
type IAVMediaSelection interface {
	objectivec.IObject
}

// An object that represents a complete rendition of media selection options on an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection
type AVMediaSelection struct {
	objectivec.Object
}

// AVMediaSelectionFrom constructs a [AVMediaSelection] from an unsafe.Pointer.
//
// An object that represents a complete rendition of media selection options on an asset.
func AVMediaSelectionFrom(ptr unsafe.Pointer) AVMediaSelection {
	return AVMediaSelection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVMediaSelectionClass) Alloc() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVMediaSelectionClass) New() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVMediaSelection) Init() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVMediaSelection) Autorelease() AVMediaSelection {
	rv := objc.Send[AVMediaSelection](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaSelection creates a new AVMediaSelection instance.
func NewAVMediaSelection() AVMediaSelection {
	return getAVMediaSelectionClass().New()
}




