// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/url"

	"github.com/anonhoarder/demeter/db"
	"github.com/anonhoarder/demeter/lib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add hosturl",
	Args:  cobra.ExactArgs(1),
	Short: "add a host to the scrape list",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := url.Parse(args[0])
		if err != nil {
			log.WithField("err", err).Error("invalid url provided")
			return
		}
		h := lib.Host{
			URL: args[0],
		}

		err = db.Conn.Save(&h)
		if err != nil {
			log.WithField("err", err).Error("could not save")
			return
		}
		fmt.Println(args[0], "added")
		log.WithFields(log.Fields{
			"id":  h.ID,
			"url": h.URL,
		}).Info("host has been added to the database")
	},
}

func init() {
	hostCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
