/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package app

import (
	"github.com/caarlos0/env/v6"
)

// Config represents a Kale bot configuration object
var Config struct {
	Token       string `env:"KALE_TOKEN"`
	Prefix      string `env:"KALE_PREFIX"`
	DatabaseURI string `env:"KALE_DATABASE_URI"`
	BotOwner    string `env:"KALE_BOT_OWNER"`
	ClientID    string `env:"KALE_CLIENT_ID"`
}

// LoadConfig loads the bots configuration from environment variables
func LoadConfig() error {
	return env.Parse(&Config)
}
