// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var fileCoordinatorClass _FileCoordinatorClass

func init() {
	fileCoordinatorClass = _FileCoordinatorClass{objc.GetClass("NSFileCoordinator")}
}

type _FileCoordinatorClass struct {
	class objc.Class
}

type FileCoordinator struct {
	objc.ID
}

func FileCoordinatorFrom(ptr unsafe.Pointer) FileCoordinator {
	return FileCoordinator{
		ID: objc.ID(ptr),
	}
}




