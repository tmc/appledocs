
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enableSnapshotRestoration] class.
var enableSnapshotRestorationClass _enableSnapshotRestorationClass

func init() {
	enableSnapshotRestorationClass = _enableSnapshotRestorationClass{objc.GetClass("enableSnapshotRestoration")}
}

type _enableSnapshotRestorationClass struct {
	objc.Class
}

// An interface definition for the [enableSnapshotRestoration] class.
type IenableSnapshotRestoration interface {
	ID() objc.ID
}

type enableSnapshotRestoration struct {
	id objc.ID
}

func enableSnapshotRestorationFrom(ptr unsafe.Pointer) enableSnapshotRestoration {
	return enableSnapshotRestoration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enableSnapshotRestoration) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enableSnapshotRestorationClass) Alloc() enableSnapshotRestoration {
	rv := objc.Send[enableSnapshotRestoration](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enableSnapshotRestorationClass) New() enableSnapshotRestoration {
	rv := objc.Send[enableSnapshotRestoration](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenableSnapshotRestoration creates and returns a new initialized instance.
func NewenableSnapshotRestoration() enableSnapshotRestoration {
	return enableSnapshotRestorationClass.New()
}

// Init initializes the instance.
func (e_ enableSnapshotRestoration) Init() enableSnapshotRestoration {
	rv := objc.Send[enableSnapshotRestoration](e_.ID(), selInit)
	return rv
}
