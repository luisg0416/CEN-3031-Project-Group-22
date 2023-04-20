# CEN-3031-Project-Group-22

Project Name: UF Studying Website

Project Description: Our plan is to make a website that gives UF students different resources to study. These resources include flashcards, study-guides, practice exams, and games. 

Team Members: Matthew Darrow (front-end), Kory Gauger (back-end), Luis Guerrero (front-end), Amari Nixon (back-end)

# Frontend
===================

- Hosts *React* on *Go*
- Uses *Go* as the webserver using *fiber* instead of the included node server within a "react app"
- Added routing using *BrowserRouter* to switch seamlessly between different pages without refreshing the browser tab
- Uses *EsBuild* as the javascript bundler

## Running the frontend

```
node ./esbuild.js --watch
```

# Backend
===================
### Packages that need to be installed
1. "github.com/gofiber/fiber/v2"
2. "github.com/gofiber/template/html"
3. "github.com/gocarina/gocsv"
4. etc.

if any problems occur during setup that involve a package missing using the command
```
go mod tidy
```
should fix any issues. If a particular package is improperly installed this way you can install it using the comman
```
go get "github.com/example"
```
where the example in the above is the package that needs to be installed

## Running the application
the command
```
go run main.go
```
launches the application
ctrl clicking on the link that appears in the consol will bring you to the right page. 
We have it configured to port 3000 if this port is occupied on your machine going into the .env file and changing the value next to 
```
PORT=
```
Will set you to a different port
