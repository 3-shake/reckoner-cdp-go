package reckonercdp

import (
	"net/url"

	"log"

	"github.com/vmihailenco/msgpack"
)

func (client *Client) Insert(src interface{}, databaseName, tableName string) error {
	values := url.Values{}
	values.Add("destination_database", databaseName)
	values.Add("destination_table", tableName)

	b, err := msgpack.Marshal([]interface{}{src})
	if err != nil {
		return err
	}
	values.Add("data", string(b))

	log.Println(b)

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

func (client *Client) BulkInsert(src []interface{}, databaseName, tableName string) error {
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
