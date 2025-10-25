// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileSecurity */


/* debug [class_header]: Header for NSFileSecurity */
// The class instance for the [FileSecurity] class.
var (
	FileSecurityClass     _FileSecurityClass
	FileSecurityClassOnce sync.Once
)

func getFileSecurityClass() _FileSecurityClass {
	FileSecurityClassOnce.Do(func() {
		FileSecurityClass = _FileSecurityClass{objc.GetClass("NSFileSecurity")}
	})
	return FileSecurityClass
}

type _FileSecurityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileSecurity */
// An interface definition for the [FileSecurity] class.
type IFileSecurity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileSecurity */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileSecurity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileSecurity */
// Alloc allocates a new instance without initialization.
func (fc _FileSecurityClass) Alloc() FileSecurity {
	rv := objc.Send[FileSecurity](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileSecurityClass) New() FileSecurity {
	rv := objc.Send[FileSecurity](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileSecurity) Init() FileSecurity {
	rv := objc.Send[FileSecurity](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileSecurity) Autorelease() FileSecurity {
	rv := objc.Send[FileSecurity](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileSecurity creates a new FileSecurity instance.
func NewFileSecurity() FileSecurity {
	return getFileSecurityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileSecurity */
// A stub class that encapsulates security information about a file.
//
// contains no methods of its own. Instead, it is transparently bridged to .


// A stub class that encapsulates security information about a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileSecurity
type FileSecurity struct {
	objectivec.Object
}

// FileSecurityFrom constructs a [FileSecurity] from an unsafe.Pointer.
//
// A stub class that encapsulates security information about a file.
func FileSecurityFrom(ptr unsafe.Pointer) FileSecurity {
	return FileSecurity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileSecurity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileSecurity/init(coder:)
func NewFileSecurityWithCoder(coder ICoder) FileSecurity {
	instance := getFileSecurityClass().Alloc()
	rv := objc.Send[FileSecurity](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileSecurityWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileSecurity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileSecurity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileSecurity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileSecurity */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileSecurity */


