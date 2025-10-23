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
	HasLocalContents() bool /* primitive/slice/pointer */
	HasThumbnail() bool /* primitive/slice/pointer */
	Conflict() bool /* primitive/slice/pointer */
	Discardable() bool /* primitive/slice/pointer */
	SetDiscardable(value bool /* primitive/slice/pointer */)
	Resolved() bool /* primitive/slice/pointer */
	SetResolved(value bool /* primitive/slice/pointer */)
	LocalizedName() string /* primitive/slice/pointer */
	LocalizedNameOfSavingComputer() string /* primitive/slice/pointer */
	ModificationDate() IDate
	OriginatorNameComponents() IPersonNameComponents
	PersistentIdentifier() objc.ID
	URL() IURL
	IsConflict() bool /* primitive/slice/pointer */
	SetIsConflict(value bool /* primitive/slice/pointer */)
	IsDiscardable() bool /* primitive/slice/pointer */
	SetIsDiscardable(value bool /* primitive/slice/pointer */)
	IsResolved() bool /* primitive/slice/pointer */
	SetIsResolved(value bool /* primitive/slice/pointer */)
	// methods:
	RemoveAndReturnError(outError IError) bool /* primitive/slice/pointer */
	ReplaceItemAtURLOptionsError(url IURL, options FileVersionReplacingOptions, error_ IError) IURL
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

// Alloc allocates a new instance without initialization.
func (fc _FileVersionClass) Alloc() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a version of the file at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/addOfItem(at:withContentsOf:options:)
func (fc _FileVersionClass) AddVersionOfItemAtURLWithContentsOfURLOptionsError(url IURL, contentsURL IURL, options FileVersionAddingOptions, outError IError) IFileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("addVersionOfItemAtURL:withContentsOfURL:options:error:"), url, contentsURL, options, outError)
	return rv
}


// Returns the most recent version object for the file at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/currentVersionOfItem(at:)
func (fc _FileVersionClass) CurrentVersionOfItemAtURL(url IURL) IFileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("currentVersionOfItemAtURL:"), url)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/getNonlocalVersionsOfItem(at:completionHandler:)
func (fc _FileVersionClass) GetNonlocalVersionsOfItemAtURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getNonlocalVersionsOfItemAtURL:completionHandler:"), url, completionHandler)
}


// Returns all versions of the specified file except the current version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/otherVersionsOfItem(at:)
func (fc _FileVersionClass) OtherVersionsOfItemAtURL(url IURL) []FileVersion /* primitive/slice/pointer */ {
	rv := objc.Send[[]FileVersion](objc.ID(fc.class), objc.Sel("otherVersionsOfItemAtURL:"), url)
	return rv
}


// Removes all versions of a file, except the current one, from the version store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/removeOtherVersionsOfItem(at:)
func (fc _FileVersionClass) RemoveOtherVersionsOfItemAtURLError(url IURL, outError IError) bool /* primitive/slice/pointer */ {
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


// Returns an array of version objects that are currently in conflict for the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/unresolvedConflictVersionsOfItem(at:)
func (fc _FileVersionClass) UnresolvedConflictVersionsOfItemAtURL(url IURL) []FileVersion /* primitive/slice/pointer */ {
	rv := objc.Send[[]FileVersion](objc.ID(fc.class), objc.Sel("unresolvedConflictVersionsOfItemAtURL:"), url)
	return rv
}


// Returns the version of the file that has the specified persistent ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/version(itemAt:forPersistentIdentifier:)
func (fc _FileVersionClass) VersionOfItemAtURLForPersistentIdentifier(url IURL, persistentIdentifier objectivec.IObject) IFileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("versionOfItemAtURL:forPersistentIdentifier:"), url, persistentIdentifier)
	return rv
}


// Remove this version object and its associated file from the version store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/remove()
func (f_ FileVersion) RemoveAndReturnError(outError IError) bool /* primitive/slice/pointer */ {
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasLocalContents
func (f_ FileVersion) HasLocalContents() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasLocalContents"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasThumbnail
func (f_ FileVersion) HasThumbnail() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasThumbnail"))
	return rv
}


// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isConflict
func (f_ FileVersion) Conflict() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("conflict"))
	return rv
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isDiscardable
func (f_ FileVersion) Discardable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("discardable"))
	return rv
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isDiscardable
func (f_ FileVersion) SetDiscardable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDiscardable:"), value)
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isResolved
func (f_ FileVersion) Resolved() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("resolved"))
	return rv
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/isResolved
func (f_ FileVersion) SetResolved(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setResolved:"), value)
}


// The string containing the user-presentable name of the file version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedName
func (f_ FileVersion) LocalizedName() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](f_.ID, objc.Sel("localizedName"))
	return rv
}


// The user-presentable name of the computer on which the revision was saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/localizedNameOfSavingComputer
func (f_ FileVersion) LocalizedNameOfSavingComputer() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](f_.ID, objc.Sel("localizedNameOfSavingComputer"))
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/originatorNameComponents
func (f_ FileVersion) OriginatorNameComponents() IPersonNameComponents {
	rv := objc.Send[PersonNameComponents](f_.ID, objc.Sel("originatorNameComponents"))
	return rv
}


// The identifier for this version of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/persistentIdentifier
func (f_ FileVersion) PersistentIdentifier() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("persistentIdentifier"))
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


// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) IsConflict() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isConflict"))
	return rv
}


// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) SetIsConflict(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsConflict:"), value)
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) IsDiscardable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDiscardable"))
	return rv
}


// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) SetIsDiscardable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDiscardable:"), value)
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) IsResolved() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("isResolved"))
	return rv
}


// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) SetIsResolved(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsResolved:"), value)
}



