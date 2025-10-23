// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Unarchiver] class.
type IUnarchiver interface {
	ICoder
	// properties:
	AtEnd() bool /* primitive/slice/pointer. */
	SystemVersion() unsafe.Pointer
	IsAtEnd() bool /* primitive/slice/pointer. */
	SetIsAtEnd(value bool /* primitive/slice/pointer. */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (uc _UnarchiverClass) Alloc() Unarchiver {
	rv := objc.Send[Unarchiver](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns an object initialized to read an archive from a given data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/init(forReadingWith:)
func NewUnarchiverForReadingWithData(data IData) Unarchiver {
	instance := getUnarchiverClass().Alloc()
	rv := objc.Send[Unarchiver](instance.ID, objc.Sel("initForReadingWithData:"), data)
	rv.Autorelease()
	return rv
}



// Returns the name of the class used when instantiating objects whose ostensible class, according to the archived data, is a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/classNameDecoded(forArchiveClassName:)-swift.type.method
func (uc _UnarchiverClass) ClassNameDecodedForArchiveClassName(inArchiveName IString) IString {
	rv := objc.Send[String](objc.ID(uc.class), objc.Sel("classNameDecodedForArchiveClassName:"), inArchiveName)
	return rv
}


// Instructs instances of to use the class with a given name when instantiating objects whose ostensible class, according to the archived data, is another given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/decodeClassName(_:asClassName:)-swift.type.method
func (uc _UnarchiverClass) DecodeClassNameAsClassName(inArchiveName IString, trueName IString) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("decodeClassName:asClassName:"), inArchiveName, trueName)
}


// Decodes and returns the object archived in a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unarchiveObject(with:)
func (uc _UnarchiverClass) UnarchiveObjectWithData(data IData) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unarchiveObjectWithData:"), data)
	return rv
}


// Decodes and returns the object archived in the file .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unarchiveObject(withFile:)
func (uc _UnarchiverClass) UnarchiveObjectWithFile(path IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unarchiveObjectWithFile:"), path)
	return rv
}


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/isAtEnd
func (u_ Unarchiver) AtEnd() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("atEnd"))
	return rv
}


// The system version number in effect when the archive was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/systemVersion-swift.property
func (u_ Unarchiver) SystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("systemVersion"))
	return rv
}


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend
func (u_ Unarchiver) IsAtEnd() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isAtEnd"))
	return rv
}


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend
func (u_ Unarchiver) SetIsAtEnd(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsAtEnd:"), value)
}


