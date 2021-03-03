/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package app

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/lukewhrit/kale/commands/ping"
)

func loadAllCommands(r *dgc.Router) {
	ping.Register(r)
}
