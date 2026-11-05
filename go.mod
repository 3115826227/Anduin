module github.com/3115826227/Anduin

go 1.16

require (
 github.com/gin-gonic/gin v1.8.2
 github.com/spf13/viper v1.13.0
 go.uber.org/zap v1.24.0
 gorm.io/driver/mysql v1.4.4
 gorm.io/gorm v1.23.8
)

// Go 1.16 compatibility: newer x/sys versions use unsafe.Slice (Go 1.17+).
replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220330033206-e17cdc41300f
