// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileCoordinator] class.
var fileCoordinatorClass = _FileCoordinatorClass{objc.GetClass("NSFileCoordinator")}

type _FileCoordinatorClass struct {
	class objc.Class
}

// An object that coordinates the reading and writing of files and directories among file presenters. [Full Topic]
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



