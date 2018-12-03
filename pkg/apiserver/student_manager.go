package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mongo "github.com/ENPM613/HOLMS/pkg/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Health(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, http.StatusOK, "Welcome")
}

func PostStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	student := mongo.Student{}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	student.StudentID = bson.NewObjectId()

	students, err := mongo.PostStudent(student)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusCreated, students)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Header)
	student := mongo.Student{}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	valid := mongo.Authenticate(student)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: student.UserName, Value: student.StudentID.String(), Expires: expiration}
	http.SetCookie(w, &cookie)
	// log.Println("after setting cookie, w.Header(): ", w.Header())
	// log.Println(cookie)
	w.Header().Add("Set-Cookie", student.UserName)
	if !valid {
		respondWithError(w, r, http.StatusBadRequest, "Username and password don't match!")
	} else {
		respondWithJSON(w, r, http.StatusCreated, valid)
		log.Println(w.Header())
	}

}
func GetStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("someone is using GetStudent")
	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	respondWithJSON(w, r, http.StatusOK, students)
}

func ListAllStudent(w http.ResponseWriter, r *http.Request) {

	students, err := mongo.ListAllStudent()
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	respondWithJSON(w, r, http.StatusOK, students)
}

func getStudent(r *http.Request) ([]mongo.Student, error) {
	student := mongo.Student{}
	studentid := r.FormValue("student_id")
	username := r.FormValue("user_name")
	password := r.FormValue("pass_word")
	if studentid != "" {
		student.StudentID = bson.ObjectIdHex(studentid)
	}
	if username != "" {
		student.UserName = username
	}
	if password != "" {
		student.PassWord = password
	}
	students, err := mongo.GetStudent(student)
	return students, err
}

func PatchStudent(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var student mongo.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := mongo.PatchStudent(student); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var student mongo.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := mongo.DeleteStudent(student); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
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
