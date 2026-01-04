package handlers

import "net/http"

type User struct {
	Name     string
	Username string
	Password string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	session, ok := GetSession(r)

	if ok {
		data := map[string]interface{}{
			"Title": "Home",
			"User":  session,
		}
		Render(w, "home.html", data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	session, ok := GetSession(r)

	if ok {
		data := map[string]interface{}{
			"Title": "Dashboard",
			"User":  session,
		}
		Render(w, "dashboard.html", data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
