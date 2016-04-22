package npm

import (
    "errors"
	"os/exec"
	"strconv"
	log "github.com/Sirupsen/logrus"
)

var Global = make(map[string]struct{})
var Local = make(map[string]struct{})

func Install() error {
	
	logger := log.WithFields(log.Fields{
		"func": "NPM Install",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	for target, _ := range Global {
		cmd := exec.Command("npm", "install", "-g", target)
		if err:= cmd.Run(); err != nil {
			logger.Error(err)
			errs++
		}
	}
	
	for target, _ := range Local {
		cmd := exec.Command("npm", "install", target)
		if err:= cmd.Run(); err != nil {
			logger.Error(err)
			errs++
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " erros")
	} else {
		return nil
	}
}