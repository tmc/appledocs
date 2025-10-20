package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/devicediscoveryextension"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("DeviceDiscoveryExtension Framework Examples")
	fmt.Println("==========================================")

	// Example 1: Framework overview
	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   DeviceDiscoveryExtension enables third-party apps to:")
	fmt.Println("   - Discover devices on local network or via Bluetooth")
	fmt.Println("   - Present devices in system device picker UI")
	fmt.Println("   - Handle device connection and communication")
	fmt.Println("   - Support media streaming protocols (DIAL, etc.)")

	// Example 2: Create DDDevice instance
	fmt.Println("\n2. Creating DDDevice Instance:")
	fmt.Println("   Note: DDDevice instances are typically created within")
	fmt.Println("   the extension context, not in standalone apps.")
	fmt.Println("   ")
	fmt.Println("   In an extension:")
	fmt.Println("   device := devicediscoveryextension.NewDDDevice()")

	// Prevent unused import error
	_ = devicediscoveryextension.NewDDDevice

	// Example 3: Device properties
	fmt.Println("\n3. DDDevice Properties:")

	properties := map[string]string{
		"displayName":        "User-visible name for the device",
		"identifier":         "Unique identifier for the device",
		"category":           "Device category (TV, speaker, etc.)",
		"protocolType":       "Communication protocol (DIAL, etc.)",
		"bluetoothIdentifier": "Bluetooth UUID for BLE devices",
		"networkEndpoint":    "Network endpoint for TCP/IP devices",
		"state":              "Device state (activated, authorized, etc.)",
		"mediaPlaybackState": "Current playback state",
		"mediaContentTitle":  "Currently playing content title",
		"mediaContentSubtitle": "Currently playing content subtitle",
	}

	for property, description := range properties {
		fmt.Printf("   %-25s: %s\n", property, description)
	}

	// Example 4: Device categories
	fmt.Println("\n4. Device Categories:")

	categories := []string{
		"accessorySetup - Generic accessory setup",
		"tv - Television or display",
		"tvWithMediaBox - TV with streaming box",
		"hifiSpeaker - Single Hi-Fi speaker",
		"hifiSpeakerMultiple - Multiple speakers",
		"laptopComputer - Laptop computer",
		"desktopComputer - Desktop computer",
	}

	for i, category := range categories {
		fmt.Printf("   %d. %s\n", i+1, category)
	}

	// Example 5: Device protocols
	fmt.Println("\n5. Supported Protocols:")

	protocols := []string{
		"DIAL (Discovery and Launch) - Media streaming protocol",
		"Custom protocols via UTType",
	}

	for i, protocol := range protocols {
		fmt.Printf("   %d. %s\n", i+1, protocol)
	}

	// Example 6: DDDeviceEvent
	fmt.Println("\n6. Device Events:")

	events := map[string]string{
		"deviceFound":   "New device discovered",
		"deviceChanged": "Device properties updated",
		"deviceLost":    "Device no longer available",
		"unknown":       "Unknown event type",
	}

	for event, description := range events {
		fmt.Printf("   %-15s: %s\n", event, description)
	}

	// Example 7: DDDiscoverySession
	fmt.Println("\n7. Discovery Session:")
	fmt.Println("   DDDiscoverySession manages device discovery lifecycle:")
	fmt.Println("   - Receives start/stop discovery requests from system")
	fmt.Println("   - Reports discovered devices via reportEvent()")
	fmt.Println("   - Updates device status changes")

	// Example 8: Extension implementation workflow
	fmt.Println("\n8. Extension Implementation Workflow:")

	workflow := []string{
		"1. Create app extension target in Xcode",
		"2. Implement DDDiscoveryExtension protocol",
		"3. Implement startDiscovery(session:) method",
		"4. Start scanning for devices (BLE or network)",
		"5. Create DDDevice instances for found devices",
		"6. Report devices via session.reportEvent()",
		"7. Implement stopDiscovery(session:) method",
		"8. Handle device state changes via didReceiveEvent()",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 9: Device state lifecycle
	fmt.Println("\n9. Device State Lifecycle:")

	states := []string{
		"invalid - Initial/invalid state",
		"activating - Device being activated",
		"activated - Device ready for use",
		"authorized - User authorized device",
		"invalidating - Device being deactivated",
	}

	for i, state := range states {
		fmt.Printf("   %d. %s\n", i+1, state)
	}

	// Example 10: Creating a device with protocol
	fmt.Println("\n10. Creating Device with Protocol:")
	fmt.Println("   // Example code structure:")
	fmt.Println("   protocolType := uniformtypeidentifiers.NewUTType()")
	fmt.Println("   // Note: In real usage, specify actual protocol UTType")
	fmt.Println("   ")
	fmt.Println("   device := devicediscoveryextension.NewDDDeviceWithDisplayNameCategoryProtocolTypeIdentifier(")
	fmt.Println("       \"Living Room TV\",")
	fmt.Println("       nil, // category (Swift enum, use unsafe.Pointer)")
	fmt.Println("       protocolType,")
	fmt.Println("       \"com.example.device.livingroom-tv\")")

	// Example 11: Device capabilities
	fmt.Println("\n11. Device Capabilities (DDDeviceSupports):")

	capabilities := []string{
		"bluetoothPairingLE - Bluetooth Low Energy pairing",
		"bluetoothHID - Human Interface Device support",
		"bluetoothTransportBridging - Transport bridging",
	}

	for i, capability := range capabilities {
		fmt.Printf("   %d. %s\n", i+1, capability)
	}

	// Example 12: Use cases
	fmt.Println("\n12. Use Cases:")

	useCases := map[string]string{
		"Media Streaming":     "Discover and stream to smart TVs/speakers",
		"Smart Home":          "Find and configure IoT devices",
		"Wireless Display":    "Connect to wireless displays/projectors",
		"Gaming":              "Discover game controllers and accessories",
		"Audio Distribution":  "Multi-room audio system discovery",
		"Screen Mirroring":    "AirPlay-like device discovery",
	}

	for useCase, description := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, description)
	}

	// Example 13: Requirements
	fmt.Println("\n13. Requirements:")

	requirements := []string{
		"macOS 13.0+ or iOS 16.0+",
		"App extension with DDDiscoveryExtension protocol",
		"Appropriate entitlements for network/Bluetooth access",
		"Info.plist configuration for extension",
		"User permission for local network access",
	}

	for i, requirement := range requirements {
		fmt.Printf("   %d. %s\n", i+1, requirement)
	}

	// Example 14: Media playback states
	fmt.Println("\n14. Media Playback States:")

	playbackStates := []string{
		"noContent - No media loaded",
		"playing - Media actively playing",
		"paused - Media paused",
	}

	for i, state := range playbackStates {
		fmt.Printf("   %d. %s\n", i+1, state)
	}

	// Example 15: Implementation tips
	fmt.Println("\n15. Implementation Tips:")

	tips := []string{
		"Use Bonjour for local network device discovery",
		"Implement proper device deduplication by identifier",
		"Update device properties when connection state changes",
		"Handle network changes gracefully (WiFi switches, etc.)",
		"Test with actual hardware devices when possible",
		"Follow Apple's extension best practices for memory/power",
		"Cache device information for faster rediscovery",
	}

	for i, tip := range tips {
		fmt.Printf("   %d. %s\n", i+1, tip)
	}

	fmt.Println("\n✓ DeviceDiscoveryExtension framework examples completed!")
	fmt.Println("\nNote: DeviceDiscoveryExtension is used in app extensions:")
	fmt.Println("  - Not typically used in main app code directly")
	fmt.Println("  - Requires app extension target configuration")
	fmt.Println("  - Extension runs when user opens device picker")
	fmt.Println("  - System manages extension lifecycle")
	fmt.Println("\nReal extensions would:")
	fmt.Println("  - Implement DDDiscoveryExtension protocol in Swift")
	fmt.Println("  - Use Core Bluetooth or Network framework for discovery")
	fmt.Println("  - Handle all required protocol methods")
	fmt.Println("  - Report devices to system via DDDiscoverySession")
	fmt.Println("  - Update device state based on system callbacks")
}
