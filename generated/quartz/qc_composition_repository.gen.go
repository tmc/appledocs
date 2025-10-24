// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCCompositionRepository */


/* debug [class_header]: Header for QCCompositionRepository */
// The class instance for the [QCCompositionRepository] class.
var (
	QCCompositionRepositoryClass     _QCCompositionRepositoryClass
	QCCompositionRepositoryClassOnce sync.Once
)

func getQCCompositionRepositoryClass() _QCCompositionRepositoryClass {
	QCCompositionRepositoryClassOnce.Do(func() {
		QCCompositionRepositoryClass = _QCCompositionRepositoryClass{objc.GetClass("QCCompositionRepository")}
	})
	return QCCompositionRepositoryClass
}

type _QCCompositionRepositoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCCompositionRepository */
// An interface definition for the [QCCompositionRepository] class.
type IQCCompositionRepository interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QCCompositionRepository */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCCompositionRepository */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCCompositionRepository */
// Alloc allocates a new instance without initialization.
func (qc _QCCompositionRepositoryClass) Alloc() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCCompositionRepositoryClass) New() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionRepository) Init() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionRepository) Autorelease() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionRepository creates a new QCCompositionRepository instance.
func NewQCCompositionRepository() QCCompositionRepository {
	return getQCCompositionRepositoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCCompositionRepository */
// The class represents a system-wide centralized repository of built-in and installed Quartz Composer compositions ( and ). The class cannot be subclassed.
//
// Compositions in the repository are represented by the class. You can use the methods of the class to fetch all compositions or only those that meet specific criteria.


// The class represents a system-wide centralized repository of built-in and installed Quartz Composer compositions ( and ). The class cannot be subclassed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionRepository
type QCCompositionRepository struct {
	objectivec.Object
}

// QCCompositionRepositoryFrom constructs a [QCCompositionRepository] from an unsafe.Pointer.
//
// The class represents a system-wide centralized repository of built-in and installed Quartz Composer compositions ( and ). The class cannot be subclassed.
func QCCompositionRepositoryFrom(ptr unsafe.Pointer) QCCompositionRepository {
	return QCCompositionRepository{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCCompositionRepository *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCCompositionRepository */

// Returns the shared instance of the composition repository.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionRepository/shared()
func (qc _QCCompositionRepositoryClass) SharedCompositionRepository() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](objc.ID(qc.class), objc.Sel("sharedCompositionRepository"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCompositionRepository) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCCompositionRepository */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCCompositionRepository */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCCompositionRepository */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCCompositionRepository */



