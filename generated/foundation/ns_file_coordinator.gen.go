// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FileCoordinator] class.
var (
	FileCoordinatorClass     _FileCoordinatorClass
	FileCoordinatorClassOnce sync.Once
)

func getFileCoordinatorClass() _FileCoordinatorClass {
	FileCoordinatorClassOnce.Do(func() {
		FileCoordinatorClass = _FileCoordinatorClass{objc.GetClass("NSFileCoordinator")}
	})
	return FileCoordinatorClass
}

type _FileCoordinatorClass struct {
	class objc.Class
}





// An interface definition for the [FileCoordinator] class.
type IFileCoordinator interface {
	objectivec.IObject
	

	// properties:
	PurposeIdentifier() IString
	SetPurposeIdentifier(value IString)
	NSUserCancelledError() int
	SetNSUserCancelledError(value int)


	

	// methods:
	Cancel()
	CoordinateReadingItemAtURLOptionsErrorByAccessor(url IURL, options FileCoordinatorReadingOptions, outError IError, reader unsafe.Pointer)
	CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL IURL, readingOptions FileCoordinatorReadingOptions, writingURL IURL, writingOptions FileCoordinatorWritingOptions, outError IError, readerWriter unsafe.Pointer)
	CoordinateAccessWithIntentsQueueByAccessor(intents []FileAccessIntent, queue IOperationQueue, accessor unsafe.Pointer)
	CoordinateWritingItemAtURLOptionsErrorByAccessor(url IURL, options FileCoordinatorWritingOptions, outError IError, writer unsafe.Pointer)
	CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 IURL, options1 FileCoordinatorWritingOptions, url2 IURL, options2 FileCoordinatorWritingOptions, outError IError, writer unsafe.Pointer)
	ItemAtURLDidChangeUbiquityAttributes(url IURL, attributes unsafe.Pointer)
	ItemAtURLDidMoveToURL(oldURL IURL, newURL IURL)
	ItemAtURLWillMoveToURL(oldURL IURL, newURL IURL)
	PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []URL, readingOptions FileCoordinatorReadingOptions, writingURLs []URL, writingOptions FileCoordinatorWritingOptions, outError IError, batchAccessor unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (fc _FileCoordinatorClass) Alloc() FileCoordinator {
	rv := objc.Send[FileCoordinator](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileCoordinatorClass) New() FileCoordinator {
	rv := objc.Send[FileCoordinator](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileCoordinator) Init() FileCoordinator {
	rv := objc.Send[FileCoordinator](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileCoordinator) Autorelease() FileCoordinator {
	rv := objc.Send[FileCoordinator](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileCoordinator creates a new FileCoordinator instance.
func NewFileCoordinator() FileCoordinator {
	return getFileCoordinatorClass().New()
}





// An object that coordinates the reading and writing of files and directories among file presenters.
//
// The class coordinates the reading and writing of files and directories among multiple processes and objects in the same process. You use instances of this class as is to read from, write to, modify the attributes of, change the location of, or delete a file or directory, but before your code to perform those actions executes, the file coordinator lets registered file presenter objects perform any tasks that they might require to ensure their own integrity. For example, if you want to change the location of a file, other objects interested in that file need to know where you intend to move it so that they can update their references. Objects that adopt the protocol must register themselves with the class to be notified of any pending changes. They do this by calling the class method. A file presenter must balance calls to with a call to before being released, even in a garbage-collected application. The file presenter class maintains a list of active file presenter objects in the current application and uses that list, plus the file coordinator classes in other processes, to deliver notifications to all of the objects interested in a particular item. Instances of are meant to be used on a per-file-operation basis, where a file operation is something like opening and reading the contents of a file or moving a batch of files and directories to a new location. There is no benefit to keeping a file coordinator object past the length of the planned operation. In fact, because file coordinators retain file presenter objects, keeping one around could prevent the file presenter objects from being released in a timely manner. For information about implementing a file presenter object to receive file-related notifications, see .


// An object that coordinates the reading and writing of files and directories among file presenters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
type FileCoordinator struct {
	objectivec.Object
}

// FileCoordinatorFrom constructs a [FileCoordinator] from an unsafe.Pointer.
//
// An object that coordinates the reading and writing of files and directories among file presenters.
func FileCoordinatorFrom(ptr unsafe.Pointer) FileCoordinator {
	return FileCoordinator{objectivec.Object{objc.ID(ptr)}}
}






// Initializes and returns a file coordinator object using the specified file presenter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/init(filePresenter:)
func NewFileCoordinatorWithFilePresenter(filePresenterOrNil unsafe.Pointer) FileCoordinator {
	instance := getFileCoordinatorClass().Alloc()
	rv := objc.Send[FileCoordinator](instance.ID, objc.Sel("initWithFilePresenter:"), filePresenterOrNil)
	rv.Autorelease()
	return rv
}







// Registers the specified file presenter object so that it can receive notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/addFilePresenter(_:)
func (fc _FileCoordinatorClass) AddFilePresenter(filePresenter unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("addFilePresenter:"), filePresenter)
}


// Unregisters the specified file presenter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/removeFilePresenter(_:)
func (fc _FileCoordinatorClass) RemoveFilePresenter(filePresenter unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeFilePresenter:"), filePresenter)
}







// Returns an array containing the currently registered file presenter objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/filePresenters
func (fc _FileCoordinatorClass) FilePresenters() []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(fc.class), objc.Sel("filePresenters"))
	return rv
}






// Cancels any active file coordination calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/cancel()
func (f_ FileCoordinator) Cancel() {
	objc.Send[objc.ID](f_.ID, objc.Sel("cancel"))
}


// Initiates a read operation on a single file or directory using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(readingItemAt:options:error:byAccessor:)
func (f_ FileCoordinator) CoordinateReadingItemAtURLOptionsErrorByAccessor(url IURL, options FileCoordinatorReadingOptions, outError IError, reader unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateReadingItemAtURL:options:error:byAccessor:"), url, options, outError, reader)
}


// Initiates a read operation that contains a follow-up write operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(readingItemAt:options:writingItemAt:options:error:byAccessor:)
func (f_ FileCoordinator) CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL IURL, readingOptions FileCoordinatorReadingOptions, writingURL IURL, writingOptions FileCoordinatorWritingOptions, outError IError, readerWriter unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateReadingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), readingURL, readingOptions, writingURL, writingOptions, outError, readerWriter)
}


// Performs a number of coordinated-read or -write operations asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(with:queue:byAccessor:)
func (f_ FileCoordinator) CoordinateAccessWithIntentsQueueByAccessor(intents []FileAccessIntent, queue IOperationQueue, accessor unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateAccessWithIntents:queue:byAccessor:"), intents, queue, accessor)
}


// Initiates a write operation on a single file or directory using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(writingItemAt:options:error:byAccessor:)
func (f_ FileCoordinator) CoordinateWritingItemAtURLOptionsErrorByAccessor(url IURL, options FileCoordinatorWritingOptions, outError IError, writer unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateWritingItemAtURL:options:error:byAccessor:"), url, options, outError, writer)
}


// Initiates a write operation that involves a secondary write operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(writingItemAt:options:writingItemAt:options:error:byAccessor:)
func (f_ FileCoordinator) CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 IURL, options1 FileCoordinatorWritingOptions, url2 IURL, options2 FileCoordinatorWritingOptions, outError IError, writer unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateWritingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), url1, options1, url2, options2, outError, writer)
}


// Tells observing file providers that the item’s ubiquity attributes have changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:didChangeUbiquityAttributes:)
func (f_ FileCoordinator) ItemAtURLDidChangeUbiquityAttributes(url IURL, attributes unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemAtURL:didChangeUbiquityAttributes:"), url, attributes)
}


// Notifies relevant file presenters that the location of a file or directory changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:didMoveTo:)
func (f_ FileCoordinator) ItemAtURLDidMoveToURL(oldURL IURL, newURL IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemAtURL:didMoveToURL:"), oldURL, newURL)
}


// Announces that your app is moving a file to a new URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:willMoveTo:)
func (f_ FileCoordinator) ItemAtURLWillMoveToURL(oldURL IURL, newURL IURL) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemAtURL:willMoveToURL:"), oldURL, newURL)
}


// Prepare to read or write from multiple files in a single batch operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/prepare(forReadingItemsAt:options:writingItemsAt:options:error:byAccessor:)
func (f_ FileCoordinator) PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []URL, readingOptions FileCoordinatorReadingOptions, writingURLs []URL, writingOptions FileCoordinatorWritingOptions, outError IError, batchAccessor unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("prepareForReadingItemsAtURLs:options:writingItemsAtURLs:options:error:byAccessor:"), readingURLs, readingOptions, writingURLs, writingOptions, outError, batchAccessor)
}







// Returns an array containing the currently registered file presenter objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/filePresenters
func (f_ FileCoordinator) FilePresenters() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("filePresenters"))
	return rv
}


// A string that uniquely identifies the file access that was performed by this file coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/purposeIdentifier
func (f_ FileCoordinator) PurposeIdentifier() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("purposeIdentifier"))
	return rv
}


// A string that uniquely identifies the file access that was performed by this file coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/purposeIdentifier
func (f_ FileCoordinator) SetPurposeIdentifier(value IString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPurposeIdentifier:"), value)
}


// The user canceled the operation (for example, by pressing Command-period).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusercancellederror-swift.var
func (f_ FileCoordinator) NSUserCancelledError() int {
	rv := objc.Send[int](f_.ID, objc.Sel("NSUserCancelledError"))
	return rv
}


// The user canceled the operation (for example, by pressing Command-period).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusercancellederror-swift.var
func (f_ FileCoordinator) SetNSUserCancelledError(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSUserCancelledError:"), value)
}







