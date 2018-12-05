package apiserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/authserver"
	mongo "github.com/Johnlovescoding/ENPM613/HOLMS/pkg/mongo"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func PostCourse(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	course := mongo.Course{}
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	courses, err := mongo.PostCourse(course)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusCreated, courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	respondWithJSON(w, r, http.StatusOK, courses)
}

func getCourse(r *http.Request) ([]mongo.Course, error) {

	params := mux.Vars(r)
	course := mongo.Course{}
	if id, ok := params["course_id"]; ok {
		course.CourseID = bson.ObjectIdHex(id)
	}
	if name, ok := params["course_name"]; ok {
		course.CourseName = name
	}
	courses, err := mongo.GetCourse(course)
	return courses, err
}

// func RegisterCourse(w http.ResponseWriter, r *http.Request) {

// 	student, err := getStudentFromToken(r)
// 	if err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
// 		return
// 	}
// 	bodyData, err := parseBody(r)
// 	if _, ok := bodyData["course_name"]; !ok || err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid payload")
// 		return
// 	}
// 	student.CourseRecords[bodyData["course_name"].(string)] = map[string]string{}

// 	if err := mongo.PatchStudent(student); err != nil {
// 		respondWithError(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
// }

func GetCourseSyllabus(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	var syllabus []string
	for _, course := range courses {
		syllabus = append(syllabus, course.Syllabus)
	}
	respondWithJSON(w, r, http.StatusOK, syllabus)
}

func GetCourseChapter(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	var chapter []map[string]string
	for _, course := range courses {
		chapter = append(chapter, course.ChapterURL)
	}
	respondWithJSON(w, r, http.StatusOK, chapter)
}

// func GetCourseRecord(w http.ResponseWriter, r *http.Request) {

// 	students, err := getStudent(r)
// 	if err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
// 		return
// 	}
// 	params := mux.Vars(r)
// 	var courseRecord []map[string]string
// 	for _, student := range students {
// 		courseRecord = append(courseRecord, student.CourseRecords[params["course_id"]])
// 	}
// 	respondWithJSON(w, r, http.StatusOK, courseRecord)
// }

func PatchCourseRecord(w http.ResponseWriter, r *http.Request) {
	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	student := students[0]

	params := mux.Vars(r)
	student.CourseRecords[params["course_id"]][params["chapter_id"]] = "true"

	if err := mongo.PatchStudent(student); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}

func PatchStudentGrade(w http.ResponseWriter, r *http.Request) {
	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	student := students[0]

	params := mux.Vars(r)
	student.Grades[params["course_id"]] = params["course_grade"]

	if err := mongo.PatchStudent(student); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}

func GetStudentGrade(w http.ResponseWriter, r *http.Request) {

	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	//params := mux.Vars(r)
	var grade []map[string]string
	for _, student := range students {
		grade = append(grade, student.Grades)
	}
	respondWithJSON(w, r, http.StatusOK, grade)
}

func GetCourseQuiz(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	var quiz []mongo.Question
	for _, course := range courses {
		quiz = course.Quiz
	}
	respondWithJSON(w, r, http.StatusOK, quiz)
}

// func GetCourseComment(w http.ResponseWriter, r *http.Request) {

// 	courses, err := getCourse(r)
// 	if err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
// 		return
// 	}
// 	var comment []mongo.Comment
// 	// for _, course := range courses {
// 	// 	comment = course.DiscussionBoard
// 	// }
// 	respondWithJSON(w, r, http.StatusOK, comment)
// }

// func PostCourseComment(w http.ResponseWriter, r *http.Request) {

// 	courses, err := getCourse(r)
// 	if err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
// 		return
// 	}
// 	course := courses[0]

// 	params := mux.Vars(r)

// 	// comment := mongo.Comment{
// 	// 	params["student_id"],
// 	// 	time.Now().String(),
// 	// 	params["content"],
// 	// }
// 	// course.DiscussionBoard = append(course.DiscussionBoard, comment)

// 	if err := mongo.PatchCourse(course); err != nil {
// 		respondWithError(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
// }

// func PatchCourseComment(w http.ResponseWriter, r *http.Request) {

// 	courses, err := getCourse(r)
// 	if err != nil {
// 		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
// 		return
// 	}
// 	course := courses[0]

// 	params := mux.Vars(r)

// 	// for _, comment := range course.DiscussionBoard {
// 	// 	if comment.PosterName == params["student_id"] && comment.PostDate == params["post_date"] {
// 	// 		comment.Content = params["content"]
// 	// 	}
// 	// }

// 	if err := mongo.PatchCourse(course); err != nil {
// 		respondWithError(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
// }

func getStudentFromToken(r *http.Request) (mongo.Student, error) {
	authToken := r.Header.Get("Cookie")
	jwtToken := authToken

	claims, err := jwt.ParseWithClaims(jwtToken, &authserver.JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(authserver.SECRET), nil
	})

	student := mongo.Student{}
	if err != nil {
		log.Println(err)
		return student, err
	}
	data := claims.Claims.(*authserver.JWTData)
	student.UserName = data.CustomClaims["user_name"]
	student.StudentID = bson.ObjectIdHex(data.CustomClaims["student_id"])

	return student, nil
}
func parseBody(r *http.Request) (map[string]interface{}, error) {
	userData := map[string]interface{}{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return userData, err
	}
	json.Unmarshal(body, &userData)
	return userData, nil
}

func ListAllComment(w http.ResponseWriter, r *http.Request) {

	comments, err := mongo.ListAllComment()
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Cann't list all comments")
	}
	respondWithJSON(w, r, http.StatusOK, comments)
}

func PostComment(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	comment := mongo.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		log.Println(comment)
		respondWithError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	comment.CommentID = bson.NewObjectId()
	comments, err := mongo.PostComment(comment)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, r, http.StatusOK, comments)
}
