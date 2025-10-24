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
	CancellationHandler() unsafe.Pointer
	SetCancellationHandler(value unsafe.Pointer)
	CompletedUnitCount() unsafe.Pointer
	SetCompletedUnitCount(value unsafe.Pointer)
	EstimatedTimeRemaining() float64
	SetEstimatedTimeRemaining(value float64)
	FileCompletedCount() int
	SetFileCompletedCount(value int)
	FileOperationKind() unsafe.Pointer
	SetFileOperationKind(value unsafe.Pointer)
	FileTotalCount() int
	SetFileTotalCount(value int)
	FileURL() IURL
	SetFileURL(value IURL)
	FractionCompleted() float64
	SetFractionCompleted(value float64)
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
	Kind() unsafe.Pointer
	SetKind(value unsafe.Pointer)
	LocalizedAdditionalDescription() IString
	SetLocalizedAdditionalDescription(value IString)
	LocalizedDescription() IString
	SetLocalizedDescription(value IString)
	PausingHandler() unsafe.Pointer
	SetPausingHandler(value unsafe.Pointer)
	ResumingHandler() unsafe.Pointer
	SetResumingHandler(value unsafe.Pointer)
	Throughput() int
	SetThroughput(value int)
	TotalUnitCount() unsafe.Pointer
	SetTotalUnitCount(value unsafe.Pointer)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
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



// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/cancellationhandler
func (p_ Progress) CancellationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cancellationHandler"))
	return rv
}


// The block to invoke when canceling progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/cancellationhandler
func (p_ Progress) SetCancellationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCancellationHandler:"), value)
}


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/completedunitcount
func (p_ Progress) CompletedUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("completedUnitCount"))
	return rv
}


// The number of completed units of work for the current job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/completedunitcount
func (p_ Progress) SetCompletedUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCompletedUnitCount:"), value)
}


// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/estimatedtimeremaining
func (p_ Progress) EstimatedTimeRemaining() float64 {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("estimatedTimeRemaining"))
	return rv
}


// A value that indicates the estimated amount of time remaining to complete the progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/estimatedtimeremaining
func (p_ Progress) SetEstimatedTimeRemaining(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEstimatedTimeRemaining:"), value)
}


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/filecompletedcount
func (p_ Progress) FileCompletedCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("fileCompletedCount"))
	return rv
}


// The number of completed files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/filecompletedcount
func (p_ Progress) SetFileCompletedCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileCompletedCount:"), value)
}


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fileoperationkind-swift.property
func (p_ Progress) FileOperationKind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fileOperationKind"))
	return rv
}


// The kind of file operation for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fileoperationkind-swift.property
func (p_ Progress) SetFileOperationKind(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileOperationKind:"), value)
}


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/filetotalcount
func (p_ Progress) FileTotalCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("fileTotalCount"))
	return rv
}


// The total number of files for a file progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/filetotalcount
func (p_ Progress) SetFileTotalCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileTotalCount:"), value)
}


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fileurl
func (p_ Progress) FileURL() IURL {
	rv := objc.Send[URL](p_.ID, objc.Sel("fileURL"))
	return rv
}


// A URL that represents the file for the current progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fileurl
func (p_ Progress) SetFileURL(value IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileURL:"), value)
}


// The fraction of the overall work that the progress object completes, including work from its suboperations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fractioncompleted
func (p_ Progress) FractionCompleted() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("fractionCompleted"))
	return rv
}


// The fraction of the overall work that the progress object completes, including work from its suboperations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/fractioncompleted
func (p_ Progress) SetFractionCompleted(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFractionCompleted:"), value)
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) IsCancellable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancellable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancellable
func (p_ Progress) SetIsCancellable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancellable:"), value)
}


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) IsCancelled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value that Indicates whether the receiver is tracking canceled work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/iscancelled
func (p_ Progress) SetIsCancelled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) IsFinished() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFinished"))
	return rv
}


// A Boolean value that indicates the progress object is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isfinished
func (p_ Progress) SetIsFinished(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFinished:"), value)
}


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) IsIndeterminate() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isIndeterminate"))
	return rv
}


// A Boolean value that indicates whether the tracked progress is indeterminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isindeterminate
func (p_ Progress) SetIsIndeterminate(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsIndeterminate:"), value)
}


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) IsOld() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOld"))
	return rv
}


// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/isold
func (p_ Progress) SetIsOld(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOld:"), value)
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) IsPausable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPausable"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking work that you can pause.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispausable
func (p_ Progress) SetIsPausable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPausable:"), value)
}


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) IsPaused() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether the receiver is tracking paused work.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/ispaused
func (p_ Progress) SetIsPaused(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsPaused:"), value)
}


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/kind
func (p_ Progress) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("kind"))
	return rv
}


// An object that represents the kind of progress for the progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/kind
func (p_ Progress) SetKind(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKind:"), value)
}


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/localizedadditionaldescription
func (p_ Progress) LocalizedAdditionalDescription() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("localizedAdditionalDescription"))
	return rv
}


// A more specific localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/localizedadditionaldescription
func (p_ Progress) SetLocalizedAdditionalDescription(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedAdditionalDescription:"), value)
}


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/localizeddescription
func (p_ Progress) LocalizedDescription() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A localized description of tracked progress for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/localizeddescription
func (p_ Progress) SetLocalizedDescription(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/pausinghandler
func (p_ Progress) PausingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pausingHandler"))
	return rv
}


// The block to invoke when pausing progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/pausinghandler
func (p_ Progress) SetPausingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPausingHandler:"), value)
}


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/resuminghandler
func (p_ Progress) ResumingHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resumingHandler"))
	return rv
}


// The block to invoke when progress resumes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/resuminghandler
func (p_ Progress) SetResumingHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResumingHandler:"), value)
}


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/throughput
func (p_ Progress) Throughput() int {
	rv := objc.Send[int](p_.ID, objc.Sel("throughput"))
	return rv
}


// A value that represents the speed of data processing, in bytes per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/throughput
func (p_ Progress) SetThroughput(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThroughput:"), value)
}


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/totalunitcount
func (p_ Progress) TotalUnitCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("totalUnitCount"))
	return rv
}


// The total number of tracked units of work for the current progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/totalunitcount
func (p_ Progress) SetTotalUnitCount(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTotalUnitCount:"), value)
}


// A dictionary of arbitrary values for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/userinfo
func (p_ Progress) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("userInfo"))
	return rv
}


// A dictionary of arbitrary values for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/progress/userinfo
func (p_ Progress) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfo:"), value)
}



