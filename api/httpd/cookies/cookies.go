package cookies

import (
	"log"
	"net/http"
	"time"
)

// cookieMaxAge is how long a cookie can be active for without having to sign back into the application.
const cookieMaxAge = 3600

const loginCookie = "UFPMP-Login"

// SetLoginCookie creates a login cookie that keeps track of if and when a user logged into the application.
func SetLoginCookie(w http.ResponseWriter, username string) {
	//Make a login cookie with the user's username.
	cookie := http.Cookie{
		Name:     loginCookie,
		Value:    username,
		Path:     "/",
		MaxAge:   cookieMaxAge,
		Expires:  time.Now().Add(time.Second * cookieMaxAge),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

// GetLoginCookie returns the login cookie.
func GetLoginCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(loginCookie)
	if err == http.ErrNoCookie {
		log.Printf("No login cookie found!")
		return nil
	} else if err != nil {
		log.Printf("Error retreiving login cookie: %v", err.Error())
	}

	return cookie
}

// ExpireLoginCookie invalidates the login cookie. This is used when the user logs out or the cookie has expired.
func ExpireLoginCookie(w http.ResponseWriter, r *http.Request) {
	//Make an expired login cookie with the user's username.
	cookie := http.Cookie{
		Name:     loginCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Now().Add(time.Second * -100),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	log.Printf("Login cookie is now expired. User must sign in again to access account-specific features.")

	log.Printf("Cookie details:\n%v\n", GetLoginCookie(r).String())
}
