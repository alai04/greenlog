package log

import (
	"fmt"
	"os"
	"path"
)

// LogRec 表示传入的log信息
type LogRec struct {
	Did      string `json:"did"`
	Category string `json:"category"`
	Time     string `json:"time"`
	Message  string `json:"message"`
}

var (
	dir string
)

func init() {
	SetDir("log_files")
}

func mustDirExists() {
	fi, err := os.Stat(dir)
	if err == nil {
		if !fi.IsDir() {
			panic(fmt.Errorf("%s is not a directory", dir))
		}
	} else {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

// SetDir 设置log文件目录
func SetDir(d string) {
	dir = d
	mustDirExists()
}

// Log2File 将log写入文件，每个设备id对应一个文件
func Log2File(rec *LogRec) error {
	file, err := os.OpenFile(path.Join(dir, rec.Did+".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s [%s]: %s\n", rec.Time, rec.Category, rec.Message)
	return nil
}
