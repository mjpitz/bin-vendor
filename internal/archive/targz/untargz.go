package targz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
)

func writeFile(filePath string, header *tar.Header, r io.Reader) error {
	outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, r)
	return err
}

func Untargz(in, out string) error {
	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	tgzr := tar.NewReader(gzr)

	for {
		header, err := tgzr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		filePath := path.Join(out, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			//
			_ = os.MkdirAll(filePath, os.ModePerm)
		case tar.TypeReg:
			//
			_ = os.MkdirAll(path.Dir(filePath), os.ModePerm)

			err = writeFile(filePath, header, tgzr)
			if err != nil {
				return nil
			}
		}
	}
}
