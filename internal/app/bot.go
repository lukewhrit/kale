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

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Start begins the program and connects to Discord
func Start() error {
	// Create a new Discord session using the token
	dg, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		return err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open websocket connection to Discord and begin listening
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

	log.Println("Bot is now running. Press CTRL-C to exit.")

	router := dgc.Create(&dgc.Router{
		Prefixes:    []string{"-"},
		BotsAllowed: false,
		Commands:    []*dgc.Command{},
		Middlewares: []dgc.Middleware{},
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong!")
		},
	})

	router.RegisterDefaultHelpCommand(dg, nil)
	loadAllCommands(router)
	router.Initialize(dg)

	return nil
}
