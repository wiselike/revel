package revel

import (
	"github.com/wiselike/revel/logger"
)

// Logger.
var (
	// The root log is what all other logs are branched from, meaning if you set the handler for the root
	// it will adjust all children.
	RootLog = logger.New()
	// This logger is the application logger, use this for your application log messages - ie jobs and startup,
	// Use Controller.Log for Controller logging
	// The requests are logged to this logger with the context of `section:requestlog`.
	AppLog = RootLog.New("module", "app")
	// This is the logger revel writes to, added log messages will have a context of module:revel in them
	// It is based off of `RootLog`.
	RevelLog = RootLog.New("module", "revel")

	// This is the handler for the RootLog, it is stored so that if the RootLog is changed it can be assigned to the
	// new RootLog.
	rootLogHandler *logger.CompositeMultiHandler

	// System logger.
	SysLog = AppLog.New("section", "system")
)

// Initialize the loggers first.
func init() {
	// RootLog.SetHandler(
	//	logger.LevelHandler(logger.LogLevel(log15.LvlDebug),
	//		logger.StreamHandler(os.Stdout, logger.TerminalFormatHandler(false, true))))
	initLoggers()
	OnAppStart(initLoggers, -5)
}

func initLoggers() {
	rootHandle := logger.InitializeFromConfig(BasePath, Config)

	// Set all the log handlers
	setRootLog(rootHandle)
}

// Set the application log and handler, if handler is nil it will
// use the same handler used to configure the application log before.
func setRootLog(rootHandle *logger.CompositeMultiHandler) {
	if rootHandle != nil {
		rootLogHandler = rootHandle
		// Set the app log and the handler for all forked loggers
		RootLog.SetHandler(rootLogHandler)

		// Set the system log handler - this sets golang writer stream to the
		// sysLog router
		logger.SetDefaultLog(SysLog)
		SysLog.SetStackDepth(5)
		SysLog.SetHandler(rootLogHandler)
	}
}
