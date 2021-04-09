/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"github.com/zekroTJA/shireikan"
)

// Stats retrieves basic general information on Kale and its operations and
// sends them to the user.
type Stats struct {
}

// GetInvokes returns the command invokes.
func (c *Stats) GetInvokes() []string {
	return []string{"stats", "botinfo", "statistics", "botinformation"}
}

// GetDescription returns the commands description.
func (c *Stats) GetDescription() string {
	return "get general information and statistics on kale."
}

// GetHelp returns the commands help text.
func (c *Stats) GetHelp() string {
	return "`-stats`"
}

// GetGroup returns the commands group.
func (c *Stats) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *Stats) GetDomainName() string {
	return "xyz.lwhr.kale.general.stats"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Stats) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Stats) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Stats) Exec(ctx shireikan.Context) error {
	_, err := ctx.Reply("Command not implemented.")

	return err
}
