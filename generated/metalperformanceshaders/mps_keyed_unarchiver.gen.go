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
}

// A keyed archiver that supports Metal Performance Shaders kernel decoding.
//
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

// Alloc allocates a new instance without initialization.
func (kc _KeyedUnarchiverClass) Alloc() KeyedUnarchiver {
	rv := objc.Send[KeyedUnarchiver](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/init(forReadingFrom:device:error:)
func NewKeyedUnarchiverForReadingFromDataDeviceError(data unsafe.Pointer, device objc.ID, error_ unsafe.Pointer) KeyedUnarchiver {
	instance := getKeyedUnarchiverClass().Alloc()
	rv := objc.Send[KeyedUnarchiver](instance.ID, objc.Sel("initForReadingFromData:device:error:"), data, device, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKeyedUnarchiver/unarchivedObject(ofClasses:from:device:)
func (kc _KeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataDeviceError(classes unsafe.Pointer, data unsafe.Pointer, device objc.ID, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(kc.class), objc.Sel("unarchivedObjectOfClasses:fromData:device:error:"), classes, data, device, error_)
	return rv
}


