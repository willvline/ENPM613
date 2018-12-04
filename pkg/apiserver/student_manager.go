package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	mongo "github.com/Johnlovescoding/ENPM613/HOLMS/pkg/mongo"
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

	student := mongo.Student{}
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	valid := mongo.Authenticate(student)

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

	student, err := getStudentFromToken(r)
	students, err := mongo.GetStudent(student)
	return students, err
}

func PatchStudent(w http.ResponseWriter, r *http.Request) {

	log.Println("someone is using PatchStudent")
	defer r.Body.Close()

	tokenStudent, _ := getStudentFromToken(r)
	students, err := mongo.GetStudent(tokenStudent)
	if err != nil || len(students) < 1 {
		respondWithError(w, r, http.StatusBadRequest, "cann't find student")
		return
	}

	tokenStudent = students[0]

	userData, err := parseBody(r)
	log.Println("userData: ", userData)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if username, ok := userData["user_name"].(string); ok {
		tokenStudent.UserName = username
	}
	if password, ok := userData["pass_word"].(string); ok && password != "" {
		log.Println(password)
		tokenStudent.PassWord = password
	}
	if email, ok := userData["email"].(string); ok {
		tokenStudent.Email = email
	}
	if grades, ok := userData["grades"].(map[string]string); ok {
		tokenStudent.Grades = grades
	}
	if courserecords, ok := userData["course_records"].(map[string]map[string]bool); ok {
		tokenStudent.CourseRecords = courserecords
	}

	if lastname, ok := userData["last_name"].(string); ok {
		tokenStudent.LastName = lastname
	}
	if firstname, ok := userData["first_name"].(string); ok {
		tokenStudent.FirstName = firstname
	}

	log.Println("after merge:", tokenStudent)

	if err := mongo.PatchStudent(tokenStudent); err != nil {
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

	w.WriteHeader(code)
	w.Write(response)
}
