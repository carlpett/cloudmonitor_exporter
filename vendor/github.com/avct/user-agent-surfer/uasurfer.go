// Package uasurfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The philosophy is to identify
// technologies that holds >1% market share, and to avoid
// expending resources and accuracy on guessing at esoteric UA
// strings.
package uasurfer

import "strings"

//go:generate stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go

// DeviceType (int) returns a constant.
type DeviceType int

// A complete list of supported devices in the
// form of constants.
const (
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceWearable
	DeviceTV
)

// BrowserName (int) returns a constant.
type BrowserName int

// A complete list of supported web browsers in the
// form of constants.
const (
	BrowserUnknown BrowserName = iota
	BrowserChrome
	BrowserIE
	BrowserSafari
	BrowserFirefox
	BrowserAndroid
	BrowserOpera
	BrowserBlackberry
	BrowserUCBrowser
	BrowserSilk
	BrowserNokia
	BrowserNetFront
	BrowserQQ
	BrowserMaxthon
	BrowserSogouExplorer
	BrowserSpotify
	BrowserBot
)

// OSName (int) returns a constant.
type OSName int

// A complete list of supported OSes in the
// form of constants. For handling particular versions
// of operating systems (e.g. Windows 2000), see
// the README.md file.
const (
	OSUnknown OSName = iota
	OSWindowsPhone
	OSWindows
	OSMacOSX
	OSiOS
	OSAndroid
	OSBlackberry
	OSChromeOS
	OSKindle
	OSWebOS
	OSLinux
	OSPlaystation
	OSXbox
	OSNintendo
	OSBot
)

// Platform (int) returns a constant.
type Platform int

// A complete list of supported platforms in the
// form of constants. Many OSes report their
// true platform, such as Android OS being Linux
// platform.
const (
	PlatformUnknown Platform = iota
	PlatformWindows
	PlatformMac
	PlatformLinux
	PlatformiPad
	PlatformiPhone
	PlatformiPod
	PlatformBlackberry
	PlatformWindowsPhone
	PlatformPlaystation
	PlatformXbox
	PlatformNintendo
	PlatformBot
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) Less(c Version) bool {
	if v.Major < c.Major {
		return true
	}

	if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	}

	if v.Minor > c.Minor {
		return false
	}

	return v.Patch < c.Patch
}

type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType DeviceType
}

type Browser struct {
	Name    BrowserName
	Version Version
}

type OS struct {
	Platform Platform
	Name     OSName
	Version  Version
}

// Parse accepts a raw user agent (string) and returns the UserAgent.
func Parse(ua string) *UserAgent {
	ua = strings.ToLower(ua)
	resp := &UserAgent{}

	switch {
	case len(ua) == 0:
		resp.OS.Platform = PlatformUnknown
		resp.OS.Name = OSUnknown
		resp.Browser.Name = BrowserUnknown
		resp.DeviceType = DeviceUnknown

	// stop on on first case returning true
	case resp.evalOS(ua):
	case resp.evalBrowserName(ua):
	default:
		resp.evalBrowserVersion(ua)
		resp.evalDevice(ua)
	}

	return resp
}
