package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"
)

// 传参接收的变量
var (
	help       bool
	interval   string
	logHome    string
	pidPath    string
	hisLogHome string
)

func init() {
	// 注意 `interval`。默认是 -i string，有了 `interval` 之后，变为 -i interval
	flag.StringVar(&interval, "i", "1d", ": set the time `interval` for splitting logs. supported format: 1d(1 day) 1M(1 mounth)")
	flag.StringVar(&logHome, "d", "/usr/local/nginx/logs/", ": set nginx `loghome`")
	flag.StringVar(&hisLogHome, "H", "/usr/local/nginx/logs/hislogs/", ": set nginx `history log home`")
	flag.StringVar(&pidPath, "p", "/usr/local/nginx/logs/nginx.pid", ": set nginx `nginxpid`")
	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分
	flag.Usage = usage
	flag.Parse()
	createHistoryLogHome()
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: ngxLogrotate [-h] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

func createHistoryLogHome() {
	_, err := os.Stat(hisLogHome)
	if err != nil {
		os.MkdirAll(hisLogHome, 0755)
	}
}

func getLogFilePath() ([]string, error) {
	files := make([]string, 0)
	infos, err := os.ReadDir(logHome)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
	}
	return files, nil
}

func logSplit() {
	files, err := getLogFilePath()
	if err != nil {
		fmt.Println(err)
	} else if len(files) == 0 {
		fmt.Printf("there is no file in directory: %s", logHome)
	}

	for {
		suffix := time.Now().Format(rotateSuffix())
		if len(suffix) != 0 {
			hisLogDir := path.Join(hisLogHome, suffix)
			_, err := os.Stat(hisLogDir)
			if err != nil {
				os.MkdirAll(hisLogDir, 0755)
				for _, oldFile := range files {
					oldLogPath := path.Join(logHome, oldFile)
					newLogPath := path.Join(hisLogDir, oldFile)
					fmt.Printf("move %s to %s\n", oldLogPath, newLogPath)
					syscall.Rename(oldLogPath, newLogPath)
				}
				reloadNginx()
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func reloadNginx() {
	pfile, err := os.Open(pidPath)
	if err != nil {
		fmt.Println("nginx pid file not found")
		return
	}
	defer pfile.Close()
	pidData, _ := ioutil.ReadAll(pfile)
	pid := string(pidData)
	pid = strings.Replace(pid, "\n", "", -1)
	cmd := exec.Command("kill", "-USR1", pid)
	_, errCmd := cmd.Output()
	if errCmd != nil {
		fmt.Println("nginx restart failed：" + errCmd.Error())
		return
	}
}

func rotateSuffix() string {
	switch interval {
	case "1d":
		if time.Now().Format("150405") == "000000" {
			return "20060102"
		}
		return ""
	case "1M":
		if time.Now().Format("02") == "01" {
			return "200601"
		}
		return ""
	default:
		return ""
	}
}

func main() {
	if help {
		flag.Usage()
	} else {
		logSplit()
	}
}
