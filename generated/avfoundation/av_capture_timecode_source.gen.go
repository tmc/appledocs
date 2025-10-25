// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureTimecodeSource */


/* debug [class_header]: Header for AVCaptureTimecodeSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureTimecodeSource */
// An interface definition for the [CaptureTimecodeSource] class.
type ICaptureTimecodeSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureTimecodeSource */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	Type() CaptureTimecodeSourceType
	Uuid() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureTimecodeSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureTimecodeSource */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureTimecodeSource */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureTimecodeSource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureTimecodeSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureTimecodeSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureTimecodeSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureTimecodeSource */

// The name of the timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/displayName
func (c_ CaptureTimecodeSource) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// The type of timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/type
func (c_ CaptureTimecodeSource) Type() CaptureTimecodeSourceType {
	rv := objc.Send[CaptureTimecodeSourceType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A unique identifier for the timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/Source/uuid
func (c_ CaptureTimecodeSource) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("uuid"))
	return rv
}/* debug [instance_properties/getter]: uuid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureTimecodeSource */



