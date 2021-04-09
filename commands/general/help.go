/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package general

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shireikan"
)

// Help will either display a list of all of Kale's commands or display help on
// a specific command
type Help struct {
}

// GetInvokes returns the command invokes.
func (c *Help) GetInvokes() []string {
	return []string{"help", "h", "?", "man", "info", "commands", "information"}
}

// GetDescription returns the commands description.
func (c *Help) GetDescription() string {
	return "display list of bot commands or get help on a specific command"
}

// GetHelp returns the commands help text.
func (c *Help) GetHelp() string {
	return "`help [command]`"
}

// GetGroup returns the commands group.
func (c *Help) GetGroup() string {
	return shireikan.GroupGeneral
}

// GetDomainName returns the commands domain name.
func (c *Help) GetDomainName() string {
	return "xyz.lwhr.kale.general.help"
}

// GetSubPermissionRules returns the commands sub permissions array.
func (c *Help) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

// IsExecutableInDMChannels returns whether the command is executable in DM
// channels.
func (c *Help) IsExecutableInDMChannels() bool {
	return true
}

// Exec is the commands execution handler.
func (c *Help) Exec(ctx shireikan.Context) error {
	emb := &discordgo.MessageEmbed{
		Color:  0x1dd1a1,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	handler, _ := ctx.GetObject("cmdhandler").(shireikan.Handler)

	if len(ctx.GetArgs()) == 0 {
		// Respond with list of all bot commands

		commands := make(map[string][]shireikan.Command)

		for _, c := range handler.GetCommandInstances() {
			group := strings.Title(strings.ToLower(c.GetGroup()))

			if _, ok := commands[group]; !ok {
				commands[group] = make([]shireikan.Command, 0)
			}

			commands[group] = append(commands[group], c)
		}

		emb.Title = "🥬 Kale Command List"

		for category, categoryCommands := range commands {
			commandHelpLines := ""

			for _, c := range categoryCommands {
				commandHelpLines += fmt.Sprintf("`%s` - *%s*\n", c.GetInvokes()[0], c.GetDescription())
			}

			emb.Fields = append(emb.Fields, &discordgo.MessageEmbedField{
				Name:  category,
				Value: commandHelpLines,
			})
		}
	} else {
		// Respond with help on specific command

		invoke := ctx.GetArgs().Get(0).AsString()
		command, ok := handler.GetCommand(invoke)

		if !ok {
			_, err := ctx.ReplyEmbedError(
				fmt.Sprintf("No command was found with the invoke `%s`.", invoke), "Error")
			return err
		}

		emb.Title = "Command Description"

		description := command.GetDescription()

		if description == "" {
			description = "`no description`"
		}

		help := command.GetHelp()

		if help == "" {
			help = "`no usage information`"
		}

		emb.Fields = []*discordgo.MessageEmbedField{
			{
				Name:   "Invocable by",
				Value:  "`" + strings.Join(command.GetInvokes(), ", ") + "`",
				Inline: true,
			},
			{
				Name:   "Group",
				Value:  strings.Title(strings.ToLower(command.GetGroup())),
				Inline: true,
			},
			{
				Name:  "Description",
				Value: command.GetDescription(),
			},
			{
				Name:  "Usage",
				Value: help,
			},
		}

		if spr := command.GetSubPermissionRules(); spr != nil {
			txt := "*`[E]` in front of permissions means `Explicit`, which means that this " +
				"permission must be explicitly allowed and can not be wild-carded.\n" +
				"`[D]` implies that wildcards will apply to this sub permission.*\n\n"

			for _, rule := range spr {
				explicit := "D"

				if rule.Explicit {
					explicit = "E"
				}

				txt = fmt.Sprintf("%s`[%s]` %s - *%s*\n",
					txt, explicit, getTermAssembly(command, rule.Term),
					rule.Description)
			}

			emb.Fields = append(emb.Fields, &discordgo.MessageEmbedField{
				Name:  "Sub Permission Rules",
				Value: txt,
			})
		}
	}

	_, err := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, emb)

	return err
}

func getTermAssembly(cmd shireikan.Command, term string) string {
	if strings.HasPrefix(term, "/") {
		return term[1:]
	}

	return cmd.GetDomainName() + "." + term
}
