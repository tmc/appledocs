// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [statistics] class.
var (
	StatisticsClass     _statisticsClass
	StatisticsClassOnce sync.Once
)

func getstatisticsClass() _statisticsClass {
	StatisticsClassOnce.Do(func() {
		StatisticsClass = _statisticsClass{objc.GetClass("statistics")}
	})
	return StatisticsClass
}

type _statisticsClass struct {
	class objc.Class
}





// An interface definition for the [statistics] class.
type Istatistics interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _statisticsClass) Alloc() statistics {
	rv := objc.Send[statistics](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _statisticsClass) New() statistics {
	rv := objc.Send[statistics](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ statistics) Init() statistics {
	rv := objc.Send[statistics](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ statistics) Autorelease() statistics {
	rv := objc.Send[statistics](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newstatistics creates a new statistics instance.
func Newstatistics() statistics {
	return getstatisticsClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/statistics-c.ivar
type statistics struct {
	objectivec.Object
}

// statisticsFrom constructs a [statistics] from an unsafe.Pointer.
func statisticsFrom(ptr unsafe.Pointer) statistics {
	return statistics{objectivec.Object{objc.ID(ptr)}}
}































