package reckonercdp

import (
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	client := NewClient(&ClientSettings{
		AccessKeyID:     "416aeb30-97eb-4fee-a4d7-0348779297e7",
		SecretAccessKey: "f2fc5ed4-e4f0-4ed3-9df9-a0df273d6afb",
	})

	res, err := client.Query("SELECT * FROM \\`reckoner-dev.load_test\\`.\\`load_test\\` LIMIT 1")
	if err != nil {
		t.Fatalf("Client#Query does not work.")
		return
	}
	log.Println(res)
}
