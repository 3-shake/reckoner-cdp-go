# Reckoner CDP API Client Library for Go

Go Client Library for Reckoner CDP services.

With this gem, you can:
* stream your data to Reckoner CDP databases.
* run query.

## Installation

    $ go get github.com/3-shake/reckoner-cdp-go

## Usage

### Initialization

Initialize a client to your reckoner cdp database:

```go
import "github.com/3-shake/reckoner-cdp-go"

client := reckonercdp.NewClient(&reckonercdp.ClientSettings{
	AccessKeyID:         "yourAccessKeyID",
	SecretAccessKey:     "yourSecretAccessKey",
})
```

You can get your `Access Key ID` and `Secret Access Key` from
`Management` page on WEB UI.

### Stream Your Data to Reckoner CDP Database

Stream Struct:

```go
type Item struct {
	Bar string
	Foo string
}

item := &Item{
		    Bar: "bar",
		    Foo: "foo",
	    }
if err := client.Insert("database_name", "table_name", item); err != nil {
	fmt.Println(err)
}
```

Stream Slice:

```go
type Item struct {
	Bar string
	Foo string
}

items := []interface{}{
	&Item{
		Bar: "bar1",
		Foo: "foo1",
	},
	&Item{
		Bar: "bar2",
		Foo: "foo2",
	},
}
if err := client.BulkInsert("database_name", "table_name", items); err != nil {
	fmt.Println(err)
}
```

Before streaming your data to reckoenr CDP, you need to create database and table which
you stream your data to on reckoner CDP. you can create database and table from WEB UI.

### Query

```go
res, err := client.Query("SELECT * FROM `team.database_name`.`table_name` LIMIT 1")
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(res.Records)
```
