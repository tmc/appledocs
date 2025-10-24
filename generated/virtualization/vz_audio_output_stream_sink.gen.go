// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZAudioOutputStreamSink */


/* debug [class_header]: Header for VZAudioOutputStreamSink */
// The class instance for the [VZAudioOutputStreamSink] class.
var (
	VZAudioOutputStreamSinkClass     _VZAudioOutputStreamSinkClass
	VZAudioOutputStreamSinkClassOnce sync.Once
)

func getVZAudioOutputStreamSinkClass() _VZAudioOutputStreamSinkClass {
	VZAudioOutputStreamSinkClassOnce.Do(func() {
		VZAudioOutputStreamSinkClass = _VZAudioOutputStreamSinkClass{objc.GetClass("VZAudioOutputStreamSink")}
	})
	return VZAudioOutputStreamSinkClass
}

type _VZAudioOutputStreamSinkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZAudioOutputStreamSink */
// An interface definition for the [VZAudioOutputStreamSink] class.
type IVZAudioOutputStreamSink interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZAudioOutputStreamSink */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZAudioOutputStreamSink */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZAudioOutputStreamSink */
// Alloc allocates a new instance without initialization.
func (vc _VZAudioOutputStreamSinkClass) Alloc() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZAudioOutputStreamSinkClass) New() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioOutputStreamSink) Init() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioOutputStreamSink) Autorelease() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioOutputStreamSink creates a new VZAudioOutputStreamSink instance.
func NewVZAudioOutputStreamSink() VZAudioOutputStreamSink {
	return getVZAudioOutputStreamSinkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZAudioOutputStreamSink */
// The base class for an audio output stream sink.
//
// An audio output stream sink defines how the host system consumes audio data from a guest. Don’t instantiate directly, use one of its subclasses, such as instead.


// The base class for an audio output stream sink.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type VZAudioOutputStreamSink struct {
	objectivec.Object
}

// VZAudioOutputStreamSinkFrom constructs a [VZAudioOutputStreamSink] from an unsafe.Pointer.
//
// The base class for an audio output stream sink.
func VZAudioOutputStreamSinkFrom(ptr unsafe.Pointer) VZAudioOutputStreamSink {
	return VZAudioOutputStreamSink{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZAudioOutputStreamSink *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZAudioOutputStreamSink */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZAudioOutputStreamSink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZAudioOutputStreamSink */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZAudioOutputStreamSink */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZAudioOutputStreamSink */



