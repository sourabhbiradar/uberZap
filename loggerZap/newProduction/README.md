# zap.NewProduction()

Configures a logger for production environments

Logs at the `Info level and above`.

NOTE: Does NOT log Debug level

Logs to both standard output and standard error.

output : time, level, message, and fields in each log entry. (`JSON`)

{"level":"info","ts":1717028449.0636532,"caller":"newProduction/main.go:15","msg":"Second zap log","user ID:":3210,"name":"Abc","error":"dummy error"}