package main

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	client, err := elasticsearch8.NewDefaultClient()

	require.NoError(t, err)
	require.NotNil(t, client)

	defer func() {
		_, err := client.Indices.Delete([]string{"my_index"})
		require.NoError(t, err)
	}()

	res, err := client.Indices.Create("my_index")
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = client.Get("my_index", "1")
	require.NoError(t, err)
	require.NotNil(t, res)

	b, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	t.Log(string(b))
	require.JSONEq(t, `{"_index":"my_index","_id":"1","found":false}`, string(b))

	document := struct {
		Name string `json:"name"`
	}{
		Name: "test_document",
	}
	data, err := json.Marshal(document)
	require.NoError(t, err)

	res, err = client.Create("my_index", "1", bytes.NewReader(data))
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = client.Get("my_index", "1")
	require.NoError(t, err)
	require.NotNil(t, res)

	b, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	t.Log(string(b))
	require.JSONEq(t, `{
  "_index": "my_index",
  "_id": "1",
  "_version": 1,
  "_seq_no": 0,
  "_primary_term": 1,
  "found": true,
  "_source": {
    "name": "test_document"
  }
}`, string(b))

}
