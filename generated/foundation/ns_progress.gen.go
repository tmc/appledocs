// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Progress] class.
var (
	ProgressClass     _ProgressClass
	ProgressClassOnce sync.Once
)

func getProgressClass() _ProgressClass {
	ProgressClassOnce.Do(func() {
		ProgressClass = _ProgressClass{objc.GetClass("NSProgress")}
	})
	return ProgressClass
}

type _ProgressClass struct {
	class objc.Class
}

// An interface definition for the [Progress] class.
type IProgress interface {
	objectivec.IObject
	// properties:
	EstimatedTimeRemaining() INumber
	SetEstimatedTimeRemaining(value INumber)
	FileCompletedCount() INumber
	SetFileCompletedCount(value INumber)
	FileTotalCount() INumber
	SetFileTotalCount(value INumber)
	Throughput() INumber
	SetThroughput(value INumber)
	CancellationHandler() unsafe.Pointer
	SetCancellationHandler(value unsafe.Pointer)
	CompletedUnitCount() unsafe.Pointer
	SetCompletedUnitCount(value unsafe.Pointer)
	FileOperationKind() ProgressFileOperationKind /* foo */
	SetFileOperationKind(value ProgressFileOperationKind /* foo */)
	FileURL() IURL
	SetFileURL(value IURL)
	FractionCompleted() float64 /* primitive/slice/pointer */
	Cancellable() bool /* primitive/slice/pointer */
	SetCancellable(value bool /* primitive/slice/pointer */)
	Cancelled() bool /* primitive/slice/pointer */
	Finished() bool /* primitive/slice/pointer */
	Indeterminate() bool /* primitive/slice/pointer */
	Old() bool /* primitive/slice/pointer */
	Pausable() bool /* primitive/slice/pointer */
	SetPausable(value bool /* primitive/slice/pointer */)
	Paused() bool /* primitive/slice/pointer */
	Kind() ProgressKind /* foo */
	SetKind(value ProgressKind /* foo */)
	LocalizedAdditionalDescription() string /* primitive/slice/pointer */
	SetLocalizedAdditionalDescription(value string /* primitive/slice/pointer */)
	LocalizedDescription() string /* primitive/slice/pointer */
	SetLocalizedDescription(value string /* primitive/slice/pointer */)
	PausingHandler() unsafe.Pointer
	SetPausingHandler(value unsafe.Pointer)
	ResumingHandler() unsafe.Pointer
	SetResumingHandler(value unsafe.Pointer)
	TotalUnitCount() unsafe.Pointer
	SetTotalUnitCount(value unsafe.Pointer)
	UserInfo() IDictionary /* already interface */
	IsCancellable() bool /* primitive/slice/pointer */
	SetIsCancellable(value bool /* primitive/slice/pointer */)
	IsCancelled() bool /* primitive/slice/pointer */
	SetIsCancelled(value bool /* primitive/slice/pointer */)
	IsFinished() bool /* primitive/slice/pointer */
	SetIsFinished(value bool /* primitive/slice/pointer */)
	IsIndeterminate() bool /* primitive/slice/pointer */
	SetIsIndeterminate(value bool /* primitive/slice/pointer */)
	IsOld() bool /* primitive/slice/pointer */
	SetIsOld(value bool /* primitive/slice/pointer */)
	IsPausable() bool /* primitive/slice/pointer */
	SetIsPausable(value bool /* primitive/slice/pointer */)
	IsPaused() bool /* primitive/slice/pointer */
	SetIsPaused(value bool /* primitive/slice/pointer */)
	// methods:
	PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount unsafe.Pointer, work unsafe.Pointer)
	AddChildWithPendingUnitCount(child IProgress, inUnitCount unsafe.Pointer)
	BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer)
	Cancel()
	Pause()
	Publish()
	ResignCurrent()
	Resume()
	SetUserInfoObjectForKey(objectOrNil objectivec.IObject, key ProgressUserInfoKey /* foo */)
	Unpublish()
}

// An object that conveys ongoing progress to the user for a specified task.
//
// The class provides a self-contained mechanism for progress reporting. It makes it easy for code that performs work to report the progress of that work, and for user interface code to observe that progress for presentation to the user. Specifically, you can use a progress object to show the user a progress bar and explanatory text that update as you do work. It also allows the user to cancel or pause work.


// An object that conveys ongoing progress to the user for a specified task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress
type Progress struct {
	objectivec.Object
}

// ProgressFrom constructs a [Progress] from an unsafe.Pointer.
//
// An object that conveys ongoing progress to the user for a specified task.
func ProgressFrom(ptr unsafe.Pointer) Progress {
	return Progress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProgressClass) Alloc() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProgressClass) New() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Progress) Init() Progress {
	rv := objc.Send[Progress](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Progress) Autorelease() Progress {
	rv := objc.Send[Progress](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProgress creates a new Progress instance.
func NewProgress() Progress {
	return getProgressClass().New()
}



// Creates a new progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil IDictionary /* already interface */) Progress {
	instance := getProgressClass().Alloc()
	rv := objc.Send[Progress](instance.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	rv.Autorelease()
	return rv
}


// Creates and returns a progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func NewProgressWithTotalUnitCount(unitCount unsafe.Pointer) Progress {
	rv := objc.Send[Progress](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}


// Creates a progress instance for the specified progress object with a unit count that’s a portion of the containing object’s total unit count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:parent:pendingUnitCount:)
func NewProgressWithTotalUnitCountParentPendingUnitCount(unitCount unsafe.Pointer, parent IProgress, portionOfParentTotalUnitCount unsafe.Pointer) Progress {
	rv := objc.Send[Progress](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:parent:pendingUnitCount:"), unitCount, parent, portionOfParentTotalUnitCount)
	return rv
}



// Registers a file URL to hear about the progress of a file operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/addSubscriber(forFileURL:withPublishingHandler:)
func (pc _ProgressClass) AddSubscriberForFileURLWithPublishingHandler(url IURL, publishingHandler ProgressPublishingHandler /* foo */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("addSubscriberForFileURL:withPublishingHandler:"), url, publishingHandler)
	return rv
}


// Returns the progress instance, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/current()
func (pc _ProgressClass) CurrentProgress() IProgress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("currentProgress"))
	return rv
}


// Creates and returns a progress instance with the specified unit count that isn’t part of any existing progress tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/discreteProgress(totalUnitCount:)
func (pc _ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("discreteProgressWithTotalUnitCount:"), unitCount)
	return rv
}


// Creates and returns a progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}


// Creates a progress instance for the specified progress object with a unit count that’s a portion of the containing object’s total unit count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:parent:pendingUnitCount:)
func (pc _ProgressClass) ProgressWithTotalUnitCountParentPendingUnitCount(unitCount unsafe.Pointer, parent IProgress, portionOfParentTotalUnitCount unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("progressWithTotalUnitCount:parent:pendingUnitCount:"), unitCount, parent, portionOfParentTotalUnitCount)
	return rv
}


// Removes a proxy progress object that the add subscriber method returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/removeSubscriber(_:)
func (pc _ProgressClass) RemoveSubscriber(subscriber objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("removeSubscriber:"), subscriber)
}


// Retrieves the current thread’s progress object, executes the specified block, and increments the progress object by the specified units of work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/performAsCurrentWithPendingUnitCount:usingBlock:
func (p_ Progress) PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount unsafe.Pointer, work unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performAsCurrentWithPendingUnitCount:usingBlock:"), unitCount, work)
}


// Adds a process object as a suboperation of a progress tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/addChild(_:withPendingUnitCount:)
func (p_ Progress) AddChildWithPendingUnitCount(child IProgress, inUnitCount unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addChild:withPendingUnitCount:"), child, inUnitCount)
}


// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/becomeCurrent(withPendingUnitCount:)
func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("becomeCurrentWithPendingUnitCount:"), unitCount)
}


// Cancels progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (p_ Progress) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}


// Pauses progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pause()
func (p_ Progress) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}


// Publishes the progress object for other processes to observe it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/publish()
func (p_ Progress) Publish() {
	objc.Send[objc.ID](p_.ID, objc.Sel("publish"))
}


// Restores the previous progress object to become the current progress object on the thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resignCurrent()
func (p_ Progress) ResignCurrent() {
	objc.Send[objc.ID](p_.ID, objc.Sel("resignCurrent"))
}


// Resumes progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resume()
func (p_ Progress) Resume() {
	objc.Send[objc.ID](p_.ID, objc.Sel("resume"))
}


// Sets a value in the user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/setUserInfoObject(_:forKey:)
func (p_ Progress) SetUserInfoObjectForKey(objectOrNil objectivec.IObject, key ProgressUserInfoKey /* foo */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfoObject:forKey:"), objectOrNil, key)
}


// Removes a progress object from publication, making it unobservable by other processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/unpublish()
func (p_ Progress) Unpublish() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unpublish"))
}


// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/estimatedTimeRemaining
func (p_ Progress) EstimatedTimeRemaining() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("estimatedTimeRemaining"))
	return rv
}


// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/estimatedTimeRemaining
func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEstimatedTimeRemaining:"), value)
}


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileCompletedCount
func (p_ Progress) FileCompletedCount() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("fileCompletedCount"))
	return rv
}


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileCompletedCount
func (p_ Progress) SetFileCompletedCount(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileCompletedCount:"), value)
}


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileTotalCount
func (p_ Progress) FileTotalCount() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("fileTotalCount"))
	return rv
}


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileTotalCount
func (p_ Progress) SetFileTotalCount(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileTotalCount:"), value)
}


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/throughput
func (p_ Progress) Throughput() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("throughput"))
	return rv
}


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/throughput
func (p_ Progress) SetThroughput(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThroughput:"), value)
}


// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancellationHandler
func (p_ Progress) CancellationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cancellationHandler"))
	return rv
}


// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancellationHandler
func (p_ Progress) SetCancellationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCancellationHandler:"), value)
}


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) CompletedUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("completedUnitCount"))
	return rv
}


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) SetCompletedUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCompletedUnitCount:"), value)
}


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
func (p_ Progress) FileOperationKind() ProgressFileOperationKind /* foo */ {
	rv := objc.Send[ProgressFileOperationKind](p_.ID, objc.Sel("fileOperationKind"))
	return rv
}


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind /* foo */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileOperationKind:"), value)
}


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileURL
func (p_ Progress) FileURL() IURL {
	rv := objc.Send[URL](p_.ID, objc.Sel("fileURL"))
	return rv
}


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileURL
func (p_ Progress) SetFileURL(value IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileURL:"), value)
}


// The fraction of the overall work that the progress object completes, including work from its suboperations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fractionCompleted
func (p_ Progress) FractionCompleted() float64 /* primitive/slice/pointer */ {
	rv := objc.Send[float64](p_.ID, objc.Sel("fractionCompleted"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancellable
func (p_ Progress) Cancellable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("cancellable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancellable
func (p_ Progress) SetCancellable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCancellable:"), value)
}


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancelled
func (p_ Progress) Cancelled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("cancelled"))
	return rv
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isFinished
func (p_ Progress) Finished() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("finished"))
	return rv
}


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isIndeterminate
func (p_ Progress) Indeterminate() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("indeterminate"))
	return rv
}


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isOld
func (p_ Progress) Old() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("old"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPausable
func (p_ Progress) Pausable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("pausable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPausable
func (p_ Progress) SetPausable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPausable:"), value)
}


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPaused
func (p_ Progress) Paused() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("paused"))
	return rv
}


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/kind
func (p_ Progress) Kind() ProgressKind /* foo */ {
	rv := objc.Send[ProgressKind](p_.ID, objc.Sel("kind"))
	return rv
}


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/kind
func (p_ Progress) SetKind(value ProgressKind /* foo */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKind:"), value)
}


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedAdditionalDescription
func (p_ Progress) LocalizedAdditionalDescription() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedAdditionalDescription"))
	return rv
}


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedAdditionalDescription
func (p_ Progress) SetLocalizedAdditionalDescription(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedAdditionalDescription:"), objc.String(value))
}


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedDescription
func (p_ Progress) LocalizedDescription() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedDescription
func (p_ Progress) SetLocalizedDescription(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pausingHandler
func (p_ Progress) PausingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pausingHandler"))
	return rv
}


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pausingHandler
func (p_ Progress) SetPausingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPausingHandler:"), value)
}


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resumingHandler
func (p_ Progress) ResumingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resumingHandler"))
	return rv
}


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resumingHandler
func (p_ Progress) SetResumingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResumingHandler:"), value)
}


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) TotalUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("totalUnitCount"))
	return rv
}


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) SetTotalUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTotalUnitCount:"), value)
}


// A dictionary of arbitrary values for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/userInfo
func (p_ Progress) UserInfo() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](p_.ID, objc.Sel("userInfo"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) IsCancellable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancellable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) SetIsCancellable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancellable:"), value)
}


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) IsCancelled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) SetIsCancelled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) IsFinished() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFinished"))
	return rv
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) SetIsFinished(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFinished:"), value)
}


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) IsIndeterminate() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) SetIsIndeterminate(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) IsOld() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOld"))
	return rv
}


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) SetIsOld(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOld:"), value)
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) IsPausable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPausable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) SetIsPausable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPausable:"), value)
}


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) IsPaused() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) SetIsPaused(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPaused:"), value)
}


