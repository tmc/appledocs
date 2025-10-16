
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [dataSource] class.
var dataSourceClass _dataSourceClass

func init() {
	dataSourceClass = _dataSourceClass{objc.GetClass("dataSource")}
}

type _dataSourceClass struct {
	objc.Class
}

// An interface definition for the [dataSource] class.
type IdataSource interface {
	ID() objc.ID
}

type dataSource struct {
	id objc.ID
}

func dataSourceFrom(ptr unsafe.Pointer) dataSource {
	return dataSource{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ dataSource) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _dataSourceClass) Alloc() dataSource {
	rv := objc.Send[dataSource](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _dataSourceClass) New() dataSource {
	rv := objc.Send[dataSource](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdataSource creates and returns a new initialized instance.
func NewdataSource() dataSource {
	return dataSourceClass.New()
}

// Init initializes the instance.
func (d_ dataSource) Init() dataSource {
	rv := objc.Send[dataSource](d_.ID(), selInit)
	return rv
}
