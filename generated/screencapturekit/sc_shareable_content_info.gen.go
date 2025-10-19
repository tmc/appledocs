// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCShareableContentInfo] class.
var (
	sCShareableContentInfoClass     _SCShareableContentInfoClass
	sCShareableContentInfoClassOnce sync.Once
)

func getSCShareableContentInfoClass() _SCShareableContentInfoClass {
	sCShareableContentInfoClassOnce.Do(func() {
		sCShareableContentInfoClass = _SCShareableContentInfoClass{objc.GetClass("SCShareableContentInfo")}
	})
	return sCShareableContentInfoClass
}

type _SCShareableContentInfoClass struct {
	class objc.Class
}

// An interface definition for the [SCShareableContentInfo] class.
type ISCShareableContentInfo interface {
	objectivec.IObject
}

// An instance that provides information for the content in a given stream.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentInfo
type SCShareableContentInfo struct {
	objectivec.Object
}

// SCShareableContentInfoFrom constructs a [SCShareableContentInfo] from an unsafe.Pointer.
//
// An instance that provides information for the content in a given stream.
func SCShareableContentInfoFrom(ptr unsafe.Pointer) SCShareableContentInfo {
	return SCShareableContentInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCShareableContentInfoClass) Alloc() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCShareableContentInfoClass) New() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCShareableContentInfo) Init() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCShareableContentInfo) Autorelease() SCShareableContentInfo {
	rv := objc.Send[SCShareableContentInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCShareableContentInfo creates a new SCShareableContentInfo instance.
func NewSCShareableContentInfo() SCShareableContentInfo {
	return getSCShareableContentInfoClass().New()
}




