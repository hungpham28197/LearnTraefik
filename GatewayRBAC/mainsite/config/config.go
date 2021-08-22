package config

import (
	"fmt"
	"os"
	"strings"

	json "github.com/goccy/go-json"
)

type Configuration struct {
	Host     string
	Port     string
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`

	Redis struct {
		Network     string `json:"network"`
		Addr        string `json:"addr"`
		Password    string `json:"password"`
		Database    string `json:"database"`
		MaxIdle     int    `json:"max_idle"`
		MaxActive   int    `json:"max_active"`
		IdleTimeout int    `json:"idle_timeout"`
		Prefix      string `json:"prefix"`
	} `json:"redis"`

	Authenticate struct {
		MaxAccessFailed      int32 `json:"max_access_failed"`
		LockoutInMinutes     int   `json:"lockout_in_minutes"`
		BlackListTokenInDays int32 `json:"black_list_token_in_days"`
	} `json:"authenticate"`

	SMTPConfig struct {
		Host     string
		From     string
		Password string
		Port     int
	} `json:"smtp"`
}

var Config *Configuration

/*
Đọc cấu hình từ file json. Nếu có bất kỳ lỗi nào thì raise panic error
baseDir rỗng đọc từ thư mục hiện tại "."
baseDir khác rỗng thì đọc từ thư mục truyền vào
*/
func ReadConfig(baseDir ...string) {
	Config = new(Configuration)

	var baseDir_ string
	if len(baseDir) == 0 {
		baseDir_ = "."
	} else {
		baseDir_ = baseDir[0]
	}

	var configFile *os.File
	var err error

	if IsAppInDebugMode() {
		fmt.Println("*** Debug Mode ***")
		if configFile, err = os.Open(baseDir_ + "/config/config.dev.json"); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("*** Production Mode ***")
		if configFile, err = os.Open(baseDir_ + "/config/config.product.json"); err != nil {
			panic(err)
		}
	}

	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(Config); err != nil {
		panic(err)
	}
}

/*
Trả về true nếu ứng dụng đang chạy ở chế độ Debug và ngược lại

*/
func IsAppInDebugMode() bool {
	appCommand := os.Args[0]
	if strings.Contains(appCommand, "debug") || strings.Contains(appCommand, "exe") {
		return true
	}
	return false
}
