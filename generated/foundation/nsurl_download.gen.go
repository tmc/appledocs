// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLDownload */


/* debug [class_header]: Header for NSURLDownload */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLDownload */
// An interface definition for the [URLDownload] class.
type IURLDownload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLDownload */
	// properties:
	DeletesFileUponFailure() bool
	SetDeletesFileUponFailure(value bool)
	Request() IURLRequest
	ResumeData() IData
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLDownload */
	// methods:
	Cancel()
	SetDestinationAllowOverwrite(path IString, allowOverwrite bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLDownload */
// Alloc allocates a new instance without initialization.
func (uc _URLDownloadClass) Alloc() URLDownload {
	rv := objc.Send[URLDownload](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLDownload */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLDownload */

// Returns an initialized URL download for a URL request and begins to download the data for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/init(request:delegate:)
func NewURLDownloadWithRequestDelegate(request IURLRequest, delegate unsafe.Pointer) URLDownload {
	instance := getURLDownloadClass().Alloc()
	rv := objc.Send[URLDownload](instance.ID, objc.Sel("initWithRequest:delegate:"), request, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLDownloadWithRequestDelegate */


// Returns an initialized NSURLDownload object that will resume downloading the specified data to the specified file and begins the download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/init(resumeData:delegate:path:)
func NewURLDownloadWithResumeDataDelegatePath(resumeData IData, delegate unsafe.Pointer, path IString) URLDownload {
	instance := getURLDownloadClass().Alloc()
	rv := objc.Send[URLDownload](instance.ID, objc.Sel("initWithResumeData:delegate:path:"), resumeData, delegate, path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLDownloadWithResumeDataDelegatePath */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLDownload */

// Returns whether a URL download object can resume a download that was decoded with the specified MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/canResumeDownloadDecoded(withEncodingMIMEType:)
func (uc _URLDownloadClass) CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType IString) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canResumeDownloadDecodedWithEncodingMIMEType:"), MIMEType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanResumeDownloadDecodedWithEncodingMIMEType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLDownload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLDownload */

// Cancels the receiver’s download and deletes the downloaded file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/cancel()
func (u_ URLDownload) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Sets the destination path of the downloaded file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/setDestination(_:allowOverwrite:)
func (u_ URLDownload) SetDestinationAllowOverwrite(path IString, allowOverwrite bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDestination:allowOverwrite:"), path, allowOverwrite)
}/* debug [instance_methods/method]: SetDestinationAllowOverwrite */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLDownload */

// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/deletesFileUponFailure
func (u_ URLDownload) DeletesFileUponFailure() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("deletesFileUponFailure"))
	return rv
}/* debug [instance_properties/getter]: deletesFileUponFailure */


// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/deletesFileUponFailure
func (u_ URLDownload) SetDeletesFileUponFailure(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeletesFileUponFailure:"), value)
}/* debug [instance_properties/setter]: deletesFileUponFailure */


// Returns the request that initiated the receiver’s download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/request
func (u_ URLDownload) Request() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}/* debug [instance_properties/getter]: request */


// Returns the resume data for a download that is not yet complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload/resumeData
func (u_ URLDownload) ResumeData() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("resumeData"))
	return rv
}/* debug [instance_properties/getter]: resumeData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLDownload */


