// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	IsAtEnd() bool
	SetIsAtEnd(value bool)
	SystemVersion() unsafe.Pointer
	SetSystemVersion(value unsafe.Pointer)
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



// Instructs instances of to use the class with a given name when instantiating objects whose ostensible class, according to the archived data, is another given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/decodeClassName(_:asClassName:)-swift.type.method

func (uc _UnarchiverClass) DecodeClassNameAsClassName(inArchiveName string, trueName string) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("decodeClassName:asClassName:"), objc.String(inArchiveName), objc.String(trueName))
}


// Decodes and returns the object archived in a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unarchiveObject(with:)

func (uc _UnarchiverClass) UnarchiveObjectWithData(data IData) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unarchiveObjectWithData:"), data)
	return rv
}


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend

func (u_ Unarchiver) IsAtEnd() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isAtEnd"))
	return rv
}


// A Boolean value that indicates whether the receiver has reached the end of the encoded data while decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/isatend

func (u_ Unarchiver) SetIsAtEnd(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsAtEnd:"), value)
}


// The system version number in effect when the archive was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/systemversion-swift.property

func (u_ Unarchiver) SystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("systemVersion"))
	return rv
}


// The system version number in effect when the archive was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunarchiver/systemversion-swift.property

func (u_ Unarchiver) SetSystemVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSystemVersion:"), value)
}


