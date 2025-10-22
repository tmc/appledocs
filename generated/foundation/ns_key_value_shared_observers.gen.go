// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [KeyValueSharedObservers] class.
var (
	KeyValueSharedObserversClass     _KeyValueSharedObserversClass
	KeyValueSharedObserversClassOnce sync.Once
)

func getKeyValueSharedObserversClass() _KeyValueSharedObserversClass {
	KeyValueSharedObserversClassOnce.Do(func() {
		KeyValueSharedObserversClass = _KeyValueSharedObserversClass{objc.GetClass("NSKeyValueSharedObservers")}
	})
	return KeyValueSharedObserversClass
}

type _KeyValueSharedObserversClass struct {
	class objc.Class
}

// An interface definition for the [KeyValueSharedObservers] class.
type IKeyValueSharedObservers interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers
type KeyValueSharedObservers struct {
	objectivec.Object
}

// KeyValueSharedObserversFrom constructs a [KeyValueSharedObservers] from an unsafe.Pointer.
func KeyValueSharedObserversFrom(ptr unsafe.Pointer) KeyValueSharedObservers {
	return KeyValueSharedObservers{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (kc _KeyValueSharedObserversClass) Alloc() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyValueSharedObserversClass) New() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyValueSharedObservers) Init() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyValueSharedObservers) Autorelease() KeyValueSharedObservers {
	rv := objc.Send[KeyValueSharedObservers](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyValueSharedObservers creates a new KeyValueSharedObservers instance.
func NewKeyValueSharedObservers() KeyValueSharedObservers {
	return getKeyValueSharedObserversClass().New()
}




