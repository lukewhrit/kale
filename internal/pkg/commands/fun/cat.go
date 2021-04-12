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
	"github.com/lukewhrit/kale/internal/pkg/domain"
	"github.com/zekroTJA/shireikan"
)

type catResponse struct {
	File string `json:"file"`
}

// Cat gets a unique image of a cat and sends a embed with the image when
// invoked.
type Cat struct {
}

// GetInvokes returns the command invokes.
func (c *Cat) GetInvokes() []string {
	return []string{"cat", "kitty"}
}

// GetDescription returns the commands description.
func (c *Cat) GetDescription() string {
	return "get a unique image of a cat."
}

// GetHelp returns the commands help text.
func (c *Cat) GetHelp() string {
	return "`-cat`"
}

// GetGroup returns the commands group.
func (c *Cat) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Cat) GetDomainName() string {
	return "xyz.lwhr.kale.fun.cat"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Cat) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Cat) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Cat) Exec(ctx shireikan.Context) error {
	var cat catResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://aws.random.cat/meow")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&cat)

	embed := &discordgo.MessageEmbed{
		Color: domain.EmbedColor,
		Title: "🐱 Meow",
		Image: &discordgo.MessageEmbedImage{
			URL: cat.File,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
