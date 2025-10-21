// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ShareableContentInfo] class.
var (
	ShareableContentInfoClass     _ShareableContentInfoClass
	ShareableContentInfoClassOnce sync.Once
)

func getShareableContentInfoClass() _ShareableContentInfoClass {
	ShareableContentInfoClassOnce.Do(func() {
		ShareableContentInfoClass = _ShareableContentInfoClass{objc.GetClass("SCShareableContentInfo")}
	})
	return ShareableContentInfoClass
}

type _ShareableContentInfoClass struct {
	class objc.Class
}

// An interface definition for the [ShareableContentInfo] class.
type IShareableContentInfo interface {
	objectivec.IObject
}

// An instance that provides information for the content in a given stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo
type ShareableContentInfo struct {
	objectivec.Object
}

// ShareableContentInfoFrom constructs a [ShareableContentInfo] from an unsafe.Pointer.
//
// An instance that provides information for the content in a given stream.
func ShareableContentInfoFrom(ptr unsafe.Pointer) ShareableContentInfo {
	return ShareableContentInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ShareableContentInfoClass) Alloc() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ShareableContentInfoClass) New() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ShareableContentInfo) Init() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ShareableContentInfo) Autorelease() ShareableContentInfo {
	rv := objc.Send[ShareableContentInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShareableContentInfo creates a new ShareableContentInfo instance.
func NewShareableContentInfo() ShareableContentInfo {
	return getShareableContentInfoClass().New()
}


// The size and location of content for the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/contentRect
func (s_ ShareableContentInfo) ContentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("contentRect"))
	return rv
}

// The current presentation style of the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo/style
func (s_ ShareableContentInfo) Style() ShareableContentStyle {
	rv := objc.Send[ShareableContentStyle](s_.ID, objc.Sel("style"))
	return rv
}

// The scaling from points to output pixel resolution for the stream.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scshareablecontentinfo/pointpixelscale
func (s_ ShareableContentInfo) PointPixelScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pointPixelScale"))
	return rv
}


// SetPointPixelScale sets the value of the pointPixelScale property.
// The scaling from points to output pixel resolution for the stream.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scshareablecontentinfo/pointpixelscale
func (s_ ShareableContentInfo) SetPointPixelScale(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPointPixelScale:"), value)
}



