// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilePromiseReceiver] class.
var (
	FilePromiseReceiverClass     _FilePromiseReceiverClass
	FilePromiseReceiverClassOnce sync.Once
)

func getFilePromiseReceiverClass() _FilePromiseReceiverClass {
	FilePromiseReceiverClassOnce.Do(func() {
		FilePromiseReceiverClass = _FilePromiseReceiverClass{objc.GetClass("NSFilePromiseReceiver")}
	})
	return FilePromiseReceiverClass
}

type _FilePromiseReceiverClass struct {
	class objc.Class
}

// An interface definition for the [FilePromiseReceiver] class.
type IFilePromiseReceiver interface {
	objectivec.IObject
	// properties:
	FileNames() objc.IObject /* cross-framework: NSString */
	SetFileNames(value objc.IObject /* cross-framework: NSString */)
	FileTypes() objc.IObject /* cross-framework: NSString */
	SetFileTypes(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object that receives a file promise from the pasteboard.
//
// Because implements the protocol, you receive all file promises on the drag pasteboard as follows: Likewise, you can enumerate dragged items by calling the following:


// An object that receives a file promise from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver
type FilePromiseReceiver struct {
	objectivec.Object
}

// FilePromiseReceiverFrom constructs a [FilePromiseReceiver] from an unsafe.Pointer.
//
// An object that receives a file promise from the pasteboard.
func FilePromiseReceiverFrom(ptr unsafe.Pointer) FilePromiseReceiver {
	return FilePromiseReceiver{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilePromiseReceiverClass) Alloc() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilePromiseReceiverClass) New() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilePromiseReceiver) Init() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilePromiseReceiver) Autorelease() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilePromiseReceiver creates a new FilePromiseReceiver instance.
func NewFilePromiseReceiver() FilePromiseReceiver {
	return getFilePromiseReceiverClass().New()
}



// An array containing names of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/filenames
func (f_ FilePromiseReceiver) FileNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("fileNames"))
	return rv
}


// An array containing names of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/filenames
func (f_ FilePromiseReceiver) SetFileNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileNames:"), value)
}


// An array containing types of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/filetypes
func (f_ FilePromiseReceiver) FileTypes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("fileTypes"))
	return rv
}


// An array containing types of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfilepromisereceiver/filetypes
func (f_ FilePromiseReceiver) SetFileTypes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileTypes:"), value)
}



