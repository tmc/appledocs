// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPipe */


/* debug [class_header]: Header for NSPipe */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Pipe */
// An interface definition for the [Pipe] class.
type IPipe interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Pipe */
	// properties:
	FileHandleForReading() IFileHandle
	SetFileHandleForReading(value IFileHandle)
	FileHandleForWriting() IFileHandle
	SetFileHandleForWriting(value IFileHandle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Pipe */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Pipe */
// Alloc allocates a new instance without initialization.
func (pc _PipeClass) Alloc() Pipe {
	rv := objc.Send[Pipe](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Pipe */
// A one-way communications channel between related processes.
//
// objects provide an object-oriented interface for accessing pipes. An object represents both ends of a pipe and enables communication through the pipe. A pipe is a one-way communications channel between related processes; one process writes data, while the other process reads that data. The data that passes through the pipe is buffered; the size of the buffer is determined by the underlying operating system. is an abstract class, the public interface of a class cluster.


// A one-way communications channel between related processes.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Pipe *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Pipe */

// Returns an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPipe/pipe
func (pc _PipeClass) Pipe() IPipe {
	rv := objc.Send[Pipe](objc.ID(pc.class), objc.Sel("pipe"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Pipe) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Pipe */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Pipe */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Pipe */

// The receiver’s read file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/pipe/filehandleforreading
func (p_ Pipe) FileHandleForReading() IFileHandle {
	rv := objc.Send[FileHandle](p_.ID, objc.Sel("fileHandleForReading"))
	return rv
}/* debug [instance_properties/getter]: fileHandleForReading */


// The receiver’s read file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/pipe/filehandleforreading
func (p_ Pipe) SetFileHandleForReading(value IFileHandle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileHandleForReading:"), value)
}/* debug [instance_properties/setter]: fileHandleForReading */


// The receiver’s write file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/pipe/filehandleforwriting
func (p_ Pipe) FileHandleForWriting() IFileHandle {
	rv := objc.Send[FileHandle](p_.ID, objc.Sel("fileHandleForWriting"))
	return rv
}/* debug [instance_properties/getter]: fileHandleForWriting */


// The receiver’s write file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/pipe/filehandleforwriting
func (p_ Pipe) SetFileHandleForWriting(value IFileHandle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileHandleForWriting:"), value)
}/* debug [instance_properties/setter]: fileHandleForWriting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPipe */



