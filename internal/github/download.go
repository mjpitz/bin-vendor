package github

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/mjpitz/bindir/internal/model"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"text/template"
)

func Download(tool *model.Tool, vendor string) (string, error) {
	name, version, gh := tool.Name, tool.Version, tool.GitHubRelease
	osName, archName := runtime.GOOS, runtime.GOARCH

	format := gh.Format
	if runtime.GOOS == "linux" && gh.Linux != "" {
		format = gh.Linux
	} else if runtime.GOOS == "darwin" && gh.OSX != "" {
		format = gh.OSX
	} else if runtime.GOOS == "windows" && gh.Windows != "" {
		format = gh.Windows
	}

	if sug := gh.Replacements[osName]; sug != "" {
		osName = sug
	}

	if sug := gh.Replacements[archName]; sug != "" {
		archName = sug
	}

	t, err := template.New(name).
		Funcs(sprig.TxtFuncMap()).
		Parse(format)

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	err = t.Execute(buf, &model.Render{
		Name:    name,
		Version: version,
		OS:      osName,
		Arch:    archName,
	})
	if err != nil {
		return "", err
	}

	fileName := buf.String()

	asset := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s",
		gh.Repository, version, fileName)

	resp, err := http.Get(asset)
	if err != nil {
		return "", err
	}

	vendor = path.Join(vendor, fileName)
	file, err := os.OpenFile(vendor, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
