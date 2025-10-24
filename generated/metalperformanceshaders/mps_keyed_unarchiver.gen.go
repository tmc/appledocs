// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [KeyedUnarchiver] class.
type IKeyedUnarchiver interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	MpsMTLDevice()


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2966644-initforreadingfromdata
func NewKeyedUnarchiverForReadingFromDataDeviceError(data foundation.Data, device unsafe.Pointer, error_ objectivec.IObject) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingFromData:device:error:"), data, device, error_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951877-initforreadingwithdata
func NewKeyedUnarchiverForReadingWithDataDevice(data foundation.Data, device unsafe.Pointer) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingWithData:device:"), data, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951874-initwithdevice
func NewKeyedUnarchiverWithDevice(device unsafe.Pointer) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951875-unarchiveobject
func (kc _KeyedUnarchiverClass) UnarchiveObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObject"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951875-unarchiveobjectwithfile
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithFileDevice(path string, device unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithFile:device:"), objc.String(path), device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951876-unarchivetoplevelobject
func (kc _KeyedUnarchiverClass) UnarchiveTopLevelObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveTopLevelObject"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951876-unarchivetoplevelobjectwithdata
func (kc _KeyedUnarchiverClass) UnarchiveTopLevelObjectWithDataDeviceError(data foundation.Data, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveTopLevelObjectWithData:device:error:"), data, device, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951881-unarchiveobjectwithdata
func (kc _KeyedUnarchiverClass) UnarchiveObjectWithDataDevice(data foundation.Data, device unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchiveObjectWithData:device:"), data, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobject
func (kc _KeyedUnarchiverClass) UnarchivedObject() {
	objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObject"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976453-unarchivedobjectofclass
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassFromDataDeviceError(cls objc.Class, data foundation.Data, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClass:fromData:device:error:"), cls, data, device, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2976454-unarchivedobjectofclasses
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceError(classes objc.Class, data foundation.Data, device unsafe.Pointer, error_ objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), classes, data, device, error_)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskeyedunarchiver/2951880-mpsmtldevice
func (k_ KeyedUnarchiver) MpsMTLDevice() {
	objc.Send[objc.ID](k_.ID, objc.Sel("mpsMTLDevice"))
}












