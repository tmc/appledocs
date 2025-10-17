// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FileCoordinator] class.
var FileCoordinatorClass objc.Class

func init() {
	FileCoordinatorClass = objc.GetClass("NSFileCoordinator")
}

type FileCoordinator struct {
	objc.ID
}

func FileCoordinatorFrom(ptr unsafe.Pointer) FileCoordinator {
	return FileCoordinator{
		ID: objc.ID(ptr),
	}
}




