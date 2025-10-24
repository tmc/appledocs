// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUnarchiver */


/* debug [class_header]: Header for NSUnarchiver */
// The class instance for the [Unarchiver] class.
var (
	UnarchiverClass     _UnarchiverClass
	UnarchiverClassOnce sync.Once
)

func getUnarchiverClass() _UnarchiverClass {
	UnarchiverClassOnce.Do(func() {
		UnarchiverClass = _UnarchiverClass{objc.GetClass("NSUnarchiver")}
	})
	return UnarchiverClass
}

type _UnarchiverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Unarchiver */
// An interface definition for the [Unarchiver] class.
type IUnarchiver interface {
	ICoder
	
/* debug [class_interface_properties]: Properties for Unarchiver */
	// properties:
	AtEnd() bool
	SystemVersion() objectivec.IObject
	IsAtEnd() bool
	SetIsAtEnd(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Unarchiver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Unarchiver */
// Alloc allocates a new instance without initialization.
func (uc _UnarchiverClass) Alloc() Unarchiver {
	rv := objc.Send[Unarchiver](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnarchiverClass) New() Unarchiver {
	rv := objc.Send[Unarchiver](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ Unarchiver) Init() Unarchiver {
	rv := objc.Send[Unarchiver](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ Unarchiver) Autorelease() Unarchiver {
	rv := objc.Send[Unarchiver](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnarchiver creates a new Unarchiver instance.
func NewUnarchiver() Unarchiver {
	return getUnarchiverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Unarchiver */
// A decoder that restores data from an archive.
//
// , a concrete subclass of , defines methods for decoding a set of Objective-C objects from an archive. Such archives are produced by objects of the class. In macOS 10.2 and later, and have been replaced by and respectively—see .


// A decoder that restores data from an archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver
type Unarchiver struct {
	Coder
}

// UnarchiverFrom constructs a [Unarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive.
func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		Coder: CoderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Unarchiver */

// Returns an object initialized to read an archive from a given data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/init(forReadingWith:)
func NewUnarchiverForReadingWithData(data IData) Unarchiver {
	instance := getUnarchiverClass().Alloc()
	rv := objc.Send[Unarchiver](instance.ID, objc.Sel("initForReadingWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUnarchiverForReadingWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Unarchiver */

// Returns the name of the class used when instantiating objects whose ostensible class, according to the archived data, is a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/classNameDecoded(forArchiveClassName:)-swift.type.method
func (uc _UnarchiverClass) ClassNameDecodedForArchiveClassName(inArchiveName IString) IString {
	rv := objc.Send[String](objc.ID(uc.class), objc.Sel("classNameDecodedForArchiveClassName:"), inArchiveName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClassNameDecodedForArchiveClassName) */


// Instructs instances of to use the class with a given name when instantiating objects whose ostensible class, according to the archived data, is another given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/decodeClassName(_:asClassName:)-swift.type.method
func (uc _UnarchiverClass) DecodeClassNameAsClassName(inArchiveName IString, trueName IString) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("decodeClassName:asClassName:"), inArchiveName, trueName)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecodeClassNameAsClassName) */


// Decodes and returns the object archived in a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unarchiveObject(with:)
func (uc _UnarchiverClass) UnarchiveObjectWithData(data IData) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unarchiveObjectWithData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveObjectWithData) */


// Decodes and returns the object archived in the file .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unarchiveObject(withFile:)
func (uc _UnarchiverClass) UnarchiveObjectWithFile(path IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unarchiveObjectWithFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnarchiveObjectWithFile) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Unarchiver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Unarchiver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Unarchiver */

// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/isAtEnd
func (u_ Unarchiver) AtEnd() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("atEnd"))
	return rv
}/* debug [instance_properties/getter]: atEnd */


// The system version number in effect when the archive was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/systemVersion-swift.property
func (u_ Unarchiver) SystemVersion() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("systemVersion"))
	return rv
}/* debug [instance_properties/getter]: systemVersion */


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend
func (u_ Unarchiver) IsAtEnd() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isAtEnd"))
	return rv
}/* debug [instance_properties/getter]: isAtEnd */


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend
func (u_ Unarchiver) SetIsAtEnd(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsAtEnd:"), value)
}/* debug [instance_properties/setter]: isAtEnd */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnarchiver */


