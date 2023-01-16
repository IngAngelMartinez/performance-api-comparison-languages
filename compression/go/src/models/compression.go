package models

import (
	"bytes"
	"compression/src/extensions"
	"encoding/json"
	"time"
)

type Compression struct {
	Id     string                 `json:"id,omitempty"`
	Name   string                 `json:"name"`
	Data   map[string]interface{} `json:"data"`
	Number int                    `json:"number"`
	Date   time.Time              `json:"date"`
}

func (c *Compression) ToDB() CompressionDb {

	b, _ := json.Marshal(c.Data)

	// if err != nil {

	// 	return CompressionDb{}, err
	// }

	var buffer bytes.Buffer

	extensions.Compress(&buffer, b)

	return CompressionDb{
		Id:     c.Id,
		Name:   c.Name,
		Data:   buffer.Bytes(),
		Number: c.Number,
		Date:   c.Date,
	}
}
