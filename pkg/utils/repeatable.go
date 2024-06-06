package repeatable

import (
	"image/jpeg"
	"io"
	"os"
	"test/pkg/logging"
	"time"

	"github.com/jdeng/goheif"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

var (
	PublicFilePath = "./../../public"
)

func CrateDir() error {
	return os.MkdirAll(PublicFilePath, os.ModePerm)
}


func HeicToJpeg(heicFormat, jpegFormat string) {
	var logger *logging.Logger
	fout := jpegFormat
	fi, err := os.Open(heicFormat)

	if err != nil {
		logger.Error("os.Open(heicFormat)", err)
	}

	defer fi.Close()

	exif, err := goheif.ExtractExif(fi)
	if err != nil {
		logger.Error("Warning: no EXIF from %s: %v\n", "", err)
	}

	img, err := goheif.Decode(fi)
	if err != nil {
		logger.Error("Failed to parse %s: %v\n", "fin", err)
	}

	fo, err := os.OpenFile(fout, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logger.Error("Failed to create output file %s: %v\n", fout, err)
	}
	defer fo.Close()

	w, _ := newWriterExif(fo, exif)
	err = jpeg.Encode(w, img, nil)
	if err != nil {
		logger.Error("Failed to encode %s: %v\n", fout, err)
	}
}

type writerSkipper struct {
	w           io.Writer
	bytesToSkip int
}

func (w *writerSkipper) Write(data []byte) (int, error) {
	if w.bytesToSkip <= 0 {
		return w.w.Write(data)
	}

	if dataLen := len(data); dataLen < w.bytesToSkip {
		w.bytesToSkip -= dataLen
		return dataLen, nil
	}

	if n, err := w.w.Write(data[w.bytesToSkip:]); err == nil {
		n += w.bytesToSkip
		w.bytesToSkip = 0
		return n, nil
	} else {
		return n, err
	}
}

func newWriterExif(w io.Writer, exif []byte) (io.Writer, error) {
	writer := &writerSkipper{w, 2}
	soi := []byte{0xff, 0xd8}
	if _, err := w.Write(soi); err != nil {
		return nil, err
	}

	if exif != nil {
		app1Marker := 0xe1
		markerlen := 2 + len(exif)
		marker := []byte{0xff, uint8(app1Marker), uint8(markerlen >> 8), uint8(markerlen & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}
