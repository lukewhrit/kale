/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lukewhrit/kale/commands/general"
	"github.com/zekroTJA/shireikan"
)

// Start begins the program and connects to Discord
func Start() error {
	// Create a new Discord session using the token
	dg, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		return err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	dg.AddHandler(onReady)

	if err = dg.Open(); err != nil {
		return err
	}

	// Wait here until CTRL-C or another term is signal is received
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		// Cleanly close the Discord session
		dg.Close()
	}()

	handler := shireikan.NewHandler(&shireikan.Config{
		GeneralPrefix:         "-",
		AllowBots:             false,
		AllowDM:               true,
		ExecuteOnEdit:         false,
		InvokeToLower:         true,
		UseDefaultHelpCommand: true,
		OnError: func(ctx shireikan.Context, typ shireikan.ErrorType, err error) {
			log.Fatalf("[ERR] [%d] %s", typ, err.Error())
		},
	})

	// Register Commands from General Group
	handler.RegisterCommand(&general.Ping{})

	handler.RegisterHandlers(dg)

	return nil
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Bot is now running. Press Ctrl-C to exit.")

	if err := s.UpdateGameStatus(0, "-help"); err != nil {
		log.Fatalf("[ERR] %s", err.Error())
	}
}
