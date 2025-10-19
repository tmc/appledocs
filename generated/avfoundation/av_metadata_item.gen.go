// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVMetadataItem] class.
var aVMetadataItemClass = _AVMetadataItemClass{objc.GetClass("AVMetadataItem")}

type _AVMetadataItemClass struct {
	class objc.Class
}

// An interface definition for the [AVMetadataItem] class.
type IAVMetadataItem interface {
	objectivec.IObject
}

// A metadata item for an audiovisual asset or one of its tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem

type AVMetadataItem struct {
	objectivec.Object
}

// AVMetadataItemFrom constructs a [AVMetadataItem] from an unsafe.Pointer.
//
// A metadata item for an audiovisual asset or one of its tracks.
func AVMetadataItemFrom(ptr unsafe.Pointer) AVMetadataItem {
	return AVMetadataItem{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVMetadataItemClass) Alloc() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVMetadataItemClass) New() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVMetadataItem) Init() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVMetadataItem) Autorelease() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataItem creates a new AVMetadataItem instance.
func NewAVMetadataItem() AVMetadataItem {
	return aVMetadataItemClass.New()
}




