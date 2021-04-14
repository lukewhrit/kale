/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	// EmbedColor holds the global color of embeds
	EmbedColor = 0x1dd1a1
)

// Logger is the Logrus instance
var Logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

// CodeifyString function encapsulates the `value` string in backticks to
// indicate code blocks in Discord.
func CodeifyString(value string) string {
	return "`" + value + "`"
}

// IntToString converts provided base-10 integers into strings use fmt.Sprintf
// method
func IntToString(value int) string {
	return fmt.Sprintf("%d", value)
}

// BytesToMegabytes performs a basic math operation to convert byte values in
// megabytes
func BytesToMegabytes(bytes uint64) uint64 {
	return bytes / 1024 / 1024
}

func ParseVersionFile() ([]string, error) {
	versionFile, err := ioutil.ReadFile("./version.txt")

	if err != nil {
		return nil, err
	}

	versionText := string(versionFile)
	versionArray := strings.Split(versionText, ".")

	return versionArray, nil
}

func VersionString() (string, error) {
	version, err := ParseVersionFile()

	return strings.Join(version, "."), err
}
