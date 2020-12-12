package st

import (
	"io"
	"log"
	"os"
	"path"
	"time"
)

const logDir = "logs"

// LogToFile - Включить логирование в файл ./logs/filename.log
func LogToFile() (err error) {
	// Создаем директорию логов, если ее нет
	if _, err = os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, 0777); err != nil {
			return
		}
	}

	// В качестве имени файла ставим текущее время
	name := time.Now().Format("02-01-2006_15-04-05") + ".log"
	var file io.Writer
	if file, err = os.Create(path.Join(logDir, name)); err != nil {
		return
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Выключаем вывод стандартного логера в созданный файл и терминал
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	return
}
