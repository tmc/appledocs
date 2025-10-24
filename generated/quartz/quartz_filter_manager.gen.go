// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QuartzFilterManager */


/* debug [class_header]: Header for QuartzFilterManager */
// The class instance for the [QuartzFilterManager] class.
var (
	QuartzFilterManagerClass     _QuartzFilterManagerClass
	QuartzFilterManagerClassOnce sync.Once
)

func getQuartzFilterManagerClass() _QuartzFilterManagerClass {
	QuartzFilterManagerClassOnce.Do(func() {
		QuartzFilterManagerClass = _QuartzFilterManagerClass{objc.GetClass("QuartzFilterManager")}
	})
	return QuartzFilterManagerClass
}

type _QuartzFilterManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuartzFilterManager */
// An interface definition for the [QuartzFilterManager] class.
type IQuartzFilterManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QuartzFilterManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuartzFilterManager */
	// methods:
	Delegate() objc.ID
	FilterPanel() appkit.Panel
	FilterView() IQuartzFilterView
	ImportFilter(filterProperties objc.IObject /* cross-framework: NSDictionary */) IQuartzFilter
	SelectFilter(filter IQuartzFilter) bool
	SelectedFilter() IQuartzFilter
	SetDelegate(aDelegate objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuartzFilterManager */
// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterManagerClass) Alloc() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QuartzFilterManagerClass) New() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilterManager) Init() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilterManager) Autorelease() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilterManager creates a new QuartzFilterManager instance.
func NewQuartzFilterManager() QuartzFilterManager {
	return getQuartzFilterManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuartzFilterManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager
type QuartzFilterManager struct {
	objectivec.Object
}

// QuartzFilterManagerFrom constructs a [QuartzFilterManager] from an unsafe.Pointer.
func QuartzFilterManagerFrom(ptr unsafe.Pointer) QuartzFilterManager {
	return QuartzFilterManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuartzFilterManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuartzFilterManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/filterManager
func (qc _QuartzFilterManagerClass) FilterManager() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](objc.ID(qc.class), objc.Sel("filterManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FilterManager) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/filters(inDomains:)
func (qc _QuartzFilterManagerClass) FiltersInDomains(domains objc.IObject /* cross-framework: NSArray */) foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(qc.class), objc.Sel("filtersInDomains:"), domains)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FiltersInDomains) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuartzFilterManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuartzFilterManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/delegate()
func (q_ QuartzFilterManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](q_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/filterPanel()
func (q_ QuartzFilterManager) FilterPanel() appkit.Panel {
	rv := objc.Send[appkit.Panel](q_.ID, objc.Sel("filterPanel"))
	return rv
}/* debug [instance_methods/method]: FilterPanel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/filterView()
func (q_ QuartzFilterManager) FilterView() IQuartzFilterView {
	rv := objc.Send[QuartzFilterView](q_.ID, objc.Sel("filterView"))
	return rv
}/* debug [instance_methods/method]: FilterView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/importFilter(_:)
func (q_ QuartzFilterManager) ImportFilter(filterProperties objc.IObject /* cross-framework: NSDictionary */) IQuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("importFilter:"), filterProperties)
	return rv
}/* debug [instance_methods/method]: ImportFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/select(_:)
func (q_ QuartzFilterManager) SelectFilter(filter IQuartzFilter) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("selectFilter:"), filter)
	return rv
}/* debug [instance_methods/method]: SelectFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/selectedFilter()
func (q_ QuartzFilterManager) SelectedFilter() IQuartzFilter {
	rv := objc.Send[QuartzFilter](q_.ID, objc.Sel("selectedFilter"))
	return rv
}/* debug [instance_methods/method]: SelectedFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager/setDelegate(_:)
func (q_ QuartzFilterManager) SetDelegate(aDelegate objc.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setDelegate:"), aDelegate)
}/* debug [instance_methods/method]: SetDelegate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuartzFilterManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QuartzFilterManager */



