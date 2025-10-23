// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ResolvedModel() IManagedObjectModel
	SetResolvedModel(value IManagedObjectModel)
	VersionChecksum() string /* primitive/slice/pointer. */
	SetVersionChecksum(value string /* primitive/slice/pointer. */)
	// methods:
}

// An object that describes a specific version of an object model.


// An object that describes a specific version of an object model.
//
// [Full Topic]
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



// The resolved object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodelreference/resolvedmodel
func (m_ ManagedObjectModelReference) ResolvedModel() IManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("resolvedModel"))
	return rv
}


// The resolved object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodelreference/resolvedmodel
func (m_ ManagedObjectModelReference) SetResolvedModel(value IManagedObjectModel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResolvedModel:"), value)
}


// The version checksum of the resolved model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodelreference/versionchecksum
func (m_ ManagedObjectModelReference) VersionChecksum() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("versionChecksum"))
	return rv
}


// The version checksum of the resolved model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectmodelreference/versionchecksum
func (m_ ManagedObjectModelReference) SetVersionChecksum(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVersionChecksum:"), objc.String(value))
}



