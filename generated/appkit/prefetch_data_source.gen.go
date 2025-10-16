
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prefetchDataSource] class.
var prefetchDataSourceClass _prefetchDataSourceClass

func init() {
	prefetchDataSourceClass = _prefetchDataSourceClass{objc.GetClass("prefetchDataSource")}
}

type _prefetchDataSourceClass struct {
	objc.Class
}

// An interface definition for the [prefetchDataSource] class.
type IprefetchDataSource interface {
	ID() objc.ID
}

type prefetchDataSource struct {
	id objc.ID
}

func prefetchDataSourceFrom(ptr unsafe.Pointer) prefetchDataSource {
	return prefetchDataSource{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prefetchDataSource) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prefetchDataSourceClass) Alloc() prefetchDataSource {
	rv := objc.Send[prefetchDataSource](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prefetchDataSourceClass) New() prefetchDataSource {
	rv := objc.Send[prefetchDataSource](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprefetchDataSource creates and returns a new initialized instance.
func NewprefetchDataSource() prefetchDataSource {
	return prefetchDataSourceClass.New()
}

// Init initializes the instance.
func (p_ prefetchDataSource) Init() prefetchDataSource {
	rv := objc.Send[prefetchDataSource](p_.ID(), selInit)
	return rv
}
