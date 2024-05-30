# zap.NewDevelopment()

Configures a logger for development environments.

Logs at the Debug level and above.

Logs to standard output.

More verbose to assist with debugging.

output : time, level, message, caller, and fields. (`only fields JSON formatted`)

2024-05-30T05:34:48.844+0530	DEBUG	newDevelopment/main.go:13	First zap log	{"zap": "NewExapmle"}