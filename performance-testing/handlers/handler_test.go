package handlers_test

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ruangguru/playground/performance-testing/api"
	"github.com/ruangguru/playground/performance-testing/handlers"
	"github.com/ruangguru/playground/performance-testing/repository"

	"log"
	"net/http"
	"net/http/httptest"
)

type Movie struct {
	Id      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

var router *gin.Engine
var _ = Describe("handlers", func() {
	When("the handlers implemented", func() {
		It("can add new movie with method POST", func() {

			w := httptest.NewRecorder()

			payload, err := json.Marshal(Movie{
				Episode: 1,
				Name:    "baru",
			})

			if err != nil {
				log.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/movie", strings.NewReader(string(payload)))
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
			// Expect(res.StatusCode).To(Equal(200))
		})

		It("can get all movies", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/movies", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			var movies []Movie
			err = json.NewDecoder(w.Body).Decode(&movies)
			if err != nil {
				log.Fatal(err)
			}
			Expect(w.Code).To(Equal(200))
			Expect(len(movies)).To(Equal(1))
		})

		It("can get a movie", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/movie/1", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)

			var movie Movie
			err = json.NewDecoder(w.Body).Decode(&movie)

			if err != nil {
				log.Fatal(err)
			}
			Expect(w.Code).To(Equal(200))
			Expect(movie.Id).To(Equal(1))
		})

		It("can test GET method", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
		})

		It("can test POST method", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("POST", "/", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
		})
	})
})

func setupRouter() *gin.Engine {
	engine := gin.New()
	repo := repository.NewRepo()
	svc := handlers.New(repo)
	serviceApi := api.New(engine, svc)
	serviceApi.InitAPI()
	return engine
}

func init() {
	router = setupRouter()
}