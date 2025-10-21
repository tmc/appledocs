// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Archiver] class.
var (
	ArchiverClass     _ArchiverClass
	ArchiverClassOnce sync.Once
)

func getArchiverClass() _ArchiverClass {
	ArchiverClassOnce.Do(func() {
		ArchiverClass = _ArchiverClass{objc.GetClass("NSArchiver")}
	})
	return ArchiverClass
}

type _ArchiverClass struct {
	class objc.Class
}

// An interface definition for the [Archiver] class.
type IArchiver interface {
	ICoder
	EncodeRootObject(rootObject objc.ID)
}

// A coder that stores an object’s data to an archive.
//
// , a concrete subclass of , provides a way to encode objects into an architecture-independent format that can be stored in a file. When you archive a graph of objects, the class information and instance variables for each object are written to the archive. The companion class decodes the data in an archive and creates a graph of objects equivalent to the original set. stores the archive data in a mutable data object ( ). After encoding the objects, you can have the object write this mutable data object immediately to a file, or you can retrieve the mutable data object for some other use. In macOS 10.2 and later, and have been replaced by and respectively—see .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver
type Archiver struct {
	Coder
}

// ArchiverFrom constructs a [Archiver] from an unsafe.Pointer.
//
// A coder that stores an object’s data to an archive.
func ArchiverFrom(ptr unsafe.Pointer) Archiver {
	return Archiver{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ArchiverClass) Alloc() Archiver {
	rv := objc.Send[Archiver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArchiverClass) New() Archiver {
	rv := objc.Send[Archiver](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Archiver) Init() Archiver {
	rv := objc.Send[Archiver](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Archiver) Autorelease() Archiver {
	rv := objc.Send[Archiver](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArchiver creates a new Archiver instance.
func NewArchiver() Archiver {
	return getArchiverClass().New()
}


// Creates a temporary instance of and archives an object graph by encoding it into a data object and writing the resulting data object to a specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/archiveRootObject(_:toFile:)
func (ac _ArchiverClass) ArchiveRootObjectToFile(rootObject objc.ID, path string) bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("archiveRootObject:toFile:"), rootObject, objc.String(path))
	return rv
}

// Archives a given object along with all the objects to which it is connected.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/encodeRootObject(_:)
func (a_ Archiver) EncodeRootObject(rootObject objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("encodeRootObject:"), rootObject)
}

// The receiver’s archive data.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarchiver/archiverdata
func (a_ Archiver) ArchiverData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("archiverData"))
	return rv
}


// SetArchiverData sets the value of the archiverData property.
// The receiver’s archive data.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarchiver/archiverdata
func (a_ Archiver) SetArchiverData(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArchiverData:"), value)
}



