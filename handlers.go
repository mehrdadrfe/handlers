// handlers.go
package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mehrdadrfe/usermanager"
)

// HomeHandler handles the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, homePage, nil)
}

// AddUserHandler handles the addition of a new user.
func AddUserHandler(userManager *usermanager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(r.FormValue("userID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		userManager.AddUser(userID)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// ApplyPenaltyHandler handles the application of a penalty.
func ApplyPenaltyHandler(userManager *usermanager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(r.FormValue("userID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		penaltyDays, err := strconv.Atoi(r.FormValue("penaltyDays"))
		if err != nil {
			http.Error(w, "Invalid penalty days", http.StatusBadRequest)
			return
		}

		penaltyAmt, err := strconv.ParseFloat(r.FormValue("penaltyAmt"), 64)
		if err != nil {
			http.Error(w, "Invalid penalty amount", http.StatusBadRequest)
			return
		}

		userManager.ApplyPenalty(userID, penaltyDays, penaltyAmt)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// CheckPenaltyHandler handles the checking of penalties.
func CheckPenaltyHandler(userManager *usermanager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(r.FormValue("userID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user := userManager.CheckPenalty(userID)
		renderTemplate(w, checkPenaltyPage, user)
	}
}

// PayPenaltyHandler handles the payment of penalties.
func PayPenaltyHandler(userManager *usermanager.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(r.FormValue("userID"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		paymentAmt, err := strconv.ParseFloat(r.FormValue("paymentAmt"), 64)
		if err != nil {
			http.Error(w, "Invalid payment amount", http.StatusBadRequest)
			return
		}

		userManager.PayPenalty(userID, paymentAmt)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// renderTemplate renders the HTML templates.
func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
