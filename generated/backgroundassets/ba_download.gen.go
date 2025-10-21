// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BADownload] class.
var (
	BADownloadClass     _BADownloadClass
	BADownloadClassOnce sync.Once
)

func getBADownloadClass() _BADownloadClass {
	BADownloadClassOnce.Do(func() {
		BADownloadClass = _BADownloadClass{objc.GetClass("BADownload")}
	})
	return BADownloadClass
}

type _BADownloadClass struct {
	class objc.Class
}

// An interface definition for the [BADownload] class.
type IBADownload interface {
	objectivec.IObject
	CopyAsNonEssential() unsafe.Pointer
}

// An object that represents an in-progress or concluded asset download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload
type BADownload struct {
	objectivec.Object
}

// BADownloadFrom constructs a [BADownload] from an unsafe.Pointer.
//
// An object that represents an in-progress or concluded asset download.
func BADownloadFrom(ptr unsafe.Pointer) BADownload {
	return BADownload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BADownloadClass) Alloc() BADownload {
	rv := objc.Send[BADownload](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BADownloadClass) New() BADownload {
	rv := objc.Send[BADownload](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BADownload) Init() BADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BADownload) Autorelease() BADownload {
	rv := objc.Send[BADownload](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBADownload creates a new BADownload instance.
func NewBADownload() BADownload {
	return getBADownloadClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/removingEssential()
func (b_ BADownload) CopyAsNonEssential() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("copyAsNonEssential"))
	return rv
}

// The app-specific string that uniquely identifies the downloadable asset.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/identifier
func (b_ BADownload) Identifier() appkit.string {
	rv := objc.Send[appkit.string](b_.ID, objc.Sel("identifier"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/isEssential
func (b_ BADownload) IsEssential() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEssential"))
	return rv
}

// The download’s execution priority.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/priority-swift.property
func (b_ BADownload) Priority() BADownloaderPriority {
	rv := objc.Send[BADownloaderPriority](b_.ID, objc.Sel("priority"))
	return rv
}

// The current state of the download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/state-swift.property
func (b_ BADownload) State() BADownloadState {
	rv := objc.Send[BADownloadState](b_.ID, objc.Sel("state"))
	return rv
}

// The system-provided string that uniquely identifies the download object.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/uniqueIdentifier
func (b_ BADownload) UniqueIdentifier() appkit.string {
	rv := objc.Send[appkit.string](b_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}



