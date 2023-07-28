package routes

import (
	m "AELS/middleware"
	ch "AELS/middleware/chain"
	"AELS/routes/api/auth"
	"AELS/routes/api/chat"
	"AELS/routes/api/chat/chatid"
	"AELS/routes/api/chat/chatid/complete"
	"AELS/routes/api/login"
	"AELS/routes/api/logout"
	"AELS/routes/api/user"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.StrictSlash(true)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server Ready to Accept Connections\n"))
	}).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	api.Handle("/login", login.LoginUser()).Methods("POST")
	api.Handle("/logout", logout.LogoutUser()).Methods("POST")
	api.Handle("/auth", ch.New(m.Auth).Then(auth.CheckAuth())).Methods("GET")

	api.Handle("/user", ch.New(m.Auth).Then(user.GetUser())).Methods("GET")
	api.Handle("/user", user.CreateUser()).Methods("POST") // Built in Key Auth

	api.Handle("/chat", ch.New(m.Auth).Then(chat.ChatList())).Methods("GET")
	api.Handle("/chat", ch.New(m.Auth).Then(chat.NewChat())).Methods("POST")
	api.Handle("/chat", ch.New(m.Auth, m.ChatOwnership).Then(chat.UpdateChat())).Methods("PATCH")
	api.Handle("/chat", ch.New(m.Auth, m.ChatOwnership).Then(chat.DeleteChat())).Methods("DELETE")

	api.Handle("/chat/{chatid:[0-9]+}", ch.New(m.Auth, m.ChatOwnership).Then(chatid.ListMessages())).Methods("GET")
	api.Handle("/chat/{chatid:[0-9]+}", ch.New(m.Auth, m.ChatOwnership).Then(chatid.NewMessage())).Methods("POST")
	api.Handle("/chat/{chatid:[0-9]+}/complete", ch.New(m.Auth, m.ChatOwnership).Then(complete.CompleteChat())).Methods("POST")

	/*
		app.Static("/", "dist")

		app.Get("/*", func(c *fiber.Ctx) error {
			return c.SendFile("./dist/index.html")
		})
	*/
}
