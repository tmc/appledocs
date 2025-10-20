// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PictureInPictureControllerContentSource] class.
var (
	PictureInPictureControllerContentSourceClass     _PictureInPictureControllerContentSourceClass
	PictureInPictureControllerContentSourceClassOnce sync.Once
)

func getPictureInPictureControllerContentSourceClass() _PictureInPictureControllerContentSourceClass {
	PictureInPictureControllerContentSourceClassOnce.Do(func() {
		PictureInPictureControllerContentSourceClass = _PictureInPictureControllerContentSourceClass{objc.GetClass("AVPictureInPictureControllerContentSource")}
	})
	return PictureInPictureControllerContentSourceClass
}

type _PictureInPictureControllerContentSourceClass struct {
	class objc.Class
}

// An interface definition for the [PictureInPictureControllerContentSource] class.
type IPictureInPictureControllerContentSource interface {
	objectivec.IObject
}

// An object that represents the source of the content to present in Picture in Picture.
//
// The system supports displaying content from an or in a Picture in Picture window. Use an instance of this class to describe the source of your app’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class
type PictureInPictureControllerContentSource struct {
	objectivec.Object
}

// PictureInPictureControllerContentSourceFrom constructs a [PictureInPictureControllerContentSource] from an unsafe.Pointer.
//
// An object that represents the source of the content to present in Picture in Picture.
func PictureInPictureControllerContentSourceFrom(ptr unsafe.Pointer) PictureInPictureControllerContentSource {
	return PictureInPictureControllerContentSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureControllerContentSourceClass) Alloc() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PictureInPictureControllerContentSourceClass) New() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PictureInPictureControllerContentSource) Init() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PictureInPictureControllerContentSource) Autorelease() PictureInPictureControllerContentSource {
	rv := objc.Send[PictureInPictureControllerContentSource](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPictureInPictureControllerContentSource creates a new PictureInPictureControllerContentSource instance.
func NewPictureInPictureControllerContentSource() PictureInPictureControllerContentSource {
	return getPictureInPictureControllerContentSourceClass().New()
}


// The view that contains the video content of the call.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureController/ContentSource-swift.class/activeVideoCallSourceView
func (p_ PictureInPictureControllerContentSource) ActiveVideoCallSourceView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("activeVideoCallSourceView"))
	return rv
}



