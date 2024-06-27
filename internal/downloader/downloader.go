package downloader

import (
	"cpk/internal/configs"
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(url string) {
	if url != "" {

		var paths = strings.Split(url, "/")
		var name = strings.Split(paths[len(paths)-1], ".")
		var path strings.Builder
		path.WriteString(configs.Cpks_Settings.Install_dir)
		// var path =

		if name[1] == "h" || name[1] == "hpp" {
			path.WriteString("headers/")
		} else {
			path.WriteString("libs/")
		}

		out, create_err := os.Create(path.String() + paths[len(paths)-1])
		resp, get_err := http.Get(url)
		n, err := io.Copy(out, resp.Body)
		if create_err != nil || get_err != nil || n < 1 {
			panic(err)
		}

		defer out.Close()
		defer resp.Body.Close()
	}
}
