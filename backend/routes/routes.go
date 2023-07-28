package routes

import (
	m "AELS/middleware"
	ch "AELS/middleware/chain"
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

	r.Handle("/api/login", login.LoginUser()).Methods("POST")
	r.Handle("/api/logout", logout.LogoutUser()).Methods("POST")
	r.Handle("/api/auth", ch.New(m.Auth).Then(user.GetUser())).Methods("GET")

	r.Handle("/api/user", ch.New(m.Auth).Then(user.GetUser())).Methods("GET")
	r.Handle("/api/user", user.CreateUser()).Methods("POST") // Built in Key Auth

	r.Handle("/api/chat", chat.ChatList()).Methods("GET")
	r.Handle("/api/chat", chat.NewChat()).Methods("POST")
	r.Handle("/api/chat", ch.New(m.Auth, m.ChatOwnership).Then(chat.UpdateChat())).Methods("PATCH")
	r.Handle("/api/chat", ch.New(m.Auth, m.ChatOwnership).Then(chat.DeleteChat())).Methods("DELETE")

	r.Handle("/chat/{chatid:[0-9]+}", ch.New(m.Auth, m.ChatOwnership).Then(chatid.ListMessages())).Methods("GET")
	r.Handle("/chat/{chatid:[0-9]+}", ch.New(m.Auth, m.ChatOwnership).Then(chatid.NewMessage())).Methods("POST")
	r.Handle("/chat/{chatid:[0-9]+}/complete", ch.New(m.Auth, m.ChatOwnership).Then(complete.CompleteChat())).Methods("POST")

	/*
		app.Static("/", "dist")

		app.Get("/*", func(c *fiber.Ctx) error {
			return c.SendFile("./dist/index.html")
		})
	*/
}
