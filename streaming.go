package reckonercdp

import (
	"net/url"

	"github.com/vmihailenco/msgpack"
)

func (client *Client) Insert(databaseName, tableName string, src interface{}) error {
	values := url.Values{}
	values.Add("destination_database", databaseName)
	values.Add("destination_table", tableName)

	b, err := msgpack.Marshal([]interface{}{src})
	if err != nil {
		return err
	}
	values.Add("data", string(b))

	res, err := client.streamingGet("/api/v1/streaming", values)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		return client.createError(res, "Insert failed")
	}

	return nil
}

func (client *Client) BulkInsert(databaseName, tableName string, src ...interface{}) error {
	values := url.Values{}
	values.Add("destination_database", databaseName)
	values.Add("destination_table", tableName)

	b, err := msgpack.Marshal(&src)
	if err != nil {
		return err
	}
	values.Add("data", string(b))

	res, err := client.streamingGet("/api/v1/streaming", values)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode/100 != 2 {
		return client.createError(res, "Bulk Insert failed")
	}

	return nil
}
