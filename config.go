package main

import (
	"fmt"

	libconfig "github.com/opensourceways/community-robot-lib/config"
)

type configuration struct {
	ConfigItems []botConfig `json:"config_items,omitempty"`
}

func (c *configuration) configFor(org, repo string) *botConfig {
	if c == nil {
		return nil
	}

	items := c.ConfigItems
	v := make([]libconfig.IPluginForRepo, len(items))

	for i := range items {
		v[i] = &items[i]
	}

	if i := libconfig.FindConfig(org, repo, v); i >= 0 {
		return &items[i]
	}

	return nil
}

func (c *configuration) Validate() error {
	if c == nil {
		return nil
	}

	items := c.ConfigItems
	for i := range items {
		if err := items[i].validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *configuration) SetDefault() {
	if c == nil {
		return
	}

	Items := c.ConfigItems
	for i := range Items {
		Items[i].setDefault()
	}
}

type botConfig struct {
	libconfig.PluginForRepo
	SwitchOfMilestone string `json:"switch_of_milestone" required:"true"`
}

func (c *botConfig) setDefault() {
}

func (c *botConfig) validate() error {
	if err := c.PluginForRepo.Validate(); err != nil {
		return err
	}

	if c.SwitchOfMilestone == "" {
		return fmt.Errorf("missing SwitchOfMilestone")
	}

	return nil
}
