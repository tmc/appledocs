# ServiceManagement Framework Overview

This example provides a comprehensive educational overview of Apple's ServiceManagement framework for managing launch agents, daemons, and login items.

## Overview

ServiceManagement is Apple's framework for managing background services and startup items. It enables:
- Launch agent and daemon registration
- Login item management (modern and legacy)
- Background service control
- System Extension registration
- XPC service discovery
- App-managed background tasks

## Building

```bash
go build .
```

## Running

```bash
# Show comprehensive framework overview
./launch-management
```

## Features Demonstrated

1. **Framework Overview** - Background service management capabilities
2. **Service Types** - Login items, launch agents, daemons, system extensions
3. **Login Items** - Modern SMAppService vs legacy LSSharedFileList
4. **Launch Agents vs Daemons** - User-level vs system-level services
5. **SMAppService** - Modern API for macOS 13+
6. **Use Cases** - Menu bar apps, sync services, monitoring, auto-update
7. **Service Lifecycle** - Registration to execution
8. **Registration Process** - SMAppService Swift examples
9. **Launchd Property List** - Traditional plist configuration
10. **plist Keys** - Common configuration options
11. **System Extensions** - DriverKit, Network, Endpoint Security
12. **XPC Services** - Inter-process communication
13. **Security Considerations** - Permissions, code signing, sandboxing
14. **Development Workflow** - From Xcode project to installed service
15. **Debugging Strategies** - Console.app, launchctl, Activity Monitor
16. **launchctl Commands** - Command-line service management
17. **Testing Strategy** - Unit, integration, manual testing
18. **Migration Guide** - From legacy APIs to SMAppService
19. **Best Practices** - Modern service development
20. **Requirements** - macOS versions, code signing, entitlements
21. **Limitations** - User approval, sandboxing, debugging

## Service Types

### Login Items
- **Modern (macOS 13+)**: Managed by SMAppService, user-visible in System Settings
- **Legacy**: Managed by LSSharedFileList (deprecated), System Preferences → Users & Groups
- User can enable/disable without app restart
- Can be background services or UI applications

### Launch Agents
- Run as the logged-in user
- Access to user's files and preferences
- Can display UI
- Location: `~/Library/LaunchAgents/`
- Use cases: Menu bar apps, user-specific services

### Launch Daemons
- Run as root or specific user
- No UI access
- System-wide services
- Location: `/Library/LaunchDaemons/`
- Use cases: System services, network servers

### System Extensions
- DriverKit - User-space device drivers
- Network Extensions - VPN, content filters, DNS proxy
- Endpoint Security - System monitoring and security
- Requires user approval
- Managed via System Preferences

## SMAppService (macOS 13+)

Modern API for managing services:

```swift
import ServiceManagement

// Register login item
let service = SMAppService.loginItem(identifier: "com.example.helper")
try service.register()

// Check status
let status = service.status
// .enabled, .requiresApproval, .notRegistered, .notFound

// Unregister
try service.unregister()
```

### Service Types
- **LoginItem** - Visible apps launched at login
- **Agent** - Background services for current user
- **Daemon** - System-wide background services

## Launchd Property List Configuration

Traditional launchd services use XML property lists:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
  "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.example.myagent</string>

    <key>ProgramArguments</key>
    <array>
        <string>/Applications/MyApp.app/Contents/Library/LoginItems/MyHelper.app/Contents/MacOS/MyHelper</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>KeepAlive</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/tmp/myagent.stdout</string>

    <key>StandardErrorPath</key>
    <string>/tmp/myagent.stderr</string>
</dict>
</plist>
```

### Common plist Keys

| Key | Description |
|-----|-------------|
| `Label` | Unique identifier for the service |
| `ProgramArguments` | Command and arguments to execute |
| `RunAtLoad` | Start service when loaded |
| `KeepAlive` | Restart if service exits |
| `StartInterval` | Run periodically (seconds) |
| `StartCalendarInterval` | Run at specific times |
| `StandardOutPath` | Redirect stdout to file |
| `StandardErrorPath` | Redirect stderr to file |
| `EnvironmentVariables` | Environment variables for service |
| `WorkingDirectory` | Working directory for service |
| `UserName` | User to run service as (daemons only) |

## Service Lifecycle

```
┌─────────────────────────────────────────────────────────────┐
│  1. App registers service with SMAppService                 │
│  2. Service added to user's login items (if login item)     │
│  3. System launches service at appropriate time             │
│  4. Service runs in background                              │
│  5. Service communicates with main app via XPC              │
│  6. User can enable/disable in System Settings              │
│  7. App can unregister service when uninstalled             │
└─────────────────────────────────────────────────────────────┘
```

## XPC Communication

XPC (Cross-Process Communication) enables:
- Type-safe communication between app and service
- Automatic lifecycle management
- Security through code signing validation
- Crash isolation

### Communication Pattern
1. Main app creates XPC connection
2. Service exports interface via protocol
3. App calls service methods remotely
4. Service responds asynchronously

## Development Workflow

1. **Create Helper App** - Add helper app or service target in Xcode
2. **Bundle Structure** - Place helper in `Contents/Library/LoginItems/`
3. **Configure Info.plist** - Add `SMLoginItemIdentifier`
4. **Register Service** - Use `SMAppService.loginItem().register()`
5. **XPC Protocol** - Define protocol for app-service communication
6. **Implement Logic** - Add service functionality in helper app
7. **Code Sign** - Sign both main app and helper with same certificate
8. **Test Registration** - Verify service appears in System Settings
9. **Test Communication** - Verify XPC calls work correctly
10. **Handle Approval** - Test user enable/disable workflow

## launchctl Commands

Manage services from the command line:

```bash
# List all services
launchctl list

# Load service (agent)
launchctl load ~/Library/LaunchAgents/com.example.myagent.plist

# Unload service
launchctl unload ~/Library/LaunchAgents/com.example.myagent.plist

# Start service
launchctl start com.example.myagent

# Stop service
launchctl stop com.example.myagent

# View service details
launchctl print gui/$(id -u)/com.example.myagent

# Bootstrap service (newer syntax)
launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.example.myagent.plist

# Bootout service (newer syntax)
launchctl bootout gui/$(id -u)/com.example.myagent
```

## Security & Permissions

### Required Elements
- Code signing with valid Developer ID
- Notarization for distribution outside App Store
- User approval for login items (macOS 13+)
- Entitlements for SMAppService registration
- Hardened runtime enabled

### Security Best Practices
1. **Validate XPC connections** - Check code signing of connecting processes
2. **Limit privileges** - Run services with minimum necessary permissions
3. **Sandbox appropriately** - Enable sandbox but allow necessary access
4. **Use TCC carefully** - Request only needed permissions
5. **Handle approval gracefully** - Guide users through enable workflow
6. **Audit logs** - Use `os_log` for security-relevant events
7. **Update securely** - Use auto-update frameworks with code signing
8. **Document behavior** - Be transparent about service purpose

## Debugging

### Logging
```swift
import os.log

let log = OSLog(subsystem: "com.example.service", category: "lifecycle")
os_log("Service started", log: log, type: .info)
```

### Tools
- **Console.app** - View system logs and service output
- **Activity Monitor** - Check if service is running
- **launchctl list** - List loaded services
- **launchctl print** - View service configuration and state
- **System Settings → Login Items** - User's view of services
- **lldb** - Attach debugger to running service
- **Instruments** - Profile service performance

### Log Locations
- Console.app - Unified logging system
- `~/Library/Logs/` - Legacy service logs
- `/var/log/` - System daemon logs (if configured)
- Custom paths via `StandardOutPath`/`StandardErrorPath`

## Testing Strategy

### Unit Testing
- Test XPC protocol implementation
- Mock service responses
- Verify error handling
- Test service logic in isolation

### Integration Testing
- Test service registration/unregistration
- Verify service launches correctly
- Test XPC communication end-to-end
- Verify service restarts after crash
- Test with different macOS versions

### Manual Testing
- Logout/login to test auto-launch
- Toggle in System Settings → Login Items
- Test app installation/uninstallation
- Verify service behavior on system restart
- Test with different user accounts

## Migration from Legacy APIs

### Deprecated (macOS 13+)
- `SMLoginItemSetEnabled()` - Legacy C API
- `LSSharedFileList` - Login items API
- `SMJobBless()` - Privileged helper installation

### Modern Replacement
- `SMAppService.loginItem()` - For login items
- `SMAppService.agent()` - For launch agents
- `SMAppService.daemon()` - For launch daemons
- User control via System Settings
- Better security model with user approval

### Migration Steps
1. Add helper app to `Contents/Library/LoginItems/`
2. Update `Info.plist` with `SMLoginItemIdentifier`
3. Replace legacy registration with `SMAppService`
4. Update UI to guide users to System Settings
5. Test on macOS 13+
6. Optionally maintain backward compatibility for older macOS

## Use Cases

### Menu Bar Applications
- Always-running menu bar utilities
- Quick access to app features
- Background operations with UI

### Background Sync Services
- Cloud sync daemons
- Backup services
- File synchronization

### System Monitoring
- Performance monitoring
- Resource usage tracking
- System logging

### Network Services
- Local web servers
- File sharing services
- Media streaming servers

### Auto-Update Services
- Check for app updates
- Download and install updates
- Version management

### Helper Applications
- Privileged operations for main app
- System-level tasks
- Administrative functions

### Developer Tools
- Build servers
- Code indexing services
- Development environment management

## Best Practices

1. **Use SMAppService** - Prefer modern API on macOS 13+
2. **Provide Uninstall** - Clean up login items when app is removed
3. **Handle Crashes** - Implement proper crash recovery
4. **XPC Error Handling** - Handle connection failures gracefully
5. **Structured Logging** - Use `os_log` with appropriate levels
6. **Lightweight Services** - Keep service CPU and memory usage low
7. **Validate Connections** - Check code signing of XPC connections
8. **User Transparency** - Clearly document service purpose
9. **Version Testing** - Test across supported macOS versions
10. **User Control** - Respect user's enable/disable choices

## Requirements

- **macOS 10.6+** - For launchd (basic functionality)
- **macOS 13+** - For SMAppService (modern API)
- **Code Signing** - Required for distribution
- **Notarization** - Required for apps outside App Store
- **Entitlements** - For system extensions
- **User Approval** - For login items (macOS 13+)
- **Xcode** - For development and debugging

## Limitations

### User Control
- Cannot force-enable login items (user approval required)
- Users can disable services at any time
- No silent background installation

### Sandboxing
- Sandbox restrictions limit service capabilities
- Cannot access arbitrary files without permissions
- Network access may require entitlements

### Debugging
- More complex than regular app debugging
- Service crashes harder to diagnose
- Limited Xcode integration for background services

### Platform Constraints
- Legacy APIs deprecated in macOS 13+
- Launch daemons cannot access UI
- System Extensions require approval workflow

## Example: Simple Login Item

### Info.plist Configuration
```xml
<key>SMLoginItemIdentifier</key>
<string>com.example.MyApp.MyHelper</string>

<key>LSMinimumSystemVersion</key>
<string>13.0</string>
```

### Main App Registration
```swift
import ServiceManagement

class AppDelegate: NSObject, NSApplicationDelegate {
    func applicationDidFinishLaunching(_ notification: Notification) {
        // Get the login item service
        let service = SMAppService.loginItem(
            identifier: "com.example.MyApp.MyHelper"
        )

        // Check current status
        print("Service status: \(service.status)")

        // Register if needed
        if service.status == .notRegistered {
            do {
                try service.register()
                print("Login item registered")
            } catch {
                print("Failed to register: \(error)")
            }
        }
    }
}
```

### Helper App
```swift
import Cocoa

@main
class AppDelegate: NSObject, NSApplicationDelegate {
    func applicationDidFinishLaunching(_ notification: Notification) {
        print("Helper app started")

        // Set up XPC listener
        let listener = NSXPCListener.service()
        listener.delegate = self
        listener.resume()
    }
}
```

## References

- [ServiceManagement Framework](https://developer.apple.com/documentation/servicemanagement)
- [SMAppService](https://developer.apple.com/documentation/servicemanagement/smappservice)
- [Updating Login Items](https://developer.apple.com/documentation/servicemanagement/updating_helper_executables_from_earlier_versions_of_macos)
- [Creating Launch Daemons and Agents](https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/)
- [launchd.info](https://www.launchd.info/) - Community documentation
- [XPC Services](https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/Chapters/CreatingXPCServices.html)

## See Also

- `../../generated/servicemanagement/` - Generated ServiceManagement bindings
- System Extensions framework for extension lifecycle
- XPC framework for inter-process communication
- Foundation framework for service utilities

## Production Considerations

When developing real background services:

1. **User Experience** - Clear communication about service purpose
2. **Performance** - Monitor CPU and memory usage
3. **Reliability** - Handle crashes and restarts gracefully
4. **Security** - Validate all IPC connections
5. **Privacy** - Minimize data collection and storage
6. **Updates** - Plan for service updates without disruption
7. **Diagnostics** - Implement comprehensive logging
8. **Support** - Provide troubleshooting documentation
9. **Compatibility** - Test across macOS versions
10. **Uninstallation** - Clean removal of all service components

## Notes

This example demonstrates framework concepts only. Real background service implementation requires:
- Helper app or service in app bundle structure
- Proper Info.plist configuration with identifiers
- SMAppService registration in main app
- XPC protocol for app-service communication
- Code signing with same certificate for both apps
- User approval workflow for login items
- Testing across macOS versions

The generated bindings in `../../generated/servicemanagement/` provide the Go interface to ServiceManagement, but actual service development is done in Swift/Objective-C with proper app bundle structure.
