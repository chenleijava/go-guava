package zinc

import (
	"bytes"
	"fmt"
	"github.com/chenleijava/go-guava/elasticsql"
	jsoniter "github.com/json-iterator/go"
	"github.com/ztrue/tracerr"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//关注文档： https://docs.zincsearch.com/api/index/delete/
//使用原生的API接口进行操作zinc-server
//关于中文分词： https://mp.weixin.qq.com/s/g9fcKNSEHqwiK8Tt3MY3GA

//  zinc server
const BASEURL = "http://localhost:4080"

//
// Delete
// @Description: delete index
// @param userName
// @param password
// @param baseUrl
// @param indexName
// @return error
//
func Delete(userName, password, baseUrl, indexName string) error {
	if baseUrl == "" {
		baseUrl = BASEURL
	}
	dd, err3 := zinc0(http.MethodDelete, userName, password, fmt.Sprintf("%s/api/index/%s", baseUrl, indexName), "")
	if err3 != nil {
		tracerr.PrintSource(err3)
	} else {
		log.Printf("delete index :%s", dd)
	}
	return err3
}

//
// Index
// @Description: create index
// @param userName
// @param password
// @param baseUrl
// @param mapping {"name":"test-bulk-example","shard_num":3,"storage_type":"disk","settings":{"analysis":{"analyzer":{"default":{"type":"gse_search"}}}},"mappings":{"properties":{"published":{"type":"date","sortable":true}}}}
// @return error
//
func Index(userName, password, baseUrl, mapping string) error {
	if baseUrl == "" {
		baseUrl = BASEURL
	}
	//index http://localhost:4080/api/index
	d, err := zinc0(http.MethodPost, userName, password,
		fmt.Sprintf("%s/api/index", baseUrl), mapping)
	if err != nil {
		tracerr.PrintSource(err, 3)
	} else {
		log.Printf("create index : %s", d)
	}
	return err
}

//
// Bulk
// @Description: bulk
// @param userName
// @param password
// @param baseUrl
// @param indexName
// @param values
func Bulk(userName, password, baseUrl, indexName string, values []interface{}) error {
	if baseUrl == "" {
		baseUrl = BASEURL
	}
	var buf bytes.Buffer
	for _, v := range values {
		marshal, err := json.Marshal(v)
		if err != nil {
			tracerr.PrintSource(err)
		}
		if buf.Len() != 0 {
			buf.WriteRune('\n')
		}
		buf.Write(marshal)
	}
	//https://docs.zincsearch.com/api/document/multi/
	dd, err := zinc0(http.MethodPost, userName, password,
		fmt.Sprintf("%s/api/%s/_multi", baseUrl, indexName), buf.String())
	if err != nil {
		tracerr.PrintSource(err, 3)
	} else {
		log.Printf("%s", dd)
	}
	return err
}

//https://docs.zincsearch.com/api-es-compatible/search/search/
//
// MSearch
// @Description: es search compatible
// @param userName
// @param password
// @param baseUrl  http://localhost:4080
// @param indexName index name
// @param dsl
//
func MSearch(userName, password, baseUrl, indexName, dsl string) *MetaHit {
	//POST http://localhost:4080/es/olympics/_search
	if baseUrl == "" {
		baseUrl = BASEURL
	}
	url := fmt.Sprintf("%s/es/%s/_search", baseUrl, indexName)
	b, err := zinc0(http.MethodPost, userName, password, url, dsl)
	//http happened err
	if err != nil {
		tracerr.PrintSource(err)
		return nil
	} else {
		metaSearchResp := &MetaSearchResponse{}
		_ = json.Unmarshal(b, metaSearchResp)
		return metaSearchResp.Hits
		////hit total
		//total := metaHits.Total
		//hits := metaHits.Hits
		//for _, metaHit := range hits {
		//	metaSource := metaHit.Source
		//	article := &Article{}
		//	marshal, _ := json.Marshal(metaSource)
		//	_ = json.Unmarshal(marshal, article)
		//	log.Printf("%s", metaSource)
		//}
	}
}

//
// map2struck
// @Description:
// @param s pointer
// @param out pointer
//
func Map2struck(s *map[string]interface{}, out interface{}) {
	marshal, _ := json.Marshal(s)
	_ = json.Unmarshal(marshal, out)
}

//
// zinc0
// @Description: zinc0 http client
// @param method  http method
// @param userName
// @param password
// @param url  see https://docs.zincsearch.com/api/index/analyze/
// @param query
// @return []byte
// @return error
//
func zinc0(method, userName, password, url, query string) ([]byte, error) {
	var req *http.Request
	var err error

	if query != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(query))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(userName, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "zinc0-http-client 0.3.3")

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		//  close body
		defer func(Body io.ReadCloser) {
			var err = Body.Close()
			if err != nil {
				log.Printf("%s", err)
				tracerr.PrintSource(err)
			}
		}(resp.Body)

		if err != nil {
			tracerr.PrintSource(err)
			return nil, err
		}

		return io.ReadAll(resp.Body)
	}
	return nil, err
}

//
//MetaSearchResponse
//@Description:  zinc0 search response
//
type MetaSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits *MetaHit `json:"hits"`
}

//
//MetaHit
//@Description:
//
type MetaHit struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []struct {
		Index     string                 `json:"_index"`
		Type      string                 `json:"_type"`
		Id        string                 `json:"_id"`
		Score     float64                `json:"_score"`
		Timestamp time.Time              `json:"@timestamp"`
		Source    map[string]interface{} `json:"_source"`
	} `json:"hits"`
}

//
// Sql2DSL
// @Description:  sql to es DSL
// @param sql
// @param pretty
// @return string
//
func Sql2DSL(sql string, pretty bool) string {
	var dsl string
	var err error
	if pretty {
		dsl, _, err = elasticsql.ConvertPretty(sql)
	} else {
		dsl, _, err = elasticsql.Convert(sql)
	}
	if err != nil {
		tracerr.PrintSourceColor(err, 3)
	}
	return dsl

}
