// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDI2DeviceInfo */


/* debug [class_header]: Header for MIDI2DeviceInfo */
// The class instance for the [MIDI2DeviceInfo] class.
var (
	MIDI2DeviceInfoClass     _MIDI2DeviceInfoClass
	MIDI2DeviceInfoClassOnce sync.Once
)

func getMIDI2DeviceInfoClass() _MIDI2DeviceInfoClass {
	MIDI2DeviceInfoClassOnce.Do(func() {
		MIDI2DeviceInfoClass = _MIDI2DeviceInfoClass{objc.GetClass("MIDI2DeviceInfo")}
	})
	return MIDI2DeviceInfoClass
}

type _MIDI2DeviceInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDI2DeviceInfo */
// An interface definition for the [MIDI2DeviceInfo] class.
type IMIDI2DeviceInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDI2DeviceInfo */
	// properties:
	Family() MIDIUInteger14 /* typedef */
	ManufacturerID() objc.IObject /* cross-framework: MIDI2DeviceManufacturer */
	ModelNumber() MIDIUInteger14 /* typedef */
	RevisionLevel() objc.IObject /* cross-framework: MIDI2DeviceRevisionLevel */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDI2DeviceInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDI2DeviceInfo */
// Alloc allocates a new instance without initialization.
func (mc _MIDI2DeviceInfoClass) Alloc() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDI2DeviceInfoClass) New() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDI2DeviceInfo) Init() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDI2DeviceInfo) Autorelease() MIDI2DeviceInfo {
	rv := objc.Send[MIDI2DeviceInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDI2DeviceInfo creates a new MIDI2DeviceInfo instance.
func NewMIDI2DeviceInfo() MIDI2DeviceInfo {
	return getMIDI2DeviceInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDI2DeviceInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo
type MIDI2DeviceInfo struct {
	objectivec.Object
}

// MIDI2DeviceInfoFrom constructs a [MIDI2DeviceInfo] from an unsafe.Pointer.
func MIDI2DeviceInfoFrom(ptr unsafe.Pointer) MIDI2DeviceInfo {
	return MIDI2DeviceInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDI2DeviceInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/init(manufacturerID:family:modelNumber:revisionLevel:)
func NewMIDI2DeviceInfoWithManufacturerIDFamilyModelNumberRevisionLevel(manufacturerID objc.IObject /* cross-framework: MIDI2DeviceManufacturer */, family MIDIUInteger14 /* typedef */, modelNumber MIDIUInteger14 /* typedef */, revisionLevel objc.IObject /* cross-framework: MIDI2DeviceRevisionLevel */) MIDI2DeviceInfo {
	instance := getMIDI2DeviceInfoClass().Alloc()
	rv := objc.Send[MIDI2DeviceInfo](instance.ID, objc.Sel("initWithManufacturerID:family:modelNumber:revisionLevel:"), manufacturerID, family, modelNumber, revisionLevel)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDI2DeviceInfoWithManufacturerIDFamilyModelNumberRevisionLevel */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDI2DeviceInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDI2DeviceInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDI2DeviceInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDI2DeviceInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/family
func (m_ MIDI2DeviceInfo) Family() MIDIUInteger14 /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("family"))
	return rv
}/* debug [instance_properties/getter]: family */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/manufacturerID
func (m_ MIDI2DeviceInfo) ManufacturerID() objc.IObject /* cross-framework: MIDI2DeviceManufacturer */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("manufacturerID"))
	return rv
}/* debug [instance_properties/getter]: manufacturerID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/modelNumber
func (m_ MIDI2DeviceInfo) ModelNumber() MIDIUInteger14 /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("modelNumber"))
	return rv
}/* debug [instance_properties/getter]: modelNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDI2DeviceInfo/revisionLevel
func (m_ MIDI2DeviceInfo) RevisionLevel() objc.IObject /* cross-framework: MIDI2DeviceRevisionLevel */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("revisionLevel"))
	return rv
}/* debug [instance_properties/getter]: revisionLevel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDI2DeviceInfo */


