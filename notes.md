## When do I consider myself "done" with learning Go?
- create __CRUD app__ in a day (using popular framework + ORM for Postgres)
    - webframework(s):
        - Gin
        - others?
    - ORM: Gorm
    - working with JSON, extensively
    - understand SDLC for Golang projects
    (e.g. go mod tidy, etc - how to build, how to deploy etc)
- can write more complex (business) logic as well
    - data transformations
        - Is Go any good here?
    - video streaming
        - pion
        -> do something with the stream:
            - store to disk
            - allow other devices to connect to stream (broadcast)
            - analyze?
- fluid and solid when it comes to syntax and implementation
    -> requires quite a bit of coding I assume
    -> however, quite a bit of coding comes naturally, and Go is not that far of Python actually


## Resources:
- Getting started tutorial: https://go.dev/doc/tutorial/getting-started
    - Golang does NOT have a "raise/except", but "panic/recover"
        -> those serve different purposes
        -> raise/except is quite common, and sometimes even used for "standard" program flow
        -> panic/recover is ONLY used in exceptional cases where we cannot continue
            -> otherwise, we return an error and expect the calling code to know how to handle that error!
    - everything is typed in Golang
        - no implicit type conversions
        - this makes typing errors basically impossible (?)
    - since everything is typed, things feel A LOT more "to the point"
        -> instead of wasting thought on how to do things "elegantly" you "just do it"
        -> this makes the language less expressive, but at the same time it could be more productive
    - lower-case functions are NOT exported from a module
        - Note: you can still use it in the same module, defined by go.mod
    - use testing framework: https://github.com/stretchr/testify
        - otherwise you will go mad with boilerplate code ...

- Main Learning Hub: https://go.dev/learn/
    - Web Dev: https://gowebexamples.com/

- https://go.dev/talks/2013/go4python.slide#1
- https://go.dev/talks/2012/zen.slide#1
- https://www.augmentedmind.de/2023/01/22/go-vs-python-senior-developers/
    - pretty good intro
    - some parts are missing tho (panic/recover)

Frameworks:
- https://github.com/beego/beego?tab=readme-ov-file

- Book "Cloud Native Go" (see "Downloads" folder)
- Go by Example: https://gobyexample.com/
- https://www.w3schools.com/go/go_getting_started.php


### Fundamental stuff:
- Generic functions:
    - on slices: https://go.dev/blog/generic-slice-functions

- https://github.com/mrekh/go-for-python-developers

- https://go.dev/blog/maps

- How to deal with JSON in a statically typed language?
    - https://medium.com/rahasak/handle-arbitrary-json-in-golang-3900b75a7fbb#id_token=eyJhbGciOiJSUzI1NiIsImtpZCI6ImQyZDQ0NGNmOGM1ZTNhZTgzODZkNjZhMTNhMzE2OTc2YWEzNjk5OTEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTQ4NTY1NDc3MjczMTA5NzcxMzIiLCJlbWFpbCI6InJlc2NoYmVyZ2VyQGdvb2dsZW1haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5iZiI6MTcyMzU3ODM1MSwibmFtZSI6IkJlbmVkaWt0IFJlc2NoYmVyZ2VyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0tTdlBwc01NR1h1RFlBYzFMVFU2U0UtSmdWTnJWcUhxdVAxZ3VMeV9lZ21mMGwxdz1zOTYtYyIsImdpdmVuX25hbWUiOiJCZW5lZGlrdCIsImZhbWlseV9uYW1lIjoiUmVzY2hiZXJnZXIiLCJpYXQiOjE3MjM1Nzg2NTEsImV4cCI6MTcyMzU4MjI1MSwianRpIjoiZWExYzkyN2NiZmNjZTU5NTU0MjQwNmNhMDkyMDMyN2QyNTk2NTU2ZiJ9.g-MaXBNq9dMOwXmNYWnHuaHKYwDR7FxJqu6SBWUX_lVfzchPcfm7sd0dP-CxDmrpaPB-Rt3XWHEldR0L_FLQqZxQ6luhx7MRMRZBXdXLJ9Nq9HohZp42hHHd2pfixMaVPHOm2bui08iB797tfZYEMwDfodyBGzRvBAJa0BolWDIEA4VFUCOERkW-3K-UdJiYUNEAIWltpsQ26-y62ODJNwzIh9QyLQxNDh-oOd1DS_pwKxTniz-41yEGlkjBjFV1RwKvEJI0ETySVvAYQUol54f7zIFC-Hlhvb1w5ooun1010Gcax5IeQ-ZWDKLX1kmjXoVTawh2ZFVb3JSKZN2Vyw
    - basically you now have to create Types for EVERYTHING
        - "There is no dict"

- https://tutorialedge.net/golang/reading-console-input-golang/


Working with Go:
- popular web server/REST API frameworks for Go
    - create a webserver including some endpoints for a CRUD app
    - implement some data processing on top of it
    - user session management for web servers

- Desktop Apps using Go and Web tech:
https://wails.io/docs/introduction

- Gorm

- data processing in Go

- Video streaming with Go
    - saving a video stream to disk



###

After many years working with SPA's via React/Vue/Svelte, I've completely switched to 
- HTMX
- Hyperscript
- TailwindCSS
- and Go for all my frontend needs.

Advantages include:

- Only one programming language to work in, no context switching between Typescript and Go, which cognitively makes me work many times faster
- No more npm/node dependency management hell
- Because I don't use npm/node, build times are now extremely fast, 5-6 seconds of compile time and I have a fully working frontend and backend, which is impossible with the current node based ecosystems
- For the rare times I need Javascript (animations, etc), Hyperscript is descriptive and doesn't generally break my "flow"
- No more syncing state between your frontend and backend -- which means no more HTTP REST API's for your frontend, translation mechanisms if you use gRPC, or mapping data structures between your backend and frontend. All you have to do is render your page using state held in Go.
- The last point is extremely important here: server <-> client state management is miserable, no matter which way you cut it. By using HTMX, each of your HTTP endpoints in Go are just paths that spit out rendered templates. All of your user session data and state is stored entirely server side without having to send state to the client -- you just send exactly (and only!) what you want the user to see. HTMX will take your server's response and replace the DOM (or a part of the DOM if you want!) with your response, without doing a full page reload.

I estimate that ever since switching to the above stack, I work a solid 5-10x faster when creating SPA's. I'll never go back to a frontend framework again.