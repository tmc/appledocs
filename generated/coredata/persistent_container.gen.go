// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentContainer] class.
var persistentContainerClass = _PersistentContainerClass{objc.GetClass("NSPersistentContainer")}

type _PersistentContainerClass struct {
	class objc.Class
}

// A container that encapsulates the Core Data stack in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer

type PersistentContainer struct {
	objectivec.Object
}

// PersistentContainerFrom constructs a [PersistentContainer] from an unsafe.Pointer.
//
// A container that encapsulates the Core Data stack in your app.
func PersistentContainerFrom(ptr unsafe.Pointer) PersistentContainer {
	return PersistentContainer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PersistentContainerClass) Alloc() PersistentContainer {
	rv := objc.Send[PersistentContainer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersistentContainerClass) New() PersistentContainer {
	rv := objc.Send[PersistentContainer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentContainer) Init() PersistentContainer {
	rv := objc.Send[PersistentContainer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentContainer) Autorelease() PersistentContainer {
	rv := objc.Send[PersistentContainer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentContainer creates a new PersistentContainer instance.
func NewPersistentContainer() PersistentContainer {
	return persistentContainerClass.New()
}
// Creates a container with the specified name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:)
func NewPersistentContainerWithName(name string) PersistentContainer {
	instance := persistentContainerClass.Alloc()
	rv := objc.Send[PersistentContainer](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}


// Loads the persistent stores. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/loadPersistentStores(completionHandler:)
func (p_ PersistentContainer) LoadPersistentStoresWithCompletionHandler(block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("loadPersistentStoresWithCompletionHandler:"), block)
}
// Returns a new managed object context that executes on a private queue. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/newBackgroundContext()
func (p_ PersistentContainer) NewBackgroundContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("newBackgroundContext"))
	return rv
}
// Executes a closure on a private queue using an ephemeral managed object context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/performBackgroundTask(_:)-39sch
func (p_ PersistentContainer) PerformBackgroundTask(block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performBackgroundTask:"), block)
}

