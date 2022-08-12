// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/spf13/cobra"
	"github.com/wangxn2015/onos-lib-go/pkg/cli"
)

//var viper = viperapi.New()

// init initializes the command line
func init() {
	cli.InitConfig("logging")
}

// Init is a hook called after cobra initialization
func Init() {
	// noop for now
}

// GetCommand returns the root command for the logging service
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log {set/get} level [args]",
		Short: "logging api commands",
	}

	cmd.AddCommand(getSetCommand())
	cmd.AddCommand(getGetCommand())

	return cmd
}
