// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLDownload] class.
var (
	URLDownloadClass     _URLDownloadClass
	URLDownloadClassOnce sync.Once
)

func getURLDownloadClass() _URLDownloadClass {
	URLDownloadClassOnce.Do(func() {
		URLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}
	})
	return URLDownloadClass
}

type _URLDownloadClass struct {
	class objc.Class
}

// An interface definition for the [URLDownload] class.
type IURLDownload interface {
	objectivec.IObject
	// properties:
	DeletesFileUponFailure() bool /* primitive/slice/pointer */
	SetDeletesFileUponFailure(value bool /* primitive/slice/pointer */)
	Request() IURLRequest
	SetRequest(value IURLRequest)
	ResumeData() IData
	SetResumeData(value IData)
	// methods:
}

// An object that downloads a resource asynchronously and saves the data to a file.
//
// The interface for provides methods to initialize a download, set the destination path and cancel loading the request. The delegate object assigned to each instance of this class should implement the methods defined by the protocol. These methods provide the delegate with the current status of in-progress asynchronous downloads and allow the delegate to customize the URL loading process. These delegate methods are called on the thread that started the asynchronous load operation for the associated object.


// An object that downloads a resource asynchronously and saves the data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload
type URLDownload struct {
	objectivec.Object
}

// URLDownloadFrom constructs a [URLDownload] from an unsafe.Pointer.
//
// An object that downloads a resource asynchronously and saves the data to a file.
func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLDownloadClass) Alloc() URLDownload {
	rv := objc.Send[URLDownload](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLDownloadClass) New() URLDownload {
	rv := objc.Send[URLDownload](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLDownload) Init() URLDownload {
	rv := objc.Send[URLDownload](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLDownload) Autorelease() URLDownload {
	rv := objc.Send[URLDownload](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLDownload creates a new URLDownload instance.
func NewURLDownload() URLDownload {
	return getURLDownloadClass().New()
}



// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/deletesfileuponfailure
func (u_ URLDownload) DeletesFileUponFailure() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("deletesFileUponFailure"))
	return rv
}


// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/deletesfileuponfailure
func (u_ URLDownload) SetDeletesFileUponFailure(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletesFileUponFailure:"), value)
}


// Returns the request that initiated the receiver’s download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/request
func (u_ URLDownload) Request() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}


// Returns the request that initiated the receiver’s download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/request
func (u_ URLDownload) SetRequest(value IURLRequest) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequest:"), value)
}


// Returns the resume data for a download that is not yet complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/resumedata
func (u_ URLDownload) ResumeData() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("resumeData"))
	return rv
}


// Returns the resume data for a download that is not yet complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurldownload/resumedata
func (u_ URLDownload) SetResumeData(value IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResumeData:"), value)
}



