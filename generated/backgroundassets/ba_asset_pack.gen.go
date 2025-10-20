// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BAAssetPack] class.
var (
	BAAssetPackClass     _BAAssetPackClass
	BAAssetPackClassOnce sync.Once
)

func getBAAssetPackClass() _BAAssetPackClass {
	BAAssetPackClassOnce.Do(func() {
		BAAssetPackClass = _BAAssetPackClass{objc.GetClass("BAAssetPack")}
	})
	return BAAssetPackClass
}

type _BAAssetPackClass struct {
	class objc.Class
}

// An interface definition for the [BAAssetPack] class.
type IBAAssetPack interface {
	objectivec.IObject
	Download() unsafe.Pointer
	DownloadForContentRequest(contentRequest unsafe.Pointer) unsafe.Pointer
}

// An archive of assets that the system downloads together.
//
// An instance of this class can be invalidated when the asset pack that it represents is updated on the server.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack
type BAAssetPack struct {
	objectivec.Object
}

// BAAssetPackFrom constructs a [BAAssetPack] from an unsafe.Pointer.
//
// An archive of assets that the system downloads together.
func BAAssetPackFrom(ptr unsafe.Pointer) BAAssetPack {
	return BAAssetPack{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackClass) Alloc() BAAssetPack {
	rv := objc.Send[BAAssetPack](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BAAssetPackClass) New() BAAssetPack {
	rv := objc.Send[BAAssetPack](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPack) Init() BAAssetPack {
	rv := objc.Send[BAAssetPack](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPack) Autorelease() BAAssetPack {
	rv := objc.Send[BAAssetPack](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPack creates a new BAAssetPack instance.
func NewBAAssetPack() BAAssetPack {
	return getBAAssetPackClass().New()
}


// Creates a download object for the asset pack that you schedule using a download manager.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/download
func (b_ BAAssetPack) Download() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("download"))
	return rv
}

// Creates a download object for the asset pack that you schedule using a download manager.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/downloadForContentRequest:
func (b_ BAAssetPack) DownloadForContentRequest(contentRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("downloadForContentRequest:"), contentRequest)
	return rv
}

// The size of the download file containing the asset pack in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/downloadSize
func (b_ BAAssetPack) DownloadSize() int {
	rv := objc.Send[int](b_.ID, objc.Sel("downloadSize"))
	return rv
}

// A unique identifier for the asset pack.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/identifier
func (b_ BAAssetPack) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("identifier"))
	return rv
}

// JSON-encoded custom information that’s associated with the asset pack.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/userInfo
func (b_ BAAssetPack) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("userInfo"))
	return rv
}

// The asset pack’s version number
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPack/version
func (b_ BAAssetPack) Version() int {
	rv := objc.Send[int](b_.ID, objc.Sel("version"))
	return rv
}



