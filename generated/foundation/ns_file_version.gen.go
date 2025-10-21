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
}

// A snapshot of a file at a specific point in time.
//
// Use the methods of this class to access, create, and manage file revisions in your app. Each file version instance contains metadata about a single revision, including the location of the associated file, the modification date of the revision, and whether the revision is discardable. In Mac apps, you can use file version objects to track changes to a local file over time and to prevent the loss of data during editing. When managing local versions, the document architecture creates versions at specific points in the lifetime of your application. Your application can also create versions explicitly at times that your application designates as appropriate. In addition to managing local files, the system also uses this class to manage cloud-based files. For files in the cloud, there is usually only one version of the file at any given time. However, additional file versions may be created in cases where two different computers attempt to save the file to the cloud at the same time. In that case, one file is chosen as the current version and any other versions are tagged as being in conflict with the original. Conflict versions are reported to the appropriate file presenter objects and should be resolved as soon as possible so that the corresponding files can be removed from the cloud.
//
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


// The string containing the user-presentable name of the file version.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/localizedname
func (f_ FileVersion) LocalizedName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("localizedName"))
	return rv
}


// SetLocalizedName sets the value of the localizedName property.
// The string containing the user-presentable name of the file version.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/localizedname
func (f_ FileVersion) SetLocalizedName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/hasthumbnail
func (f_ FileVersion) HasThumbnail() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasThumbnail"))
	return rv
}


// SetHasThumbnail sets the value of the hasThumbnail property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/hasthumbnail
func (f_ FileVersion) SetHasThumbnail(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setHasThumbnail:"), value)
}

// A Boolean value that specifies whether the system can delete the associated file at some future time.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) IsDiscardable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDiscardable"))
	return rv
}


// SetIsDiscardable sets the value of the isDiscardable property.
// A Boolean value that specifies whether the system can delete the associated file at some future time.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isdiscardable
func (f_ FileVersion) SetIsDiscardable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDiscardable:"), value)
}

// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) IsConflict() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isConflict"))
	return rv
}


// SetIsConflict sets the value of the isConflict property.
// A Boolean value indicating whether the contents of the version are in conflict with the contents of another version.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isconflict
func (f_ FileVersion) SetIsConflict(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsConflict:"), value)
}

// The user-presentable name of the computer on which the revision was saved.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/localizednameofsavingcomputer
func (f_ FileVersion) LocalizedNameOfSavingComputer() string {
	rv := objc.Send[string](f_.ID, objc.Sel("localizedNameOfSavingComputer"))
	return rv
}


// SetLocalizedNameOfSavingComputer sets the value of the localizedNameOfSavingComputer property.
// The user-presentable name of the computer on which the revision was saved.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/localizednameofsavingcomputer
func (f_ FileVersion) SetLocalizedNameOfSavingComputer(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLocalizedNameOfSavingComputer:"), objc.String(value))
}

// A Boolean value that indicates if the version object is in conflict or not.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) IsResolved() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isResolved"))
	return rv
}


// SetIsResolved sets the value of the isResolved property.
// A Boolean value that indicates if the version object is in conflict or not.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/isresolved
func (f_ FileVersion) SetIsResolved(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsResolved:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/originatornamecomponents
func (f_ FileVersion) OriginatorNameComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("originatorNameComponents"))
	return rv
}


// SetOriginatorNameComponents sets the value of the originatorNameComponents property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/originatornamecomponents
func (f_ FileVersion) SetOriginatorNameComponents(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOriginatorNameComponents:"), value)
}

// The modification date of the version.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/modificationdate
func (f_ FileVersion) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("modificationDate"))
	return rv
}


// SetModificationDate sets the value of the modificationDate property.
// The modification date of the version.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfileversion/modificationdate
func (f_ FileVersion) SetModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setModificationDate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/hasLocalContents
func (f_ FileVersion) HasLocalContents() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasLocalContents"))
	return rv
}

// The identifier for this version of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/persistentIdentifier
func (f_ FileVersion) PersistentIdentifier() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("persistentIdentifier"))
	return rv
}

// The URL identifying the location of the file associated with the file version object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion/url
func (f_ FileVersion) URL() URL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URL"))
	return rv
}



