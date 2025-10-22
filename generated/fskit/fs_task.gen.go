// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSTask] class.
var (
	FSTaskClass     _FSTaskClass
	FSTaskClassOnce sync.Once
)

func getFSTaskClass() _FSTaskClass {
	FSTaskClassOnce.Do(func() {
		FSTaskClass = _FSTaskClass{objc.GetClass("FSTask")}
	})
	return FSTaskClass
}

type _FSTaskClass struct {
	class objc.Class
}

// An interface definition for the [FSTask] class.
type IFSTask interface {
	objectivec.IObject
	DidCompleteWithError(error_ foundation.IError)
	LogMessage(str string)
	CancellationHandler() unsafe.Pointer
	SetCancellationHandler(value unsafe.Pointer)
}

// A class that enables a file system module to pass log messages and completion notifications to clients.
//
// FSKit creates an instance of this class for each long-running operations.


// A class that enables a file system module to pass log messages and completion notifications to clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask

type FSTask struct {
	objectivec.Object
}

// FSTaskFrom constructs a [FSTask] from an unsafe.Pointer.
//
// A class that enables a file system module to pass log messages and completion notifications to clients.
func FSTaskFrom(ptr unsafe.Pointer) FSTask {
	return FSTask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSTaskClass) Alloc() FSTask {
	rv := objc.Send[FSTask](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSTaskClass) New() FSTask {
	rv := objc.Send[FSTask](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSTask) Init() FSTask {
	rv := objc.Send[FSTask](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSTask) Autorelease() FSTask {
	rv := objc.Send[FSTask](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTask creates a new FSTask instance.
func NewFSTask() FSTask {
	return getFSTaskClass().New()
}




// Informs the client that the task completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/didComplete(error:)

func (f_ FSTask) DidCompleteWithError(error_ foundation.IError) {
	objc.Send[objc.ID](f_.ID, objc.Sel("didCompleteWithError:"), error_)
}



// Logs the given string to the initiating client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/logMessage(_:)

func (f_ FSTask) LogMessage(str string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("logMessage:"), objc.String(str))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/cancellationHandler

func (f_ FSTask) CancellationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("cancellationHandler"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTask/cancellationHandler

func (f_ FSTask) SetCancellationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCancellationHandler:"), value)
}



