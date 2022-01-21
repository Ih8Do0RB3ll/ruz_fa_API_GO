package ruz_fa_API_GO

type TimeTable []struct {
	Auditorium                string      `json:"auditorium"`
	AuditoriumAmount          int         `json:"auditoriumAmount"`
	AuditoriumOid             int         `json:"auditoriumOid"`
	Author                    string      `json:"author"`
	BeginLesson               string      `json:"beginLesson"`
	Building                  string      `json:"building"`
	BuildingGid               int         `json:"buildingGid"`
	BuildingOid               int         `json:"buildingOid"`
	ContentOfLoadOid          int         `json:"contentOfLoadOid"`
	ContentOfLoadUID          interface{} `json:"contentOfLoadUID"`
	ContentTableOfLessonsName string      `json:"contentTableOfLessonsName"`
	ContentTableOfLessonsOid  int         `json:"contentTableOfLessonsOid"`
	Createddate               string      `json:"createddate"`
	Date                      string      `json:"date"`
	DateOfNest                string      `json:"dateOfNest"`
	DayOfWeek                 int         `json:"dayOfWeek"`
	DayOfWeekString           string      `json:"dayOfWeekString"`
	DetailInfo                string      `json:"detailInfo"`
	Discipline                string      `json:"discipline"`
	DisciplineOid             int         `json:"disciplineOid"`
	Disciplineinplan          interface{} `json:"disciplineinplan"`
	Disciplinetypeload        int         `json:"disciplinetypeload"`
	Duration                  float64     `json:"duration"`
	EndLesson                 string      `json:"endLesson"`
	Group                     string      `json:"group"`
	GroupOid                  int         `json:"groupOid"`
	GroupUID                  string      `json:"groupUID"`
	GroupFacultyoid           int         `json:"group_facultyoid"`
	Hideincapacity            int         `json:"hideincapacity"`
	IsBan                     bool        `json:"isBan"`
	KindOfWork                string      `json:"kindOfWork"`
	KindOfWorkComplexity      int         `json:"kindOfWorkComplexity"`
	KindOfWorkOid             int         `json:"kindOfWorkOid"`
	KindOfWorkUID             string      `json:"kindOfWorkUid"`
	Lecturer                  string      `json:"lecturer"`
	LecturerCustomUID         string      `json:"lecturerCustomUID"`
	LecturerEmail             string      `json:"lecturerEmail"`
	LecturerOid               int         `json:"lecturerOid"`
	LecturerUID               string      `json:"lecturerUID"`
	LecturerRank              string      `json:"lecturer_rank"`
	LecturerTitle             string      `json:"lecturer_title"`
	LessonNumberEnd           int         `json:"lessonNumberEnd"`
	LessonNumberStart         int         `json:"lessonNumberStart"`
	LessonOid                 int         `json:"lessonOid"`
	ListOfLecturers           []struct {
		Lecturer          string `json:"lecturer"`
		LecturerCustomUID string `json:"lecturerCustomUID"`
		LecturerEmail     string `json:"lecturerEmail"`
		LecturerOid       int    `json:"lecturerOid"`
		LecturerUID       string `json:"lecturerUID"`
		LecturerRank      string `json:"lecturer_rank"`
		LecturerTitle     string `json:"lecturer_title"`
	} `json:"listOfLecturers"`
	Modifieddate       string      `json:"modifieddate"`
	Note               string      `json:"note"`
	NoteDescription    string      `json:"note_description"`
	Parentschedule     string      `json:"parentschedule"`
	Replaces           interface{} `json:"replaces"`
	Stream             interface{} `json:"stream"`
	StreamOid          int         `json:"streamOid"`
	StreamFacultyoid   int         `json:"stream_facultyoid"`
	SubGroup           interface{} `json:"subGroup"`
	SubGroupOid        int         `json:"subGroupOid"`
	SubgroupFacultyoid int         `json:"subgroup_facultyoid"`
	TableofLessonsName string      `json:"tableofLessonsName"`
	TableofLessonsOid  int         `json:"tableofLessonsOid"`
	URL1               string      `json:"url1"`
	URL1Description    string      `json:"url1_description"`
	URL2               interface{} `json:"url2"`
	URL2Description    interface{} `json:"url2_description"`
}

type Group []struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Teacher []struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
