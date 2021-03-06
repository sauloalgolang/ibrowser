package endpoints

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// router.HandleFunc(DATABASE_ENDPOINT+"/{database}/chromosome", endpoints.Chromosomes).Methods("GET")                                                         //.HeadersRegexp("Content-Type", "application/json")
// router.HandleFunc(DATABASE_ENDPOINT+"/{database}/chromosome/{chromosome}", endpoints.Chromosome).Methods("GET")                                             //.HeadersRegexp("Content-Type", "application/json")
// router.HandleFunc(DATABASE_ENDPOINT+"/{database}/chromosome/{chromosome}/summary", endpoints.ChromosomeSummary).Methods("GET")                              //.HeadersRegexp("Content-Type", "application/json")
// router.HandleFunc(DATABASE_ENDPOINT+"/{database}/chromosome/{chromosome}/summary/matrix", endpoints.ChromosomeSummaryMatrix).Methods("GET")                 //.HeadersRegexp("Content-Type", "application/json")
// router.HandleFunc(DATABASE_ENDPOINT+"/{database}/chromosome/{chromosome}/summary/table", endpoints.ChromosomeSummaryMatrixTable).Methods("GET")             //.HeadersRegexp("Content-Type", "application/json")

func Chromosomes(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Chromosomes %#v", r)

	params := mux.Vars(r)
	database := params["database"]

	db, ok := databases.GetChromosomes(database)

	if !ok {
		resp := Message(false, "fail")
		resp["data"] = "No such database: " + database
		Respond(w, resp)
		return
	}

	resp := Message(true, "success")
	resp["data"] = db

	Respond(w, resp)
}

func Chromosome(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Chromosome %#v", r)

	params := mux.Vars(r)
	database := params["database"]
	chromosome := params["chromosome"]

	db, ok := databases.GetChromosome(database, chromosome)

	if !ok {
		resp := Message(false, "fail")
		resp["data"] = "No such chromosome: " + chromosome + " in database " + database
		Respond(w, resp)
		return
	}

	resp := Message(true, "success")
	resp["data"] = db

	Respond(w, resp)
}

func ChromosomeSummary(w http.ResponseWriter, r *http.Request) {
	log.Tracef("ChromosomeSummary %#v", r)

	params := mux.Vars(r)
	database := params["database"]
	chromosome := params["chromosome"]

	db, ok := databases.GetChromosomeSummaryBlock(database, chromosome)

	if !ok {
		resp := Message(false, "fail")
		resp["data"] = "No such chromosome: " + chromosome + " in database " + database
		Respond(w, resp)
		return
	}

	resp := Message(true, "success")
	resp["data"] = db

	Respond(w, resp)
}

func ChromosomeSummaryMatrix(w http.ResponseWriter, r *http.Request) {
	log.Tracef("ChromosomeSummaryMatrix %#v", r)

	params := mux.Vars(r)
	database := params["database"]
	chromosome := params["chromosome"]

	db, ok := databases.GetChromosomeSummaryBlockMatrix(database, chromosome)

	if !ok {
		resp := Message(false, "fail")
		resp["data"] = "No such chromosome: " + chromosome + " in database " + database
		Respond(w, resp)
		return
	}

	resp := Message(true, "success")
	resp["data"] = db

	Respond(w, resp)
}

func ChromosomeSummaryMatrixTable(w http.ResponseWriter, r *http.Request) {
	log.Tracef("ChromosomeSummaryMatrixTable %#v", r)

	params := mux.Vars(r)
	database := params["database"]
	chromosome := params["chromosome"]

	db, ok := databases.GetChromosomeSummaryBlockMatrixTable(database, chromosome)

	if !ok {
		resp := Message(false, "fail")
		resp["data"] = "No such chromosome: " + chromosome + " in database " + database
		Respond(w, resp)
		return
	}

	resp := Message(true, "success")
	resp["data"] = db

	Respond(w, resp)
}
