
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [jobDisposition] class.
var jobDispositionClass _jobDispositionClass

func init() {
	jobDispositionClass = _jobDispositionClass{objc.GetClass("jobDisposition")}
}

type _jobDispositionClass struct {
	objc.Class
}

// An interface definition for the [jobDisposition] class.
type IjobDisposition interface {
	ID() objc.ID
}

type jobDisposition struct {
	id objc.ID
}

func jobDispositionFrom(ptr unsafe.Pointer) jobDisposition {
	return jobDisposition{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (j_ jobDisposition) ID() objc.ID {
	return j_.id
}

// Alloc allocates a new instance without initialization.
func (jc _jobDispositionClass) Alloc() jobDisposition {
	rv := objc.Send[jobDisposition](objc.ID(jc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (jc _jobDispositionClass) New() jobDisposition {
	rv := objc.Send[jobDisposition](objc.ID(jc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewjobDisposition creates and returns a new initialized instance.
func NewjobDisposition() jobDisposition {
	return jobDispositionClass.New()
}

// Init initializes the instance.
func (j_ jobDisposition) Init() jobDisposition {
	rv := objc.Send[jobDisposition](j_.ID(), selInit)
	return rv
}
