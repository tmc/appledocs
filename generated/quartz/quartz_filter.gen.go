// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QuartzFilter */


/* debug [class_header]: Header for QuartzFilter */
// The class instance for the [QuartzFilter] class.
var (
	QuartzFilterClass     _QuartzFilterClass
	QuartzFilterClassOnce sync.Once
)

func getQuartzFilterClass() _QuartzFilterClass {
	QuartzFilterClassOnce.Do(func() {
		QuartzFilterClass = _QuartzFilterClass{objc.GetClass("QuartzFilter")}
	})
	return QuartzFilterClass
}

type _QuartzFilterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuartzFilter */
// An interface definition for the [QuartzFilter] class.
type IQuartzFilter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QuartzFilter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuartzFilter */
	// methods:
	ApplyToContext(aContext ContextRef /* not a class type */) bool
	LocalizedName() foundation.String
	Properties() foundation.Dictionary
	RemoveFromContext(aContext ContextRef /* not a class type */)
	Url() foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuartzFilter */
// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterClass) Alloc() QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QuartzFilterClass) New() QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilter) Init() QuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilter) Autorelease() QuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilter creates a new QuartzFilter instance.
func NewQuartzFilter() QuartzFilter {
	return getQuartzFilterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuartzFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter
type QuartzFilter struct {
	objectivec.Object
}

// QuartzFilterFrom constructs a [QuartzFilter] from an unsafe.Pointer.
func QuartzFilterFrom(ptr unsafe.Pointer) QuartzFilter {
	return QuartzFilter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuartzFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(outputIntents:)
func NewQuartzFilterWithOutputIntents(outputIntents objc.IObject /* cross-framework: NSArray */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(getQuartzFilterClass().class), objc.Sel("quartzFilterWithOutputIntents:"), outputIntents)
	return rv
}/* debug [class_init_methods/constructor]: NewQuartzFilterWithOutputIntents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(properties:)
func NewQuartzFilterWithProperties(properties objc.IObject /* cross-framework: NSDictionary */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(getQuartzFilterClass().class), objc.Sel("quartzFilterWithProperties:"), properties)
	return rv
}/* debug [class_init_methods/constructor]: NewQuartzFilterWithProperties */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(url:)
func NewQuartzFilterWithURL(aURL objc.IObject /* cross-framework: NSURL */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(getQuartzFilterClass().class), objc.Sel("quartzFilterWithURL:"), aURL)
	return rv
}/* debug [class_init_methods/constructor]: NewQuartzFilterWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuartzFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(outputIntents:)
func (qc _QuartzFilterClass) QuartzFilterWithOutputIntents(outputIntents objc.IObject /* cross-framework: NSArray */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("quartzFilterWithOutputIntents:"), outputIntents)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuartzFilterWithOutputIntents) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(properties:)
func (qc _QuartzFilterClass) QuartzFilterWithProperties(properties objc.IObject /* cross-framework: NSDictionary */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("quartzFilterWithProperties:"), properties)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuartzFilterWithProperties) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/init(url:)
func (qc _QuartzFilterClass) QuartzFilterWithURL(aURL objc.IObject /* cross-framework: NSURL */) QuartzFilter {
	rv := objc.Send[QuartzFilter](objc.ID(qc.class), objc.Sel("quartzFilterWithURL:"), aURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuartzFilterWithURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuartzFilter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuartzFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/apply(to:)
func (q_ QuartzFilter) ApplyToContext(aContext ContextRef /* not a class type */) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("applyToContext:"), aContext)
	return rv
}/* debug [instance_methods/method]: ApplyToContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/localizedName()
func (q_ QuartzFilter) LocalizedName() foundation.String {
	rv := objc.Send[foundation.String](q_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_methods/method]: LocalizedName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/properties()
func (q_ QuartzFilter) Properties() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](q_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_methods/method]: Properties */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/remove(from:)
func (q_ QuartzFilter) RemoveFromContext(aContext ContextRef /* not a class type */) {
	objc.Send[objc.ID](q_.ID, objc.Sel("removeFromContext:"), aContext)
}/* debug [instance_methods/method]: RemoveFromContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilter/url()
func (q_ QuartzFilter) Url() foundation.URL {
	rv := objc.Send[foundation.URL](q_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_methods/method]: Url */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuartzFilter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QuartzFilter */


