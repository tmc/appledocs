// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SortDescriptor] class.
var (
	SortDescriptorClass     _SortDescriptorClass
	SortDescriptorClassOnce sync.Once
)

func getSortDescriptorClass() _SortDescriptorClass {
	SortDescriptorClassOnce.Do(func() {
		SortDescriptorClass = _SortDescriptorClass{objc.GetClass("NSSortDescriptor")}
	})
	return SortDescriptorClass
}

type _SortDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [SortDescriptor] class.
type ISortDescriptor interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other CloudKit classes.


// A parent class referenced by other CloudKit classes. [Full Topic]
type SortDescriptor struct {
	objectivec.Object
}

// SortDescriptorFrom constructs a [SortDescriptor] from an unsafe.Pointer.
//
// A parent class referenced by other CloudKit classes.
func SortDescriptorFrom(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.Send[SortDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := objc.Send[SortDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SortDescriptor) Init() SortDescriptor {
	rv := objc.Send[SortDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SortDescriptor) Autorelease() SortDescriptor {
	rv := objc.Send[SortDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSortDescriptor creates a new SortDescriptor instance.
func NewSortDescriptor() SortDescriptor {
	return getSortDescriptorClass().New()
}




