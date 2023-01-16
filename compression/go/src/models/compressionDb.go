package models

import (
	"bytes"
	"compression/src/extensions"
	"encoding/json"
	"time"
)

type CompressionDb struct {
	Id     string    `bson:"_id,omitempty"`
	Name   string    `bson:"name"`
	Data   []byte    `bson:"data"`
	Number int       `bson:"number"`
	Date   time.Time `bson:"date"`
}

func (c *CompressionDb) ToView() Compression {

	var buffer bytes.Buffer

	extensions.Decompress(&buffer, c.Data)

	var data map[string]interface{}

	json.Unmarshal(buffer.Bytes(), &data)

	return Compression{
		Id:     c.Id,
		Name:   c.Name,
		Data:   data,
		Number: c.Number,
		Date:   c.Date,
	}
}
