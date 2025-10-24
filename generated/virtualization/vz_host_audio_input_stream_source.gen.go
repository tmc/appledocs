// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZHostAudioInputStreamSource */


/* debug [class_header]: Header for VZHostAudioInputStreamSource */
// The class instance for the [VZHostAudioInputStreamSource] class.
var (
	VZHostAudioInputStreamSourceClass     _VZHostAudioInputStreamSourceClass
	VZHostAudioInputStreamSourceClassOnce sync.Once
)

func getVZHostAudioInputStreamSourceClass() _VZHostAudioInputStreamSourceClass {
	VZHostAudioInputStreamSourceClassOnce.Do(func() {
		VZHostAudioInputStreamSourceClass = _VZHostAudioInputStreamSourceClass{objc.GetClass("VZHostAudioInputStreamSource")}
	})
	return VZHostAudioInputStreamSourceClass
}

type _VZHostAudioInputStreamSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZHostAudioInputStreamSource */
// An interface definition for the [VZHostAudioInputStreamSource] class.
type IVZHostAudioInputStreamSource interface {
	IVZAudioInputStreamSource
	
/* debug [class_interface_properties]: Properties for VZHostAudioInputStreamSource */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZHostAudioInputStreamSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZHostAudioInputStreamSource */
// Alloc allocates a new instance without initialization.
func (vc _VZHostAudioInputStreamSourceClass) Alloc() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZHostAudioInputStreamSourceClass) New() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZHostAudioInputStreamSource) Init() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZHostAudioInputStreamSource) Autorelease() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioInputStreamSource creates a new VZHostAudioInputStreamSource instance.
func NewVZHostAudioInputStreamSource() VZHostAudioInputStreamSource {
	return getVZHostAudioInputStreamSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZHostAudioInputStreamSource */
// The host audio input stream source that provides audio from the host system’s default input device.
//
// The host input data comes from the same device that uses.


// The host audio input stream source that provides audio from the host system’s default input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZHostAudioInputStreamSource
type VZHostAudioInputStreamSource struct {
	VZAudioInputStreamSource
}

// VZHostAudioInputStreamSourceFrom constructs a [VZHostAudioInputStreamSource] from an unsafe.Pointer.
//
// The host audio input stream source that provides audio from the host system’s default input device.
func VZHostAudioInputStreamSourceFrom(ptr unsafe.Pointer) VZHostAudioInputStreamSource {
	return VZHostAudioInputStreamSource{
		VZAudioInputStreamSource: VZAudioInputStreamSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZHostAudioInputStreamSource */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZHostAudioInputStreamSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZHostAudioInputStreamSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZHostAudioInputStreamSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZHostAudioInputStreamSource */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZHostAudioInputStreamSource */


