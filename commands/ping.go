/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package commands

import (
	"github.com/zekroTJA/shireikan"
)

// Ping is a command responding with a Ping
// message in the commands channel.
type Ping struct {
}

// GetInvokes returns the command invokes.
func (c *Ping) GetInvokes() []string {
	return []string{"ping", "p"}
}

// GetDescription returns the commands description.
func (c *Ping) GetDescription() string {
	return "ping pong"
}

// GetHelp returns the commands help text.
func (c *Ping) GetHelp() string {
	return "`ping` - Ping"
}

// GetGroup returns the commands group.
func (c *Ping) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Ping) GetDomainName() string {
	return "kale.fun.ping"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Ping) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Ping) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Ping) Exec(ctx shireikan.Context) error {
	_, err := ctx.Reply("Ping! :ping_pong:")

	return err
}
