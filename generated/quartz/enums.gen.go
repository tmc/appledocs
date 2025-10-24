// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

/* debug [enums.gen.go]: Generating 5 enums for Quartz */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IKCameraDeviceViewDisplayMode (3 cases) */
// IKCameraDeviceViewDisplayMode - These constants specify the display mode used by the camera view. These constants are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewDisplayMode
type IKCameraDeviceViewDisplayMode uint

const (
	// IKCameraDeviceViewDisplayModeIcon - Display the devices as icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewDisplayMode/icon
	IKCameraDeviceViewDisplayModeIcon IKCameraDeviceViewDisplayMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewDisplayMode/none
	IKCameraDeviceViewDisplayModeNone IKCameraDeviceViewDisplayMode = 0
	// IKCameraDeviceViewDisplayModeTable - Display the devices in as a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewDisplayMode/table
	IKCameraDeviceViewDisplayModeTable IKCameraDeviceViewDisplayMode = 0
)

/* debug [enums.gen.go]: Processing enum IKCameraDeviceViewTransferMode (2 cases) */
// IKCameraDeviceViewTransferMode - These constants specify the transfer mode used by the camera view. These constants are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewTransferMode
type IKCameraDeviceViewTransferMode uint

const (
	// IKCameraDeviceViewTransferModeFileBased - Transferred files will be saved to disk by the delegate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewTransferMode/fileBased
	IKCameraDeviceViewTransferModeFileBased IKCameraDeviceViewTransferMode = 0
	// IKCameraDeviceViewTransferModeMemoryBased - Transferred files will be supplied to the delegate as an   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKCameraDeviceViewTransferMode/memoryBased
	IKCameraDeviceViewTransferModeMemoryBased IKCameraDeviceViewTransferMode = 0
)

/* debug [enums.gen.go]: Processing enum IKDeviceBrowserViewDisplayMode (3 cases) */
// IKDeviceBrowserViewDisplayMode - These constants specify the display mode of the device browser.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserViewDisplayMode
type IKDeviceBrowserViewDisplayMode uint

const (
	// IKDeviceBrowserViewDisplayModeIcon - The devices are displayed as icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserViewDisplayMode/icon
	IKDeviceBrowserViewDisplayModeIcon IKDeviceBrowserViewDisplayMode = 0
	// IKDeviceBrowserViewDisplayModeOutline - The devices are displayed in an outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserViewDisplayMode/outline
	IKDeviceBrowserViewDisplayModeOutline IKDeviceBrowserViewDisplayMode = 0
	// IKDeviceBrowserViewDisplayModeTable - The devices are displayed in a table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKDeviceBrowserViewDisplayMode/table
	IKDeviceBrowserViewDisplayModeTable IKDeviceBrowserViewDisplayMode = 0
)

/* debug [enums.gen.go]: Processing enum IKScannerDeviceViewDisplayMode (3 cases) */
// IKScannerDeviceViewDisplayMode - These constants specify the display mode the scanner view will use. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewDisplayMode
type IKScannerDeviceViewDisplayMode uint

const (
	// IKScannerDeviceViewDisplayModeAdvanced - The view will display in advanced mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewDisplayMode/advanced
	IKScannerDeviceViewDisplayModeAdvanced IKScannerDeviceViewDisplayMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewDisplayMode/none
	IKScannerDeviceViewDisplayModeNone IKScannerDeviceViewDisplayMode = 0
	// IKScannerDeviceViewDisplayModeSimple - The view will display in simple mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewDisplayMode/simple
	IKScannerDeviceViewDisplayModeSimple IKScannerDeviceViewDisplayMode = 0
)

/* debug [enums.gen.go]: Processing enum IKScannerDeviceViewTransferMode (2 cases) */
// IKScannerDeviceViewTransferMode - These constants determine how the scanner data is returned to the delegate. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewTransferMode
type IKScannerDeviceViewTransferMode uint

const (
	// IKScannerDeviceViewTransferModeFileBased - The scanned content will be saved to the specified download directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewTransferMode/fileBased
	IKScannerDeviceViewTransferModeFileBased IKScannerDeviceViewTransferMode = 0
	// IKScannerDeviceViewTransferModeMemoryBased - The scanned data is returned to the delegate as a   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKScannerDeviceViewTransferMode/memoryBased
	IKScannerDeviceViewTransferModeMemoryBased IKScannerDeviceViewTransferMode = 0
)


