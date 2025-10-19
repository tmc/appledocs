// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaption] class.
var aVCaptionClass = _AVCaptionClass{objc.GetClass("AVCaption")}

type _AVCaptionClass struct {
	class objc.Class
}

// An interface definition for the [AVCaption] class.
type IAVCaption interface {
	objectivec.IObject
}

// An object that represents text to present over a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption

type AVCaption struct {
	objectivec.Object
}

// AVCaptionFrom constructs a [AVCaption] from an unsafe.Pointer.
//
// An object that represents text to present over a time range.
func AVCaptionFrom(ptr unsafe.Pointer) AVCaption {
	return AVCaption{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptionClass) Alloc() AVCaption {
	rv := objc.Send[AVCaption](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptionClass) New() AVCaption {
	rv := objc.Send[AVCaption](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaption) Init() AVCaption {
	rv := objc.Send[AVCaption](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaption) Autorelease() AVCaption {
	rv := objc.Send[AVCaption](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaption creates a new AVCaption instance.
func NewAVCaption() AVCaption {
	return aVCaptionClass.New()
}




