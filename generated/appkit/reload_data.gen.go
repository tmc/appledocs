
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [reloadData] class.
var reloadDataClass _reloadDataClass

func init() {
	reloadDataClass = _reloadDataClass{objc.GetClass("reloadData")}
}

type _reloadDataClass struct {
	objc.Class
}

// An interface definition for the [reloadData] class.
type IreloadData interface {
	ID() objc.ID
}

type reloadData struct {
	id objc.ID
}

func reloadDataFrom(ptr unsafe.Pointer) reloadData {
	return reloadData{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ reloadData) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _reloadDataClass) Alloc() reloadData {
	rv := objc.Send[reloadData](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _reloadDataClass) New() reloadData {
	rv := objc.Send[reloadData](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewreloadData creates and returns a new initialized instance.
func NewreloadData() reloadData {
	return reloadDataClass.New()
}

// Init initializes the instance.
func (r_ reloadData) Init() reloadData {
	rv := objc.Send[reloadData](r_.ID(), selInit)
	return rv
}
