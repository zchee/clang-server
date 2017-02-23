// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xdgbasedir implements a freedesktop.org XDG Base Directory Specification.
// XDG Base Directory Specification:
//  https://specifications.freedesktop.org/basedir-spec/0.8/
//
// The XDG Base Directory Specification is based on the following concepts:
//
// There is a single base directory relative to which user-specific data files should be written. This directory is defined by the environment variable $XDG_DATA_HOME.
//
// There is a single base directory relative to which user-specific configuration files should be written. This directory is defined by the environment variable $XDG_CONFIG_HOME.
//
// There is a set of preference ordered base directories relative to which data files should be searched. This set of directories is defined by the environment variable $XDG_DATA_DIRS.
//
// There is a set of preference ordered base directories relative to which configuration files should be searched. This set of directories is defined by the environment variable $XDG_CONFIG_DIRS.
//
// There is a single base directory relative to which user-specific non-essential (cached) data should be written. This directory is defined by the environment variable $XDG_CACHE_HOME.
package xdgbasedir

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/zchee/go-xdgbasedir/home"
)

// DataHome return the XDG_DATA_HOME based directory path.
//
// $XDG_DATA_HOME defines the base directory relative to which user specific data files should be stored.
// If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.
func DataHome() string {
	if dataHome := os.Getenv("XDG_DATA_HOME"); dataHome != "" {
		return dataHome
	}
	return filepath.Join(home.Dir(), ".local", "share")
}

// ConfigHome return the XDG_CONFIG_HOME based directory path.
//
// $XDG_CONFIG_HOME defines the base directory relative to which user specific configuration files should be stored.
// If $XDG_CONFIG_HOME is either not set or empty, a default equal to $HOME/.config should be used.
func ConfigHome() string {
	if configHome := os.Getenv("XDG_CONFIG_HOME"); configHome != "" {
		return configHome
	}
	return filepath.Join(home.Dir(), ".config")
}

// DataDirs return the XDG_DATA_DIRS based directory path.
//
// $XDG_DATA_DIRS defines the preference-ordered set of base directories to search for data files in addition
// to the $XDG_DATA_HOME base directory. The directories in $XDG_DATA_DIRS should be seperated with a colon ':'.
// If $XDG_DATA_DIRS is either not set or empty, a value equal to /usr/local/share/:/usr/share/ should be used.
// TODO(zchee): XDG_DATA_DIRS should be seperated with a colon, We should change return type to the []string
// which colon seperated value instead of string?
func DataDirs() string {
	if dataDirs := os.Getenv("XDG_DATA_DIRS"); dataDirs != "" {
		return dataDirs
	}
	return filepath.Join(string(filepath.Separator), "usr", "local", "share", string(filepath.ListSeparator), "usr", "share")
}

// ConfigDirs return the XDG_CONFIG_DIRS based directory path.
//
// $XDG_CONFIG_DIRS defines the preference-ordered set of base directories to search for configuration files in addition
// to the $XDG_CONFIG_HOME base directory. The directories in $XDG_CONFIG_DIRS should be seperated with a colon ':'.
// If $XDG_CONFIG_DIRS is either not set or empty, a value equal to /etc/xdg should be used.
// TODO(zchee): XDG_CONFIG_DIRS should be seperated with a colon, We should change return type to the []string
// which colon seperated value instead of string?
func ConfigDirs() string {
	if configDirs := os.Getenv("XDG_CONFIG_DIRS"); configDirs != "" {
		return configDirs
	}
	return filepath.Join(string(filepath.Separator), "etc", "xdg")
}

// CacheHome return the XDG_CACHE_HOME based directory path.
//
// $XDG_CACHE_HOME defines the base directory relative to which user specific non-essential data files should be stored.
// If $XDG_CACHE_HOME is either not set or empty, a default equal to $HOME/.cache should be used.
//
// TODO(zchee): In macOS, Is it better to use the ~/Library/Caches directory? Or add the configurable by users setting?
// Apple's "File System Programming Guide" describe the this directory should be used if users cache files.
// However, some user who is using the macOS as Unix-like prefers $HOME/.cache.
// xref:
//  https://developer.apple.com/library/content/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1
func CacheHome() string {
	if cacheHome := os.Getenv("XDG_CACHE_HOME"); cacheHome != "" {
		return cacheHome
	}
	return filepath.Join(home.Dir(), ".cache")
}

// RuntimeDir return the XDG_RUNTIME_DIR based directory path.
//
// $XDG_RUNTIME_DIR defines the base directory relative to which user-specific non-essential runtime files and
// other file objects (such as sockets, named pipes, ...) should be stored. The directory MUST be owned by the user,
// and he MUST be the only one having read and write access to it. Its Unix access mode MUST be 0700.
//
// TODO(zchee): XDG_RUNTIME_DIR seems to change depending on the each distro or init system such as systemd.
// Also In macOS, normal user haven't permission for write to this directory.
// xref:
//	http://serverfault.com/questions/388840/good-default-for-xdg-runtime-dir/727994#727994
func RuntimeDir() string {
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		return runtimeDir
	}
	return filepath.Join(string(filepath.Separator), "run", "user", strconv.Itoa(os.Getuid()))
}
