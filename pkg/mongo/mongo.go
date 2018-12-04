package mongo

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	STUDENTCOLLECTION = "student"
	COURSECOLLECTION  = "course"
	ADMINCOLLECTION   = "admin"
	COMMENTCOLLECTION = "comment"
)

// Establish a connection to database
func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.Server)
	//session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Println(err)
	}
	log.Println("MongoDB is connected successfully!")
	db = session.DB(m.Database)
}

// AddStudent create a new student in mongo
func PostStudent(student Student) ([]Student, error) {

	students, err := GetStudentByName(student)
	if err != nil {
		return students, err
	}
	if len(students) != 0 {
		// found duplication, need to de-dup
		err = fmt.Errorf("student %s already exsits", student.UserName)
		return students, err
	}
	err = db.C(STUDENTCOLLECTION).Insert(student)
	if err != nil {
		log.Println(err)
	}
	students, err = GetStudent(student)
	return students, err
}

func ListAllStudent() ([]Student, error) {
	students := []Student{}
	err := db.C(STUDENTCOLLECTION).Find(bson.M{}).All(&students)
	return students, err
}

// GetStudent search student in mongo
func GetStudent(student Student) ([]Student, error) {
	if !student.StudentID.Valid() {
		return GetStudentByName(student)
	}
	students := []Student{}
	err := db.C(STUDENTCOLLECTION).FindId(student.StudentID).All(&students)
	return students, err
}

func GetStudentByName(student Student) ([]Student, error) {

	students := []Student{}
	log.Println("input student:", student)
	err := db.C(STUDENTCOLLECTION).Find(bson.M{"user_name": student.UserName}).All(&students)
	log.Println("found student:", students)
	return students, err
}

func Authenticate(student Student) bool {
	foundStudent, err := GetStudent(student)
	if err != nil || len(foundStudent) == 0 {
		return false
	}
	for _, stu := range foundStudent {
		if stu.PassWord != student.PassWord {
			return false
		}
	}
	return true
}

// PatchStudent update the whole student structure
func PatchStudent(student Student) error {

	_, err := db.C(STUDENTCOLLECTION).UpsertId(student.StudentID, student)
	return err
}

// DeleteStudent delete the student by student_id
func DeleteStudent(student Student) error {
	err := db.C(STUDENTCOLLECTION).RemoveId(student.StudentID)
	return err
}

// AddCourse create a new student in mongo
func PostCourse(course Course) ([]Course, error) {

	courses, err := GetCourse(course)
	if err != nil {
		return courses, err
	}
	if len(courses) != 0 {
		// found duplication, need to de-dup
		err = fmt.Errorf("student %s already exsits", course.CourseName)
		return courses, err
	}
	err = db.C(COURSECOLLECTION).Insert(course)
	courses, err = GetCourse(course)
	return courses, err
}

// GetCourse search student in mongo
func GetCourse(course Course) ([]Course, error) {

	courses := []Course{}
	err := db.C(COURSECOLLECTION).FindId(course.CourseID).All(courses)
	return courses, err
}

// PatchCourse update the whole student structure
func PatchCourse(course Course) error {

	_, err := db.C(COURSECOLLECTION).UpsertId(course.CourseID, course)
	return err
}

// DeleteCourse delete the student by student_id
func DeleteCourse(course Course) error {
	err := db.C(COURSECOLLECTION).RemoveId(course.CourseID)
	return err
}

// AddAdmin create a new student in mongo
func PostAdmin(admin Admin) ([]Admin, error) {

	admins, err := GetAdmin(admin)
	if err != nil {
		return admins, err
	}
	if len(admins) != 0 {
		// found duplication, need to de-dup
		err = fmt.Errorf("student %s already exsits", admin.UserName)
		return admins, err
	}
	err = db.C(COURSECOLLECTION).Insert(admin)
	admins, err = GetAdmin(admin)
	return admins, err
}

// GetCourse search student in mongo
func GetAdmin(admin Admin) ([]Admin, error) {

	admins := []Admin{}
	err := db.C(ADMINCOLLECTION).FindId(admin.AdminID).All(admins)
	return admins, err
}

// PatchCourse update the whole student structure
func PatchAdmin(admin Admin) error {

	_, err := db.C(ADMINCOLLECTION).UpsertId(admin.AdminID, admin)
	return err
}

// DeleteCourse delete the student by student_id
func DeleteAdmin(admin Admin) error {
	err := db.C(ADMINCOLLECTION).RemoveId(admin.AdminID)
	return err
}

// AddStudent create a new student in mongo
func PostComment(comment Comment) ([]Comment, error) {

	err := db.C(COMMENTCOLLECTION).Insert(comment)
	if err != nil {
		log.Println(err)
	}

	comments, err := GetComment(comment)
	return comments, err
}

// GetStudent search student in mongo
func GetComment(comment Comment) ([]Comment, error) {
	comments := []Comment{}

	if !comment.CommentID.Valid() {
		return comments, nil
	}

	err := db.C(COMMENTCOLLECTION).FindId(comment.CommentID).All(&comments)
	return comments, err
}

func ListAllComment() ([]Comment, error) {
	comments := []Comment{}
	err := db.C(COMMENTCOLLECTION).Find(bson.M{}).All(&comments)
	return comments, err
}
