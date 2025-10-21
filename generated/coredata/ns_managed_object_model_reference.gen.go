// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ManagedObjectModelReference] class.
var (
	ManagedObjectModelReferenceClass     _ManagedObjectModelReferenceClass
	ManagedObjectModelReferenceClassOnce sync.Once
)

func getManagedObjectModelReferenceClass() _ManagedObjectModelReferenceClass {
	ManagedObjectModelReferenceClassOnce.Do(func() {
		ManagedObjectModelReferenceClass = _ManagedObjectModelReferenceClass{objc.GetClass("NSManagedObjectModelReference")}
	})
	return ManagedObjectModelReferenceClass
}

type _ManagedObjectModelReferenceClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectModelReference] class.
type IManagedObjectModelReference interface {
	objectivec.IObject
}

// An object that describes a specific version of an object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference
type ManagedObjectModelReference struct {
	objectivec.Object
}

// ManagedObjectModelReferenceFrom constructs a [ManagedObjectModelReference] from an unsafe.Pointer.
//
// An object that describes a specific version of an object model.
func ManagedObjectModelReferenceFrom(ptr unsafe.Pointer) ManagedObjectModelReference {
	return ManagedObjectModelReference{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectModelReferenceClass) Alloc() ManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectModelReferenceClass) New() ManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObjectModelReference) Init() ManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObjectModelReference) Autorelease() ManagedObjectModelReference {
	rv := objc.Send[ManagedObjectModelReference](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObjectModelReference creates a new ManagedObjectModelReference instance.
func NewManagedObjectModelReference() ManagedObjectModelReference {
	return getManagedObjectModelReferenceClass().New()
}




