// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSDirectoryEnumerator */


/* debug [class_header]: Header for NSDirectoryEnumerator */
// The class instance for the [DirectoryEnumerator] class.
var (
	DirectoryEnumeratorClass     _DirectoryEnumeratorClass
	DirectoryEnumeratorClassOnce sync.Once
)

func getDirectoryEnumeratorClass() _DirectoryEnumeratorClass {
	DirectoryEnumeratorClassOnce.Do(func() {
		DirectoryEnumeratorClass = _DirectoryEnumeratorClass{objc.GetClass("NSDirectoryEnumerator")}
	})
	return DirectoryEnumeratorClass
}

type _DirectoryEnumeratorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DirectoryEnumerator */
// An interface definition for the [DirectoryEnumerator] class.
type IDirectoryEnumerator interface {
	IEnumerator
	
/* debug [class_interface_properties]: Properties for DirectoryEnumerator */
	// properties:
	DirectoryAttributes() IDictionary
	FileAttributes() IDictionary
	IsEnumeratingDirectoryPostOrder() bool
	Level() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DirectoryEnumerator */
	// methods:
	SkipDescendants()
	SkipDescendents()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DirectoryEnumerator */
// Alloc allocates a new instance without initialization.
func (dc _DirectoryEnumeratorClass) Alloc() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DirectoryEnumeratorClass) New() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DirectoryEnumerator) Init() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DirectoryEnumerator) Autorelease() DirectoryEnumerator {
	rv := objc.Send[DirectoryEnumerator](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDirectoryEnumerator creates a new DirectoryEnumerator instance.
func NewDirectoryEnumerator() DirectoryEnumerator {
	return getDirectoryEnumeratorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DirectoryEnumerator */
// An object that enumerates the contents of a directory.
//
// You obtain a directory enumerator using ’s method. The enumeration provides the pathnames of all files and directories contained within that directory. These pathnames are relative to the directory. An enumeration is recursive, including the files of all subdirectories, and crosses device boundaries. An enumeration does not resolve symbolic links, or attempt to traverse symbolic links that point to directories.


// An object that enumerates the contents of a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator
type DirectoryEnumerator struct {
	Enumerator
}

// DirectoryEnumeratorFrom constructs a [DirectoryEnumerator] from an unsafe.Pointer.
//
// An object that enumerates the contents of a directory.
func DirectoryEnumeratorFrom(ptr unsafe.Pointer) DirectoryEnumerator {
	return DirectoryEnumerator{
		Enumerator: EnumeratorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DirectoryEnumerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DirectoryEnumerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DirectoryEnumerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DirectoryEnumerator */

// Causes the receiver to skip recursion into the most recently obtained subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/skipDescendants()
func (d_ DirectoryEnumerator) SkipDescendants() {
	objc.Send[objc.ID](d_.ID, objc.Sel("skipDescendants"))
}/* debug [instance_methods/method]: SkipDescendants */


// Causes the receiver to skip recursion into the most recently obtained subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/skipDescendents()
func (d_ DirectoryEnumerator) SkipDescendents() {
	objc.Send[objc.ID](d_.ID, objc.Sel("skipDescendents"))
}/* debug [instance_methods/method]: SkipDescendents */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DirectoryEnumerator */

// A dictionary with the attributes of the directory at which enumeration started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/directoryAttributes
func (d_ DirectoryEnumerator) DirectoryAttributes() IDictionary {
	rv := objc.Send[Dictionary](d_.ID, objc.Sel("directoryAttributes"))
	return rv
}/* debug [instance_properties/getter]: directoryAttributes */


// A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/fileAttributes
func (d_ DirectoryEnumerator) FileAttributes() IDictionary {
	rv := objc.Send[Dictionary](d_.ID, objc.Sel("fileAttributes"))
	return rv
}/* debug [instance_properties/getter]: fileAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/isEnumeratingDirectoryPostOrder
func (d_ DirectoryEnumerator) IsEnumeratingDirectoryPostOrder() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEnumeratingDirectoryPostOrder"))
	return rv
}/* debug [instance_properties/getter]: isEnumeratingDirectoryPostOrder */


// The number of levels deep the current object is in the directory hierarchy being enumerated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/level
func (d_ DirectoryEnumerator) Level() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDirectoryEnumerator */



