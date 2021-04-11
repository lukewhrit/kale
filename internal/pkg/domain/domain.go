/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package domain

const (
	// EmbedColor holds the global color of embeds
	EmbedColor = 0x1dd1a1
)

// CodeifyString function encapsulates the `value` string in backticks to
// indicate code blocks in Discord.
func CodeifyString(value string) string {
	return "`" + value + "`"
}
