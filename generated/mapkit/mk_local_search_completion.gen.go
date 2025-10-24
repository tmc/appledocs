// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalSearchCompletion */


/* debug [class_header]: Header for MKLocalSearchCompletion */
// The class instance for the [MKLocalSearchCompletion] class.
var (
	MKLocalSearchCompletionClass     _MKLocalSearchCompletionClass
	MKLocalSearchCompletionClassOnce sync.Once
)

func getMKLocalSearchCompletionClass() _MKLocalSearchCompletionClass {
	MKLocalSearchCompletionClassOnce.Do(func() {
		MKLocalSearchCompletionClass = _MKLocalSearchCompletionClass{objc.GetClass("MKLocalSearchCompletion")}
	})
	return MKLocalSearchCompletionClass
}

type _MKLocalSearchCompletionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalSearchCompletion */
// An interface definition for the [MKLocalSearchCompletion] class.
type IMKLocalSearchCompletion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalSearchCompletion */
	// properties:
	Subtitle() objc.IObject /* cross-framework: NSString */
	SubtitleHighlightRanges() []foundation.Value
	Title() objc.IObject /* cross-framework: NSString */
	TitleHighlightRanges() []foundation.Value
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalSearchCompletion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalSearchCompletion */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchCompletionClass) Alloc() MKLocalSearchCompletion {
	rv := objc.Send[MKLocalSearchCompletion](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalSearchCompletionClass) New() MKLocalSearchCompletion {
	rv := objc.Send[MKLocalSearchCompletion](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalSearchCompletion) Init() MKLocalSearchCompletion {
	rv := objc.Send[MKLocalSearchCompletion](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalSearchCompletion) Autorelease() MKLocalSearchCompletion {
	rv := objc.Send[MKLocalSearchCompletion](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalSearchCompletion creates a new MKLocalSearchCompletion instance.
func NewMKLocalSearchCompletion() MKLocalSearchCompletion {
	return getMKLocalSearchCompletionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalSearchCompletion */
// A fully formed string that completes a partial string.
//
// You don’t create instances of this class directly. Instead, you use an to initiate a search based on a set of partial search strings. That object stores any matches in its results property. Retrieve any objects from that property and display the search terms in your interface, or use one to initiate a search for content based on that search term. When displaying text completions for a partial search term in your user interface, you might want to use a bold version of a font or add some other highlighting to the portion of the completion string that causes it to match the partial search term. To help you add this styling, the completion object includes highlight ranges for the title and subtitle strings.


// A fully formed string that completes a partial string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompletion
type MKLocalSearchCompletion struct {
	objectivec.Object
}

// MKLocalSearchCompletionFrom constructs a [MKLocalSearchCompletion] from an unsafe.Pointer.
//
// A fully formed string that completes a partial string.
func MKLocalSearchCompletionFrom(ptr unsafe.Pointer) MKLocalSearchCompletion {
	return MKLocalSearchCompletion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalSearchCompletion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalSearchCompletion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalSearchCompletion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalSearchCompletion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalSearchCompletion */

// The subtitle (if any) associated with the point of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompletion/subtitle
func (m_ MKLocalSearchCompletion) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The ranges of characters to highlight in the subtitle string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompletion/subtitleHighlightRanges
func (m_ MKLocalSearchCompletion) SubtitleHighlightRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("subtitleHighlightRanges"))
	return rv
}/* debug [instance_properties/getter]: subtitleHighlightRanges */


// The title string associated with the point of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompletion/title
func (m_ MKLocalSearchCompletion) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The ranges of characters to highlight in the title string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompletion/titleHighlightRanges
func (m_ MKLocalSearchCompletion) TitleHighlightRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("titleHighlightRanges"))
	return rv
}/* debug [instance_properties/getter]: titleHighlightRanges */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalSearchCompletion */



