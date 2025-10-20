# HIDDriverKit Framework Overview

This example provides a comprehensive educational overview of Apple's HIDDriverKit framework for developing Human Interface Device (HID) drivers.

## Overview

HIDDriverKit is part of Apple's DriverKit framework for developing user-space device drivers. It enables:
- Development of HID device drivers (keyboards, mice, game controllers)
- User-space driver implementation (safer than kernel extensions)
- Custom hardware support
- System Extension-based architecture
- Event handling and device communication
- HID report parsing and generation

## Building

```bash
go build .
```

## Running

```bash
# Show comprehensive framework overview
./driver-overview
```

## Features Demonstrated

1. **Framework Overview** - HID driver capabilities and architecture
2. **Device Types** - Keyboards, pointing devices, game controllers, digitizers, sensors
3. **DriverKit Architecture** - System Extension-based driver model
4. **Core Classes** - IOUserHIDEventService, IOHIDEventService, IOHIDInterface
5. **Use Cases** - Custom input devices, industrial equipment, accessibility, gaming
6. **HID Report Descriptor** - Device capability description format
7. **Event Types** - Keyboard, pointer, digitizer, game controller, consumer, sensor events
8. **Driver Lifecycle** - From activation to event dispatch
9. **Report Processing** - Input/output report handling
10. **Development Workflow** - From Xcode project to installed driver
11. **Device Matching** - Info.plist configuration for hardware matching
12. **Performance Best Practices** - Efficient event processing
13. **Security & Entitlements** - Code signing and user approval
14. **Debugging Strategies** - Logs, packet analyzers, HID tools
15. **Testing Approach** - Unit, integration, and performance testing
16. **Common Patterns** - Button tracking, coordinate mapping, gesture recognition
17. **Requirements** - macOS version, Xcode, Apple Developer account
18. **Limitations** - User approval, sandbox restrictions, debugging challenges
19. **HID Specification** - Usage pages and common device types
20. **Resources** - Documentation, tools, sample code

## DriverKit Architecture

```
┌─────────────────────────────────────┐
│     User Application                │
│  (IOHIDManager, IOHIDDevice)        │
└────────────────┬────────────────────┘
                 │ HID Events
┌────────────────┴────────────────────┐
│  System Extension (Your Driver)     │
│  IOUserHIDEventService              │
└────────────────┬────────────────────┘
                 │ USB/Bluetooth/I2C
┌────────────────┴────────────────────┐
│      Hardware Device                │
│  (Keyboard, Mouse, Controller)      │
└─────────────────────────────────────┘
```

## Supported Device Types

- **Keyboards** - Standard and custom keyboard devices
- **Pointing Devices** - Mice, trackpads, trackballs
- **Game Controllers** - Gamepads, joysticks, steering wheels
- **Digitizers** - Graphics tablets, pen displays, touch screens
- **Sensors** - Accelerometers, gyroscopes, environmental sensors
- **Custom HID Devices** - Specialized input devices
- **Multi-Function Devices** - Combined input/output devices

## Core Classes

### IOUserHIDEventService
- Main driver service class
- Handles device lifecycle (start, stop)
- Processes HID reports from hardware
- Dispatches events to the system
- Manages device state

### IOHIDEventService
- Event generation and dispatch
- Coordinate space management
- Event filtering and processing
- Property management

### IOHIDInterface
- Hardware interface abstraction
- Report descriptor parsing
- Communication with physical device
- Transport-agnostic (USB, Bluetooth, I2C)

## HID Report Descriptor

The report descriptor describes device capabilities in a standardized format:

### Components
- **Usage Pages** - Define device category (Generic Desktop, Consumer, Digitizer)
- **Usages** - Specific controls (X axis, Y axis, Button 1, Volume)
- **Report IDs** - Support multiple report types from one device
- **Input/Output/Feature Reports** - Data direction and purpose
- **Logical/Physical Min/Max** - Value ranges for each control
- **Unit/Unit Exponent** - Physical units (cm, degrees, etc.)

### Example (Simple Mouse)
```
Usage Page: Generic Desktop (0x01)
Usage: Mouse (0x02)
Collection: Application
  Usage: Pointer (0x01)
  Collection: Physical
    Usage Page: Button (0x09)
    Usage Minimum: Button 1 (0x01)
    Usage Maximum: Button 3 (0x03)
    Logical Minimum: 0
    Logical Maximum: 1
    Report Count: 3
    Report Size: 1
    Input: (Data, Variable, Absolute)  ; 3 buttons

    Usage Page: Generic Desktop (0x01)
    Usage: X (0x30)
    Usage: Y (0x31)
    Logical Minimum: -127
    Logical Maximum: 127
    Report Size: 8
    Report Count: 2
    Input: (Data, Variable, Relative)  ; X, Y movement
  End Collection
End Collection
```

## Event Processing Flow

### Input Report Processing (Device → System)
1. Hardware generates data (button press, stick movement)
2. Device sends HID report via USB/Bluetooth/I2C
3. Driver receives raw report bytes
4. Driver parses report using cached descriptor info
5. Driver creates IOHIDEvent objects (keyboard, pointer, etc.)
6. Events dispatched to system via `dispatchEvent()`
7. System routes events to applications

### Output Report Processing (System → Device)
1. System requests device state change (keyboard LED)
2. Driver receives request
3. Driver formats HID output report
4. Driver sends report to device via transport
5. Device updates state (LED turns on)

## Development Workflow

1. **Create DriverKit Extension** - Add System Extension target in Xcode
2. **Subclass IOUserHIDEventService** - Implement driver logic
3. **Configure Info.plist** - Set device matching criteria
4. **Implement Start()** - Initialize driver resources
5. **Parse Report Descriptor** - Understand device capabilities
6. **Implement handleReport()** - Process incoming HID reports
7. **Generate Events** - Create IOHIDEvent objects
8. **Dispatch Events** - Call `dispatchEvent()` to send to system
9. **Implement Stop()** - Clean up resources
10. **Code Sign** - Sign with valid Developer ID
11. **Install Extension** - Activate System Extension
12. **Test** - Verify with hardware device

## Device Matching Example

Configure Info.plist to match your hardware:

```xml
<key>IOKitPersonalities</key>
<dict>
    <key>MyCustomKeyboard</key>
    <dict>
        <key>CFBundleIdentifier</key>
        <string>com.example.MyDriver</string>
        <key>IOClass</key>
        <string>MyDriverClass</string>
        <key>IOProviderClass</key>
        <string>IOUSBHostHIDDevice</string>
        <key>IOUserClass</key>
        <string>MyDriverClass</string>
        <key>Transport</key>
        <string>USB</string>
        <key>idVendor</key>
        <integer>0x1234</integer>
        <key>idProduct</key>
        <integer>0x5678</integer>
        <key>DeviceUsagePage</key>
        <integer>1</integer>
        <key>DeviceUsage</key>
        <integer>6</integer>
    </dict>
</dict>
```

## Performance Best Practices

1. **Minimize handleReport() processing** - Keep event processing fast
2. **Cache descriptor information** - Parse once, reuse often
3. **Batch events when possible** - Reduce dispatch overhead
4. **Avoid blocking operations** - Use async I/O
5. **Optimize memory allocation** - Reuse event objects
6. **Monitor CPU usage** - Profile driver performance
7. **Handle high event rates** - Gaming mice can send 1000+ events/sec
8. **Implement error handling** - Don't crash on malformed reports

## Security & Entitlements

### Required Entitlements
```xml
<key>com.apple.developer.system-extension.install</key>
<true/>
<key>com.apple.developer.driverkit</key>
<true/>
<key>com.apple.developer.driverkit.family.hid.device</key>
<true/>
<key>com.apple.developer.driverkit.transport.usb</key>
<true/>
```

### Security Model
- System Extension runs in user space (not kernel)
- Sandboxed environment with limited privileges
- User approval required via System Preferences
- Cannot be installed silently
- Code signing with Developer ID required
- Communication validated by system

## Debugging

### Logging
```swift
import os.log

let log = OSLog(subsystem: "com.example.MyDriver", category: "HID")
os_log("Received report: %{public}@", log: log, type: .debug, reportData)
```

### Tools
- **Console.app** - View driver logs in real-time
- **Wireshark** - Capture USB/Bluetooth packets
- **USB Prober** - Enumerate devices and view descriptors
- **Xcode Instruments** - Profile CPU and memory usage
- **lldb** - Attach debugger to System Extension

### Common Issues
- Driver not loading → Check entitlements, code signing
- Events not delivered → Verify `dispatchEvent()` calls
- High CPU usage → Profile `handleReport()` performance
- Device not matching → Verify Info.plist matching criteria

## Testing Strategy

### Unit Testing
- Test report parsing logic in isolation
- Validate event generation from reports
- Test edge cases (invalid reports, overflow)
- Verify coordinate transformations

### Integration Testing
- Test with real hardware devices
- Verify event delivery to test applications
- Test device connect/disconnect cycles
- Test multiple simultaneous devices
- Verify output report handling

### Performance Testing
- High event rate scenarios (1000 Hz gaming mice)
- CPU usage profiling under load
- Memory leak detection over extended use
- Latency measurement (input to event delivery)

## Common Implementation Patterns

### Button State Tracking
```swift
var previousButtonState: UInt8 = 0

func handleReport(_ report: Data) {
    let currentButtons = report[0]
    let pressed = currentButtons & ~previousButtonState
    let released = previousButtonState & ~currentButtons

    // Generate button events...
    previousButtonState = currentButtons
}
```

### Coordinate Mapping
```swift
func mapDeviceToScreen(x: Int16, y: Int16) -> (CGFloat, CGFloat) {
    let screenX = CGFloat(x) / deviceMaxX * screenWidth
    let screenY = CGFloat(y) / deviceMaxY * screenHeight
    return (screenX, screenY)
}
```

### Dead Zone Handling
```swift
func applyDeadZone(_ value: Int16, threshold: Int16) -> Int16 {
    if abs(value) < threshold {
        return 0
    }
    return value
}
```

## HID Specification Basics

### Usage Pages
- `0x01` - Generic Desktop (mouse, keyboard, joystick)
- `0x02` - Simulation Controls (flight stick, steering wheel)
- `0x03` - VR Controls (head tracker, gloves)
- `0x04` - Sport Controls (rowing, skiing)
- `0x05` - Game Controls (gamepad, pinball)
- `0x0C` - Consumer (media controls, volume)
- `0x0D` - Digitizer (pen, touch, multi-touch)

### Common Device Usages
- `0x01/0x02` - Mouse
- `0x01/0x06` - Keyboard
- `0x01/0x04` - Joystick
- `0x01/0x05` - Game Pad
- `0x0D/0x02` - Pen
- `0x0D/0x04` - Touch Screen

## Requirements

- macOS 10.15+ (Catalina or later)
- Xcode 11+ with DriverKit SDK
- Apple Developer account ($99/year)
- Valid Developer ID certificate
- Hardware device for testing
- System Extension approval workflow
- Understanding of HID specification

## Limitations

This example is educational only. HIDDriverKit requires:

### Development Constraints
- Cannot run in Simulator (real hardware required)
- Requires user approval (cannot be silent)
- More complex than app development
- Limited to HID-class devices
- Cannot directly access non-HID hardware

### Sandbox Restrictions
- Limited file system access
- Restricted network access
- Cannot spawn processes
- Limited IPC options

### Debugging Challenges
- Cannot attach debugger during activation
- Logs require Console.app
- Crashes harder to diagnose than apps
- Performance profiling more complex

## Use Cases

### Custom Input Devices
- Specialized keyboards with custom layouts
- Game controllers with unique button configurations
- Control panels for audio/video equipment

### Industrial Equipment
- Data entry terminals
- Industrial HMI devices
- Manufacturing control interfaces

### Medical Devices
- Patient monitoring input devices
- Medical imaging control panels
- Diagnostic equipment interfaces

### Accessibility Devices
- Adaptive keyboards for disabilities
- Custom switch interfaces
- Alternative input methods

### Gaming Peripherals
- Custom game controllers
- Racing wheels and pedals
- Flight simulation controls

### Professional Tools
- Graphics tablets
- 3D input devices (SpaceNavigator)
- Audio control surfaces

## References

- [HIDDriverKit Documentation](https://developer.apple.com/documentation/hiddriverkit)
- [DriverKit Framework](https://developer.apple.com/documentation/driverkit)
- [USB HID Specification](https://www.usb.org/hid)
- [HID Usage Tables](https://www.usb.org/document-library/hid-usage-tables-13)
- [System Extensions](https://developer.apple.com/system-extensions/)
- [DriverKit Tutorials](https://developer.apple.com/tutorials/driverkit)

## Additional Resources

### Tools
- **Wireshark** - USB/Bluetooth packet analyzer
- **USB Prober** - Device enumeration and descriptor viewing
- **Xcode Instruments** - Performance profiling
- **Console.app** - System log viewer
- **HID Descriptor Tool** - Parse and validate descriptors

### Sample Code
- Apple's DriverKit sample projects in Xcode
- Open-source HID driver examples on GitHub
- Community driver implementations

## See Also

- `../../generated/hiddriverkit/` - Generated HIDDriverKit bindings
- DriverKit framework for driver development
- IOKit for hardware interaction
- System Extensions for extension lifecycle

## Production Considerations

When developing real HID drivers:

1. **Thorough Testing** - Test with multiple hardware variations
2. **Error Handling** - Handle malformed reports gracefully
3. **Performance** - Optimize for high event rates
4. **User Experience** - Smooth installation and update process
5. **Documentation** - Document supported hardware clearly
6. **Support** - Plan for user support and bug reports
7. **Updates** - Strategy for driver updates and maintenance
8. **Compatibility** - Test across macOS versions

## Notes

This example demonstrates framework concepts only. Real HID driver implementation requires:
- System Extension project in Xcode with DriverKit target
- IOUserHIDEventService subclass with event handling
- Proper entitlements and code signing
- Hardware device for testing and debugging
- User approval workflow for installation
- Understanding of HID specification and report descriptors

The generated bindings in `../../generated/hiddriverkit/` provide the Go interface to HIDDriverKit, but actual driver development must be done in C++ within a DriverKit System Extension.
