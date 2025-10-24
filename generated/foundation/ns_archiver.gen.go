// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSArchiver */


/* debug [class_header]: Header for NSArchiver */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Archiver */
// An interface definition for the [Archiver] class.
type IArchiver interface {
	ICoder
	
/* debug [class_interface_properties]: Properties for Archiver */
	// properties:
	ArchiverData() IMutableData
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Archiver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Archiver */
// Alloc allocates a new instance without initialization.
func (ac _ArchiverClass) Alloc() Archiver {
	rv := objc.Send[Archiver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Archiver */
// A coder that stores an object’s data to an archive.
//
// , a concrete subclass of , provides a way to encode objects into an architecture-independent format that can be stored in a file. When you archive a graph of objects, the class information and instance variables for each object are written to the archive. The companion class decodes the data in an archive and creates a graph of objects equivalent to the original set. stores the archive data in a mutable data object ( ). After encoding the objects, you can have the object write this mutable data object immediately to a file, or you can retrieve the mutable data object for some other use. In macOS 10.2 and later, and have been replaced by and respectively—see .


// A coder that stores an object’s data to an archive.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Archiver */

// Returns an archiver, initialized to encode stream and version information into a given mutable data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/init(forWritingWith:)
func NewArchiverForWritingWithMutableData(mdata IMutableData) Archiver {
	instance := getArchiverClass().Alloc()
	rv := objc.Send[Archiver](instance.ID, objc.Sel("initForWritingWithMutableData:"), mdata)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewArchiverForWritingWithMutableData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Archiver */

// Creates a temporary instance of and archives an object graph by encoding it into a data object and writing the resulting data object to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/archiveRootObject(_:toFile:)
func (ac _ArchiverClass) ArchiveRootObjectToFile(rootObject objc.IObject, path IString) bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("archiveRootObject:toFile:"), rootObject, path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArchiveRootObjectToFile) */


// Returns a data object containing the encoded form of the object graph whose root object is given.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/archivedData(withRootObject:)
func (ac _ArchiverClass) ArchivedDataWithRootObject(rootObject objc.IObject) IData {
	rv := objc.Send[Data](objc.ID(ac.class), objc.Sel("archivedDataWithRootObject:"), rootObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArchivedDataWithRootObject) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Archiver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Archiver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Archiver */

// The receiver’s archive data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/archiverData
func (a_ Archiver) ArchiverData() IMutableData {
	rv := objc.Send[MutableData](a_.ID, objc.Sel("archiverData"))
	return rv
}/* debug [instance_properties/getter]: archiverData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSArchiver */


