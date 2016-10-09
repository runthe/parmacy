package main

import (
	"net/http"
	"fmt"

	"github.com/urfave/negroni"
	"github.com/unrolled/render"
	"github.com/danryan/env"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/url"
	"encoding/xml"
	"encoding/json"
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	Header  Header `xml:"header"`
	Body    Body `xml:"body"`
}

type Header struct {
	ResultCode string `xml:"resultCode"`
	ResultMsg  string `xml:"resultMsg"`
}

type Body struct {
	Items      []Item `xml:"items>item"`
	NumOfRows  int `xml:"numOfRows"`
	PageNo     int `xml:"pageNo"`
	TotalCount int `xml:"totalCount"`
}

type Item struct {
	DutyName   string `xml:"dutyName"`   //기관명
	PostCdn1   string `xml:"postCdn1"`   //우편번호1
	PostCdn2   string `xml:"postCdn2"`   //우편번호2
	DutyAddr   string `xml:"dutyAddr"`   //주소
	DutyTel1   string `xml:"dutyTel1"`   //대표전화1
	Wgs84Lan   string `xml:"wgs84Lon"`   //병원경도
	Wgs84Lat   string `xml:"wgs84Lat"`   //병원위도
	DutyEtc    string `xml:"dutyEtc"`    //비고
	DutyInf    string `xml:"dutyInf"`    //기관설명
	DutyMapimg string `xml:"dutyMapimg"` //간이약도

	DutyTime1c string `xml:"dutyTime1c"` //진료시간(월요일)C
	DutyTime2c string `xml:"dutyTime2c"` //진료시간(화요일)C
	DutyTime3c string `xml:"dutyTime3c"` //진료시간(수요일)C
	DutyTime4c string `xml:"dutyTime4c"` //진료시간(목요일)C
	DutyTime5c string `xml:"dutyTime5c"` //진료시간(금요일)C
	DutyTime6c string `xml:"dutyTime6c"` //진료시간(토요일)C
	DutyTime7c string `xml:"dutyTime7c"` //진료시간(일요일)C
	DutyTime8c string `xml:"dutyTime8c"` //진료시간(공휴일)C
	DutyTime1s string `xml:"dutyTime1s"` //진료시간(월요일)S
	DutyTime2s string `xml:"dutyTime2s"` //진료시간(월\화요일)S
	DutyTime3s string `xml:"dutyTime3s"` //진료시간(수요일)S
	DutyTime4s string `xml:"dutyTime4s"` //진료시간(목요일)S
	DutyTime5s string `xml:"dutyTime5s"` //진료시간(금요일)S
	DutyTime6s string `xml:"dutyTime6s"` //진료시간(토요일)S
	DutyTime7s string `xml:"dutyTime7s"` //진료시간(일요일)S
}

//PORT=8080 bin/soo
type Config struct {
	Environment string `env:"key=ENVIRONMENT default=development"`
	Port        string `env:"key=PORT default=9000`
	EnableCors  string `env:"key=ENABLE_CORS default=false`
}

var (
	renderer *render.Render
	config *Config
)

func init() {
	var option render.Options
	config = &Config{}
	if err := env.Process(config); err != nil {
		fmt.Println(err)
	}
	if config.Environment == "development" {
		option.IndentJSON = true
	}

	//https://github.com/unrolled/render
	option.Directory = "public"
	option.Extensions = []string{".tmpl", ".html"}

	renderer = render.New(option);
}

func index(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index", map[string]interface{}{"host": r.Host})
}

func App() http.Handler {
	//router
	r := mux.NewRouter()

	//r.HandleFunc("/", index)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "index", map[string]interface{}{"host": r.Host})
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		var Url *url.URL
		Url, err := url.Parse("http://openapi.e-gen.or.kr")

		if err != nil {
			panic("boom")
		}

		Url.Path += "/openapi/service/rest/ErmctInsttInfoInqireService/getParmacyListInfoInqire"
		rawQuery := "?ServiceKey=LHpVSLmmvdL6ZAlqOvkmQMwYHK5i6O2%2Bf4ASrHDzDH6f9UQ4YVOROW2NwsLwxQmd%2F2LpDRr0HFKVNbrPMnbW3A%3D%3D&Q0=%EC%84%9C%EC%9A%B8%ED%8A%B9%EB%B3%84%EC%8B%9C&Q1=%EC%84%B1%EB%B6%81%EA%B5%AC&QT=8&ORD=ADDR&numOfRows=999&pageSize=999&pageNo=1&startPage=1"

		res, _ := http.Get(Url.String() + rawQuery)
		fmt.Printf("Not Encoded URL is %q%s\n", Url.String(), rawQuery)

		defer res.Body.Close()

		var response Response
		if err := xml.NewDecoder(res.Body).Decode(&response); err != nil {
			fmt.Printf("error is : %v", err)
			return
		}

		fmt.Printf("%#v\n", response)
		renderer.XML(w, http.StatusOK, response)
	})

	//middleware
	//return New(NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
	n := negroni.Classic() // Includes some default middlewares

	//enable CORS
	if config.EnableCors == "true" {
		c := cors.New(cors.Options{})
		n.Use(c)
	}

	// add handler
	n.UseHandler(r)

	return n
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	//listen
	http.ListenAndServe(":3000", App())
}
