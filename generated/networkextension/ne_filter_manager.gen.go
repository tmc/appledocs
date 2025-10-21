// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEFilterManager] class.
var (
	NEFilterManagerClass     _NEFilterManagerClass
	NEFilterManagerClassOnce sync.Once
)

func getNEFilterManagerClass() _NEFilterManagerClass {
	NEFilterManagerClassOnce.Do(func() {
		NEFilterManagerClass = _NEFilterManagerClass{objc.GetClass("NEFilterManager")}
	})
	return NEFilterManagerClass
}

type _NEFilterManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterManager] class.
type INEFilterManager interface {
	objectivec.IObject
}

// An object to create and manage a content filter’s configuration.
//
// Each app is allowed to create a single filter configuration. The class has a class method ( ) that provides access to a single instance. This single instance corresponds to a single filter configuration. The filter configuration is stored in the Network Extension preferences which are managed by the Network Extension framework. The filter configuration must be explicitly loaded into memory from the Network Extension preferences before it can be used, and any changes must be explicitly saved to the Network Extension preferences before taking effect on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager
type NEFilterManager struct {
	objectivec.Object
}

// NEFilterManagerFrom constructs a [NEFilterManager] from an unsafe.Pointer.
//
// An object to create and manage a content filter’s configuration.
func NEFilterManagerFrom(ptr unsafe.Pointer) NEFilterManager {
	return NEFilterManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterManagerClass) Alloc() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterManagerClass) New() NEFilterManager {
	rv := objc.Send[NEFilterManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterManager) Init() NEFilterManager {
	rv := objc.Send[NEFilterManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterManager) Autorelease() NEFilterManager {
	rv := objc.Send[NEFilterManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterManager creates a new NEFilterManager instance.
func NewNEFilterManager() NEFilterManager {
	return getNEFilterManagerClass().New()
}




