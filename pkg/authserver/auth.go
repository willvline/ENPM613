package authserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/mongo"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	PORT   = "1337"
	SECRET = "42isTheAnswer"
)

var TokenPool = map[string]bool{}

type JWTData struct {
	// Standard claims are the standard jwt claims from the IETF standard
	// https://tools.ietf.org/html/rfc7519
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}

// type Account struct {
// 	Email    string  `json:"email"`
// 	Balance  float64 `json:"balance"`
// 	Currency string  `json:"currency"`
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	mux.HandleFunc("/login", login)
// 	mux.HandleFunc("/account", account)

// 	handler := cors.Default().Handler(mux)

// 	log.Println("Listening for connections on port: ", PORT)
// 	log.Fatal(http.ListenAndServe(":"+PORT, handler))
// }

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}

	var userData map[string]string
	json.Unmarshal(body, &userData)

	// Demo - in real case scenario you'd check this against your database
	student := mongo.Student{}
	if _, ok := userData["user_name"]; ok {
		student.UserName = userData["user_name"]
	}
	if _, ok := userData["pass_word"]; ok {
		student.PassWord = userData["pass_word"]
	}
	// Validate the account against Database
	valid := mongo.Authenticate(student)

	if valid {
		log.Println(student)
		students, err := mongo.GetStudent(student)
		student = students[0]
		claims := JWTData{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},

			CustomClaims: map[string]string{
				"user_name":  student.UserName,
				"student_id": student.StudentID.Hex(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(SECRET))
		TokenPool[tokenString] = true
		if err != nil {
			log.Println(err)
			http.Error(w, "Login failed!", http.StatusUnauthorized)
			return
		}

		jsondata, err := json.Marshal(struct {
			Token string `json:"token"`
		}{
			tokenString,
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Login failed!", http.StatusUnauthorized)
			return
		}

		w.Header().Add("Set-Cookie", tokenString)

		respondWithJSON(w, r, http.StatusOK, jsondata)
	} else {
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}
}

func Authorize(w http.ResponseWriter, r *http.Request) (bool, string, int) {
	authToken := r.Header.Get("Cookie")

	jwtToken := authToken

	if _, ok := TokenPool[jwtToken]; !ok {
		return false, "Invalid Token!", http.StatusUnauthorized
	}

	claims, err := jwt.ParseWithClaims(jwtToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		log.Println(err)
		return false, "Request failed!", http.StatusUnauthorized
	}

	data := claims.Claims.(*JWTData)

	if data.ExpiresAt < time.Now().Unix() {
		return false, "Token expired!", http.StatusUnauthorized
	}
	return true, "Authorized!", http.StatusAccepted

}
func Account(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Cookie")

	jwtToken := authToken

	claims, err := jwt.ParseWithClaims(jwtToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
	}

	data := claims.Claims.(*JWTData)

	student := mongo.Student{}
	student.UserName = data.CustomClaims["user_name"]
	student.StudentID = bson.ObjectIdHex(data.CustomClaims["student_id"])

	// fetch some data based on the userID and then send that data back to the user in JSON format
	students, err := mongo.GetStudent(student)
	log.Println(students)
	json, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
	}
	respondWithJSON(w, r, http.StatusOK, json)
}

func respondWithError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	respondWithJSON(w, r, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	// if origin := r.Header.Get("Origin"); origin != "" {
	// 	w.Header().Set("Access-Control-Allow-Origin", origin)
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	w.Header().Set("Access-Control-Allow-Headers",
	// 		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// }
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

	w.WriteHeader(code)
	w.Write(response)
}
