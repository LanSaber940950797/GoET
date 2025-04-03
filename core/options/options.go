package options

type AppType uint32

const (
	Server AppType = iota
	Watcher
	GameTool
	ExcelExporter
	Proto2CS
	BenchmarkClient
	BenchmarkServer
	Demo
	LockStep
)

type Options struct {
	AppType     AppType
	StartConfig string
	Process     int32
	Develop     int32
	LogLevel    int32
	Console     int32
}

var options = &Options{
	AppType:     Server,
	StartConfig: "config.json",
	Process:     1,
	Develop:     0,
	LogLevel:    0,
	Console:     0,
}

func GetAppType() AppType {
	return options.AppType
}

func GetStartConfig() string {
	return options.StartConfig
}

func GetProcess() int32 {
	return options.Process
}

func GetDevelop() int32 {
	return options.Develop
}

func GetLogLevel() int32 {
	return options.LogLevel
}

func GetConsole() int32 {
	return options.Console
}
