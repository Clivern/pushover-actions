// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clivern/pushover-actions/internal/app/handler"
	"github.com/clivern/pushover-actions/internal/app/module"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	configFile := "/github/workflow/event.json"

	eventsData, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while reading event file [%s]: %s\n",
			configFile,
			err.Error(),
		))
	}

	watchTower := module.WatchTower{}
	watchTower.Add(handler.NewIssue{})
	watchTower.Add(handler.EditIssue{})
	watchTower.Notify(string(eventsData))

	httpClient := module.NewHTTPClient()

	response, error := httpClient.Post(
		context.Background(),
		"https://api.pushover.net/1/messages.json",
		fmt.Sprintf(
			"token=%s&user=%s&title=%s&sound=none&message=%s",
			os.Getenv("PUSHOVER_TOKEN"),
			os.Getenv("PUSHOVER_USER"),
			"Hooray!",
			"Something Just Happened.",
		),
		map[string]string{},
		map[string]string{},
	)

	if error != nil {
		panic(fmt.Sprintf(
			"Error while calling pushover: %s\n",
			error.Error(),
		))
	}

	fmt.Printf(httpClient.ToString(response))
}
