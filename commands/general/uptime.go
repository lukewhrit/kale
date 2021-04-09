/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

// Uptime is a command responding with the long form of the bots uptime, already
// available in stats.
type Uptime struct {
}

// GetInvokes returns the command invokes.
func (c *Uptime) GetInvokes() []string {
	return []string{"uptime"}
}

// GetDescription returns the commands description.
func (c *Uptime) GetDescription() string {
	return "retrieve the number of minutes and hours kale has been online."
}

// GetHelp returns the commands help text.
func (c *Uptime) GetHelp() string {
	return "`-uptime`"
}

// GetGroup returns the commands group.
func (c *Uptime) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *Uptime) GetDomainName() string {
	return "xyz.lwhr.kale.general.uptime"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Uptime) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Uptime) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Uptime) Exec(ctx shireikan.Context) error {
	uptime := time.Since(startTime)

	embed := &discordgo.MessageEmbed{
		Color: 0x1dd1a1,
		Title: "⌚ Kale Uptime",
		Description: "Kale has been online for `" +
			uptime.Truncate(time.Second).String() + "`.",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
