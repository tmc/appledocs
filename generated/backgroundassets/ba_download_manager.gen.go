// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BADownloadManager] class.
var (
	BADownloadManagerClass     _BADownloadManagerClass
	BADownloadManagerClassOnce sync.Once
)

func getBADownloadManagerClass() _BADownloadManagerClass {
	BADownloadManagerClassOnce.Do(func() {
		BADownloadManagerClass = _BADownloadManagerClass{objc.GetClass("BADownloadManager")}
	})
	return BADownloadManagerClass
}

type _BADownloadManagerClass struct {
	class objc.Class
}

// An interface definition for the [BADownloadManager] class.
type IBADownloadManager interface {
	objectivec.IObject
	CancelDownloadError(download IBADownload, error_ unsafe.Pointer) bool
	FetchCurrentDownloads(error_ unsafe.Pointer) []BADownload
	FetchCurrentDownloadsWithCompletionHandler(completionHandler unsafe.Pointer)
	ScheduleDownloadError(download IBADownload, error_ unsafe.Pointer) bool
	StartForegroundDownloadError(download IBADownload, error_ unsafe.Pointer) bool
	PerformWithExclusiveControl(performHandler unsafe.Pointer)
	PerformWithExclusiveControlBeforeDatePerformHandler(date foundation.IDate, performHandler unsafe.Pointer)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
}

// An object that manages the queue of scheduled asset downloads.
//
// Use to schedule and cancel asset downloads, monitor their progress, and access the queue of pending downloads. You don’t create instances of this class directly; instead, use the property to access the framework’s singleton that it shares between your app and the app’s extension. Because the download manager is a shared resource, prevent race conditions by using the and methods to assume absolute control of the manager before you schedule asset downloads or manipulate those already in the manager’s queue. To respond to asset download events and process concluded downloads, create a type that conforms to the protocol and assign an instance of it to the download manager’s property. The following example shows how to create an asset download, acquire exclusive control of the shared download manager, and then use the manager to schedule the download:


// An object that manages the queue of scheduled asset downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager

type BADownloadManager struct {
	objectivec.Object
}

// BADownloadManagerFrom constructs a [BADownloadManager] from an unsafe.Pointer.
//
// An object that manages the queue of scheduled asset downloads.
func BADownloadManagerFrom(ptr unsafe.Pointer) BADownloadManager {
	return BADownloadManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BADownloadManagerClass) Alloc() BADownloadManager {
	rv := objc.Send[BADownloadManager](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BADownloadManagerClass) New() BADownloadManager {
	rv := objc.Send[BADownloadManager](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BADownloadManager) Init() BADownloadManager {
	rv := objc.Send[BADownloadManager](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BADownloadManager) Autorelease() BADownloadManager {
	rv := objc.Send[BADownloadManager](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBADownloadManager creates a new BADownloadManager instance.
func NewBADownloadManager() BADownloadManager {
	return getBADownloadManagerClass().New()
}



// The download manager that both the app and the extension share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/shared

func (bc _BADownloadManagerClass) SharedManager() BADownloadManager {
	rv := objc.Send[BADownloadManager](objc.ID(bc.class), objc.Sel("sharedManager"))
	return rv
}


// Cancels an asset download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/cancel(_:)

func (b_ BADownloadManager) CancelDownloadError(download IBADownload, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("cancelDownload:error:"), download, error_)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/fetchCurrentDownloads()

func (b_ BADownloadManager) FetchCurrentDownloads(error_ unsafe.Pointer) []BADownload {
	rv := objc.Send[[]BADownload](b_.ID, objc.Sel("fetchCurrentDownloads:"), error_)
	return rv
}



// Fetches the contents of the manager’s download queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/fetchCurrentDownloads(completionHandler:)

func (b_ BADownloadManager) FetchCurrentDownloadsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("fetchCurrentDownloadsWithCompletionHandler:"), completionHandler)
}



// Schedules an asset download to execute in the background at a nonspecific time in the future.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/scheduleDownload(_:)

func (b_ BADownloadManager) ScheduleDownloadError(download IBADownload, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("scheduleDownload:error:"), download, error_)
	return rv
}



// Schedules an asset download that executes immediately in the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/startForegroundDownload(_:)

func (b_ BADownloadManager) StartForegroundDownloadError(download IBADownload, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("startForegroundDownload:error:"), download, error_)
	return rv
}



// Attempts to acquire immediate, exclusive access to the download manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/withExclusiveControl(_:)

func (b_ BADownloadManager) PerformWithExclusiveControl(performHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("performWithExclusiveControl:"), performHandler)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/withExclusiveControl(beforeDate:perform:)

func (b_ BADownloadManager) PerformWithExclusiveControlBeforeDatePerformHandler(date foundation.IDate, performHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("performWithExclusiveControlBeforeDate:performHandler:"), date, performHandler)
}


// The download manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/delegate

func (b_ BADownloadManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// The download manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/delegate

func (b_ BADownloadManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}


// The download manager that both the app and the extension share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownloadManager/shared

func (b_ BADownloadManager) SharedManager() BADownloadManager {
	rv := objc.Send[BADownloadManager](b_.ID, objc.Sel("sharedManager"))
	return rv
}



