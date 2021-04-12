/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lukewhrit/kale/internal/pkg/domain"
	"github.com/zekroTJA/shireikan"
)

// ServerInfo retrieves and outputs general information on the current guild
type ServerInfo struct {
}

// GetInvokes returns the command invokes.
func (c *ServerInfo) GetInvokes() []string {
	return []string{"serverinfo", "sinfo", "ginfo", "guildinfo", "si"}
}

// GetDescription returns the commands description.
func (c *ServerInfo) GetDescription() string {
	return "get general information on the current guild."
}

// GetHelp returns the commands help text.
func (c *ServerInfo) GetHelp() string {
	return "`-serverinfo`"
}

// GetGroup returns the commands group.
func (c *ServerInfo) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *ServerInfo) GetDomainName() string {
	return "xyz.lwhr.kale.general.serverinfo"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *ServerInfo) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *ServerInfo) IsExecutableInDMChannels() bool {
	return false
}

// Exec is the commands execution handler.
func (c *ServerInfo) Exec(ctx shireikan.Context) error {
	var (
		mfaLevel           bool
		verificationLevel  string
		voiceChannelsCount []*discordgo.Channel
		textChannelsCount  []*discordgo.Channel
	)

	guild := ctx.GetGuild()

	switch guild.MfaLevel {
	case discordgo.MfaLevelNone:
		mfaLevel = false
	case discordgo.MfaLevelElevated:
		mfaLevel = true
	}

	switch guild.VerificationLevel {
	case discordgo.VerificationLevelNone:
		verificationLevel = "None"
	case discordgo.VerificationLevelLow:
		verificationLevel = "Low"
	case discordgo.VerificationLevelMedium:
		verificationLevel = "Medium"
	case discordgo.VerificationLevelHigh:
		verificationLevel = "High"
	case discordgo.VerificationLevelVeryHigh:
		verificationLevel = "Highest"
	}

	for _, channel := range guild.Channels {
		switch channel.Type {
		case discordgo.ChannelTypeGuildText:
			textChannelsCount = append(textChannelsCount, channel)
		case discordgo.ChannelTypeGuildVoice:
			voiceChannelsCount = append(voiceChannelsCount, channel)
		}
	}

	creationDate, err := discordgo.SnowflakeTimestamp(guild.ID)

	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		Color:       domain.EmbedColor,
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       fmt.Sprintf("📜 Info on %s", guild.Name),
		Description: guild.Description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Owner",
				Value:  fmt.Sprintf("<@%s>", guild.OwnerID),
				Inline: true,
			},
			{
				Name:   "Guild ID",
				Value:  guild.ID,
				Inline: true,
			},
			{
				Name:   "Total Members",
				Value:  fmt.Sprintf("%d/%d", guild.MemberCount, guild.MaxMembers),
				Inline: true,
			},
			{
				Name:   "Voice Region",
				Value:  domain.CodeifyString(ctx.GetGuild().Region),
				Inline: true,
			},
			{
				Name:   "Bot Prefix",
				Value:  domain.CodeifyString("-"),
				Inline: true,
			},
			{
				Name:   "Created On",
				Value:  creationDate.Format(time.RFC822),
				Inline: true,
			},
			{
				Name:   "2FA Enabled",
				Value:  strings.Title(strconv.FormatBool(mfaLevel)),
				Inline: true,
			},
			{
				Name:   "Verification Level",
				Value:  verificationLevel,
				Inline: true,
			},
			{
				Name: "Channels",
				Value: fmt.Sprintf("%d Text Channels\n %d Voice Channels",
					len(textChannelsCount), len(voiceChannelsCount)),
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(),
		},
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
