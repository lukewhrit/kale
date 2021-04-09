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

// Ping is a command responding with a "Ping" message in the execution channel.
type Ping struct {
}

// GetInvokes returns the command invokes.
func (c *Ping) GetInvokes() []string {
	return []string{"ping", "p"}
}

// GetDescription returns the commands description.
func (c *Ping) GetDescription() string {
	return "test kale's connection to Discord by sending a ping."
}

// GetHelp returns the commands help text.
func (c *Ping) GetHelp() string {
	return "`-ping`"
}

// GetGroup returns the commands group.
func (c *Ping) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *Ping) GetDomainName() string {
	return "xyz.lwhr.kale.general.ping"
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
	embed := &discordgo.MessageEmbed{
		Color:     0x1dd1a1,
		Fields:    make([]*discordgo.MessageEmbedField, 0),
		Title:     "🏓 Pinging...",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	message, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	if err != nil {
		return err
	}

	timestamp, err := message.Timestamp.Parse()

	if err != nil {
		return err
	}

	embed.Title = "🏓 Pong!"
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Message Latency",
		Value:  "`" + time.Since(timestamp).Truncate(time.Millisecond).String() + "`",
		Inline: true,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Heartbeat Latency",
		Value:  "`" + ctx.GetSession().HeartbeatLatency().Truncate(time.Millisecond).String() + "`",
		Inline: true,
	})

	_, err = ctx.GetSession().ChannelMessageEditEmbed(ctx.GetChannel().ID, message.ID, embed)

	return err
}
