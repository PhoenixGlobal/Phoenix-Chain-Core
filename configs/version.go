package configs

import (
	"fmt"
)

const (
	//These versions are meaning the current code version.
	VersionMajor = 1          // Major version component of the current release
	VersionMinor = 1          // Minor version component of the current release
	VersionPatch = 1          // Patch version component of the current release
	VersionMeta  = "unstable" // Version metadata to append to the version string

	//CAUTION: DO NOT MODIFY THIS ONCE THE CHAIN HAS BEEN INITIALIZED!!!
	GenesisVersion = uint32(1<<16 | 0<<8 | 0)
)

func CodeVersion() uint32 {
	return uint32(VersionMajor<<16 | VersionMinor<<8 | VersionPatch)
}

func LtMinorVersion(version uint32) bool {
	localVersion := CodeVersion() >> 8
	version = version >> 8
	return localVersion < version
}

// Version holds the textual version string.
var Version = func() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}()

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := Version
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()

func FormatVersion(version uint32) string {
	if version == 0 {
		return "0.0.0"
	}
	major := version << 8
	major = major >> 24

	minor := version << 16
	minor = minor >> 24

	patch := version << 24
	patch = patch >> 24

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// ArchiveVersion holds the textual version string used for PhoenixChain archives.
// e.g. "1.8.11-dea1ce05" for stable releases, or
//      "1.8.13-unstable-21c059b6" for unstable releases
func ArchiveVersion(gitCommit string) string {
	vsn := Version
	if VersionMeta != "stable" {
		vsn += "-" + VersionMeta
	}
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}

func VersionWithCommit(gitCommit, gitDate string) string {
	vsn := VersionWithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	if (VersionMeta != "stable") && (gitDate != "") {
		vsn += "-" + gitDate
	}
	return vsn
}

type ProgramVersion struct {
	Version uint32
	Sign    string
}
