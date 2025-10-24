// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEDecodeFrameOptions */


/* debug [class_header]: Header for MEDecodeFrameOptions */
// The class instance for the [MEDecodeFrameOptions] class.
var (
	MEDecodeFrameOptionsClass     _MEDecodeFrameOptionsClass
	MEDecodeFrameOptionsClassOnce sync.Once
)

func getMEDecodeFrameOptionsClass() _MEDecodeFrameOptionsClass {
	MEDecodeFrameOptionsClassOnce.Do(func() {
		MEDecodeFrameOptionsClass = _MEDecodeFrameOptionsClass{objc.GetClass("MEDecodeFrameOptions")}
	})
	return MEDecodeFrameOptionsClass
}

type _MEDecodeFrameOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEDecodeFrameOptions */
// An interface definition for the [MEDecodeFrameOptions] class.
type IMEDecodeFrameOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEDecodeFrameOptions */
	// properties:
	DoNotOutputFrame() bool
	SetDoNotOutputFrame(value bool)
	RealTimePlayback() bool
	SetRealTimePlayback(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEDecodeFrameOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEDecodeFrameOptions */
// Alloc allocates a new instance without initialization.
func (mc _MEDecodeFrameOptionsClass) Alloc() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEDecodeFrameOptionsClass) New() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEDecodeFrameOptions) Init() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEDecodeFrameOptions) Autorelease() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodeFrameOptions creates a new MEDecodeFrameOptions instance.
func NewMEDecodeFrameOptions() MEDecodeFrameOptions {
	return getMEDecodeFrameOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEDecodeFrameOptions */
// An object that guides the video decoder operation on a per-frame basis.


// An object that guides the video decoder operation on a per-frame basis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions
type MEDecodeFrameOptions struct {
	objectivec.Object
}

// MEDecodeFrameOptionsFrom constructs a [MEDecodeFrameOptions] from an unsafe.Pointer.
//
// An object that guides the video decoder operation on a per-frame basis.
func MEDecodeFrameOptionsFrom(ptr unsafe.Pointer) MEDecodeFrameOptions {
	return MEDecodeFrameOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEDecodeFrameOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEDecodeFrameOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEDecodeFrameOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEDecodeFrameOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEDecodeFrameOptions */

// A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/doNotOutputFrame
func (m_ MEDecodeFrameOptions) DoNotOutputFrame() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("doNotOutputFrame"))
	return rv
}/* debug [instance_properties/getter]: doNotOutputFrame */


// A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/doNotOutputFrame
func (m_ MEDecodeFrameOptions) SetDoNotOutputFrame(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoNotOutputFrame:"), value)
}/* debug [instance_properties/setter]: doNotOutputFrame */


// A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/realTimePlayback
func (m_ MEDecodeFrameOptions) RealTimePlayback() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("realTimePlayback"))
	return rv
}/* debug [instance_properties/getter]: realTimePlayback */


// A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/realTimePlayback
func (m_ MEDecodeFrameOptions) SetRealTimePlayback(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRealTimePlayback:"), value)
}/* debug [instance_properties/setter]: realTimePlayback */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEDecodeFrameOptions */



