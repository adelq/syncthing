// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package model

import (
	"github.com/syncthing/syncthing/lib/config"
)

func init() {
	// WO folders are really just RW folders where we reject local changes...
	folderFactories[config.FolderTypeReceiveOnly] = newSendReceiveFolder
}