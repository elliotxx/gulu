<div align = "center">
<p>
    <img width="160" src="https://github.com/elliotxx/gulu/blob/master/golang-logo.png?sanitize=true">
</p>
<h2>A Golang util (轱辘)</h2>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/gulu"><img src="https://pkg.go.dev/badge/github.com/elliotxx/gulu.svg"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/elliotxx/gulu"><img src="https://goreportcard.com/badge/github.com/elliotxx/gulu?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/gulu?branch=master"><img src="https://img.shields.io/coveralls/github/elliotxx/gulu/master"></a>
<a title="Code Size" target="_blank" href="https://github.com/elliotxx/gulu"><img src="https://img.shields.io/github/languages/code-size/elliotxx/gulu.svg?style=flat-square"></a>
<br>
<a title="GitHub release" target="_blank" href="https://github.com/elliotxx/gulu/releases"><img src="https://img.shields.io/github/release/elliotxx/gulu.svg"></a>
<a title="License" target="_blank" href="https://github.com/elliotxx/gulu/blob/master/LICENSE"><img src="https://img.shields.io/github/license/elliotxx/gulu"></a>
<a title="GitHub Commits" target="_blank" href="https://github.com/elliotxx/gulu/commits/master"><img src="https://img.shields.io/github/commit-activity/m/elliotxx/gulu.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/elliotxx/gulu/commits/master"><img src="https://img.shields.io/github/last-commit/elliotxx/gulu.svg?style=flat-square&color=FF9900"></a>
</p>
</div>

Gulu is a Golang util, `gulu` inspired by [88250/gulu](https://github.com/88250/gulu), cool!

## 📜 Language

[English](https://github.com/elliotxx/gulu/blob/master/README.md) | [简体中文](https://github.com/elliotxx/gulu/blob/master/README-zh.md)

## ⚡ Usage

```
go get -u github.com/elliotxx/gulu
```

## ✨ Features

### archive

```go
func UnpackArchive(targetDir, archiveFile string) error
func UnpackTarGz(targetDir, archiveFile string) error
func UnpackZip(targetDir, archiveFile string) error
```

### cmdutil

```go
func CheckErr(err error)
```

Example:

```go
func main() {
	cmd := NewRootCommand()

	if err := cmd.Execute(); err != nil {
		cmdutil.CheckErr(err)
	}
}
```

### configutil

```go
func FromConfigString(config, configType string, data interface{}) error
func FromFile(fs afero.Fs, filename string, data interface{}) error
func GetFileExtension(filename string) string
func IsValidConfigFilename(filename string) bool
```

Example:

```go
configData := &Configuration{}
if !configutil.IsValidConfigFilename(configFile) {
    return fmt.Errorf("invalid config file: %s", configFile)
}
err := configutil.FromFile(afero.NewOsFs(), configFile, configData)
if err != nil {
    return err
}
```

### gitutil

```go
func GetHeadHash() (sha string, err error)
func GetHeadHashShort() (sha string, err error)
func GetLatestTag() (string, error)
func GetLatestTagFromLocal() (tag string, err error)
func GetLatestTagFromRemote() (tag string, err error)
func GetRemoteURL() (string, error)
func GetTagCommitSha(tag string) (sha string, err error)
func GetTagCommitShaFromLocal(tag string) (sha string, err error)
func GetTagCommitShaFromRemote(_ string) (string, error)
func GetTagList() (tags []string, err error)
func GetTagListFromRemote(remoteURL string, reverse bool) (tags []string, err error)
func IsDirty() (dirty bool, err error)
func IsHeadAtTag(tag string) (bool, error)
```

### json

```go
func MustMarshal(v interface{}) []byte
func MustMarshalString(v interface{}) string
func MustPrettyMarshal(v interface{}) []byte
func MustPrettyMarshalString(v interface{}) string
```

### log

```go
func Debug(args ...interface{})
func Debugf(format string, args ...interface{})
func Error(args ...interface{})
func Errorf(format string, args ...interface{})
func Fatal(args ...interface{})
func Fatalf(format string, args ...interface{})
func Info(args ...interface{})
func Infof(format string, args ...interface{})
func Panic(args ...interface{})
func Panicf(format string, args ...interface{})
func SetLevel(level Level)
func SetLogger(config *Configuration)
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder)
func Warn(args ...interface{})
func Warnf(format string, args ...interface{})
type Configuration
func NewDefaultConfiguration() *Configuration
type Level
func GetLevelFromStr(level string) Level
type LogFile
func GetLogFile() LogFile
type Logger
```

### misc

```go
func GetKeys(m map[string]interface{}) []string
```

### os

```go
func Exists(path string) bool
func GetGoModRootDir(curDir string) (string, error)
func GetParentDir(dirctory string) string
func IsDir(path string) bool
func IsFile(path string) bool
func ReadLines(path string) ([]string, error)
```

### version

```go
func JSON() string
func ReleaseVersion() string
func ShortString() string
func String() string
func YAML() string
type BuildInfo
type GitInfo
type Info
func NewDefaultVersionInfo() *Info
func NewInfo() (*Info, error)
func NewMainOrDefaultVersionInfo() *Info
func (v *Info) JSON() string
func (v *Info) ShortString() string
func (v *Info) String() string
func (v *Info) YAML() string
```
