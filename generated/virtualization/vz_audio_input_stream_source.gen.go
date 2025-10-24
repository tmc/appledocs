// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZAudioInputStreamSource */

/* debug [class_header]: Header for VZAudioInputStreamSource */
// The class instance for the [VZAudioInputStreamSource] class.
var (
	VZAudioInputStreamSourceClass     _VZAudioInputStreamSourceClass
	VZAudioInputStreamSourceClassOnce sync.Once
)

func getVZAudioInputStreamSourceClass() _VZAudioInputStreamSourceClass {
	VZAudioInputStreamSourceClassOnce.Do(func() {
		VZAudioInputStreamSourceClass = _VZAudioInputStreamSourceClass{objc.GetClass("VZAudioInputStreamSource")}
	})
	return VZAudioInputStreamSourceClass
}

type _VZAudioInputStreamSourceClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZAudioInputStreamSource */
// An interface definition for the [VZAudioInputStreamSource] class.
type IVZAudioInputStreamSource interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZAudioInputStreamSource */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZAudioInputStreamSource */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZAudioInputStreamSource */
// Alloc allocates a new instance without initialization.
func (vc _VZAudioInputStreamSourceClass) Alloc() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZAudioInputStreamSourceClass) New() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioInputStreamSource) Init() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioInputStreamSource) Autorelease() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioInputStreamSource creates a new VZAudioInputStreamSource instance.
func NewVZAudioInputStreamSource() VZAudioInputStreamSource {
	return getVZAudioInputStreamSourceClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZAudioInputStreamSource */
// The base class for an audio input stream source.
//
// An audio input stream source defines how th guest produces audio input data on the host system. Don’t instantiate directly, use one of its subclasses such as instead.

// The base class for an audio input stream source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type VZAudioInputStreamSource struct {
	objectivec.Object
}

// VZAudioInputStreamSourceFrom constructs a [VZAudioInputStreamSource] from an unsafe.Pointer.
//
// The base class for an audio input stream source.
func VZAudioInputStreamSourceFrom(ptr unsafe.Pointer) VZAudioInputStreamSource {
	return VZAudioInputStreamSource{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZAudioInputStreamSource */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZAudioInputStreamSource */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZAudioInputStreamSource */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZAudioInputStreamSource */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZAudioInputStreamSource */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZAudioInputStreamSource */
