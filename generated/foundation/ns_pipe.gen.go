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
	PipeClass     _PipeClass
	PipeClassOnce sync.Once
)

func getPipeClass() _PipeClass {
	PipeClassOnce.Do(func() {
		PipeClass = _PipeClass{objc.GetClass("NSPipe")}
	})
	return PipeClass
}

type _PipeClass struct {
	class objc.Class
}

// An interface definition for the [Pipe] class.
type IPipe interface {
	objectivec.IObject
}

// A one-way communications channel between related processes.
//
// objects provide an object-oriented interface for accessing pipes. An object represents both ends of a pipe and enables communication through the pipe. A pipe is a one-way communications channel between related processes; one process writes data, while the other process reads that data. The data that passes through the pipe is buffered; the size of the buffer is determined by the underlying operating system. is an abstract class, the public interface of a class cluster.
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


// Returns an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPipe/pipe
func (pc _PipeClass) Pipe() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pipe"))
	return rv
}

// The receiver’s read file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Pipe/fileHandleForReading
func (p_ Pipe) FileHandleForReading() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fileHandleForReading"))
	return rv
}

// The receiver’s write file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Pipe/fileHandleForWriting
func (p_ Pipe) FileHandleForWriting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fileHandleForWriting"))
	return rv
}



