/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"context"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/lukewhrit/kale/internal/pkg/domain"
)

// Ping is a command responding with a "Ping" message in the execution channel.
type Ping struct {
	gommand.CommandBasics
}

func (p *Ping) Init() {
	p.Name = "ping"
	p.Aliases = []string{"pong", "p"}
	p.Description = "Test Kale's connection to Discord by sending a ping."
}

func (p *Ping) CommandFunction(ctx *gommand.Context) error {
	// timestamp, err := message.Timestamp.Parse()

	embed := &disgord.Embed{
		Color:  domain.EmbedColor,
		Fields: make([]*disgord.EmbedField, 0),
		Title:  "🏓 Pinging...",
		Timestamp: disgord.Time{
			Time: time.Now(),
		},
	}

	_, err := ctx.Session.SendMsg(ctx.Message.ChannelID, embed)

	if err != nil {
		return err
	}

	timestamp, err := time.Parse(time.RFC3339, ctx.Message.Timestamp.String())

	if err != nil {
		return err
	}

	embed.Title = "🏓 Pong!"
	embed.Fields = append(embed.Fields, &disgord.EmbedField{
		Name:   "Message Latency",
		Value:  domain.CodeifyString(time.Since(timestamp).Truncate(time.Millisecond).String()),
		Inline: true,
	})

	heartbeatLatency, err := ctx.Session.AvgHeartbeatLatency()

	if err != nil {
		return err
	}

	embed.Fields = append(embed.Fields, &disgord.EmbedField{
		Name:   "Heartbeat Latency",
		Value:  domain.CodeifyString(heartbeatLatency.Truncate(time.Millisecond).String()),
		Inline: true,
	})

	_, err = ctx.Message.Reply(context.Background(), ctx.Session, &disgord.CreateMessageParams{
		Embed: embed,
	})

	return err
}
