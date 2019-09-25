package main
import (
	"io"
	"log"
	"net/http"
	"time"
	apachelog "github.com/lestrrat-go/apache-logformat"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})

	// WithMaxAge和WithRotationCount二者只能设置一个，

	// WithMaxAge设置文件清理前的最长保存时间，
	//rotatelogs.WithMaxAge(time.Hour*24),

	// WithRotationCount设置文件清理前最多保存的个数。
	//rotatelogs.WithRotationCount(maxRemainCnt),
	logf, err := rotatelogs.New(
		"/tmp/access_log.%Y%m%d%H%M",

		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName("/tmp/access_log.txt"),

		//保留24小时的数据内的日志文件
		rotatelogs.WithMaxAge(24 * time.Hour),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	// Now you must write to logf. apache-logformat library can create
	// a http.Handler that only writes the approriate logs for the request
	// to the given handle
	http.ListenAndServe(":8080", apachelog.CombinedLog.Wrap(mux, logf))
}