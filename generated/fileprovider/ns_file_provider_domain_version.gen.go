// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderDomainVersion */


/* debug [class_header]: Header for NSFileProviderDomainVersion */
// The class instance for the [FileProviderDomainVersion] class.
var (
	FileProviderDomainVersionClass     _FileProviderDomainVersionClass
	FileProviderDomainVersionClassOnce sync.Once
)

func getFileProviderDomainVersionClass() _FileProviderDomainVersionClass {
	FileProviderDomainVersionClassOnce.Do(func() {
		FileProviderDomainVersionClass = _FileProviderDomainVersionClass{objc.GetClass("NSFileProviderDomainVersion")}
	})
	return FileProviderDomainVersionClass
}

type _FileProviderDomainVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderDomainVersion */
// An interface definition for the [FileProviderDomainVersion] class.
type IFileProviderDomainVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderDomainVersion */
	// properties:
	DomainVersion() IFileProviderDomainVersion
	SetDomainVersion(value IFileProviderDomainVersion)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderDomainVersion */
	// methods:
	Compare(otherVersion IFileProviderDomainVersion) ComparisonResult /* not a class type */
	Next() IFileProviderDomainVersion
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderDomainVersion */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderDomainVersionClass) Alloc() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderDomainVersionClass) New() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderDomainVersion) Init() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderDomainVersion) Autorelease() FileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderDomainVersion creates a new FileProviderDomainVersion instance.
func NewFileProviderDomainVersion() FileProviderDomainVersion {
	return getFileProviderDomainVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderDomainVersion */
// An opaque object that identifies a specific version of a domain.
//
// The file provider extension is responsible for assigning and updating the domain version. To specify the domain version, adopt the protocol. The system then calls your extension’s method to read the current version. The system reads the domain version after you call: The completion handler The completion handler The completion handler The completion handler The or method when enumerating the materialized set. The system always reads the domain version on the same dispatch queue as the completion handler. Your extension defines when the domain version changes. When you update the version, call the and passing the constant as the property. This notifies the system of the update. The system ignores any lower versions. When the system discovers a change on disk, it associates that change with the current domain version. It then includes the version in the object passed to the file provider extension. Only file provider extensions based on the use instances of this class. Each version object is immutable. You can use them as keys in a dictionary.


// An opaque object that identifies a specific version of a domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion
type FileProviderDomainVersion struct {
	objectivec.Object
}

// FileProviderDomainVersionFrom constructs a [FileProviderDomainVersion] from an unsafe.Pointer.
//
// An opaque object that identifies a specific version of a domain.
func FileProviderDomainVersionFrom(ptr unsafe.Pointer) FileProviderDomainVersion {
	return FileProviderDomainVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderDomainVersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderDomainVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderDomainVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderDomainVersion */

// Compares another domain version with this one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion/compare:
func (f_ FileProviderDomainVersion) Compare(otherVersion IFileProviderDomainVersion) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](f_.ID, objc.Sel("compare:"), otherVersion)
	return rv
}/* debug [instance_methods/method]: Compare */


// Creates a new version that supersedes the current version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion/next()
func (f_ FileProviderDomainVersion) Next() IFileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("next"))
	return rv
}/* debug [instance_methods/method]: Next */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderDomainVersion */

// An opaque object that uniquely identifies the domain’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/domainversion
func (f_ FileProviderDomainVersion) DomainVersion() IFileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("domainVersion"))
	return rv
}/* debug [instance_properties/getter]: domainVersion */


// An opaque object that uniquely identifies the domain’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/domainversion
func (f_ FileProviderDomainVersion) SetDomainVersion(value IFileProviderDomainVersion) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDomainVersion:"), value)
}/* debug [instance_properties/setter]: domainVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderDomainVersion */



