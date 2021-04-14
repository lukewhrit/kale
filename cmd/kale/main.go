/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/joho/godotenv"
	"github.com/lukewhrit/kale/internal/app"
	"github.com/lukewhrit/kale/internal/pkg/domain"
)

func init() {
	// Load values in dotenv file to environment
	if err := godotenv.Load(); err != nil {
		domain.Logger.Error(err)
	}

	// Load environment variables into a golang struct
	if err := app.LoadConfig(); err != nil {
		domain.Logger.Error(err)
	}
}

func main() {
	if err := app.Start(); err != nil {
		domain.Logger.Error(err)
	}
}
