// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariViewControllerPrewarmingToken */


/* debug [class_header]: Header for SFSafariViewControllerPrewarmingToken */
// The class instance for the [SFSafariViewControllerPrewarmingToken] class.
var (
	SFSafariViewControllerPrewarmingTokenClass     _SFSafariViewControllerPrewarmingTokenClass
	SFSafariViewControllerPrewarmingTokenClassOnce sync.Once
)

func getSFSafariViewControllerPrewarmingTokenClass() _SFSafariViewControllerPrewarmingTokenClass {
	SFSafariViewControllerPrewarmingTokenClassOnce.Do(func() {
		SFSafariViewControllerPrewarmingTokenClass = _SFSafariViewControllerPrewarmingTokenClass{objc.GetClass("SFSafariViewControllerPrewarmingToken")}
	})
	return SFSafariViewControllerPrewarmingTokenClass
}

type _SFSafariViewControllerPrewarmingTokenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariViewControllerPrewarmingToken */
// An interface definition for the [SFSafariViewControllerPrewarmingToken] class.
type ISFSafariViewControllerPrewarmingToken interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariViewControllerPrewarmingToken */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariViewControllerPrewarmingToken */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariViewControllerPrewarmingToken */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerPrewarmingTokenClass) Alloc() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariViewControllerPrewarmingTokenClass) New() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerPrewarmingToken) Init() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerPrewarmingToken) Autorelease() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerPrewarmingToken creates a new SFSafariViewControllerPrewarmingToken instance.
func NewSFSafariViewControllerPrewarmingToken() SFSafariViewControllerPrewarmingToken {
	return getSFSafariViewControllerPrewarmingTokenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariViewControllerPrewarmingToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/PrewarmingToken
type SFSafariViewControllerPrewarmingToken struct {
	objectivec.Object
}

// SFSafariViewControllerPrewarmingTokenFrom constructs a [SFSafariViewControllerPrewarmingToken] from an unsafe.Pointer.
func SFSafariViewControllerPrewarmingTokenFrom(ptr unsafe.Pointer) SFSafariViewControllerPrewarmingToken {
	return SFSafariViewControllerPrewarmingToken{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariViewControllerPrewarmingToken *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariViewControllerPrewarmingToken */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariViewControllerPrewarmingToken */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariViewControllerPrewarmingToken */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariViewControllerPrewarmingToken */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariViewControllerPrewarmingToken */


