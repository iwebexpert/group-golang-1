package main

import (
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type DB map[string]string
var UserDB map[string]DB

const (
	ARG_KEY = "key"
	ARG_VALUE = "value"
	COOKIE_KEY = "cookie"
)

/*func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./pages/index.html")
	if err != nil {
		log.Println(err)
	}
	w.Write(file)
}*/

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	userKey := CookieControl(w, r)

	key := chi.URLParam(r, ARG_KEY)

	DB, exists := UserDB[userKey]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	value, exists := DB[key]
	log.Println("Create", userKey, key, value)

	if exists {
		w.Write([]byte(value))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func PostIndexHandler(w http.ResponseWriter, r *http.Request){
	userKey := CookieControl(w, r)

	key := chi.URLParam(r, ARG_KEY)
	value := r.FormValue(ARG_VALUE)

	UserDB[userKey] = DB{}
	UserDB[userKey][key] = value

	log.Println("Create", userKey, key, value)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(value))
}

//middleware
func CookieControl (w http.ResponseWriter, r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_KEY)
	if cookie == nil {
		cookie = &http.Cookie{
			Name: COOKIE_KEY,
		}
	}

	userKey := cookie.Value

	if userKey != "" {
		return userKey
	}

	cookie.Value = uuid.Must(uuid.NewV4()).String()
	http.SetCookie(w, cookie)
	return cookie.Value
}

func main() {
	stopChan := make(chan os.Signal)
	UserDB = map[string]DB{}

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router){
		r.Get("/{key}", GetIndexHandler)
		r.Post("/{key}", PostIndexHandler)
	})

	go func() {
		log.Println("Server starts")
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)
	}()

	//ожидание сигнала на выключение сервера
	signal.Notify(stopChan, os.Kill, os.Interrupt)
	<- stopChan

	log.Println("Server stop")
}