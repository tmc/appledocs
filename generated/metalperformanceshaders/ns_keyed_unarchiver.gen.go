// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [KeyedUnarchiver] class.
var (
	KeyedUnarchiverClass     _KeyedUnarchiverClass
	KeyedUnarchiverClassOnce sync.Once
)

func getKeyedUnarchiverClass() _KeyedUnarchiverClass {
	KeyedUnarchiverClassOnce.Do(func() {
		KeyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}
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

// A parent class referenced by other MetalPerformanceShaders classes.
type KeyedUnarchiver struct {
	objectivec.Object
}

// KeyedUnarchiverFrom constructs a [KeyedUnarchiver] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
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




