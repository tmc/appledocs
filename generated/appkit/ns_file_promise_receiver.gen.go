// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFilePromiseReceiver */


/* debug [class_header]: Header for NSFilePromiseReceiver */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FilePromiseReceiver */
// An interface definition for the [FilePromiseReceiver] class.
type IFilePromiseReceiver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FilePromiseReceiver */
	// properties:
	FileNames() []string
	FileTypes() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FilePromiseReceiver */
	// methods:
	ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir objc.IObject /* cross-framework: NSURL */, options objc.IObject /* cross-framework: NSDictionary */, operationQueue foundation.OperationQueue, reader unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FilePromiseReceiver */
// Alloc allocates a new instance without initialization.
func (fc _FilePromiseReceiverClass) Alloc() FilePromiseReceiver {
	rv := objc.Send[FilePromiseReceiver](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FilePromiseReceiver */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FilePromiseReceiver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FilePromiseReceiver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FilePromiseReceiver */

// An array containing dragged file types that are readable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/readableDraggedTypes
func (fc _FilePromiseReceiverClass) ReadableDraggedTypes() []string {
	rv := objc.Send[[]string](objc.ID(fc.class), objc.Sel("readableDraggedTypes"))
	return rv
}/* debug [class_properties_class/property]: readableDraggedTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FilePromiseReceiver */

// Fulfills the promises at the specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/receivePromisedFiles(atDestination:options:operationQueue:reader:)
func (f_ FilePromiseReceiver) ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir objc.IObject /* cross-framework: NSURL */, options objc.IObject /* cross-framework: NSDictionary */, operationQueue foundation.OperationQueue, reader unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("receivePromisedFilesAtDestination:options:operationQueue:reader:"), destinationDir, options, operationQueue, reader)
}/* debug [instance_methods/method]: ReceivePromisedFilesAtDestinationOptionsOperationQueueReader */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FilePromiseReceiver */

// An array containing names of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/fileNames
func (f_ FilePromiseReceiver) FileNames() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("fileNames"))
	return rv
}/* debug [instance_properties/getter]: fileNames */


// An array containing types of the promised files being written to the destination location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/fileTypes
func (f_ FilePromiseReceiver) FileTypes() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("fileTypes"))
	return rv
}/* debug [instance_properties/getter]: fileTypes */


// An array containing dragged file types that are readable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/readableDraggedTypes
func (f_ FilePromiseReceiver) ReadableDraggedTypes() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("readableDraggedTypes"))
	return rv
}/* debug [instance_properties/getter]: readableDraggedTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFilePromiseReceiver */



