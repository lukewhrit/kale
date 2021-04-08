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

type dogResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Dog gets a unique image of a dog and sends a embed with the image when
// invoked.
type Dog struct {
}

// GetInvokes returns the command invokes.
func (c *Dog) GetInvokes() []string {
	return []string{"dog", "doggy", "puppy", "pup", "hound"}
}

// GetDescription returns the commands description.
func (c *Dog) GetDescription() string {
	return "get a unique image of a dog."
}

// GetHelp returns the commands help text.
func (c *Dog) GetHelp() string {
	return "`-dog`"
}

// GetGroup returns the commands group.
func (c *Dog) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Dog) GetDomainName() string {
	return "xyz.lwhr.kale.fun.dog"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Dog) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Dog) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Dog) Exec(ctx shireikan.Context) error {
	var dog dogResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://dog.ceo/api/breeds/image/random")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&dog)

	embed := &discordgo.MessageEmbed{
		Color: 0x1dd1a1,
		Title: "🐶 Woof",
		Image: &discordgo.MessageEmbedImage{
			URL: dog.Message,
		},
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
