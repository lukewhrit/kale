/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package commands

import (
	"github.com/zekroTJA/shireikan"
)

// Pong is a command responding with a Pong
// message in the commands channel.
type Pong struct {
}

// GetInvokes returns the command invokes.
func (c *Pong) GetInvokes() []string {
	return []string{"pong"}
}

// GetDescription returns the commands description.
func (c *Pong) GetDescription() string {
	return "pong ping"
}

// GetHelp returns the commands help text.
func (c *Pong) GetHelp() string {
	return "`pong` - pong"
}

// GetGroup returns the commands group.
func (c *Pong) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Pong) GetDomainName() string {
	return "kale.fun.pong"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Pong) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Pong) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Pong) Exec(ctx shireikan.Context) error {
	_, err := ctx.Reply("Pong! :ping_pong:")

	return err
}
