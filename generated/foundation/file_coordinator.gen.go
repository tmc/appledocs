// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileCoordinator] class.
var (
	fileCoordinatorClass     _FileCoordinatorClass
	fileCoordinatorClassOnce sync.Once
)

func getFileCoordinatorClass() _FileCoordinatorClass {
	fileCoordinatorClassOnce.Do(func() {
		fileCoordinatorClass = _FileCoordinatorClass{objc.GetClass("NSFileCoordinator")}
	})
	return fileCoordinatorClass
}

type _FileCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [FileCoordinator] class.
type IFileCoordinator interface {
	objectivec.IObject
}

// An object that coordinates the reading and writing of files and directories among file presenters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
type FileCoordinator struct {
	objectivec.Object
}

// FileCoordinatorFrom constructs a [FileCoordinator] from an unsafe.Pointer.
//
// An object that coordinates the reading and writing of files and directories among file presenters.
func FileCoordinatorFrom(ptr unsafe.Pointer) FileCoordinator {
	return FileCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileCoordinatorClass) Alloc() FileCoordinator {
	rv := objc.Send[FileCoordinator](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileCoordinatorClass) New() FileCoordinator {
	rv := objc.Send[FileCoordinator](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileCoordinator) Init() FileCoordinator {
	rv := objc.Send[FileCoordinator](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileCoordinator) Autorelease() FileCoordinator {
	rv := objc.Send[FileCoordinator](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileCoordinator creates a new FileCoordinator instance.
func NewFileCoordinator() FileCoordinator {
	return getFileCoordinatorClass().New()
}




