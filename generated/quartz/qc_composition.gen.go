// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCComposition */


/* debug [class_header]: Header for QCComposition */
// The class instance for the [QCComposition] class.
var (
	QCCompositionClass     _QCCompositionClass
	QCCompositionClassOnce sync.Once
)

func getQCCompositionClass() _QCCompositionClass {
	QCCompositionClassOnce.Do(func() {
		QCCompositionClass = _QCCompositionClass{objc.GetClass("QCComposition")}
	})
	return QCCompositionClass
}

type _QCCompositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCComposition */
// An interface definition for the [QCComposition] class.
type IQCComposition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QCComposition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCComposition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCComposition */
// Alloc allocates a new instance without initialization.
func (qc _QCCompositionClass) Alloc() QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCCompositionClass) New() QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCComposition) Init() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCComposition) Autorelease() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCComposition creates a new QCComposition instance.
func NewQCComposition() QCComposition {
	return getQCCompositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCComposition */
// The class represents a Quartz Composer composition that either:
//
// comes from the system-wide composition repository ( and ) where it can be accessed by any application through the methods of the class is created from an arbitrary source (typically a file on disk) using one of its methods This class cannot be subclassed. A object has the following information associated with it and that you can obtain by using the appropriate method of the class: Attributes include the name and description of the composition, copyright information, and whether or not its provided by macOS (built-in). The protocols that the composition conforms to. A defines a set of required and optional input parameters and output results. Many methods of the , , and classes take a object as a parameter.


// The class represents a Quartz Composer composition that either:
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition
type QCComposition struct {
	objectivec.Object
}

// QCCompositionFrom constructs a [QCComposition] from an unsafe.Pointer.
//
// The class represents a Quartz Composer composition that either:
func QCCompositionFrom(ptr unsafe.Pointer) QCComposition {
	return QCComposition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCComposition */

// Returns a composition object initialized with the contents of a Quartz Composer composition file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition/init(data:)
func NewQCCompositionWithData(data objc.IObject /* cross-framework: NSData */) QCComposition {
	rv := objc.Send[QCComposition](objc.ID(getQCCompositionClass().class), objc.Sel("compositionWithData:"), data)
	return rv
}/* debug [class_init_methods/constructor]: NewQCCompositionWithData */


// Returns a composition object initialized with a Quartz Composer composition file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition/init(file:)
func NewQCCompositionWithFile(path objc.IObject /* cross-framework: NSString */) QCComposition {
	rv := objc.Send[QCComposition](objc.ID(getQCCompositionClass().class), objc.Sel("compositionWithFile:"), path)
	return rv
}/* debug [class_init_methods/constructor]: NewQCCompositionWithFile */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCComposition */

// Returns a composition object initialized with the contents of a Quartz Composer composition file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition/init(data:)
func (qc _QCCompositionClass) CompositionWithData(data objc.IObject /* cross-framework: NSData */) QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("compositionWithData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionWithData) */


// Returns a composition object initialized with a Quartz Composer composition file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCComposition/init(file:)
func (qc _QCCompositionClass) CompositionWithFile(path objc.IObject /* cross-framework: NSString */) QCComposition {
	rv := objc.Send[QCComposition](objc.ID(qc.class), objc.Sel("compositionWithFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionWithFile) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCComposition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCComposition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCComposition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCComposition */


