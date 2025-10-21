// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Architecture] class.
var (
	ArchitectureClass     _ArchitectureClass
	ArchitectureClassOnce sync.Once
)

func getArchitectureClass() _ArchitectureClass {
	ArchitectureClassOnce.Do(func() {
		ArchitectureClass = _ArchitectureClass{objc.GetClass("MTLArchitecture")}
	})
	return ArchitectureClass
}

type _ArchitectureClass struct {
	class objc.Class
}

// An interface definition for the [Architecture] class.
type IArchitecture interface {
	objectivec.IObject
}

// A class that contains the architectural details of a GPU device.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArchitecture
type Architecture struct {
	objectivec.Object
}

// ArchitectureFrom constructs a [Architecture] from an unsafe.Pointer.
//
// A class that contains the architectural details of a GPU device.
func ArchitectureFrom(ptr unsafe.Pointer) Architecture {
	return Architecture{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ArchitectureClass) Alloc() Architecture {
	rv := objc.Send[Architecture](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArchitectureClass) New() Architecture {
	rv := objc.Send[Architecture](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Architecture) Init() Architecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Architecture) Autorelease() Architecture {
	rv := objc.Send[Architecture](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArchitecture creates a new Architecture instance.
func NewArchitecture() Architecture {
	return getArchitectureClass().New()
}


// The name of a GPU device’s architecture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArchitecture/name
func (a_ Architecture) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}



