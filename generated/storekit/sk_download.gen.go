// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (dc _DownloadClass) ContentURLForProductID(productID string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("contentURLForProductID:"), objc.String(productID))
	return rv
}

// Deletes the previously downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/deleteContent(forProductID:)
func (dc _DownloadClass) DeleteContentForProductID(productID string) {
	objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("deleteContentForProductID:"), objc.String(productID))
}

// The local location of the downloaded file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/contentURL
func (d_ Download) ContentURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contentURL"))
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
func (d_ Download) TimeRemaining() TimeInterval {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeRemaining"))
	return rv
}

// The transaction associated with the downloadable file.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownload/transaction
func (d_ Download) Transaction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("transaction"))
	return rv
}



