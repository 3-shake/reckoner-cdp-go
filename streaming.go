package reckonercdp

import (
	"net/url"

	"reflect"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/vmihailenco/msgpack"
)

func (client *Client) Insert(databaseName, tableName string, src interface{}) error {
	guid := xid.New()
	values := url.Values{}
	values.Add("destination_database", databaseName)
	values.Add("destination_table", tableName)
	values.Add("record_id", guid.String())

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Slice {
		return errors.New("src is not Slice")
	}

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

func (client *Client) BulkInsert(databaseName, tableName string, src interface{}) error {
	guid := xid.New()
	values := url.Values{}
	values.Add("destination_database", databaseName)
	values.Add("destination_table", tableName)
	values.Add("record_id", guid.String())

	srcVal := reflect.ValueOf(src)
	switch srcVal.Kind() {
	case reflect.Slice:
		for i := 0; i < srcVal.Len(); i++ {
			if srcVal.Index(i).Elem().Elem().Kind() != reflect.Struct && srcVal.Index(i).Elem().Elem().Kind() != reflect.Map {
				return errors.New("src is not Slice of Struct or Map")
			}
		}
	default:
		return errors.New("src is not Slice")
	}

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
