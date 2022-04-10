package control

import (
	"fmt"
	"log"
	"time"

	// "strings"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// ---------------------------------------------------
// Book table contains the information for each book and language, editor and author

type Jperson struct {
	Id     uint32 `db:"id" bson:"id,omitempty"`
	Cuenta string `db:"cuenta" bson:"cuenta"`
	Uuid   string `db:"uuid" bson:"uuid,omitempty"`
	Nivel  uint32 `db:"nivel" bson:"nivel"`
	Email  string `db:"email" bson:"email"`
}

type User struct {
	Cuenta   string `json:"cuenta"`
	Password string `json:"password"`
}

type Server struct {
	Hostname string `json:"hostname"`
}

type BookZ struct {
	Title    string `db:"title" bson:"title"`
	Comment  string `db:"comment" bson:"comment"`
	Year     uint32 `db:"year" bson:"year"`
	Author   string `db:"author" bson:"author,omitempty"`
	Editor   string `db:"editor" bson:"editor,omitempty"`
	Language string `db:"language" bson:"language,omitempty"`
}

type JLanguage struct {
	Id          uint32        `db:"id" bson:"id,omitempty"`
	Name      string        `db:"name" bson:"name"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
}

type JEditor struct {
	Id          uint32        `db:"id" bson:"id,omitempty"`
	Name      string        `db:"name" bson:"name"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
}

type JAuthor struct {
	Id          uint32        `db:"id" bson:"id,omitempty"`
	Name      string        `db:"name" bson:"name"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
}


var (
	/*formato = "2006-01-02"
	// ErrCode is a config or an internal error
	ErrCode = errors.New("Sentencia Case en codigo no es correcta.")
	// ErrNoResult is a not results error
	ErrNoResult = errors.New("Result  no encontrado.")
	// ErrUnavailable is a database not available error
	ErrUnauthorized = errors.New("Usuario sin permiso para realizar esta operacion.")*/
	bookClient      http.Client
)

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Creando Cookie %s\n", err.Error())
	}
	bookClient = http.Client{
		Jar:     jar,
		Timeout: time.Second * 2,
	}
}

// getBody -> from server
func getBody(url string) (body []byte) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("getBody 1", url, err)
		log.Fatal(err)
	}
	res, getErr := bookClient.Do(req)
	if getErr != nil {
		fmt.Println("getBody 2", url, err)
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("getBody 3", url, err)
		log.Fatal(readErr)
	}

	return
}

// JLoginGET : get user id
func JLoginGET(server Server, user User) (id uint32) {
	cuenta := user.Cuenta
	passwd := user.Password
	encPass := base64.StdEncoding.EncodeToString([]byte(passwd))
	url := server.Hostname + "/jlogin/" + cuenta + "/" + encPass
	body := getBody(url)
	pers := Jperson{}
	jsonErr := json.Unmarshal(body, &pers)
	if jsonErr != nil {
		fmt.Println("JLoginGet", jsonErr)
		log.Fatal(jsonErr)
	}
	id = pers.Id
	return
}


func JBook(server Server, st string) (id uint32) {
	//cuenta := user.Cuenta
	//passwd := user.Password
	//encPass := base64.StdEncoding.EncodeToString([]byte(passwd))
	url := server.Hostname + "/biblos/jbook/" + st
	body := getBody(url)
	pers := BookZ{}
	jsonErr := json.Unmarshal(body, &pers)
	if jsonErr != nil {
		fmt.Println("JBook", jsonErr)
		log.Fatal(jsonErr)
	}
	fmt.Println(pers.Title)
	//id = pers.Title
	return
}

func JLang(server Server, st string) (id uint32) {
	//cuenta := user.Cuenta
	//passwd := user.Password
	//encPass := base64.StdEncoding.EncodeToString([]byte(passwd))
	url := server.Hostname + "/biblos/jlang/" + st
	body := getBody(url)
	pers := JLanguage{}
	jsonErr := json.Unmarshal(body, &pers)
	if jsonErr != nil {
		fmt.Println("JLang", jsonErr)
		log.Fatal(jsonErr)
	}
	id = pers.Id
	return
}

func JAuth(server Server, st string) (id uint32) {
	//cuenta := user.Cuenta
	//passwd := user.Password
	//encPass := base64.StdEncoding.EncodeToString([]byte(passwd))
	url := server.Hostname + "/biblos/jauthor/" + st
	body := getBody(url)
	pers := JAuthor{}
	jsonErr := json.Unmarshal(body, &pers)
	if jsonErr != nil {
		fmt.Println("JAuth", jsonErr)
		log.Fatal(jsonErr)
	}
	id = pers.Id
	return
}


func JEdit(server Server, st string) (id uint32) {
	//cuenta := user.Cuenta
	//passwd := user.Password
	//encPass := base64.StdEncoding.EncodeToString([]byte(passwd))
	url := server.Hostname + "/biblos/jeditor/" + st
	body := getBody(url)
	pers := JEditor{}
	jsonErr := json.Unmarshal(body, &pers)
	if jsonErr != nil {
		fmt.Println("JEdit", jsonErr)
		log.Fatal(jsonErr)
	}
	id = pers.Id
	return
}
