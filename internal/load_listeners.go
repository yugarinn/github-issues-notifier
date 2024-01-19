package internal

import (
	"io/ioutil"

	"github.com/yugarinn/github-issues-notifier/core"

	"gopkg.in/yaml.v3"
)


func LoadListeners(app *core.App) ([]Listener, error){
    var listenersWrapper ListenersWrapper
    listenersFilePath := app.Config.ListenersFilePath

    repositoriesData, err := ioutil.ReadFile(listenersFilePath)
    if err != nil {
        return []Listener{}, err
    }

    if err := yaml.Unmarshal(repositoriesData, &listenersWrapper); err != nil {
        return []Listener{}, err
    }

    listeners := listenersWrapper.Listeners
    activeListeners := listeners[:0]

    for _, listener := range listeners {
        if listener.IsActive {
            activeListeners = append(activeListeners, listener)
        }
    }

    return activeListeners, nil
}
