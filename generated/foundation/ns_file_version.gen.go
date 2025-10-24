// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FileVersion] class.
var (
	FileVersionClass     _FileVersionClass
	FileVersionClassOnce sync.Once
)

func getFileVersionClass() _FileVersionClass {
	FileVersionClassOnce.Do(func() {
		FileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}
	})
	return FileVersionClass
}

type _FileVersionClass struct {
	class objc.Class
}





// An interface definition for the [FileVersion] class.
type IFileVersion interface {
	objectivec.IObject
	

	// properties:
	Discardable() bool
	SetDiscardable(value bool)
	LocalizedName() IString
	LocalizedNameOfSavingComputer() IString
	ModificationDate() IDate
	PersistentIdentifier() unsafe.Pointer
	URL() IURL
	HasLocalContents() bool
	SetHasLocalContents(value bool)
	HasThumbnail() bool
	SetHasThumbnail(value bool)
	IsConflict() bool
	SetIsConflict(value bool)
	IsDiscardable() bool
	SetIsDiscardable(value bool)
	IsResolved() bool
	SetIsResolved(value bool)
	OriginatorNameComponents() IPersonNameComponents
	SetOriginatorNameComponents(value IPersonNameComponents)


	

	// methods:
	RemoveAndReturnError(outError IError) bool
	ReplaceItemAtURLOptionsError(url IURL, options FileVersionReplacingOptions, error_ IError) IURL


}





// Alloc allocates a new instance without initialization.
func (fc _FileVersionClass) Alloc() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileVersionClass) New() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileVersion) Init() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileVersion) Autorelease() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileVersion creates a new FileVersion instance.
func NewFileVersion() FileVersion {
	return getFileVersionClass().New()
}





// A snapshot of a file at a specific point in time.
//
// Use the methods of this class to access, create, and manage file revisions in your app. Each file version instance contains metadata about a single revision, including the location of the associated file, the modification date of the revision, and whether the revision is discardable. In Mac apps, you can use file version objects to track changes to a local file over time and to prevent the loss of data during editing. When managing local versions, the document architecture creates versions at specific points in the lifetime of your application. Your application can also create versions explicitly at times that your application designates as appropriate. In addition to managing local files, the system also uses this class to manage cloud-based files. For files in the cloud, there is usually only one version of the file at any given time. However, additional file versions may be created in cases where two different computers attempt to save the file to the cloud at the same time. In that case, one file is chosen as the current version and any other versions are tagged as being in conflict with the original. Conflict versions are reported to the appropriate file presenter objects and should be resolved as soon as possible so that the corresponding files can be removed from the cloud.


// A snapshot of a file at a specific point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion
type FileVersion struct {
	objectivec.Object
}

// FileVersionFrom constructs a [FileVersion] from an unsafe.Pointer.
//
// A snapshot of a file at a specific point in time.
func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{objectivec.Object{objc.ID(ptr)}}
}










// Returns the most recent version object for the file at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/currentVersionOfItem(at:)
func (fc _FileVersionClass) CurrentVersionOfItemAtURL(url IURL) IFileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("currentVersionOfItemAtURL:"), url)
	return rv
}


// Returns all versions of the specified file except the current version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/otherVersionsOfItem(at:)
func (fc _FileVersionClass) OtherVersionsOfItemAtURL(url IURL) []FileVersion {
	rv := objc.Send[[]FileVersion](objc.ID(fc.class), objc.Sel("otherVersionsOfItemAtURL:"), url)
	return rv
}


// Removes all versions of a file, except the current one, from the version store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/removeOtherVersionsOfItem(at:)
func (fc _FileVersionClass) RemoveOtherVersionsOfItemAtURLError(url IURL, outError IError) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("removeOtherVersionsOfItemAtURL:error:"), url, outError)
	return rv
}


// Creates and returns a temporary directory to use for saving the contents of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/temporaryDirectoryURLForNewVersionOfItem(at:)
func (fc _FileVersionClass) TemporaryDirectoryURLForNewVersionOfItemAtURL(url IURL) IURL {
	rv := objc.Send[URL](objc.ID(fc.class), objc.Sel("temporaryDirectoryURLForNewVersionOfItemAtURL:"), url)
	return rv
}


// Returns the version of the file that has the specified persistent ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/version(itemAt:forPersistentIdentifier:)
func (fc _FileVersionClass) VersionOfItemAtURLForPersistentIdentifier(url IURL, persistentIdentifier objc.IObject) IFileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("versionOfItemAtURL:forPersistentIdentifier:"), url, persistentIdentifier)
	return rv
}












// Remove this version object and its associated file from the version store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/remove()
func (f_ FileVersion) RemoveAndReturnError(outError IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("removeAndReturnError:"), outError)
	return rv
}


// Replace the contents of the specified file with the contents of the current version’s file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/replaceItem(at:options:)
func (f_ FileVersion) ReplaceItemAtURLOptionsError(url IURL, options FileVersionReplacingOptions, error_ IError) IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("replaceItemAtURL:options:error:"), url, options, error_)
	return rv
}







// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isDiscardable
func (f_ FileVersion) Discardable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("discardable"))
	return rv
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isDiscardable
func (f_ FileVersion) SetDiscardable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDiscardable:"), value)
}


// The string containing the user-presentable name of the file version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedName
func (f_ FileVersion) LocalizedName() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("localizedName"))
	return rv
}


// The user-presentable name of the computer on which the revision was saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedNameOfSavingComputer
func (f_ FileVersion) LocalizedNameOfSavingComputer() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("localizedNameOfSavingComputer"))
	return rv
}


// The modification date of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/modificationDate
func (f_ FileVersion) ModificationDate() IDate {
	rv := objc.Send[Date](f_.ID, objc.Sel("modificationDate"))
	return rv
}


// The identifier for this version of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/persistentIdentifier
func (f_ FileVersion) PersistentIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("persistentIdentifier"))
	return rv
}


// The URL identifying the location of the file associated with the file version object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/url
func (f_ FileVersion) URL() IURL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/haslocalcontents
func (f_ FileVersion) HasLocalContents() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasLocalContents"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/haslocalcontents
func (f_ FileVersion) SetHasLocalContents(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHasLocalContents:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/hasthumbnail
func (f_ FileVersion) HasThumbnail() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasThumbnail"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/hasthumbnail
func (f_ FileVersion) SetHasThumbnail(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHasThumbnail:"), value)
}


// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) IsConflict() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isConflict"))
	return rv
}


// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) SetIsConflict(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsConflict:"), value)
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) IsDiscardable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDiscardable"))
	return rv
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) SetIsDiscardable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDiscardable:"), value)
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) IsResolved() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isResolved"))
	return rv
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) SetIsResolved(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsResolved:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/originatornamecomponents
func (f_ FileVersion) OriginatorNameComponents() IPersonNameComponents {
	rv := objc.Send[PersonNameComponents](f_.ID, objc.Sel("originatorNameComponents"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/originatornamecomponents
func (f_ FileVersion) SetOriginatorNameComponents(value IPersonNameComponents) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOriginatorNameComponents:"), value)
}








