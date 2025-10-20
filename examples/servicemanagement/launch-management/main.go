package main

import (
	"flag"
	"fmt"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	flag.Parse()

	fmt.Println("ServiceManagement Framework Overview")
	fmt.Println("====================================")

	// Example 1: Framework Overview
	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   ServiceManagement enables:")
	fmt.Println("   - Launch agent and daemon registration")
	fmt.Println("   - Login item management")
	fmt.Println("   - Background service control")
	fmt.Println("   - System Extension registration")
	fmt.Println("   - XPC service discovery")
	fmt.Println("   - App-managed background tasks")

	// Example 2: Service Types
	fmt.Println("\n2. Service Types:")
	serviceTypes := []string{
		"Login Items - Apps/services launched at user login",
		"Launch Agents - User-level background services",
		"Launch Daemons - System-level background services (root)",
		"System Extensions - DriverKit and Network extensions",
		"XPC Services - Inter-process communication services",
		"Helper Tools - Privileged helper applications",
	}
	for i, service := range serviceTypes {
		fmt.Printf("   %d. %s\n", i+1, service)
	}

	// Example 3: Login Items
	fmt.Println("\n3. Login Items:")
	fmt.Println("   Modern login items (macOS 13+):")
	fmt.Println("   - Managed by ServiceManagement framework")
	fmt.Println("   - User-visible in System Settings")
	fmt.Println("   - Enable/disable without app restart")
	fmt.Println("   - Background or UI applications")
	fmt.Println("   ")
	fmt.Println("   Legacy login items:")
	fmt.Println("   - Managed by LaunchServices (deprecated)")
	fmt.Println("   - System Preferences → Users & Groups")
	fmt.Println("   - Being phased out in favor of SMAppService")

	// Example 4: Launch Agents vs Launch Daemons
	fmt.Println("\n4. Launch Agents vs Launch Daemons:")
	fmt.Println("   Launch Agents:")
	fmt.Println("   - Run as logged-in user")
	fmt.Println("   - Access to user's files and preferences")
	fmt.Println("   - Can display UI")
	fmt.Println("   - Location: ~/Library/LaunchAgents/")
	fmt.Println("   - Use cases: Menu bar apps, user services")
	fmt.Println("   ")
	fmt.Println("   Launch Daemons:")
	fmt.Println("   - Run as root or specific user")
	fmt.Println("   - No UI access")
	fmt.Println("   - System-wide services")
	fmt.Println("   - Location: /Library/LaunchDaemons/")
	fmt.Println("   - Use cases: System services, network servers")

	// Example 5: SMAppService
	fmt.Println("\n5. SMAppService (macOS 13+):")
	fmt.Println("   Modern API for managing login items and agents")
	fmt.Println("   ")
	fmt.Println("   Key features:")
	fmt.Println("   - Register/unregister login items")
	fmt.Println("   - Check service status")
	fmt.Println("   - User control from System Settings")
	fmt.Println("   - Automatic lifecycle management")
	fmt.Println("   ")
	fmt.Println("   Service types:")
	fmt.Println("   - LoginItem - Visible apps launched at login")
	fmt.Println("   - Agent - Background services for current user")
	fmt.Println("   - Daemon - System-wide background services")

	// Example 6: Use Cases
	fmt.Println("\n6. Use Cases:")
	useCases := map[string]string{
		"Menu Bar Apps":          "Always-running menu bar utilities",
		"Background Sync":        "Cloud sync, backup services",
		"System Monitoring":      "Performance monitoring, logging",
		"Network Services":       "Local web servers, file sharing",
		"Auto-Update Services":   "Check and install app updates",
		"Helper Applications":    "Privileged operations for main app",
		"Media Servers":          "Music/video streaming services",
		"Communication Tools":    "Messaging, VoIP background services",
		"Developer Tools":        "Build servers, code indexing",
	}
	for useCase, description := range useCases {
		fmt.Printf("   %-25s: %s\n", useCase, description)
	}

	// Example 7: Service Lifecycle
	fmt.Println("\n7. Service Lifecycle:")
	lifecycle := []string{
		"1. App registers service with SMAppService",
		"2. Service added to user's login items (if login item)",
		"3. System launches service at appropriate time",
		"4. Service runs in background",
		"5. Service communicates with main app via XPC",
		"6. User can enable/disable in System Settings",
		"7. App can unregister service when uninstalled",
	}
	for _, step := range lifecycle {
		fmt.Printf("   %s\n", step)
	}

	// Example 8: Registration Process
	fmt.Println("\n8. Service Registration (SMAppService):")
	fmt.Println("   Swift example:")
	fmt.Println("   ```swift")
	fmt.Println("   import ServiceManagement")
	fmt.Println("   ")
	fmt.Println("   // Register login item")
	fmt.Println("   let service = SMAppService.loginItem(identifier: \"com.example.helper\")")
	fmt.Println("   try service.register()")
	fmt.Println("   ")
	fmt.Println("   // Check status")
	fmt.Println("   let status = service.status  // .enabled, .requiresApproval, .notRegistered")
	fmt.Println("   ")
	fmt.Println("   // Unregister")
	fmt.Println("   try service.unregister()")
	fmt.Println("   ```")

	// Example 9: Launchd Property List
	fmt.Println("\n9. Launchd Property List (plist):")
	fmt.Println("   Traditional launchd service configuration")
	fmt.Println("   ")
	fmt.Println("   Example (Launch Agent):")
	fmt.Println("   ```xml")
	fmt.Println("   <?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	fmt.Println("   <!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\"")
	fmt.Println("     \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">")
	fmt.Println("   <plist version=\"1.0\">")
	fmt.Println("   <dict>")
	fmt.Println("       <key>Label</key>")
	fmt.Println("       <string>com.example.myagent</string>")
	fmt.Println("       <key>ProgramArguments</key>")
	fmt.Println("       <array>")
	fmt.Println("           <string>/Applications/MyApp.app/Contents/Library/LoginItems/MyHelper.app/Contents/MacOS/MyHelper</string>")
	fmt.Println("       </array>")
	fmt.Println("       <key>RunAtLoad</key>")
	fmt.Println("       <true/>")
	fmt.Println("   </dict>")
	fmt.Println("   </plist>")
	fmt.Println("   ```")

	// Example 10: plist Keys
	fmt.Println("\n10. Common plist Keys:")
	plistKeys := map[string]string{
		"Label":             "Unique identifier for the service",
		"ProgramArguments":  "Command and arguments to execute",
		"RunAtLoad":         "Start service when loaded",
		"KeepAlive":         "Restart if service exits",
		"StartInterval":     "Run periodically (seconds)",
		"StartCalendarInterval": "Run at specific times",
		"StandardOutPath":   "Redirect stdout to file",
		"StandardErrorPath": "Redirect stderr to file",
		"EnvironmentVariables": "Environment variables for service",
		"WorkingDirectory":  "Working directory for service",
		"UserName":          "User to run service as (daemons only)",
	}
	for key, description := range plistKeys {
		fmt.Printf("   %-25s: %s\n", key, description)
	}

	// Example 11: System Extensions
	fmt.Println("\n11. System Extensions:")
	fmt.Println("   ServiceManagement also handles System Extensions")
	fmt.Println("   ")
	fmt.Println("   Extension types:")
	fmt.Println("   - DriverKit - User-space device drivers")
	fmt.Println("   - Network Extensions - VPN, content filters, DNS proxy")
	fmt.Println("   - Endpoint Security - System monitoring and security")
	fmt.Println("   ")
	fmt.Println("   Activation:")
	fmt.Println("   - OSSystemExtensionRequest.activationRequest()")
	fmt.Println("   - User approval required")
	fmt.Println("   - Managed via System Preferences")

	// Example 12: XPC Services
	fmt.Println("\n12. XPC Service Communication:")
	fmt.Println("   XPC (Cross-Process Communication) enables:")
	fmt.Println("   - Type-safe communication between app and service")
	fmt.Println("   - Automatic lifecycle management")
	fmt.Println("   - Security through code signing validation")
	fmt.Println("   - Crash isolation")
	fmt.Println("   ")
	fmt.Println("   Communication pattern:")
	fmt.Println("   1. Main app creates XPC connection")
	fmt.Println("   2. Service exports interface via protocol")
	fmt.Println("   3. App calls service methods remotely")
	fmt.Println("   4. Service responds asynchronously")

	// Example 13: Security Considerations
	fmt.Println("\n13. Security & Permissions:")
	security := []string{
		"User approval required for login items (macOS 13+)",
		"Code signing required for all services",
		"Hardened runtime for notarized apps",
		"Entitlements for SMAppService registration",
		"Sandboxing affects service capabilities",
		"TCC (Transparency, Consent, Control) permissions",
		"Validate XPC connections with code signing",
		"Limit privilege escalation attack surface",
	}
	for i, item := range security {
		fmt.Printf("   %d. %s\n", i+1, item)
	}

	// Example 14: Development Workflow
	fmt.Println("\n14. Development Workflow:")
	workflow := []string{
		"1. Create helper app or service target in Xcode",
		"2. Add helper to main app bundle (Contents/Library/LoginItems/)",
		"3. Configure Info.plist with SMLoginItemIdentifier",
		"4. Implement SMAppService registration in main app",
		"5. Set up XPC protocol for communication",
		"6. Implement service logic in helper app",
		"7. Code sign both main app and helper",
		"8. Test registration and communication",
		"9. Handle user approval workflow",
		"10. Test enable/disable from System Settings",
	}
	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 15: Debugging
	fmt.Println("\n15. Debugging Strategies:")
	debugging := []string{
		"Console.app - View service logs",
		"launchctl list - Check loaded services",
		"launchctl print - View service details",
		"Activity Monitor - Check if service is running",
		"System Settings → General → Login Items - User view",
		"lldb - Attach debugger to running service",
		"os_log - Structured logging from service",
		"Check ~/Library/Logs/ for service logs",
	}
	for i, strategy := range debugging {
		fmt.Printf("   %d. %s\n", i+1, strategy)
	}

	// Example 16: Common launchctl Commands
	fmt.Println("\n16. Common launchctl Commands:")
	fmt.Println("   Manage services from Terminal:")
	fmt.Println("   ")
	fmt.Println("   List all services:")
	fmt.Println("   $ launchctl list")
	fmt.Println("   ")
	fmt.Println("   Load service:")
	fmt.Println("   $ launchctl load ~/Library/LaunchAgents/com.example.myagent.plist")
	fmt.Println("   ")
	fmt.Println("   Unload service:")
	fmt.Println("   $ launchctl unload ~/Library/LaunchAgents/com.example.myagent.plist")
	fmt.Println("   ")
	fmt.Println("   Start service:")
	fmt.Println("   $ launchctl start com.example.myagent")
	fmt.Println("   ")
	fmt.Println("   Stop service:")
	fmt.Println("   $ launchctl stop com.example.myagent")
	fmt.Println("   ")
	fmt.Println("   View service details:")
	fmt.Println("   $ launchctl print gui/$(id -u)/com.example.myagent")

	// Example 17: Testing
	fmt.Println("\n17. Testing Strategy:")
	fmt.Println("   Unit Testing:")
	fmt.Println("   - Test XPC protocol implementation")
	fmt.Println("   - Mock service responses")
	fmt.Println("   - Verify error handling")
	fmt.Println("   ")
	fmt.Println("   Integration Testing:")
	fmt.Println("   - Test service registration/unregistration")
	fmt.Println("   - Verify service launches correctly")
	fmt.Println("   - Test XPC communication end-to-end")
	fmt.Println("   - Verify service restarts after crash")
	fmt.Println("   ")
	fmt.Println("   Manual Testing:")
	fmt.Println("   - Logout/login to test auto-launch")
	fmt.Println("   - Toggle in System Settings")
	fmt.Println("   - Test app uninstallation cleanup")

	// Example 18: Migration from LSSharedFileList
	fmt.Println("\n18. Migration from Legacy APIs:")
	fmt.Println("   Deprecated (macOS 13+):")
	fmt.Println("   - SMLoginItemSetEnabled() - Legacy C API")
	fmt.Println("   - LSSharedFileList - Login items API")
	fmt.Println("   ")
	fmt.Println("   Modern replacement:")
	fmt.Println("   - SMAppService.loginItem() - Swift/Objective-C")
	fmt.Println("   - User control via System Settings")
	fmt.Println("   - Better security model")
	fmt.Println("   ")
	fmt.Println("   Migration steps:")
	fmt.Println("   1. Add helper app to Contents/Library/LoginItems/")
	fmt.Println("   2. Update Info.plist with SMLoginItemIdentifier")
	fmt.Println("   3. Replace legacy registration code with SMAppService")
	fmt.Println("   4. Test on macOS 13+")
	fmt.Println("   5. Maintain backward compatibility if needed")

	// Example 19: Best Practices
	fmt.Println("\n19. Best Practices:")
	bestPractices := []string{
		"Use SMAppService for new development (macOS 13+)",
		"Provide uninstall method to remove login items",
		"Handle service crashes gracefully",
		"Implement proper XPC error handling",
		"Use os_log for structured logging",
		"Keep service lightweight and efficient",
		"Validate all XPC connections",
		"Document service purpose for users",
		"Test across macOS versions",
		"Provide user control over service behavior",
	}
	for i, practice := range bestPractices {
		fmt.Printf("   %d. %s\n", i+1, practice)
	}

	// Example 20: Requirements
	fmt.Println("\n20. Requirements:")
	requirements := []string{
		"macOS 10.6+ for launchd (basic)",
		"macOS 13+ for SMAppService (modern API)",
		"Code signing required for distribution",
		"Notarization for apps outside App Store",
		"Entitlements for system extensions",
		"User approval for login items (macOS 13+)",
		"Xcode for development",
	}
	for i, req := range requirements {
		fmt.Printf("   %d. %s\n", i+1, req)
	}

	// Example 21: Limitations
	fmt.Println("\n21. Framework Limitations:")
	limitations := []string{
		"User approval required (cannot force enable)",
		"Sandboxing limits service capabilities",
		"Cannot access UI from launch daemons",
		"System Extensions require user approval",
		"Legacy APIs deprecated in macOS 13+",
		"Limited control over service lifecycle",
		"Debugging more complex than regular apps",
	}
	for i, limitation := range limitations {
		fmt.Printf("   %d. %s\n", i+1, limitation)
	}

	fmt.Println("\n✓ ServiceManagement framework overview completed!")
	fmt.Println("\nNote: This is an educational overview only.")
	fmt.Println("  ServiceManagement requires:")
	fmt.Println("  - Helper app or service in app bundle")
	fmt.Println("  - Proper Info.plist configuration")
	fmt.Println("  - Code signing and notarization")
	fmt.Println("  - XPC protocol for communication")
	fmt.Println("  - User approval for login items (macOS 13+)")
	fmt.Println("\nFor production service development:")
	fmt.Println("  - Use SMAppService on macOS 13+")
	fmt.Println("  - Implement proper XPC communication")
	fmt.Println("  - Handle service crashes and restarts")
	fmt.Println("  - Provide user control and transparency")
	fmt.Println("  - Test across macOS versions")
	fmt.Println("  - Document service behavior clearly")
	fmt.Println("\nReferences:")
	fmt.Println("  - https://developer.apple.com/documentation/servicemanagement")
	fmt.Println("  - https://developer.apple.com/documentation/servicemanagement/smappservice")
	fmt.Println("  - https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/")
	fmt.Println("  - https://www.launchd.info/ (launchd documentation)")
}
