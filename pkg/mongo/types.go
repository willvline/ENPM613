package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	StudentID     bson.ObjectId                     `bson:"_id"       json:"student_id"`
	UserName      string                            `bson:"user_name" json:"user_name"`
	PassWord      string                            `bson:"pass_word" json:"pass_word"`
	Email         string                            `bson:"email"     json:"email"`
	Grades        map[string]string                 `bson:"grades"    json:"grades"`
	CourseRecords map[string]map[string]interface{} `bson:"course_records"   json:"course_records"`
	LastName      string                            `bson:"first_name"       json:"first_name"`
	FirstName     string                            `bson:"last_name"       json:"last_name"`
}

type Admin struct {
	AdminID   bson.ObjectId `bson:"_id" json:"admin_id"`
	UserName  string        `bson:"user_name" json:"user_name"`
	PassWord  string        `bson:"pass_word" json:"pass_word"`
	Email     string        `bson:"email" json:"email"`
	Privilege string        `bson:"privilege" json:"privilege"`
}

type Course struct {
	CourseID   bson.ObjectId     `bson:"_id" json:"course_id"`
	CourseName string            `bson:"course_name" json:"course_name"`
	Syllabus   string            `bson:"syllabus" json:"syllabus"`
	StudyLevel string            `bson:"study_level" json:"study_level"`
	ChapterURL map[string]string `bson:"chapter_url" json:"chapter_url"`
	Quiz       []Question        `bson:"quiz" json:"quiz"`
	//DiscussionBoard []Comment         `bson:"discussion_board" json:"discussion_board"`
}

type Question struct {
	QuestionID string `bson:"question_id" json:"question_id"`
	Statement  string `bson:"statement"  json:"statement"`
	Answer     string `bson:"answer"     json:"answer"`
}

type Comment struct {
	CommentID  bson.ObjectId `bson:"_id"           json:"comment_id"`
	PosterName string        `bson:"poster_name"   json:"poster_name"`
	PostDate   string        `bson:"post_date"     json:"post_date"`
	Content    string        `bson:"content"       json:"content"`
}
