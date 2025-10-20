// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentCloudKitContainerOptions] class.
var (
	PersistentCloudKitContainerOptionsClass     _PersistentCloudKitContainerOptionsClass
	PersistentCloudKitContainerOptionsClassOnce sync.Once
)

func getPersistentCloudKitContainerOptionsClass() _PersistentCloudKitContainerOptionsClass {
	PersistentCloudKitContainerOptionsClassOnce.Do(func() {
		PersistentCloudKitContainerOptionsClass = _PersistentCloudKitContainerOptionsClass{objc.GetClass("NSPersistentCloudKitContainerOptions")}
	})
	return PersistentCloudKitContainerOptionsClass
}

type _PersistentCloudKitContainerOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerOptions] class.
type IPersistentCloudKitContainerOptions interface {
	objectivec.IObject
}

// An object that customizes how a store description aligns with a CloudKit database.
//
// Use to customize the behavior of an or to create additional store descriptions that sync to other containers. For more information about setting up multiple stores, see .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPersistentCloudKitContainerOptionsClass().New()
}


// Initializes container options using the given CloudKit container identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/init(containerIdentifier:)
func NewPersistentCloudKitContainerOptionsWithContainerIdentifier(containerIdentifier string) PersistentCloudKitContainerOptions {
	instance := getPersistentCloudKitContainerOptionsClass().Alloc()
	rv := objc.Send[PersistentCloudKitContainerOptions](instance.ID, objc.Sel("initWithContainerIdentifier:"), objc.String(containerIdentifier))
	rv.Autorelease()
	return rv
}


// The identifier of the CloudKit container associated with a given store description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/containerIdentifier
func (p_ PersistentCloudKitContainerOptions) ContainerIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("containerIdentifier"))
	return rv
}

// The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/databaseScope-2784h
func (p_ PersistentCloudKitContainerOptions) DatabaseScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("databaseScope"))
	return rv
}


// SetDatabaseScope sets the value of the databaseScope property.
// The database scope — public, private, or shared — to use for a specified store in a persistent CloudKit container.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerOptions/databaseScope-2784h
func (p_ PersistentCloudKitContainerOptions) SetDatabaseScope(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDatabaseScope:"), value)
}

