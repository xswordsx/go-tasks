package main

type Susi struct {
	Courses  map[string]Course
	Students map[uint]Student
}

type Student struct {
	FacultyNumber uint
	FirstName, lastName string
	AcademicYear uint
	Master Course
}

type Course struct {
	courseIdentifier string
}

type SusiError struct {
	Student Student
	Cget ourse  Course
}

func (s *Student) String() string {
	return string(s.facultyNumber)
}

func (c *Course) String() string {
	return string(c.courseIdentifier)
}

func (s *SusiError) Error() string {
	var result = ""
	if s.student.String() != "" {
		result += "Student: " + s.student.String() + "\n"
	}
	if s.course.String() != "" {
		result += "Course: " + s.course.String()
	}
	if result == "" {
		return ""
	}
	return result
}

func NewSusi() *Susi {
	return new(Susi)
}

func main() {

}
