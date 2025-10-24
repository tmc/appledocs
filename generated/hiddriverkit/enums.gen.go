// Code generated from Apple documentation for HIDDriverKit. DO NOT EDIT.

package hiddriverkit

/* debug [enums.gen.go]: Generating 3 enums for HIDDriverKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IOHIDElementCollectionType (7 cases) */
// IOHIDElementCollectionType - Constants that indicate the types of relationships that exist between two or more elements.
//
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType
type IOHIDElementCollectionType uint

const (
	// kIOHIDElementCollectionTypeApplication - A collection in which the child elements serve different purposes in a single device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeApplication
	kIOHIDElementCollectionTypeApplication IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypeLogical - A collection in which the child elements form a composite data structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeLogical
	kIOHIDElementCollectionTypeLogical IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypeNamedArray - A collection in which the elements are an array of selector usages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeNamedArray
	kIOHIDElementCollectionTypeNamedArray IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypePhysical - A collection in which the child elements are data points collected at one geometric point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypePhysical
	kIOHIDElementCollectionTypePhysical IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypeReport - A collection that wraps all the other elements in a report.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeReport
	kIOHIDElementCollectionTypeReport IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypeUsageModifier - A collection that modifies the meaning of the usage attached to the encompassing collection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeUsageModifier
	kIOHIDElementCollectionTypeUsageModifier IOHIDElementCollectionType = 0
	// kIOHIDElementCollectionTypeUsageSwitch - A collection that modifies the meaning of the usage it contains.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementCollectionType/kIOHIDElementCollectionTypeUsageSwitch
	kIOHIDElementCollectionTypeUsageSwitch IOHIDElementCollectionType = 0
)

/* debug [enums.gen.go]: Processing enum IOHIDElementType (8 cases) */
// IOHIDElementType - The types of HID elements that you can examine.
//
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType
type IOHIDElementType uint

const (
	// kIOHIDElementTypeCollection - The element acts as a parent container for two or more related elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeCollection
	kIOHIDElementTypeCollection IOHIDElementType = 0
	// kIOHIDElementTypeFeature - The element contains input and output data not intended for consumption by the end user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeFeature
	kIOHIDElementTypeFeature IOHIDElementType = 0
	// kIOHIDElementTypeInput_Axis - The element contains an axis of movement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Axis
	kIOHIDElementTypeInput_Axis IOHIDElementType = 0
	// kIOHIDElementTypeInput_Button - The element contains a one-bit input data field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Button
	kIOHIDElementTypeInput_Button IOHIDElementType = 0
	// kIOHIDElementTypeInput_Misc - The element contains an input data field of varying size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Misc
	kIOHIDElementTypeInput_Misc IOHIDElementType = 0
	// kIOHIDElementTypeInput_NULL - The element signals the end of an input data field in an input report.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_NULL
	kIOHIDElementTypeInput_NULL IOHIDElementType = 0
	// kIOHIDElementTypeInput_ScanCodes - The element contains an input field with a scan code or usage selector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_ScanCodes
	kIOHIDElementTypeInput_ScanCodes IOHIDElementType = 0
	// kIOHIDElementTypeOutput - The element contains a data field in an output report.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeOutput
	kIOHIDElementTypeOutput IOHIDElementType = 0
)

/* debug [enums.gen.go]: Processing enum IOHIDReportType (4 cases) */
// IOHIDReportType - Describes the different types of HID reports.
//
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType
type IOHIDReportType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeCount
	kIOHIDReportTypeCount IOHIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeFeature
	kIOHIDReportTypeFeature IOHIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeInput
	kIOHIDReportTypeInput IOHIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeOutput
	kIOHIDReportTypeOutput IOHIDReportType = 0
)


