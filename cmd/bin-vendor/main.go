package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mjpitz/bindir/internal/archive/targz"
	"github.com/mjpitz/bindir/internal/archive/zip"
	"github.com/mjpitz/bindir/internal/github"
	"github.com/mjpitz/bindir/internal/model"

	"github.com/ghodss/yaml"
)

func fatal(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func setupFileSystem(binDirectory, vendorDirectory string) error {
	_, binErr := os.Stat(binDirectory)
	_, vendorErr := os.Stat(vendorDirectory)

	if os.IsExist(binErr) && !os.IsExist(vendorErr) {
		return fmt.Errorf("bin directory already exists and isn't vendored")
	}

	_ = os.MkdirAll(vendorDirectory, 0755)
	return nil
}

func main() {
	flags := flag.NewFlagSet("bin-vendor", flag.ExitOnError)

	configFile := flags.String("config", "bin.yaml", "override the default config file.")

	_ = flags.Parse(os.Args)

	fileContents, err := ioutil.ReadFile(*configFile)
	fatal(err)

	config := &model.Config{}
	err = yaml.Unmarshal(fileContents, config)
	fatal(err)

	binDirectory := path.Join("bin")
	vendorDirectory := path.Join(binDirectory, "vendor")

	err = setupFileSystem(binDirectory, vendorDirectory)
	fatal(err)

	for _, tool := range config.Tools {
		var fileName string

		safeVersion := strings.ReplaceAll(tool.Version, "/", "_")
		safeVersion = strings.ReplaceAll(safeVersion, "\\", "_")

		ref := fmt.Sprintf("%s@%s", tool.Name, safeVersion)

		log.Println("fetching", ref)
		if tool.GitHubRelease != nil {
			fileName, err = github.Download(tool, vendorDirectory)
		}

		fatal(err)

		packed := path.Join(vendorDirectory, fileName)
		unpacked := path.Join(vendorDirectory, ref)

		if strings.HasSuffix(fileName, ".tar.gz") {
			err = targz.Untargz(packed, unpacked)
		} else if strings.HasSuffix(fileName, ".tgz") {
			err = targz.Untargz(packed, unpacked)
		} else if strings.HasSuffix(fileName, ".zip") {
			err = zip.Unzip(packed, unpacked)
		}

		fatal(err)

		binaryLink, err := filepath.Abs(path.Join(binDirectory, tool.Name))
		fatal(err)

		err = filepath.Walk(unpacked, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return err
			}

			// if its' executable, link it to the
			if strings.Contains(path, tool.Name) && info.Mode()&0010 != 0 {
				path, err = filepath.Abs(path)
				if err != nil {
					return err
				}

				_ = os.Remove(binaryLink)
				_ = os.Symlink(path, binaryLink)
			}
			return nil
		})
		fatal(err)
	}
}
