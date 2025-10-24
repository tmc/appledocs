// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [QCCompositionRepository] class.
type IQCCompositionRepository interface {
	objectivec.IObject
	// properties:
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionRepositoryClass) Alloc() QCCompositionRepository {
	rv := objc.Send[QCCompositionRepository](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




