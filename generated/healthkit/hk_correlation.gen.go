// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKCorrelation */


/* debug [class_header]: Header for HKCorrelation */
// The class instance for the [HKCorrelation] class.
var (
	HKCorrelationClass     _HKCorrelationClass
	HKCorrelationClassOnce sync.Once
)

func getHKCorrelationClass() _HKCorrelationClass {
	HKCorrelationClassOnce.Do(func() {
		HKCorrelationClass = _HKCorrelationClass{objc.GetClass("HKCorrelation")}
	})
	return HKCorrelationClass
}

type _HKCorrelationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCorrelation */
// An interface definition for the [HKCorrelation] class.
type IHKCorrelation interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKCorrelation */
	// properties:
	CorrelationType() IHKCorrelationType
	Objects() unsafe.Pointer
	HKMetadataKeyFoodType() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathCorrelation() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCorrelation */
	// methods:
	ObjectsForType(objectType IHKObjectType) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCorrelation */
// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationClass) Alloc() HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKCorrelationClass) New() HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCorrelation) Init() HKCorrelation {
	rv := objc.Send[HKCorrelation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCorrelation) Autorelease() HKCorrelation {
	rv := objc.Send[HKCorrelation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCorrelation creates a new HKCorrelation instance.
func NewHKCorrelation() HKCorrelation {
	return getHKCorrelationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCorrelation */
// A sample that groups multiple related samples into a single entry.
//
// HealthKit uses correlations to represent both blood pressure and food. Blood pressure correlations always include two quantity samples, representing the systolic and diastolic values. Food correlations can contain a wide range of dietary information about the food, including information about the fat, protein, carbohydrates, energy, and vitamins consumed. In general, a food correlation should include at least a sample. You can also add nutritional quantity samples for any other items you want to track. Use the key to indicate the food’s name. The class is a concrete subclass of the class. Correlations are immutable: You set the correlation’s properties when the object is first created, and they cannot change.


// A sample that groups multiple related samples into a single entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation
type HKCorrelation struct {
	HKSample
}

// HKCorrelationFrom constructs a [HKCorrelation] from an unsafe.Pointer.
//
// A sample that groups multiple related samples into a single entry.
func HKCorrelationFrom(ptr unsafe.Pointer) HKCorrelation {
	return HKCorrelation{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCorrelation */

// Instantiates and returns a new correlation instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:)
func NewHKCorrelationWithTypeStartDateEndDateObjects(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer) HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(getHKCorrelationClass().class), objc.Sel("correlationWithType:startDate:endDate:objects:"), correlationType, startDate, endDate, objects)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCorrelationWithTypeStartDateEndDateObjects */


// Instantiates and returns a new correlation instance with the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:device:metadata:)
func NewHKCorrelationWithTypeStartDateEndDateObjectsDeviceMetadata(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer, device IHKDevice, metadata foundation.IDictionary) HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(getHKCorrelationClass().class), objc.Sel("correlationWithType:startDate:endDate:objects:device:metadata:"), correlationType, startDate, endDate, objects, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCorrelationWithTypeStartDateEndDateObjectsDeviceMetadata */


// Instantiates and returns a new correlation instance with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:metadata:)
func NewHKCorrelationWithTypeStartDateEndDateObjectsMetadata(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer, metadata foundation.IDictionary) HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(getHKCorrelationClass().class), objc.Sel("correlationWithType:startDate:endDate:objects:metadata:"), correlationType, startDate, endDate, objects, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCorrelationWithTypeStartDateEndDateObjectsMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCorrelation */

// Instantiates and returns a new correlation instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:)
func (hc _HKCorrelationClass) CorrelationWithTypeStartDateEndDateObjects(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("correlationWithType:startDate:endDate:objects:"), correlationType, startDate, endDate, objects)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrelationWithTypeStartDateEndDateObjects) */


// Instantiates and returns a new correlation instance with the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:device:metadata:)
func (hc _HKCorrelationClass) CorrelationWithTypeStartDateEndDateObjectsDeviceMetadata(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("correlationWithType:startDate:endDate:objects:device:metadata:"), correlationType, startDate, endDate, objects, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrelationWithTypeStartDateEndDateObjectsDeviceMetadata) */


// Instantiates and returns a new correlation instance with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/init(type:start:end:objects:metadata:)
func (hc _HKCorrelationClass) CorrelationWithTypeStartDateEndDateObjectsMetadata(correlationType IHKCorrelationType, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, objects unsafe.Pointer, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("correlationWithType:startDate:endDate:objects:metadata:"), correlationType, startDate, endDate, objects, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrelationWithTypeStartDateEndDateObjectsMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCorrelation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCorrelation */

// Returns a set containing all the objects of the specified type in the correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/objects(for:)
func (h_ HKCorrelation) ObjectsForType(objectType IHKObjectType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("objectsForType:"), objectType)
	return rv
}/* debug [instance_methods/method]: ObjectsForType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCorrelation */

// The type for this correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/correlationType
func (h_ HKCorrelation) CorrelationType() IHKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("correlationType"))
	return rv
}/* debug [instance_properties/getter]: correlationType */


// The set of sample objects that make up the correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation/objects
func (h_ HKCorrelation) Objects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("objects"))
	return rv
}/* debug [instance_properties/getter]: objects */


// The type of food that the HealthKit object represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyfoodtype
func (h_ HKCorrelation) HKMetadataKeyFoodType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKMetadataKeyFoodType"))
	return rv
}/* debug [instance_properties/getter]: HKMetadataKeyFoodType */


// The key path for accessing the object’s correlation inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcorrelation
func (h_ HKCorrelation) HKPredicateKeyPathCorrelation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCorrelation"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCorrelation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCorrelation */


