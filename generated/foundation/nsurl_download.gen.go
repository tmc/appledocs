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
	Cancel()
	SetDestinationAllowOverwrite(path string, allowOverwrite bool)
	DeletesFileUponFailure() bool
	SetDeletesFileUponFailure(value bool)
	Request() NSURLRequest
	ResumeData() NSData
}

// An object that downloads a resource asynchronously and saves the data to a file.
//
// The interface for provides methods to initialize a download, set the destination path and cancel loading the request. The delegate object assigned to each instance of this class should implement the methods defined by the protocol. These methods provide the delegate with the current status of in-progress asynchronous downloads and allow the delegate to customize the URL loading process. These delegate methods are called on the thread that started the asynchronous load operation for the associated object.
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





// Returns an initialized URL download for a URL request and begins to download the data for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/init(request:delegate:)

func NewURLDownloadWithRequestDelegate(request IURLRequest, delegate objectivec.IObject) URLDownload {
	instance := getURLDownloadClass().Alloc()
	rv := objc.Send[URLDownload](instance.ID, objc.Sel("initWithRequest:delegate:"), request, delegate)
	rv.Autorelease()
	return rv
}




// Returns an initialized NSURLDownload object that will resume downloading the specified data to the specified file and begins the download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/init(resumeData:delegate:path:)

func NewURLDownloadWithResumeDataDelegatePath(resumeData IData, delegate objectivec.IObject, path string) URLDownload {
	instance := getURLDownloadClass().Alloc()
	rv := objc.Send[URLDownload](instance.ID, objc.Sel("initWithResumeData:delegate:path:"), resumeData, delegate, objc.String(path))
	rv.Autorelease()
	return rv
}



// Returns whether a URL download object can resume a download that was decoded with the specified MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/canResumeDownloadDecoded(withEncodingMIMEType:)

func (uc _URLDownloadClass) CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType string) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canResumeDownloadDecodedWithEncodingMIMEType:"), objc.String(MIMEType))
	return rv
}

// Cancels the receiver’s download and deletes the downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/cancel()
func (u_ URLDownload) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}

// Sets the destination path of the downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/setDestination(_:allowOverwrite:)
func (u_ URLDownload) SetDestinationAllowOverwrite(path string, allowOverwrite bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDestination:allowOverwrite:"), objc.String(path), allowOverwrite)
}

// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/deletesFileUponFailure
func (u_ URLDownload) DeletesFileUponFailure() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("deletesFileUponFailure"))
	return rv
}


// SetDeletesFileUponFailure sets the value of the deletesFileUponFailure property.
// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/deletesFileUponFailure
func (u_ URLDownload) SetDeletesFileUponFailure(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletesFileUponFailure:"), value)
}

// Returns the request that initiated the receiver’s download.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/request
func (u_ URLDownload) Request() NSURLRequest {
	rv := objc.Send[NSURLRequest](u_.ID, objc.Sel("request"))
	return rv
}

// Returns the resume data for a download that is not yet complete.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/resumeData
func (u_ URLDownload) ResumeData() NSData {
	rv := objc.Send[NSData](u_.ID, objc.Sel("resumeData"))
	return rv
}


