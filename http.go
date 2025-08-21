package cron_manager

import (
	"html/template"
	"log"
	"net/http"
)

func Handler() http.Handler {
	log.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeee1")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeee2")
		action := r.URL.Query().Get("action")
		jobId := r.URL.Query().Get("job_id")

		if action == "run" {
			// todo check if jobId is not empty
			if err := manager.runJobNow(jobId); err != nil {
				http.Error(w, err.Error(), 500)
				log.Println("runJobNow error:", err)
				return
			}

			w.WriteHeader(200)
			w.Write([]byte("success"))
			return
		}

		tmpl, err := template.ParseFiles("./front/index.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("template error:", err)
			return
		}

		if err = tmpl.Execute(w, manager.retrieveJobsData()); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Println("execute error:", err)

		}
	})

}

//data := MainPage{
//	TotalJobs: 3,
//	Running:   0,
//	Stopped:   0,
//	Jobs: []JobPage{{
//		Id:      "1",
//		Name:    "Job Number ONE",
//		LastRun: "Diruz",
//		NextRun: "Pagoh",
//	}, {
//		Id:      "1",
//		Name:    "Job Number ONE",
//		LastRun: "Diruz",
//		NextRun: "Pagoh",
//	}, {
//		Id:      "1",
//		Name:    "Job Number ONE",
//		LastRun: "Diruz",
//		NextRun: "Pagoh",
//	}},
//}
