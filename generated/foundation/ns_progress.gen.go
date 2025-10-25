// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSProgress */


/* debug [class_header]: Header for NSProgress */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Progress */
// An interface definition for the [Progress] class.
type IProgress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Progress */
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
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	FileOperationKind() ProgressFileOperationKind
	SetFileOperationKind(value ProgressFileOperationKind)
	FileURL() IURL
	SetFileURL(value IURL)
	FractionCompleted() float64
	Cancellable() bool
	SetCancellable(value bool)
	Cancelled() bool
	Finished() bool
	Indeterminate() bool
	Old() bool
	Pausable() bool
	SetPausable(value bool)
	Paused() bool
	Kind() ProgressKind
	SetKind(value ProgressKind)
	LocalizedAdditionalDescription() IString
	SetLocalizedAdditionalDescription(value IString)
	LocalizedDescription() IString
	SetLocalizedDescription(value IString)
	PausingHandler() unsafe.Pointer
	SetPausingHandler(value unsafe.Pointer)
	ResumingHandler() unsafe.Pointer
	SetResumingHandler(value unsafe.Pointer)
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	UserInfo() IDictionary
	IsCancellable() bool
	SetIsCancellable(value bool)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsFinished() bool
	SetIsFinished(value bool)
	IsIndeterminate() bool
	SetIsIndeterminate(value bool)
	IsOld() bool
	SetIsOld(value bool)
	IsPausable() bool
	SetIsPausable(value bool)
	IsPaused() bool
	SetIsPaused(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Progress */
	// methods:
	Cancel()
	Pause()
	Publish()
	Resume()
	SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey)
	Unpublish()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Progress */
// Alloc allocates a new instance without initialization.
func (pc _ProgressClass) Alloc() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Progress */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Progress */

// Creates a new progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil IDictionary) Progress {
	instance := getProgressClass().Alloc()
	rv := objc.Send[Progress](instance.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewProgressWithParentUserInfo */


// Creates and returns a progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func NewProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.Send[Progress](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}/* debug [class_init_methods/constructor]: NewProgressWithTotalUnitCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Progress */

// Registers a file URL to hear about the progress of a file operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/addSubscriber(forFileURL:withPublishingHandler:)
func (pc _ProgressClass) AddSubscriberForFileURLWithPublishingHandler(url IURL, publishingHandler ProgressPublishingHandler /* not a class type */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("addSubscriberForFileURL:withPublishingHandler:"), url, publishingHandler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddSubscriberForFileURLWithPublishingHandler) */


// Creates and returns a progress instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount int64) IProgress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ProgressWithTotalUnitCount) */


// Removes a proxy progress object that the add subscriber method returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/removeSubscriber(_:)
func (pc _ProgressClass) RemoveSubscriber(subscriber objc.IObject) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("removeSubscriber:"), subscriber)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveSubscriber) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Progress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Progress */

// Cancels progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (p_ Progress) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Pauses progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pause()
func (p_ Progress) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Publishes the progress object for other processes to observe it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/publish()
func (p_ Progress) Publish() {
	objc.Send[objc.ID](p_.ID, objc.Sel("publish"))
}/* debug [instance_methods/method]: Publish */


// Resumes progress tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resume()
func (p_ Progress) Resume() {
	objc.Send[objc.ID](p_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Sets a value in the user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/setUserInfoObject(_:forKey:)
func (p_ Progress) SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfoObject:forKey:"), objectOrNil, key)
}/* debug [instance_methods/method]: SetUserInfoObjectForKey */


// Removes a progress object from publication, making it unobservable by other processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/unpublish()
func (p_ Progress) Unpublish() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unpublish"))
}/* debug [instance_methods/method]: Unpublish */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Progress */

// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/estimatedTimeRemaining
func (p_ Progress) EstimatedTimeRemaining() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("estimatedTimeRemaining"))
	return rv
}/* debug [instance_properties/getter]: estimatedTimeRemaining */


// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/estimatedTimeRemaining
func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEstimatedTimeRemaining:"), value)
}/* debug [instance_properties/setter]: estimatedTimeRemaining */


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileCompletedCount
func (p_ Progress) FileCompletedCount() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("fileCompletedCount"))
	return rv
}/* debug [instance_properties/getter]: fileCompletedCount */


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileCompletedCount
func (p_ Progress) SetFileCompletedCount(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileCompletedCount:"), value)
}/* debug [instance_properties/setter]: fileCompletedCount */


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileTotalCount
func (p_ Progress) FileTotalCount() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("fileTotalCount"))
	return rv
}/* debug [instance_properties/getter]: fileTotalCount */


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/fileTotalCount
func (p_ Progress) SetFileTotalCount(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileTotalCount:"), value)
}/* debug [instance_properties/setter]: fileTotalCount */


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/throughput
func (p_ Progress) Throughput() INumber {
	rv := objc.Send[Number](p_.ID, objc.Sel("throughput"))
	return rv
}/* debug [instance_properties/getter]: throughput */


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProgress/throughput
func (p_ Progress) SetThroughput(value INumber) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThroughput:"), value)
}/* debug [instance_properties/setter]: throughput */


// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancellationHandler
func (p_ Progress) CancellationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cancellationHandler"))
	return rv
}/* debug [instance_properties/getter]: cancellationHandler */


// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/cancellationHandler
func (p_ Progress) SetCancellationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCancellationHandler:"), value)
}/* debug [instance_properties/setter]: cancellationHandler */


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) CompletedUnitCount() int64 {
	rv := objc.Send[int64](p_.ID, objc.Sel("completedUnitCount"))
	return rv
}/* debug [instance_properties/getter]: completedUnitCount */


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p_ Progress) SetCompletedUnitCount(value int64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCompletedUnitCount:"), value)
}/* debug [instance_properties/setter]: completedUnitCount */


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
func (p_ Progress) FileOperationKind() ProgressFileOperationKind {
	rv := objc.Send[ProgressFileOperationKind](p_.ID, objc.Sel("fileOperationKind"))
	return rv
}/* debug [instance_properties/getter]: fileOperationKind */


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileOperationKind:"), value)
}/* debug [instance_properties/setter]: fileOperationKind */


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileURL
func (p_ Progress) FileURL() IURL {
	rv := objc.Send[URL](p_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fileURL
func (p_ Progress) SetFileURL(value IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileURL:"), value)
}/* debug [instance_properties/setter]: fileURL */


// The fraction of the overall work that the progress object completes, including work from its suboperations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/fractionCompleted
func (p_ Progress) FractionCompleted() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("fractionCompleted"))
	return rv
}/* debug [instance_properties/getter]: fractionCompleted */


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancellable
func (p_ Progress) Cancellable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cancellable"))
	return rv
}/* debug [instance_properties/getter]: cancellable */


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancellable
func (p_ Progress) SetCancellable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCancellable:"), value)
}/* debug [instance_properties/setter]: cancellable */


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isCancelled
func (p_ Progress) Cancelled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isFinished
func (p_ Progress) Finished() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finished"))
	return rv
}/* debug [instance_properties/getter]: finished */


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isIndeterminate
func (p_ Progress) Indeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("indeterminate"))
	return rv
}/* debug [instance_properties/getter]: indeterminate */


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isOld
func (p_ Progress) Old() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("old"))
	return rv
}/* debug [instance_properties/getter]: old */


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPausable
func (p_ Progress) Pausable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pausable"))
	return rv
}/* debug [instance_properties/getter]: pausable */


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPausable
func (p_ Progress) SetPausable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPausable:"), value)
}/* debug [instance_properties/setter]: pausable */


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/isPaused
func (p_ Progress) Paused() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/kind
func (p_ Progress) Kind() ProgressKind {
	rv := objc.Send[ProgressKind](p_.ID, objc.Sel("kind"))
	return rv
}/* debug [instance_properties/getter]: kind */


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/kind
func (p_ Progress) SetKind(value ProgressKind) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKind:"), value)
}/* debug [instance_properties/setter]: kind */


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedAdditionalDescription
func (p_ Progress) LocalizedAdditionalDescription() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("localizedAdditionalDescription"))
	return rv
}/* debug [instance_properties/getter]: localizedAdditionalDescription */


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedAdditionalDescription
func (p_ Progress) SetLocalizedAdditionalDescription(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedAdditionalDescription:"), value)
}/* debug [instance_properties/setter]: localizedAdditionalDescription */


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedDescription
func (p_ Progress) LocalizedDescription() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("localizedDescription"))
	return rv
}/* debug [instance_properties/getter]: localizedDescription */


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/localizedDescription
func (p_ Progress) SetLocalizedDescription(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedDescription:"), value)
}/* debug [instance_properties/setter]: localizedDescription */


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pausingHandler
func (p_ Progress) PausingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pausingHandler"))
	return rv
}/* debug [instance_properties/getter]: pausingHandler */


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/pausingHandler
func (p_ Progress) SetPausingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPausingHandler:"), value)
}/* debug [instance_properties/setter]: pausingHandler */


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resumingHandler
func (p_ Progress) ResumingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resumingHandler"))
	return rv
}/* debug [instance_properties/getter]: resumingHandler */


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/resumingHandler
func (p_ Progress) SetResumingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResumingHandler:"), value)
}/* debug [instance_properties/setter]: resumingHandler */


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) TotalUnitCount() int64 {
	rv := objc.Send[int64](p_.ID, objc.Sel("totalUnitCount"))
	return rv
}/* debug [instance_properties/getter]: totalUnitCount */


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p_ Progress) SetTotalUnitCount(value int64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTotalUnitCount:"), value)
}/* debug [instance_properties/setter]: totalUnitCount */


// A dictionary of arbitrary values for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Progress/userInfo
func (p_ Progress) UserInfo() IDictionary {
	rv := objc.Send[Dictionary](p_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) IsCancellable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancellable"))
	return rv
}/* debug [instance_properties/getter]: isCancellable */


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) SetIsCancellable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancellable:"), value)
}/* debug [instance_properties/setter]: isCancellable */


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) IsCancelled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) SetIsCancelled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) IsFinished() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFinished"))
	return rv
}/* debug [instance_properties/getter]: isFinished */


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) SetIsFinished(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFinished:"), value)
}/* debug [instance_properties/setter]: isFinished */


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) IsIndeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}/* debug [instance_properties/getter]: isIndeterminate */


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) SetIsIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}/* debug [instance_properties/setter]: isIndeterminate */


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) IsOld() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOld"))
	return rv
}/* debug [instance_properties/getter]: isOld */


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) SetIsOld(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOld:"), value)
}/* debug [instance_properties/setter]: isOld */


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) IsPausable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPausable"))
	return rv
}/* debug [instance_properties/getter]: isPausable */


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) SetIsPausable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPausable:"), value)
}/* debug [instance_properties/setter]: isPausable */


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) IsPaused() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) SetIsPaused(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSProgress */


