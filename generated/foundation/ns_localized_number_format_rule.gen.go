// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLocalizedNumberFormatRule */


/* debug [class_header]: Header for NSLocalizedNumberFormatRule */
// The class instance for the [LocalizedNumberFormatRule] class.
var (
	LocalizedNumberFormatRuleClass     _LocalizedNumberFormatRuleClass
	LocalizedNumberFormatRuleClassOnce sync.Once
)

func getLocalizedNumberFormatRuleClass() _LocalizedNumberFormatRuleClass {
	LocalizedNumberFormatRuleClassOnce.Do(func() {
		LocalizedNumberFormatRuleClass = _LocalizedNumberFormatRuleClass{objc.GetClass("NSLocalizedNumberFormatRule")}
	})
	return LocalizedNumberFormatRuleClass
}

type _LocalizedNumberFormatRuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LocalizedNumberFormatRule */
// An interface definition for the [LocalizedNumberFormatRule] class.
type ILocalizedNumberFormatRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LocalizedNumberFormatRule */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LocalizedNumberFormatRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LocalizedNumberFormatRule */
// Alloc allocates a new instance without initialization.
func (lc _LocalizedNumberFormatRuleClass) Alloc() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LocalizedNumberFormatRuleClass) New() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalizedNumberFormatRule) Init() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalizedNumberFormatRule) Autorelease() LocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalizedNumberFormatRule creates a new LocalizedNumberFormatRule instance.
func NewLocalizedNumberFormatRule() LocalizedNumberFormatRule {
	return getLocalizedNumberFormatRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LocalizedNumberFormatRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule
type LocalizedNumberFormatRule struct {
	objectivec.Object
}

// LocalizedNumberFormatRuleFrom constructs a [LocalizedNumberFormatRule] from an unsafe.Pointer.
func LocalizedNumberFormatRuleFrom(ptr unsafe.Pointer) LocalizedNumberFormatRule {
	return LocalizedNumberFormatRule{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LocalizedNumberFormatRule *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LocalizedNumberFormatRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule/automatic
func (lc _LocalizedNumberFormatRuleClass) Automatic() ILocalizedNumberFormatRule {
	rv := objc.Send[LocalizedNumberFormatRule](objc.ID(lc.class), objc.Sel("automatic"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Automatic) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LocalizedNumberFormatRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LocalizedNumberFormatRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LocalizedNumberFormatRule */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLocalizedNumberFormatRule */



