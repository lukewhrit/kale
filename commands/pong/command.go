/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package pong

import (
	"time"

	"github.com/Lukaesebrot/dgc"
)

var pongCmd = &dgc.Command{
	Name:        "pong",
	Description: "Ping",
	Usage:       "pong",
	Flags:       []string{},
	IgnoreCase:  true,
	SubCommands: []*dgc.Command{},
	RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.RespondText("You are being rate limited!")
	}),
	Handler: func(ctx *dgc.Ctx) {
		ctx.RespondText("Ping!")
	},
}

// Register registers the Ping command with dgc
func Register(r *dgc.Router) {
	r.RegisterCmd(pongCmd)
}
