// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BAURLDownload] class.
var (
	BAURLDownloadClass     _BAURLDownloadClass
	BAURLDownloadClassOnce sync.Once
)

func getBAURLDownloadClass() _BAURLDownloadClass {
	BAURLDownloadClassOnce.Do(func() {
		BAURLDownloadClass = _BAURLDownloadClass{objc.GetClass("BAURLDownload")}
	})
	return BAURLDownloadClass
}

type _BAURLDownloadClass struct {
	class objc.Class
}

// An interface definition for the [BAURLDownload] class.
type IBAURLDownload interface {
	IBADownload
	// properties:
	// methods:
}

// An object that represents a remote asset to download.


// An object that represents a remote asset to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload
type BAURLDownload struct {
	BADownload
}

// BAURLDownloadFrom constructs a [BAURLDownload] from an unsafe.Pointer.
//
// An object that represents a remote asset to download.
func BAURLDownloadFrom(ptr unsafe.Pointer) BAURLDownload {
	return BAURLDownload{
		BADownload: BADownloadFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BAURLDownloadClass) Alloc() BAURLDownload {
	rv := objc.Send[BAURLDownload](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BAURLDownloadClass) New() BAURLDownload {
	rv := objc.Send[BAURLDownload](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAURLDownload) Init() BAURLDownload {
	rv := objc.Send[BAURLDownload](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAURLDownload) Autorelease() BAURLDownload {
	rv := objc.Send[BAURLDownload](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAURLDownload creates a new BAURLDownload instance.
func NewBAURLDownload() BAURLDownload {
	return getBAURLDownloadClass().New()
}



// Creates a download that uses the specified identifier and App Group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:applicationGroupIdentifier:)
func NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifier(identifier objc.IObject /* cross-framework: NSString */, request objc.IObject /* cross-framework: URLRequest */, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:applicationGroupIdentifier:"), identifier, request, applicationGroupIdentifier)
	rv.Autorelease()
	return rv
}


// Creates a prioritized download that uses the specified identifier and App Group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:applicationGroupIdentifier:priority:)
func NewBAURLDownloadWithIdentifierRequestApplicationGroupIdentifierPriority(identifier objc.IObject /* cross-framework: NSString */, request objc.IObject /* cross-framework: URLRequest */, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */, priority BADownloaderPriority /* typedef */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:applicationGroupIdentifier:priority:"), identifier, request, applicationGroupIdentifier, priority)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:essential:fileSize:applicationGroupIdentifier:priority:)
func NewBAURLDownloadWithIdentifierRequestEssentialFileSizeApplicationGroupIdentifierPriority(identifier objc.IObject /* cross-framework: NSString */, request objc.IObject /* cross-framework: URLRequest */, essential bool, fileSize uint, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */, priority BADownloaderPriority /* typedef */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:essential:fileSize:applicationGroupIdentifier:priority:"), identifier, request, essential, fileSize, applicationGroupIdentifier, priority)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAURLDownload/init(identifier:request:fileSize:applicationGroupIdentifier:)
func NewBAURLDownloadWithIdentifierRequestFileSizeApplicationGroupIdentifier(identifier objc.IObject /* cross-framework: NSString */, request objc.IObject /* cross-framework: URLRequest */, fileSize uint, applicationGroupIdentifier objc.IObject /* cross-framework: NSString */) BAURLDownload {
	instance := getBAURLDownloadClass().Alloc()
	rv := objc.Send[BAURLDownload](instance.ID, objc.Sel("initWithIdentifier:request:fileSize:applicationGroupIdentifier:"), identifier, request, fileSize, applicationGroupIdentifier)
	rv.Autorelease()
	return rv
}



