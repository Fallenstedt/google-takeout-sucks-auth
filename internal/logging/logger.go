package logging

import (
	"log"
	"os"
)

var InfoLog = log.New(os.Stdout, "INFO ", log.LstdFlags)
var ErrorLog = log.New(os.Stderr, "ERROR ", log.LstdFlags)
