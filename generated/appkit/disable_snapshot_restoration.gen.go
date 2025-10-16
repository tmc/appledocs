
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [disableSnapshotRestoration] class.
var disableSnapshotRestorationClass _disableSnapshotRestorationClass

func init() {
	disableSnapshotRestorationClass = _disableSnapshotRestorationClass{objc.GetClass("disableSnapshotRestoration")}
}

type _disableSnapshotRestorationClass struct {
	objc.Class
}

// An interface definition for the [disableSnapshotRestoration] class.
type IdisableSnapshotRestoration interface {
	ID() objc.ID
}

type disableSnapshotRestoration struct {
	id objc.ID
}

func disableSnapshotRestorationFrom(ptr unsafe.Pointer) disableSnapshotRestoration {
	return disableSnapshotRestoration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ disableSnapshotRestoration) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _disableSnapshotRestorationClass) Alloc() disableSnapshotRestoration {
	rv := objc.Send[disableSnapshotRestoration](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _disableSnapshotRestorationClass) New() disableSnapshotRestoration {
	rv := objc.Send[disableSnapshotRestoration](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisableSnapshotRestoration creates and returns a new initialized instance.
func NewdisableSnapshotRestoration() disableSnapshotRestoration {
	return disableSnapshotRestorationClass.New()
}

// Init initializes the instance.
func (d_ disableSnapshotRestoration) Init() disableSnapshotRestoration {
	rv := objc.Send[disableSnapshotRestoration](d_.ID(), selInit)
	return rv
}
