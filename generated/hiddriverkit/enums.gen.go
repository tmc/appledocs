// Code generated from Apple documentation for HIDDriverKit. DO NOT EDIT.

package hiddriverkit

// Enum types and constants
// IOHIDElementType - The types of HID elements that you can examine.
//
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType
type HIDElementType uint

const (
	// kIOHIDElementTypeCollection - The element acts as a parent container for two or more related elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeCollection
	kIOHIDElementTypeCollection HIDElementType = 0
	// kIOHIDElementTypeFeature - The element contains input and output data not intended for consumption by the end user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeFeature
	kIOHIDElementTypeFeature HIDElementType = 0
	// kIOHIDElementTypeInput_Axis - The element contains an axis of movement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Axis
	kIOHIDElementTypeInput_Axis HIDElementType = 0
	// kIOHIDElementTypeInput_Button - The element contains a one-bit input data field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Button
	kIOHIDElementTypeInput_Button HIDElementType = 0
	// kIOHIDElementTypeInput_Misc - The element contains an input data field of varying size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_Misc
	kIOHIDElementTypeInput_Misc HIDElementType = 0
	// kIOHIDElementTypeInput_NULL - The element signals the end of an input data field in an input report.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_NULL
	kIOHIDElementTypeInput_NULL HIDElementType = 0
	// kIOHIDElementTypeInput_ScanCodes - The element contains an input field with a scan code or usage selector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeInput_ScanCodes
	kIOHIDElementTypeInput_ScanCodes HIDElementType = 0
	// kIOHIDElementTypeOutput - The element contains a data field in an output report.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElementType/kIOHIDElementTypeOutput
	kIOHIDElementTypeOutput HIDElementType = 0
)

// IOHIDReportType - Describes the different types of HID reports.
//
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType
type HIDReportType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeCount
	kIOHIDReportTypeCount HIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeFeature
	kIOHIDReportTypeFeature HIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeInput
	kIOHIDReportTypeInput HIDReportType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDReportType/kIOHIDReportTypeOutput
	kIOHIDReportTypeOutput HIDReportType = 0
)


