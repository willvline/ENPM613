package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	mongo "github.com/ENPM613/HOLMS/pkg/mongo"
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
	course.CourseID = bson.ObjectIdHex(params["course_id"])
	courses, err := mongo.GetCourse(course)
	return courses, err
}

func RegisterCourse(w http.ResponseWriter, r *http.Request) {

	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	student := students[0]

	params := mux.Vars(r)
	student.CourseRecords[params["course_id"]] = map[string]bool{}

	if err := mongo.PatchStudent(student); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}

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

func GetCourseRecord(w http.ResponseWriter, r *http.Request) {

	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	params := mux.Vars(r)
	var courseRecord []map[string]bool
	for _, student := range students {
		courseRecord = append(courseRecord, student.CourseRecords[params["course_id"]])
	}
	respondWithJSON(w, r, http.StatusOK, courseRecord)
}

func PatchCourseRecord(w http.ResponseWriter, r *http.Request) {
	students, err := getStudent(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid student ID")
		return
	}
	student := students[0]

	params := mux.Vars(r)
	student.CourseRecords[params["course_id"]][params["chapter_id"]] = true

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

func GetCourseComment(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	var comment []mongo.Comment
	for _, course := range courses {
		comment = course.DiscussionBoard
	}
	respondWithJSON(w, r, http.StatusOK, comment)
}

func PostCourseComment(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	course := courses[0]

	params := mux.Vars(r)

	comment := mongo.Comment{
		params["student_id"],
		time.Now().String(),
		params["content"],
	}
	course.DiscussionBoard = append(course.DiscussionBoard, comment)

	if err := mongo.PatchCourse(course); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}

func PatchCourseComment(w http.ResponseWriter, r *http.Request) {

	courses, err := getCourse(r)
	if err != nil {
		respondWithError(w, r, http.StatusBadRequest, "Invalid Course ID")
		return
	}
	course := courses[0]

	params := mux.Vars(r)

	for _, comment := range course.DiscussionBoard {
		if comment.PosterName == params["student_id"] && comment.PostDate == params["post_date"] {
			comment.Content = params["content"]
		}
	}

	if err := mongo.PatchCourse(course); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, r, http.StatusOK, map[string]string{"result": "success"})
}
