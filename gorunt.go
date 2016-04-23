package gorunt

import (
	"errors"
	"strconv"
	"os/exec"
	"path/filepath"
	log "github.com/Sirupsen/logrus"
)

func Copy(targets FileMap) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Copy",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	logger.Debugf("%v", targets)
	
	// do not glob keys, as destinations may not exist
	targets.GlobValues()
	
	logger.Debugf("%v", targets)
	
	for dest, src := range targets {
		for _, src := range src {
			cmd := exec.Command("cp", "-r", "-L", "-T", src, dest)
			logger.Debugf("S: %s -> D: %s", src, dest)
			if err:= cmd.Run(); err != nil {
				logger.Error(err)
			}
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " erros")
	} else {
		return nil
	}
}

func Clean(targets FileList) error {
	
	logger := log.WithFields(log.Fields{
		"func": "Clean",
	})
	
	logger.Info("Started")
	
	errs := 0
	
	logger.Debugf("%v", targets)
	
	targets.Glob()
	
	logger.Debugf("%v", targets)
	
	for _, target := range targets {
		logger.Debugf("Removing: %s", target)
		cmd := exec.Command("rm", "-r", target)
		if err:= cmd.Run(); err != nil {
			logger.Error(err)
		}
	}
	
	if errs > 0 {
		return errors.New(strconv.Itoa(errs) + " erros")
	} else {
		return nil
	}
}

type FileList []string

func (g *FileList)Glob() {
	
	o := make(FileList, 0, len(*g))
	
	for _, val := range *g {
		vals, err := filepath.Glob(val)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		for _, val := range vals {
			o = append(o, val)
		}
	}
	
	*g = o
}

type FileMap map[string]FileList

func (g *FileMap)Glob() {
	g.GlobKeys()
	g.GlobValues()
}

func (g *FileMap)GlobKeys() {
	o := make(FileMap)
	for key, val := range *g {
		// delete(*g, key)
		keys, err := filepath.Glob(key)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		// var val1 FileList = val
		for _, key := range keys {
			o[key] = val
		}
	}
	
	*g = o
}

func (g *FileMap)GlobValues() {
	for key, val := range *g {
		val.Glob()
		(*g)[key] = val
	}
}