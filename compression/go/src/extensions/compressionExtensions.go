package extensions

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

func Compress(writer io.Writer, data []byte) error {

	gw, err := gzip.NewWriterLevel(writer, gzip.BestCompression)

	defer gw.Close()

	gw.Write(data)

	return err
}

func Decompress(writer io.Writer, data []byte) error {

	gr, err := gzip.NewReader(bytes.NewBuffer(data))

	defer gr.Close()

	data, err = ioutil.ReadAll(gr)

	if err != nil {

		return err
	}

	writer.Write(data)

	return nil
}
