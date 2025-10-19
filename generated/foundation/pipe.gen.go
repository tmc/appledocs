// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pipe] class.
var (
	pipeClass     _PipeClass
	pipeClassOnce sync.Once
)

func getPipeClass() _PipeClass {
	pipeClassOnce.Do(func() {
		pipeClass = _PipeClass{objc.GetClass("NSPipe")}
	})
	return pipeClass
}

type _PipeClass struct {
	class objc.Class
}

// An interface definition for the [Pipe] class.
type IPipe interface {
	objectivec.IObject
}

// A one-way communications channel between related processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Pipe
type Pipe struct {
	objectivec.Object
}

// PipeFrom constructs a [Pipe] from an unsafe.Pointer.
//
// A one-way communications channel between related processes.
func PipeFrom(ptr unsafe.Pointer) Pipe {
	return Pipe{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PipeClass) Alloc() Pipe {
	rv := objc.Send[Pipe](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PipeClass) New() Pipe {
	rv := objc.Send[Pipe](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pipe) Init() Pipe {
	rv := objc.Send[Pipe](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pipe) Autorelease() Pipe {
	rv := objc.Send[Pipe](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPipe creates a new Pipe instance.
func NewPipe() Pipe {
	return getPipeClass().New()
}




