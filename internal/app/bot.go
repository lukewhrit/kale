/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lukewhrit/kale/internal/pkg/commands/fun"
	"github.com/lukewhrit/kale/internal/pkg/commands/general"
	"github.com/lukewhrit/kale/internal/pkg/domain"
	"github.com/zekroTJA/shireikan"
)

// Start begins the program and connects to Discord
func Start() error {
	// Create a new Discord session using the token
	dg, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		return err
	}

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers

	dg.AddHandler(onReady)

	if err = dg.Open(); err != nil {
		return err
	}

	// Wait here until CTRL-C or another term signal is received
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		// Cleanly close the Discord session
		dg.Close()
	}()

	handler := shireikan.NewHandler(&shireikan.Config{
		GeneralPrefix:         Config.Prefix,
		AllowBots:             false,
		AllowDM:               true,
		ExecuteOnEdit:         false,
		InvokeToLower:         true,
		UseDefaultHelpCommand: false,
		OnError: func(ctx shireikan.Context, typ shireikan.ErrorType, err error) {
			// Write the error to the log
			log.Printf("[ERR] [%d] %s", typ, err.Error())

			// Send an embed to the channel where the error ocurred
			embed := &discordgo.MessageEmbed{
				Color:       domain.EmbedColor,
				Title:       "🔴 An error occurred while executing this command",
				Description: fmt.Sprintf("```%s```\n You can report this error and get support for it on [Github](https://github.com/lukewhrit/kale/issues).", err.Error()),
				Timestamp:   time.Now().Format(time.RFC3339),
			}

			_, msgError := ctx.GetSession().ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)

			if msgError != nil {
				log.Printf("[ERR] [%d] %s", typ, msgError.Error())
			}
		},
		GuildPrefixGetter: func(guildID string) (string, error) {
			return "!", nil
		},
	})

	// Register Commands from General Group
	handler.RegisterCommand(&general.Ping{})
	handler.RegisterCommand(&general.Help{})
	handler.RegisterCommand(&general.Invite{})
	handler.RegisterCommand(&general.Stats{})
	handler.RegisterCommand(&general.Uptime{})
	handler.RegisterCommand(&general.ServerInfo{})

	// Register commands from Fun group
	handler.RegisterCommand(&fun.Bird{})
	handler.RegisterCommand(&fun.Bunny{})
	handler.RegisterCommand(&fun.Cat{})
	handler.RegisterCommand(&fun.Dog{})
	handler.RegisterCommand(&fun.Duck{})

	handler.RegisterHandlers(dg)

	return nil
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Bot is now running. Press Ctrl-C to exit.")

	if err := s.UpdateGameStatus(0, "-help"); err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}
}
