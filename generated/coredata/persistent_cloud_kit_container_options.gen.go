// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentCloudKitContainerOptions] class.
var persistentCloudKitContainerOptionsClass = _PersistentCloudKitContainerOptionsClass{objc.GetClass("NSPersistentCloudKitContainerOptions")}

type _PersistentCloudKitContainerOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerOptions] class.
type IPersistentCloudKitContainerOptions interface {
	objectivec.IObject
}

// An object that customizes how a store description aligns with a CloudKit database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions

type PersistentCloudKitContainerOptions struct {
	objectivec.Object
}

// PersistentCloudKitContainerOptionsFrom constructs a [PersistentCloudKitContainerOptions] from an unsafe.Pointer.
//
// An object that customizes how a store description aligns with a CloudKit database.
func PersistentCloudKitContainerOptionsFrom(ptr unsafe.Pointer) PersistentCloudKitContainerOptions {
	return PersistentCloudKitContainerOptions{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PersistentCloudKitContainerOptionsClass) Alloc() PersistentCloudKitContainerOptions {
	rv := objc.Send[PersistentCloudKitContainerOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersistentCloudKitContainerOptionsClass) New() PersistentCloudKitContainerOptions {
	rv := objc.Send[PersistentCloudKitContainerOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentCloudKitContainerOptions) Init() PersistentCloudKitContainerOptions {
	rv := objc.Send[PersistentCloudKitContainerOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentCloudKitContainerOptions) Autorelease() PersistentCloudKitContainerOptions {
	rv := objc.Send[PersistentCloudKitContainerOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentCloudKitContainerOptions creates a new PersistentCloudKitContainerOptions instance.
func NewPersistentCloudKitContainerOptions() PersistentCloudKitContainerOptions {
	return persistentCloudKitContainerOptionsClass.New()
}


// Initializes container options using the given CloudKit container identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/init(containerIdentifier:)
func NewPersistentCloudKitContainerOptionsWithContainerIdentifier(containerIdentifier string) PersistentCloudKitContainerOptions {
	instance := persistentCloudKitContainerOptionsClass.Alloc()
	rv := objc.Send[PersistentCloudKitContainerOptions](instance.ID, objc.Sel("initWithContainerIdentifier:"), containerIdentifier)
	rv.Autorelease()
	return rv
}



