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
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	Header  Header `xml:"header"`
}

type Header struct {
	ResultCode string `xml:"resultCode"`
	ResultMsg  string `xml:"resultMsg"`
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
	option.Directory = "templates"
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
		req, _ := http.NewRequest("GET", "http://openapi.e-gen.or.kr", nil)

		q := req.URL.Query()
		q.Add("ServiceKey", "LHpVSLmmvdL6ZAlqOvkmQMwYHK5i6O2%2Bf4ASrHDzDH6f9UQ4YVOROW2NwsLwxQmd%2F2LpDRr0HFKVNbrPMnbW3A%3D%3D")
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		var Url *url.URL
		Url, err := url.Parse("http://openapi.e-gen.or.kr")

		if err != nil {
			panic("boom")
		}

		Url.Path += "/openapi/service/rest/ErmctInsttInfoInqireService/getParmacyListInfoInqire"
		parameters := url.Values{}
		parameters.Add("ServiceKey", "LHpVSLmmvdL6ZAlqOvkmQMwYHK5i6O2%2Bf4ASrHDzDH6f9UQ4YVOROW2NwsLwxQmd%2F2LpDRr0HFKVNbrPMnbW3A%3D%3D")
		//parameters.Add("Q0", "서울특별시")
		//parameters.Add("Q1", "-")
		//parameters.Add("QT", "월~일요일, 공휴일 : 1~8")
		//parameters.Add("ORD", "병원명정렬기준(주소:ADDR/기관명:NAME)")
		//parameters.Add("numOfRows", "999")
		//parameters.Add("pageNo", "1")

		Url.RawQuery = parameters.Encode()

		/*
		serviceKey := url.QueryEscape("LHpVSLmmvdL6ZAlqOvkmQMwYHK5i6O2%2Bf4ASrHDzDH6f9UQ4YVOROW2NwsLwxQmd%2F2LpDRr0HFKVNbrPMnbW3A%3D%3D")
		hpid := url.QueryEscape("N0002117")
		rows := url.QueryEscape("999")
		pageSize := url.QueryEscape("999")
		pageNo := url.QueryEscape("1")
		startPage := url.QueryEscape("1")
		requestUrl := fmt.Sprintf("http://openapi.e-gen.or.kr/openapi/service/rest/ErmctInsttInfoInqireService/getParmacyBassInfoInqire?ServiceKey=%s&HPID=%s&numOfRows=%s&pageSize=%s&pageNo=%s&startPage=%s",
			serviceKey, hpid, rows, pageSize, pageNo, startPage)
		*/

		res, _ := http.Get(Url.String())
		fmt.Printf("Encoded URL is %q\n", Url.String())

		defer res.Body.Close()

		var response Response
		if err := xml.NewDecoder(res.Body).Decode(&response); err != nil {
			fmt.Printf("error is : %v", err)

			return
		}

		fmt.Printf("%#v\n", response)

		/*
		Xml파싱 방법
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		if err := xml.Unmarshal(bodyBytes, &response); err != nil {
			log.Fatalln(err)
		}
		*/

		renderer.Text(w, http.StatusOK, "pong")
	})

	//middleware
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

func main() {
	//listen
	http.ListenAndServe(":3000", App())
}
