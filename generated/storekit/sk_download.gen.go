// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKDownload */

/* debug [class_header]: Header for SKDownload */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for Download */
// An interface definition for the [Download] class.
type IDownload interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for Download */
	// properties:
	ContentIdentifier() objc.IObject /* cross-framework: NSString */
	ContentLength() unsafe.Pointer
	ContentURL() objc.IObject     /* cross-framework: NSURL */
	ContentVersion() objc.IObject /* cross-framework: NSString */
	Error() objc.IObject          /* cross-framework: Error */
	ExpectedContentLength() unsafe.Pointer
	Progress() float32
	State() DownloadState
	TimeRemaining() float64
	Transaction() ISKPaymentTransaction
	SKDownloadTimeRemainingUnknown() float64
	SetSKDownloadTimeRemainingUnknown(value float64)
	Downloads() ISKDownload
	SetDownloads(value ISKDownload)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for Download */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for Download */
// Alloc allocates a new instance without initialization.
func (dc _DownloadClass) Alloc() Download {
	rv := objc.Send[Download](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for Download */
// Downloadable content associated with a product.
//
// When you create a product in App Store Connect, you can associate one or more pieces of downloadable content with it. At runtime, when a product is purchased by a user, your app uses objects to download the content from the App Store. Your app never directly creates a object. Instead, after a payment is processed, your app reads the transaction object’s property to retrieve an array of objects associated with the transaction. To download the content, you queue a download object on the payment queue and wait for the content to be downloaded. After a download completes, read the download object’s property to get a URL to the downloaded content. Your app must process the downloaded file before completing the transaction. For example, it might copy the file into a directory whose contents are persistent. When all downloads are complete, you finish the transaction. After the transaction is finished, the download objects cannot be queued to the payment queue and any URLs to the downloaded content are invalid.

// Downloadable content associated with a product.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for Download */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for Download */

// Returns the local location for the previously downloaded flie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentURL(forProductID:)
func (dc _DownloadClass) ContentURLForProductID(productID objc.IObject /* cross-framework: NSString */) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(dc.class), objc.Sel("contentURLForProductID:"), productID)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ContentURLForProductID) */

// Deletes the previously downloaded file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/deleteContent(forProductID:)
func (dc _DownloadClass) DeleteContentForProductID(productID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("deleteContentForProductID:"), productID)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=DeleteContentForProductID) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for Download */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for Download */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for Download */

// A string that uniquely identifies the downloadable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentIdentifier
func (d_ Download) ContentIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("contentIdentifier"))
	return rv
} /* debug [instance_properties/getter]: contentIdentifier */

// The length of the downloadable content, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentLength
func (d_ Download) ContentLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contentLength"))
	return rv
} /* debug [instance_properties/getter]: contentLength */

// The local location of the downloaded file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentURL
func (d_ Download) ContentURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("contentURL"))
	return rv
} /* debug [instance_properties/getter]: contentURL */

// A string that identifies which version of the content is available for download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentVersion
func (d_ Download) ContentVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("contentVersion"))
	return rv
} /* debug [instance_properties/getter]: contentVersion */

// The error that prevented the content from being downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/error
func (d_ Download) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](d_.ID, objc.Sel("error"))
	return rv
} /* debug [instance_properties/getter]: error */

// The length of the downloadable content, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/expectedContentLength
func (d_ Download) ExpectedContentLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("expectedContentLength"))
	return rv
} /* debug [instance_properties/getter]: expectedContentLength */

// A value that indicates how much of the file has been downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/progress
func (d_ Download) Progress() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("progress"))
	return rv
} /* debug [instance_properties/getter]: progress */

// The current state of the download object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/state
func (d_ Download) State() DownloadState {
	rv := objc.Send[DownloadState](d_.ID, objc.Sel("state"))
	return rv
} /* debug [instance_properties/getter]: state */

// An estimated time, in seconds, to finish downloading the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/timeRemaining
func (d_ Download) TimeRemaining() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("timeRemaining"))
	return rv
} /* debug [instance_properties/getter]: timeRemaining */

// The transaction associated with the downloadable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/transaction
func (d_ Download) Transaction() ISKPaymentTransaction {
	rv := objc.Send[PaymentTransaction](d_.ID, objc.Sel("transaction"))
	return rv
} /* debug [instance_properties/getter]: transaction */

// Indicates that the system cannot determine how much time is needed to finish downloading the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownloadtimeremainingunknown
func (d_ Download) SKDownloadTimeRemainingUnknown() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("SKDownloadTimeRemainingUnknown"))
	return rv
} /* debug [instance_properties/getter]: SKDownloadTimeRemainingUnknown */

// Indicates that the system cannot determine how much time is needed to finish downloading the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skdownloadtimeremainingunknown
func (d_ Download) SetSKDownloadTimeRemainingUnknown(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSKDownloadTimeRemainingUnknown:"), value)
} /* debug [instance_properties/setter]: SKDownloadTimeRemainingUnknown */

// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (d_ Download) Downloads() ISKDownload {
	rv := objc.Send[Download](d_.ID, objc.Sel("downloads"))
	return rv
} /* debug [instance_properties/getter]: downloads */

// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (d_ Download) SetDownloads(value ISKDownload) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDownloads:"), value)
} /* debug [instance_properties/setter]: downloads */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKDownload */
