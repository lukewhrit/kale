/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fun

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

type birdResponse struct {
	Link string `json:"link"`
}

// Bird gets a unique image of a bird and sends a embed with the image when
// invoked.
type Bird struct {
}

// GetInvokes returns the command invokes.
func (c *Bird) GetInvokes() []string {
	return []string{"bird"}
}

// GetDescription returns the commands description.
func (c *Bird) GetDescription() string {
	return "get a unique image of a bird."
}

// GetHelp returns the commands help text.
func (c *Bird) GetHelp() string {
	return "`-bird`"
}

// GetGroup returns the commands group.
func (c *Bird) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Bird) GetDomainName() string {
	return "xyz.lwhr.kale.fun.bird"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Bird) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Bird) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Bird) Exec(ctx shireikan.Context) error {
	var bird birdResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://some-random-api.ml/img/birb")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&bird)

	embed := &discordgo.MessageEmbed{
		Color: 0x1dd1a1,
		Title: "🐦 Birds, birds, birds.",
		Image: &discordgo.MessageEmbedImage{
			URL: bird.Link,
		},
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
