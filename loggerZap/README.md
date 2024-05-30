# Zap logger
# Uber Zap

A structured logging library developer by `Uber`.

two types : 
1) Plain Zap
2) Sugared Zap

`logger.Sync()` : important as it ensures all logs are recorded regardless how program exits.

`cmd to install` : go get -u go.uber.org/zap

zap supports various logging levels:  
`Debug, Info, Warn, Error, DPanic, Panic, and Fatal`

zap provides 3 functions for creating logs tailored for various environments and use cases:  
`zap.NewProduction, zap.NewExample, and zap.NewDevelopment`
