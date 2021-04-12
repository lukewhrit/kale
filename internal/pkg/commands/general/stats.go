/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lukewhrit/kale/internal/pkg/domain"
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
	client, err := ctx.GetSession().User("@me")
	var (
		userCount int
		m         runtime.MemStats
	)

	runtime.ReadMemStats(&m)

	for _, guild := range ctx.GetSession().State.Guilds {
		userCount += guild.MemberCount
	}

	if err != nil {
		return err
	}

	version, err := domain.VersionString()

	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		Color:     domain.EmbedColor,
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "🥬 Kale Statistics",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Kale Version",
				Value:  "v" + version,
				Inline: true,
			},
			{
				Name:   "Golang Version",
				Value:  domain.CodeifyString(runtime.Version()),
				Inline: true,
			},
			{
				Name:   "Application ID",
				Value:  client.ID,
				Inline: true,
			},
			{
				Name:   "discordgo Version",
				Value:  "v0.0.0",
				Inline: true,
			},
			{
				Name:   "shireikan Version",
				Value:  "v0.0.0",
				Inline: true,
			},
			{
				Name:   "Server Count",
				Value:  domain.IntToString(len(ctx.GetSession().State.Guilds)),
				Inline: true,
			},
			{
				Name:   "User Count",
				Value:  domain.IntToString(userCount),
				Inline: true,
			},
			{
				Name:   "Contributors and Thanks",
				Value:  "<@189850839660101632> - Main Developer\n<@233064645571772419> - Helped a bunch",
				Inline: true,
			},
			{
				Name:   "Application Uptime",
				Value:  domain.CodeifyString(time.Since(startTime).Truncate(time.Second).String()),
				Inline: true,
			},
			{
				Name: "System Information",
				Value: fmt.Sprintf("```Memory Usage: %v MiB\nCPU Usage: %s\nOperating System: %s```",
					domain.BytesToMegabytes(m.TotalAlloc),
					"...", strings.Title(runtime.GOOS)),
				Inline: false,
			},
		},
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
