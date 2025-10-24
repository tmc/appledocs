// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureTimecodeSource] class.
var (
	CaptureTimecodeSourceClass     _CaptureTimecodeSourceClass
	CaptureTimecodeSourceClassOnce sync.Once
)

func getCaptureTimecodeSourceClass() _CaptureTimecodeSourceClass {
	CaptureTimecodeSourceClassOnce.Do(func() {
		CaptureTimecodeSourceClass = _CaptureTimecodeSourceClass{objc.GetClass("AVCaptureTimecodeSource")}
	})
	return CaptureTimecodeSourceClass
}

type _CaptureTimecodeSourceClass struct {
	class objc.Class
}





// An interface definition for the [CaptureTimecodeSource] class.
type ICaptureTimecodeSource interface {
	objectivec.IObject
	

	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	Type() CaptureTimecodeSourceType
	Uuid() foundation.UUID


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureTimecodeSourceClass) Alloc() CaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureTimecodeSourceClass) New() CaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureTimecodeSource) Init() CaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureTimecodeSource) Autorelease() CaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureTimecodeSource creates a new CaptureTimecodeSource instance.
func NewCaptureTimecodeSource() CaptureTimecodeSource {
	return getCaptureTimecodeSourceClass().New()
}





// Describes a timecode source that a timecode generator can synchronize to.
//
// provides information about a specific timecode source available for synchronization in . It includes metadata such as the source’s name, type, and unique identifier.


// Describes a timecode source that a timecode generator can synchronize to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source
type CaptureTimecodeSource struct {
	objectivec.Object
}

// CaptureTimecodeSourceFrom constructs a [CaptureTimecodeSource] from an unsafe.Pointer.
//
// Describes a timecode source that a timecode generator can synchronize to.
func CaptureTimecodeSourceFrom(ptr unsafe.Pointer) CaptureTimecodeSource {
	return CaptureTimecodeSource{objectivec.Object{objc.ID(ptr)}}
}

























// The name of the timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/displayName
func (c_ CaptureTimecodeSource) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("displayName"))
	return rv
}


// The type of timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/type
func (c_ CaptureTimecodeSource) Type() CaptureTimecodeSourceType {
	rv := objc.Send[CaptureTimecodeSourceType](c_.ID, objc.Sel("type"))
	return rv
}


// A unique identifier for the timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/uuid
func (c_ CaptureTimecodeSource) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("uuid"))
	return rv
}








