// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BAAssetPackManifest] class.
var (
	BAAssetPackManifestClass     _BAAssetPackManifestClass
	BAAssetPackManifestClassOnce sync.Once
)

func getBAAssetPackManifestClass() _BAAssetPackManifestClass {
	BAAssetPackManifestClassOnce.Do(func() {
		BAAssetPackManifestClass = _BAAssetPackManifestClass{objc.GetClass("BAAssetPackManifest")}
	})
	return BAAssetPackManifestClass
}

type _BAAssetPackManifestClass struct {
	class objc.Class
}

// An interface definition for the [BAAssetPackManifest] class.
type IBAAssetPackManifest interface {
	objectivec.IObject
	AllDownloads() unsafe.Pointer
	AllDownloadsForContentRequest(contentRequest unsafe.Pointer) unsafe.Pointer
}

// A representation of a manifest that lists asset packs that are available to download.
//
// This class applies only when you want to manage your asset packs manually. Don’t use this class if you want to opt in to automatic management of asset packs.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest
type BAAssetPackManifest struct {
	objectivec.Object
}

// BAAssetPackManifestFrom constructs a [BAAssetPackManifest] from an unsafe.Pointer.
//
// A representation of a manifest that lists asset packs that are available to download.
func BAAssetPackManifestFrom(ptr unsafe.Pointer) BAAssetPackManifest {
	return BAAssetPackManifest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackManifestClass) Alloc() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BAAssetPackManifestClass) New() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPackManifest) Init() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPackManifest) Autorelease() BAAssetPackManifest {
	rv := objc.Send[BAAssetPackManifest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPackManifest creates a new BAAssetPackManifest instance.
func NewBAAssetPackManifest() BAAssetPackManifest {
	return getBAAssetPackManifestClass().New()
}


// Initializes a representation of a manifest in memory given a URL to the manifest’s representation as a JSON file on disk.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest/initWithContentsOfURL:applicationGroupIdentifier:error:
func NewBAAssetPackManifestWithContentsOfURLApplicationGroupIdentifierError(URL unsafe.Pointer, applicationGroupIdentifier string, error_ unsafe.Pointer) BAAssetPackManifest {
	instance := getBAAssetPackManifestClass().Alloc()
	rv := objc.Send[BAAssetPackManifest](instance.ID, objc.Sel("initWithContentsOfURL:applicationGroupIdentifier:error:"), URL, objc.String(applicationGroupIdentifier), error_)
	rv.Autorelease()
	return rv
}

// Initializes a representation of a manifest in memory from JSON-encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest/initFromData:applicationGroupIdentifier:error:
func NewBAAssetPackManifestFromDataApplicationGroupIdentifierError(data unsafe.Pointer, applicationGroupIdentifier string, error_ unsafe.Pointer) BAAssetPackManifest {
	instance := getBAAssetPackManifestClass().Alloc()
	rv := objc.Send[BAAssetPackManifest](instance.ID, objc.Sel("initFromData:applicationGroupIdentifier:error:"), data, objc.String(applicationGroupIdentifier), error_)
	rv.Autorelease()
	return rv
}


// Creates download objects for every asset pack in this manifest.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest/allDownloads
func (b_ BAAssetPackManifest) AllDownloads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("allDownloads"))
	return rv
}

// Creates download objects for every asset pack in this manifest.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest/allDownloadsForContentRequest:
func (b_ BAAssetPackManifest) AllDownloadsForContentRequest(contentRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("allDownloadsForContentRequest:"), contentRequest)
	return rv
}

// The asset packs that are available to download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManifest/assetPacks
func (b_ BAAssetPackManifest) AssetPacks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("assetPacks"))
	return rv
}


