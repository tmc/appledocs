// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class streamerVersion */


/* debug [class_header]: Header for streamerVersion */
// The class instance for the [streamerVersion] class.
var (
	StreamerVersionClass     _streamerVersionClass
	StreamerVersionClassOnce sync.Once
)

func getstreamerVersionClass() _streamerVersionClass {
	StreamerVersionClassOnce.Do(func() {
		StreamerVersionClass = _streamerVersionClass{objc.GetClass("streamerVersion")}
	})
	return StreamerVersionClass
}

type _streamerVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for streamerVersion */
// An interface definition for the [streamerVersion] class.
type IstreamerVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for streamerVersion */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for streamerVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for streamerVersion */
// Alloc allocates a new instance without initialization.
func (sc _streamerVersionClass) Alloc() streamerVersion {
	rv := objc.Send[streamerVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _streamerVersionClass) New() streamerVersion {
	rv := objc.Send[streamerVersion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ streamerVersion) Init() streamerVersion {
	rv := objc.Send[streamerVersion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ streamerVersion) Autorelease() streamerVersion {
	rv := objc.Send[streamerVersion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewstreamerVersion creates a new streamerVersion instance.
func NewstreamerVersion() streamerVersion {
	return getstreamerVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for streamerVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/streamerVersion
type streamerVersion struct {
	objectivec.Object
}

// streamerVersionFrom constructs a [streamerVersion] from an unsafe.Pointer.
func streamerVersionFrom(ptr unsafe.Pointer) streamerVersion {
	return streamerVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for streamerVersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for streamerVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for streamerVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for streamerVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for streamerVersion */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class streamerVersion */



