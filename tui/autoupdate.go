// Copyright © 2020 Intel Corporation
//
// SPDX-License-Identifier: GPL-3.0-only

package tui

import (
	"github.com/VladimirMarkelov/clui"
	"github.com/clearlinux/clr-installer/swupd"
)

// AutoUpdatePage is the Page implementation for the auto update enable configuration page
type AutoUpdatePage struct {
	BasePage
}

// GetConfiguredValue Returns the string representation of currently value set
func (aup *AutoUpdatePage) GetConfiguredValue() string {
	if aup.getModel().AutoUpdate.Value() {
		return "Enabled"
	}
	return "Disabled"
}

func newAutoUpdatePage(tui *Tui) (Page, error) {
	page := &AutoUpdatePage{}

	page.setupMenu(tui, TuiPageAutoUpdate, swupd.AutoUpdateTitle,
		BackButton|ConfirmButton, TuiPageMenu)

	desc := swupd.AutoUpdateTitle
	desc += "\n\n"
	desc += swupd.AutoUpdateDesc1
	desc += "\n"
	desc += swupd.AutoUpdateDesc2
	desc += "\n\n"
	desc += swupd.AutoUpdateDesc3 + " " + swupd.AutoUpdateCommand
	desc += "\n\n"
	desc += "See" + "\n" + swupd.AutoUpdateLink + "\n" + "for more information."
	desc += "\n\n"
	desc += swupd.AutoUpdateWarning1 + "\n" + swupd.AutoUpdateWarning2

	lbl := clui.CreateLabel(page.content, 2, 16, desc, Fixed)
	lbl.SetMultiline(true)

	page.backBtn.SetTitle("No [Disable]")
	page.backBtn.SetSize(11, 1)

	page.confirmBtn.SetTitle("Yes [Enable, Default]")
	page.confirmBtn.SetSize(21, 1)

	return page, nil
}

// DeActivate sets the model value and adjusts the "confirm" flag for this page
func (aup *AutoUpdatePage) DeActivate() {
	model := aup.getModel()

	if aup.action == ActionConfirmButton {
		model.AutoUpdate.SetValue(true)
	} else if aup.action == ActionBackButton {
		model.AutoUpdate.SetValue(false)
	}
}

// Activate activates the proper button depending on the current model value.
// If Auto Update is enabled in the data model then the Confirm button will be active
// otherwise the Back button will be activated.
func (aup *AutoUpdatePage) Activate() {
	if aup.getModel().AutoUpdate.Value() {
		aup.activated = aup.confirmBtn
	} else {
		aup.activated = aup.backBtn
	}
}

// GetConfigDefinition returns if the config was interactively defined by the user,
// was loaded from a config file or if the config is not set.
func (aup *AutoUpdatePage) GetConfigDefinition() int {
	if aup.getModel().AutoUpdate.Value() {
		return ConfigDefinedByConfig
	}

	return ConfigDefinedByUser
}
