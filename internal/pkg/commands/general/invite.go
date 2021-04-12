/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lukewhrit/kale/internal/pkg/domain"
	"github.com/zekroTJA/shireikan"
)

// Invite is a command responding with a link to invite the bot.
type Invite struct {
}

// GetInvokes returns the command invokes.
func (c *Invite) GetInvokes() []string {
	return []string{"invite", "i", "add", "addbot", "links", "support"}
}

// GetDescription returns the commands description.
func (c *Invite) GetDescription() string {
	return "test kale's connection to Discord by sending a ping."
}

// GetHelp returns the commands help text.
func (c *Invite) GetHelp() string {
	return "`-invite`"
}

// GetGroup returns the commands group.
func (c *Invite) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *Invite) GetDomainName() string {
	return "xyz.lwhr.kale.general.invite"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Invite) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Invite) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Invite) Exec(ctx shireikan.Context) error {
	inviteLink := "https://discord.com/api/oauth2/authorize?client_id=763845128497922079&permissions=939904086&scope=bot"

	embed := &discordgo.MessageEmbed{
		Color: domain.EmbedColor,
		Title: "🔗 Invite Kale",
		Description: fmt.Sprintf(
			"Want Kale in your server? [Just click this to invite me](%s).",
			inviteLink),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
