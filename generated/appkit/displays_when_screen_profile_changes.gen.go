
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [displaysWhenScreenProfileChanges] class.
var displaysWhenScreenProfileChangesClass _displaysWhenScreenProfileChangesClass

func init() {
	displaysWhenScreenProfileChangesClass = _displaysWhenScreenProfileChangesClass{objc.GetClass("displaysWhenScreenProfileChanges")}
}

type _displaysWhenScreenProfileChangesClass struct {
	objc.Class
}

// An interface definition for the [displaysWhenScreenProfileChanges] class.
type IdisplaysWhenScreenProfileChanges interface {
	ID() objc.ID
}

type displaysWhenScreenProfileChanges struct {
	id objc.ID
}

func displaysWhenScreenProfileChangesFrom(ptr unsafe.Pointer) displaysWhenScreenProfileChanges {
	return displaysWhenScreenProfileChanges{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ displaysWhenScreenProfileChanges) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _displaysWhenScreenProfileChangesClass) Alloc() displaysWhenScreenProfileChanges {
	rv := objc.Send[displaysWhenScreenProfileChanges](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _displaysWhenScreenProfileChangesClass) New() displaysWhenScreenProfileChanges {
	rv := objc.Send[displaysWhenScreenProfileChanges](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisplaysWhenScreenProfileChanges creates and returns a new initialized instance.
func NewdisplaysWhenScreenProfileChanges() displaysWhenScreenProfileChanges {
	return displaysWhenScreenProfileChangesClass.New()
}

// Init initializes the instance.
func (d_ displaysWhenScreenProfileChanges) Init() displaysWhenScreenProfileChanges {
	rv := objc.Send[displaysWhenScreenProfileChanges](d_.ID(), selInit)
	return rv
}
