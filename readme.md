![Imgur](http://i.imgur.com/S97Ecpr.png)

Native http mux perk

> Legend tells us of a man, a hero in a tortured land, where Señoritas lived in fear. Their lonely nights in deep despair, he was EL BURRO! (Hee-Haw, Hee-Haw). Across the fields, across the plains. He ran so fast he dodged the rain. He was El Burro! He hurried in to save the day, gun in hand, and thrice they say. He was strong like a Mule, he was stubborn like a Mule, he even kicked like a Mule, El Burro! (El Burro!). A man of equal soul they say. But some men more, it's just the way! He was El Burro! He was EL BURRO! EL BURRO!

note: this package could've been called `wunderwaffe`.. But `mulekick` seems most appropriate given ze choice.

As the name suggests, `mule-kick` gives your router (gorilla/mux in this case) power.

## Features

- Middleware
- Convenience `.get`, `.post`, ...
- Sub-routing, sub-middleware declarations

## Demo usage

```go
r := mulekick.NewRouter(mux.NewRouter(), corsMiddleware)
r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

r.Get("/ping", mulekick.PongHandler)

// Authentication
func(api mulekick.Router) {
	api.Post("/password", login)
	api.Post("/register", signup)
}(r.Group("/auth"))

// Limited to requests with sessions
api := mulekick.NewRouter(r.Router, corsMiddleware, secureMiddleware)

// Fetch user
api.Get("/user", getUser)
// Update user details
api.Post("/user", updateUser)
// Update user password
api.Post("/user/password", updateUserPassword)
```
