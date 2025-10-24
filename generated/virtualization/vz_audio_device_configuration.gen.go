// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZAudioDeviceConfiguration */


/* debug [class_header]: Header for VZAudioDeviceConfiguration */
// The class instance for the [VZAudioDeviceConfiguration] class.
var (
	VZAudioDeviceConfigurationClass     _VZAudioDeviceConfigurationClass
	VZAudioDeviceConfigurationClassOnce sync.Once
)

func getVZAudioDeviceConfigurationClass() _VZAudioDeviceConfigurationClass {
	VZAudioDeviceConfigurationClassOnce.Do(func() {
		VZAudioDeviceConfigurationClass = _VZAudioDeviceConfigurationClass{objc.GetClass("VZAudioDeviceConfiguration")}
	})
	return VZAudioDeviceConfigurationClass
}

type _VZAudioDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZAudioDeviceConfiguration */
// An interface definition for the [VZAudioDeviceConfiguration] class.
type IVZAudioDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZAudioDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZAudioDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZAudioDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZAudioDeviceConfigurationClass) Alloc() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZAudioDeviceConfigurationClass) New() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioDeviceConfiguration) Init() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioDeviceConfiguration) Autorelease() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDeviceConfiguration creates a new VZAudioDeviceConfiguration instance.
func NewVZAudioDeviceConfiguration() VZAudioDeviceConfiguration {
	return getVZAudioDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZAudioDeviceConfiguration */
// The base class for an audio device configuration.
//
// Don’t instantiate this abstract class directly. Instead, instantiate one of its subclasses such as .


// The base class for an audio device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type VZAudioDeviceConfiguration struct {
	objectivec.Object
}

// VZAudioDeviceConfigurationFrom constructs a [VZAudioDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for an audio device configuration.
func VZAudioDeviceConfigurationFrom(ptr unsafe.Pointer) VZAudioDeviceConfiguration {
	return VZAudioDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZAudioDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZAudioDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZAudioDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZAudioDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZAudioDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZAudioDeviceConfiguration */



