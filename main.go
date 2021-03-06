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

/**
	https://github.com/unrolled/render
*/

//PORT=8080 bin/soo
type Config struct {
	Environment string `env:"key=ENVIRONMENT default=development"`
	Port        string `env:"key=PORT default=9000"`
	EnableCors  string `env:"key=ENABLE_CORS default=false"`
}

const (
	ApiUrl = "http://openapi.e-gen.or.kr"
	PharmacyApiPath = "/openapi/service/rest/ErmctInsttInfoInqireService/getParmacyListInfoInqire"
	Dev = "development"
)

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
	if config.Environment == Dev {
		option.IndentJSON = true
	}

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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "index", map[string]interface{}{"host": r.Host})
	})

	r.HandleFunc("/api/parmacies", func(w http.ResponseWriter, r *http.Request) {
		var Url *url.URL
		Url, err := url.Parse(ApiUrl)

		if err != nil {
			panic("boom")
		}

		Url.Path += PharmacyApiPath
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

func main() {
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), App())
}
