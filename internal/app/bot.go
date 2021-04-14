/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package app

import (
	"context"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/lukewhrit/kale/internal/pkg/commands/general"
	"github.com/lukewhrit/kale/internal/pkg/domain"
)

var (
	noCtx  = context.Background()
	router = gommand.NewRouter(&gommand.RouterConfig{
		PrefixCheck: gommand.MultiplePrefixCheckers(
			gommand.StaticPrefix("-"),
			gommand.MentionPrefix,
		),
	})
)

// Start begins the program and connects to Discord
func Start() error {
	client := disgord.New(disgord.Config{
		ProjectName: "Kale",
		BotToken:    Config.Token,
		Logger:      domain.Logger,
		RejectEvents: []string{
			disgord.EvtTypingStart,
			disgord.EvtPresenceUpdate,
		},
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: Config.Prefix + "help",
			},
		},
	})

	// Start the client
	defer client.Gateway().StayConnectedUntilInterrupted()

	// Register commands
	router.SetCommand(&general.Ping{})

	// Connect the gommand router to disgord
	router.Hook(client)

	// Log message when bot is connected to Discord
	client.Gateway().BotReady(func() {
		domain.Logger.Info("Bot is now running. Press Ctrl-C to exit.")
	})

	return nil
}
