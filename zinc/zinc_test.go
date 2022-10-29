package zinc

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

type Article struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Published *time.Time `json:"published"`
	Author    *Author    `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
}

var srcNames []string

func randName() string {
	if srcNames == nil {
		return ""
	}
	return srcNames[rand.Intn(len(srcNames))]
}

//
// TestZincAPI
// @Description:
// @param t
//
func TestZincAPI(t *testing.T) {
	indexName := "zinc-example"
	userName := "admin"
	pass := "admin"

	//err := Delete(userName, pass, "", indexName)
	//if err != nil {
	//	log.Fatalf("delete err:%s", err)
	//}
	//readFile, err := ioutil.ReadFile("mapping.json")
	//if err != nil {
	//	log.Fatalf("read file err:%s", err)
	//}
	//err = Index(userName, pass, "", string(readFile))
	//if err != nil {
	//	log.Fatalf("index err:%s", err)
	//}
	//
	////gen data
	//
	////rand name
	//dd, readErr := ioutil.ReadFile("nickname.json")
	//if readErr == nil {
	//	unmarshalErr := json.Unmarshal(dd, &srcNames)
	//	if unmarshalErr != nil {
	//		tracerr.PrintSource(unmarshalErr, 3)
	//	}
	//} else {
	//	tracerr.PrintSource(readErr)
	//}
	//
	//var articles []interface{}
	//var numItems = 3000
	//names := []string{"Alice", "John", "Mary"}
	//for i := 1; i <= numItems; i++ {
	//	add := time.Now().Add(time.Hour)
	//	articles = append(articles, &Article{
	//		ID:        i,
	//		Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
	//		Body:      randName() + "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片.",
	//		Published: &add,
	//		Author: &Author{
	//			FirstName: names[rand.Intn(len(names))],
	//			LastName:  "Smith",
	//			FullName:  randName(),
	//		},
	//	})
	//}
	//log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))
	//
	////bulk
	//err = Bulk(userName, pass, "", indexName, articles)
	//if err != nil {
	//	log.Fatalf("bulk err:%s", err)
	//}

	//search
	sql2DSL := Sql2DSL("select body from x where body='陈景润' order by id desc limit 1", true)
	log.Printf("DSL:\n%s", sql2DSL)
	metaHit := MSearch(userName, pass, "", indexName, sql2DSL)
	if metaHit != nil {
		total := metaHit.Total.Value
		hits := metaHit.Hits
		if total != 0 {
			log.Printf("total: %d", total)
			for _, hit := range hits {
				source := hit.Source
				article := &Article{}
				Map2struck(&source, article)
				log.Printf("body:%s", article.Body)
			}
		}
	}

}
