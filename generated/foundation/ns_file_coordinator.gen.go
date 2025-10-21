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
	CoordinateAccessWithIntentsQueueByAccessor(intents []FileAccessIntent, queue IOperationQueue, accessor unsafe.Pointer)
}

// An object that coordinates the reading and writing of files and directories among file presenters.
//
// The class coordinates the reading and writing of files and directories among multiple processes and objects in the same process. You use instances of this class as is to read from, write to, modify the attributes of, change the location of, or delete a file or directory, but before your code to perform those actions executes, the file coordinator lets registered file presenter objects perform any tasks that they might require to ensure their own integrity. For example, if you want to change the location of a file, other objects interested in that file need to know where you intend to move it so that they can update their references. Objects that adopt the protocol must register themselves with the class to be notified of any pending changes. They do this by calling the class method. A file presenter must balance calls to with a call to before being released, even in a garbage-collected application. The file presenter class maintains a list of active file presenter objects in the current application and uses that list, plus the file coordinator classes in other processes, to deliver notifications to all of the objects interested in a particular item. Instances of are meant to be used on a per-file-operation basis, where a file operation is something like opening and reading the contents of a file or moving a batch of files and directories to a new location. There is no benefit to keeping a file coordinator object past the length of the planned operation. In fact, because file coordinators retain file presenter objects, keeping one around could prevent the file presenter objects from being released in a timely manner. For information about implementing a file presenter object to receive file-related notifications, see .
//
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

// Alloc allocates a new instance without initialization.
func (fc _FileCoordinatorClass) Alloc() FileCoordinator {
	rv := objc.Send[FileCoordinator](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes and returns a file coordinator object using the specified file presenter.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/init(filePresenter:)
func NewFileCoordinatorWithFilePresenter(filePresenterOrNil objectivec.IObject) FileCoordinator {
	instance := getFileCoordinatorClass().Alloc()
	rv := objc.Send[FileCoordinator](instance.ID, objc.Sel("initWithFilePresenter:"), filePresenterOrNil)
	rv.Autorelease()
	return rv
}


// Unregisters the specified file presenter object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/removeFilePresenter(_:)
func (fc _FileCoordinatorClass) RemoveFilePresenter(filePresenter objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeFilePresenter:"), filePresenter)
}

// Returns an array containing the currently registered file presenter objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/filePresenters
func (fc _FileCoordinatorClass) FilePresenters() []objc.ID {
	rv := objc.Send[[]objc.ID](objc.ID(fc.class), objc.Sel("filePresenters"))
	return rv
}
// Performs a number of coordinated-read or -write operations asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(with:queue:byAccessor:)
func (f_ FileCoordinator) CoordinateAccessWithIntentsQueueByAccessor(intents []FileAccessIntent, queue IOperationQueue, accessor unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("coordinateAccessWithIntents:queue:byAccessor:"), intents, queue, accessor)
}

// Returns an array containing the currently registered file presenter objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/filePresenters
func (f_ FileCoordinator) FilePresenters() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("filePresenters"))
	return rv
}

// A string that uniquely identifies the file access that was performed by this file coordinator.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/purposeidentifier
func (f_ FileCoordinator) PurposeIdentifier() appkit.string {
	rv := objc.Send[appkit.string](f_.ID, objc.Sel("purposeIdentifier"))
	return rv
}


// SetPurposeIdentifier sets the value of the purposeIdentifier property.
// A string that uniquely identifies the file access that was performed by this file coordinator.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilecoordinator/purposeidentifier
func (f_ FileCoordinator) SetPurposeIdentifier(value appkit.string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPurposeIdentifier:"), value)
}

// The user canceled the operation (for example, by pressing Command-period).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusercancellederror-swift.var
func (f_ FileCoordinator) NSUserCancelledError() int {
	rv := objc.Send[int](f_.ID, objc.Sel("NSUserCancelledError"))
	return rv
}


// SetNSUserCancelledError sets the value of the NSUserCancelledError property.
// The user canceled the operation (for example, by pressing Command-period).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusercancellederror-swift.var
func (f_ FileCoordinator) SetNSUserCancelledError(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSUserCancelledError:"), value)
}


