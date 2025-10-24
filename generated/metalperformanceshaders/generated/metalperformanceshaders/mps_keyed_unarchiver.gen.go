// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSKeyedUnarchiver */


/* debug [class_header]: Header for MPSKeyedUnarchiver */
// The class instance for the [KeyedUnarchiver] class.
var (
	KeyedUnarchiverClass     _KeyedUnarchiverClass
	KeyedUnarchiverClassOnce sync.Once
)

func getKeyedUnarchiverClass() _KeyedUnarchiverClass {
	KeyedUnarchiverClassOnce.Do(func() {
		KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("MPSKeyedUnarchiver")}
	})
	return KeyedUnarchiverClass
}

type _KeyedUnarchiverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for KeyedUnarchiver */
// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for KeyedUnarchiver */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for KeyedUnarchiver */
	// methods:
	MpsMTLDevice()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for KeyedUnarchiver */
// Alloc allocates a new instance without initialization.
func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyedUnarchiverClass) New() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyedUnarchiver) Init() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyedUnarchiver) Autorelease() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyedUnarchiver creates a new KeyedUnarchiver instance.
func NewKeyedUnarchiver() KeyedUnarchiver {
	return getKeyedUnarchiverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for KeyedUnarchiver */
// A keyed archiver that supports Metal Performance Shaders kernel decoding.


// A keyed archiver that supports Metal Performance Shaders kernel decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver
type KeyedUnarchiver struct {
	objectivec.Object
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A keyed archiver that supports Metal Performance Shaders kernel decoding.
func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for KeyedUnarchiver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2966644-initforreadingfromdata
func NewKeyedUnarchiverForReadingFromDataDeviceError(data objc.IObject /* cross-framework: Data */, device unsafe.Pointer, error_ objectivec.IObject) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingFromData:device:error:"), data, device, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyedUnarchiverForReadingFromDataDeviceError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951877-initforreadingwithdata
func NewKeyedUnarchiverForReadingWithDataDevice(data objc.IObject /* cross-framework: Data */, device unsafe.Pointer) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingWithData:device:"), data, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyedUnarchiverForReadingWithDataDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951874-initwithdevice
func NewKeyedUnarchiverWithDevice(device unsafe.Pointer) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKeyedUnarchiverWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for KeyedUnarchiver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951875-unarchiveobject
func (kc _KeyedUnarchiverClass) UnarchiveObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObject"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveObject) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951875-unarchiveobjectwithfile
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithFileDevice(path string, device unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithFile:device:"), objc.String(path), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveObjectWithFileDevice) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951876-unarchivetoplevelobject
func (kc _KeyedUnarchiverClass) UnarchiveTopLevelObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveTopLevelObject"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveTopLevelObject) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951876-unarchivetoplevelobjectwithdata
func (kc _KeyedUnarchiverClass) UnarchiveTopLevelObjectWithDataDeviceError(data objc.IObject /* cross-framework: Data */, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveTopLevelObjectWithData:device:error:"), data, device, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveTopLevelObjectWithDataDeviceError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951881-unarchiveobjectwithdata
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithDataDevice(data objc.IObject /* cross-framework: Data */, device unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithData:device:"), data, device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveObjectWithDataDevice) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobject
func (kc _KeyedUnarchiverClass) UnarchivedObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObject"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchivedObject) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataDeviceError(cls objc.Class, data objc.IObject /* cross-framework: Data */, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClass:fromData:device:error:"), cls, data, device, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchivedObjectOfClassFromDataDeviceError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceError(classes objc.Class, data objc.IObject /* cross-framework: Data */, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), classes, data, device, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchivedObjectOfClassesFromDataDeviceError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for KeyedUnarchiver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for KeyedUnarchiver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951880-mpsmtldevice
func (k_ KeyedUnarchiver) MpsMTLDevice() {
	objc.Send[objc.ID](k_.ID, objc.Sel("mpsMTLDevice"))
}/* debug [instance_methods/method]: MpsMTLDevice */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for KeyedUnarchiver */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSKeyedUnarchiver */


