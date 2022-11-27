package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// STRUCT TEMPLATE
type Project struct {
	ProjectName         string
	ProjectStartDate    string
	ProjectEndDate      string
	ProjectDuration     string
	ProjectDescription  string
	ProjectTechnologies [4]string
}

// LOCAL DATABASE
var ProjectList = []Project{
	{
		ProjectName:         "Testing Doang",
		ProjectStartDate:    "01 October 2022",
		ProjectEndDate:      "01 December 2022",
		ProjectDuration:     "2 Months",
		ProjectDescription:  "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.",
		ProjectTechnologies: [4]string{"checked", "checked", "checked", "checked"},
	},
	{
		ProjectName:         "Lgi-lagi Testing Doang",
		ProjectStartDate:    "20 October 2022",
		ProjectEndDate:      "21 November 2022",
		ProjectDuration:     "1 Month",
		ProjectDescription:  "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.",
		ProjectTechnologies: [4]string{"checked", "", "", "checked"},
	},
}

func main() {
	route := mux.NewRouter()

	// route path folder for public folder
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/detail-project/{index}", ProjectDetail).Methods("GET")
	// ADD PROJECT
	route.HandleFunc("/project/add-project", AddProject).Methods("POST")
	// EDIT PROJECT
	route.HandleFunc("/edit-project/{index}", EditProject).Methods("GET")
	// DELETE PROJECT
	route.HandleFunc("/delete-project/{index}", DeleteProject).Methods("GET")

	fmt.Println("Server running on port 5005")
	http.ListenAndServe("localhost:5005", route)
}

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func project(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

// RENDER PROJECT DETAIL
func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		var renderDetail = Project{}
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		for i, data := range ProjectList {
			if index == i {
				renderDetail = Project{
					ProjectName:         data.ProjectName,
					ProjectStartDate:    data.ProjectStartDate,
					ProjectEndDate:      data.ProjectEndDate,
					ProjectDuration:     data.ProjectDuration,
					ProjectDescription:  data.ProjectDescription,
					ProjectTechnologies: data.ProjectTechnologies,
				}
			}
		}
		data := map[string]interface{}{
			"renderDetail": renderDetail,
		}
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	}
}

// CREATE PROJECT
func AddProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	} else {
		projectName := r.PostForm.Get("project-name")
		projectStartDate := r.PostForm.Get("date-start")
		projectEndDate := r.PostForm.Get("date-end")
		projectDescription := r.PostForm.Get("project-description")
		projectUseNodeJS := r.PostForm.Get("nodejs")
		projectUseReactJS := r.PostForm.Get("reactjs")
		projectUseGolang := r.PostForm.Get("golang")
		projectUseTypeScript := r.PostForm.Get("typescript")

		var newProject = Project{
			ProjectName:         projectName,
			ProjectStartDate:    FormatDate(projectStartDate),
			ProjectEndDate:      FormatDate(projectEndDate),
			ProjectDuration:     GetDuration(projectStartDate, projectEndDate),
			ProjectDescription:  projectDescription,
			ProjectTechnologies: [4]string{projectUseNodeJS, projectUseReactJS, projectUseGolang, projectUseTypeScript},
		}

		ProjectList = append(ProjectList, newProject)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// UPDATE PROJECT
func EditProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/edit-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		var editData = Project{}
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		for i, data := range ProjectList {
			if index == i {
				editData = Project{
					ProjectName:         data.ProjectName,
					ProjectStartDate:    ReturnDate(data.ProjectStartDate),
					ProjectEndDate:      ReturnDate(data.ProjectEndDate),
					ProjectDescription:  data.ProjectDescription,
					ProjectTechnologies: data.ProjectTechnologies,
				}
				ProjectList = append(ProjectList[:index], ProjectList[index+1:]...)
			}
		}
		data := map[string]interface{}{
			"editData": editData,
		}
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	}
}

// DELETE PROJECT
func DeleteProject(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	ProjectList = append(ProjectList[:index], ProjectList[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

// ADDITIONAL FUNCTION

// GET DURATION
func GetDuration(startDate string, endDate string) string {

	layout := "2006-01-02"

	date1, _ := time.Parse(layout, startDate)
	date2, _ := time.Parse(layout, endDate)

	margin := date2.Sub(date1).Hours() / 24
	var duration string

	if margin > 30 {
		if (margin / 30) <= 1 {
			duration = "1 Month"
		} else {
			duration = strconv.Itoa(int(margin)/30) + " Months"
		}
	} else {
		if margin <= 1 {
			duration = "1 Day"
		} else {
			duration = strconv.Itoa(int(margin)) + " Days"
		}
	}

	return duration
}

// CHANGE DATE FORMAT
func FormatDate(InputDate string) string {

	layout := "2006-01-02"
	t, _ := time.Parse(layout, InputDate)

	Formated := t.Format("02 January 2006")

	return Formated
}

// RETURN DATE FORMAT
func ReturnDate(InputDate string) string {

	layout := "02 January 2006"
	t, _ := time.Parse(layout, InputDate)

	Formated := t.Format("2006-01-02")

	return Formated
}
