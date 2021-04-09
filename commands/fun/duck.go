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

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

type duckResponse struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

// Duck gets a unique image of a duck and sends a embed with the image when
// invoked.
type Duck struct {
}

// GetInvokes returns the command invokes.
func (c *Duck) GetInvokes() []string {
	return []string{"duck", "quack"}
}

// GetDescription returns the commands description.
func (c *Duck) GetDescription() string {
	return "get a unique image of a duck."
}

// GetHelp returns the commands help text.
func (c *Duck) GetHelp() string {
	return "`-duck`"
}

// GetGroup returns the commands group.
func (c *Duck) GetGroup() string {
	return shireikan.GroupFun
}

// GetDomainName returns the commands domain name.
func (c *Duck) GetDomainName() string {
	return "xyz.lwhr.kale.fun.duck"
}

// GetSubPermissionRules returns the commands sub
// permissions array.
func (c *Duck) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether
// the command is executable in DM channels.
func (c *Duck) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Duck) Exec(ctx shireikan.Context) error {
	var duck duckResponse

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://random-d.uk/api/v2/random")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&duck)

	embed := &discordgo.MessageEmbed{
		Color: 0x1dd1a1,
		Title: "🦆 Quack!",
		Image: &discordgo.MessageEmbedImage{
			URL: duck.URL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

	return err
}
