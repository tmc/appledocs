// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioSoundDeviceStreamConfiguration */


/* debug [class_header]: Header for VZVirtioSoundDeviceStreamConfiguration */
// The class instance for the [VZVirtioSoundDeviceStreamConfiguration] class.
var (
	VZVirtioSoundDeviceStreamConfigurationClass     _VZVirtioSoundDeviceStreamConfigurationClass
	VZVirtioSoundDeviceStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceStreamConfigurationClass() _VZVirtioSoundDeviceStreamConfigurationClass {
	VZVirtioSoundDeviceStreamConfigurationClassOnce.Do(func() {
		VZVirtioSoundDeviceStreamConfigurationClass = _VZVirtioSoundDeviceStreamConfigurationClass{objc.GetClass("VZVirtioSoundDeviceStreamConfiguration")}
	})
	return VZVirtioSoundDeviceStreamConfigurationClass
}

type _VZVirtioSoundDeviceStreamConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioSoundDeviceStreamConfiguration */
// An interface definition for the [VZVirtioSoundDeviceStreamConfiguration] class.
type IVZVirtioSoundDeviceStreamConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZVirtioSoundDeviceStreamConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioSoundDeviceStreamConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioSoundDeviceStreamConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceStreamConfigurationClass) Alloc() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioSoundDeviceStreamConfigurationClass) New() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSoundDeviceStreamConfiguration) Init() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSoundDeviceStreamConfiguration) Autorelease() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceStreamConfiguration creates a new VZVirtioSoundDeviceStreamConfiguration instance.
func NewVZVirtioSoundDeviceStreamConfiguration() VZVirtioSoundDeviceStreamConfiguration {
	return getVZVirtioSoundDeviceStreamConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioSoundDeviceStreamConfiguration */
// An object that defines a Virtio sound device stream configuration.
//
// A object represents a PCM stream of audio data. Don’t instantiate this class directly. Instead, instantiate one of its subclasses such as or .


// An object that defines a Virtio sound device stream configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceStreamConfiguration
type VZVirtioSoundDeviceStreamConfiguration struct {
	objectivec.Object
}

// VZVirtioSoundDeviceStreamConfigurationFrom constructs a [VZVirtioSoundDeviceStreamConfiguration] from an unsafe.Pointer.
//
// An object that defines a Virtio sound device stream configuration.
func VZVirtioSoundDeviceStreamConfigurationFrom(ptr unsafe.Pointer) VZVirtioSoundDeviceStreamConfiguration {
	return VZVirtioSoundDeviceStreamConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioSoundDeviceStreamConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioSoundDeviceStreamConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioSoundDeviceStreamConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioSoundDeviceStreamConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioSoundDeviceStreamConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioSoundDeviceStreamConfiguration */



