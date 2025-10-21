// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Download] class.
var (
	DownloadClass     _DownloadClass
	DownloadClassOnce sync.Once
)

func getDownloadClass() _DownloadClass {
	DownloadClassOnce.Do(func() {
		DownloadClass = _DownloadClass{objc.GetClass("SKDownload")}
	})
	return DownloadClass
}

type _DownloadClass struct {
	class objc.Class
}

// An interface definition for the [Download] class.
type IDownload interface {
	objectivec.IObject
}

// Downloadable content associated with a product.
//
// When you create a product in App Store Connect, you can associate one or more pieces of downloadable content with it. At runtime, when a product is purchased by a user, your app uses objects to download the content from the App Store. Your app never directly creates a object. Instead, after a payment is processed, your app reads the transaction object’s property to retrieve an array of objects associated with the transaction. To download the content, you queue a download object on the payment queue and wait for the content to be downloaded. After a download completes, read the download object’s property to get a URL to the downloaded content. Your app must process the downloaded file before completing the transaction. For example, it might copy the file into a directory whose contents are persistent. When all downloads are complete, you finish the transaction. After the transaction is finished, the download objects cannot be queued to the payment queue and any URLs to the downloaded content are invalid.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload
type Download struct {
	objectivec.Object
}

// DownloadFrom constructs a [Download] from an unsafe.Pointer.
//
// Downloadable content associated with a product.
func DownloadFrom(ptr unsafe.Pointer) Download {
	return Download{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DownloadClass) Alloc() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DownloadClass) New() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Download) Init() Download {
	rv := objc.Send[Download](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Download) Autorelease() Download {
	rv := objc.Send[Download](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDownload creates a new Download instance.
func NewDownload() Download {
	return getDownloadClass().New()
}


// Returns the local location for the previously downloaded flie.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentURL(forProductID:)
func (dc _DownloadClass) ContentURLForProductID(productID string) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(dc.class), objc.Sel("contentURLForProductID:"), objc.String(productID))
	return rv
}

// Deletes the previously downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/deleteContent(forProductID:)
func (dc _DownloadClass) DeleteContentForProductID(productID string) {
	objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("deleteContentForProductID:"), objc.String(productID))
}

// The current state of the download object.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/state
func (d_ Download) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The current state of the download object.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/state
func (d_ Download) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setState:"), value)
}

// A string that uniquely identifies the downloadable content.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentidentifier
func (d_ Download) ContentIdentifier() string {
	rv := objc.Send[string](d_.ID, objc.Sel("contentIdentifier"))
	return rv
}


// SetContentIdentifier sets the value of the contentIdentifier property.
// A string that uniquely identifies the downloadable content.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentidentifier
func (d_ Download) SetContentIdentifier(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentIdentifier:"), objc.String(value))
}

// The length of the downloadable content, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentlength
func (d_ Download) ContentLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contentLength"))
	return rv
}


// SetContentLength sets the value of the contentLength property.
// The length of the downloadable content, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentlength
func (d_ Download) SetContentLength(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentLength:"), value)
}

// A string that identifies which version of the content is available for download.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentversion
func (d_ Download) ContentVersion() string {
	rv := objc.Send[string](d_.ID, objc.Sel("contentVersion"))
	return rv
}


// SetContentVersion sets the value of the contentVersion property.
// A string that identifies which version of the content is available for download.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/contentversion
func (d_ Download) SetContentVersion(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentVersion:"), objc.String(value))
}

// Indicates that the system cannot determine how much time is needed to finish downloading the content.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownloadtimeremainingunknown
func (d_ Download) SKDownloadTimeRemainingUnknown() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("SKDownloadTimeRemainingUnknown"))
	return rv
}


// SetSKDownloadTimeRemainingUnknown sets the value of the SKDownloadTimeRemainingUnknown property.
// Indicates that the system cannot determine how much time is needed to finish downloading the content.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownloadtimeremainingunknown
func (d_ Download) SetSKDownloadTimeRemainingUnknown(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSKDownloadTimeRemainingUnknown:"), value)
}

// The length of the downloadable content, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/expectedcontentlength
func (d_ Download) ExpectedContentLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("expectedContentLength"))
	return rv
}


// SetExpectedContentLength sets the value of the expectedContentLength property.
// The length of the downloadable content, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownload/expectedcontentlength
func (d_ Download) SetExpectedContentLength(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setExpectedContentLength:"), value)
}

// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (d_ Download) Downloads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("downloads"))
	return rv
}


// SetDownloads sets the value of the downloads property.
// An array of download objects representing the downloadable content associated with the transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (d_ Download) SetDownloads(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDownloads:"), value)
}

// The local location of the downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentURL
func (d_ Download) ContentURL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("contentURL"))
	return rv
}

// The current state of the download object.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/downloadState
func (d_ Download) DownloadState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("downloadState"))
	return rv
}

// The error that prevented the content from being downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/error
func (d_ Download) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("error"))
	return rv
}

// A value that indicates how much of the file has been downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/progress
func (d_ Download) Progress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("progress"))
	return rv
}

// An estimated time, in seconds, to finish downloading the content.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/timeRemaining
func (d_ Download) TimeRemaining() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](d_.ID, objc.Sel("timeRemaining"))
	return rv
}

// The transaction associated with the downloadable file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/transaction
func (d_ Download) Transaction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("transaction"))
	return rv
}



