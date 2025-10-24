// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZHostAudioOutputStreamSink */


/* debug [class_header]: Header for VZHostAudioOutputStreamSink */
// The class instance for the [VZHostAudioOutputStreamSink] class.
var (
	VZHostAudioOutputStreamSinkClass     _VZHostAudioOutputStreamSinkClass
	VZHostAudioOutputStreamSinkClassOnce sync.Once
)

func getVZHostAudioOutputStreamSinkClass() _VZHostAudioOutputStreamSinkClass {
	VZHostAudioOutputStreamSinkClassOnce.Do(func() {
		VZHostAudioOutputStreamSinkClass = _VZHostAudioOutputStreamSinkClass{objc.GetClass("VZHostAudioOutputStreamSink")}
	})
	return VZHostAudioOutputStreamSinkClass
}

type _VZHostAudioOutputStreamSinkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZHostAudioOutputStreamSink */
// An interface definition for the [VZHostAudioOutputStreamSink] class.
type IVZHostAudioOutputStreamSink interface {
	IVZAudioOutputStreamSink
	
/* debug [class_interface_properties]: Properties for VZHostAudioOutputStreamSink */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZHostAudioOutputStreamSink */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZHostAudioOutputStreamSink */
// Alloc allocates a new instance without initialization.
func (vc _VZHostAudioOutputStreamSinkClass) Alloc() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZHostAudioOutputStreamSinkClass) New() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZHostAudioOutputStreamSink) Init() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZHostAudioOutputStreamSink) Autorelease() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioOutputStreamSink creates a new VZHostAudioOutputStreamSink instance.
func NewVZHostAudioOutputStreamSink() VZHostAudioOutputStreamSink {
	return getVZHostAudioOutputStreamSinkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZHostAudioOutputStreamSink */
// Host audio output stream sink plays audio to the host system’s default output device.
//
// Host output data goes to the same device that uses.


// Host audio output stream sink plays audio to the host system’s default output device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZHostAudioOutputStreamSink
type VZHostAudioOutputStreamSink struct {
	VZAudioOutputStreamSink
}

// VZHostAudioOutputStreamSinkFrom constructs a [VZHostAudioOutputStreamSink] from an unsafe.Pointer.
//
// Host audio output stream sink plays audio to the host system’s default output device.
func VZHostAudioOutputStreamSinkFrom(ptr unsafe.Pointer) VZHostAudioOutputStreamSink {
	return VZHostAudioOutputStreamSink{
		VZAudioOutputStreamSink: VZAudioOutputStreamSinkFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZHostAudioOutputStreamSink */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZHostAudioOutputStreamSink */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZHostAudioOutputStreamSink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZHostAudioOutputStreamSink */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZHostAudioOutputStreamSink */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZHostAudioOutputStreamSink */


