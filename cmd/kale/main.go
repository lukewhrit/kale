/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"log"

	"github.com/lukewhrit/kale/internal/app"
)

func init() {
	if err := app.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
