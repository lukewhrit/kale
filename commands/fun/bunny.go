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

type bunnyResponse struct {
	Id          string `json:"id"`
	Source      string `json:"source"`
	ThisServed  int    `json:"thisServed"`
	TotalServed int    `json:"totalServed"`
	Media       struct {
		Gif    string `json:"gif"`
		Poster string `json:"poster"`
	} `json:"media"`
}

// Bunny gets a unique image of a bunny and sends a embed with the image when
// invoked.
type Bunny struct {
}

// GetInvokes returns the command invokes.
func (c *Bunny) GetInvokes() []string {
	return []string{"bunny", "rabit", "hare"}
}

// GetDescription returns the commands description.
func (c *Bunny) GetDescription() string {
	return "get a unique image of a bunny."
}

// GetHelp returns the commands help text.
func (c *Bunny) GetHelp() string {
	return "`-bunny`"
}

// GetGroup returns the commands group.
func (c *Bunny) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Bunny) GetDomainName() string {
	return "xyz.lwhr.kale.fun.bunny"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Bunny) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Bunny) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Bunny) Exec(ctx shireikan.Context) error {
	var bunny bunnyResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://api.bunnies.io/v2/loop/random/?media=gif,png")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&bunny)

	embed := &discordgo.MessageEmbed{
		Color: 0x1dd1a1,
		Title: "🐇 Adorable little creatures",
		Image: &discordgo.MessageEmbedImage{
			URL: bunny.Media.Gif,
		},
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
