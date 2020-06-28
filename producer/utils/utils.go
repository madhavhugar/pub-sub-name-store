package utils

import log "github.com/sirupsen/logrus"

// HandleError logs a given message and the error object
func HandleError(err error, msg string) {
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error(msg)
	}
}
