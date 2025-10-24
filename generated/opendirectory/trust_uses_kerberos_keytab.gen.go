// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustUsesKerberosKeytab */


/* debug [class_header]: Header for trustUsesKerberosKeytab */
// The class instance for the [trustUsesKerberosKeytab] class.
var (
	TrustUsesKerberosKeytabClass     _trustUsesKerberosKeytabClass
	TrustUsesKerberosKeytabClassOnce sync.Once
)

func gettrustUsesKerberosKeytabClass() _trustUsesKerberosKeytabClass {
	TrustUsesKerberosKeytabClassOnce.Do(func() {
		TrustUsesKerberosKeytabClass = _trustUsesKerberosKeytabClass{objc.GetClass("trustUsesKerberosKeytab")}
	})
	return TrustUsesKerberosKeytabClass
}

type _trustUsesKerberosKeytabClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustUsesKerberosKeytab */
// An interface definition for the [trustUsesKerberosKeytab] class.
type ItrustUsesKerberosKeytab interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustUsesKerberosKeytab */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustUsesKerberosKeytab */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustUsesKerberosKeytab */
// Alloc allocates a new instance without initialization.
func (tc _trustUsesKerberosKeytabClass) Alloc() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustUsesKerberosKeytabClass) New() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesKerberosKeytab) Init() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesKerberosKeytab) Autorelease() trustUsesKerberosKeytab {
	rv := objc.Send[trustUsesKerberosKeytab](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesKerberosKeytab creates a new trustUsesKerberosKeytab instance.
func NewtrustUsesKerberosKeytab() trustUsesKerberosKeytab {
	return gettrustUsesKerberosKeytabClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustUsesKerberosKeytab */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-c.ivar
type trustUsesKerberosKeytab struct {
	objectivec.Object
}

// trustUsesKerberosKeytabFrom constructs a [trustUsesKerberosKeytab] from an unsafe.Pointer.
func trustUsesKerberosKeytabFrom(ptr unsafe.Pointer) trustUsesKerberosKeytab {
	return trustUsesKerberosKeytab{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustUsesKerberosKeytab *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustUsesKerberosKeytab */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustUsesKerberosKeytab */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustUsesKerberosKeytab */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustUsesKerberosKeytab */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustUsesKerberosKeytab */



