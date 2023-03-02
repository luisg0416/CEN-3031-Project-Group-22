# Features
- Hosts *React* on *Go*
- Uses *Go* as the webserver using *fiber* instead of the included node server within a "react app"
- Added routing using *BrowserRouter* to switch seamlessly between different pages without refreshing the browser tab
- Uses *EsBuild* as the javascript bundler

### Frontend
- Make frontend changes in **"frontend/application.tsx"**
- Add/modify pages in **"frontend/pages/"**
### Backend
- Make changes to main.go

# Commands
front-end: 
```
node ./esbuild.js --watch
```
back-end:
``` back-end
go run main.go
```

# Setup
If you're insterested in how it was setup:

https://www.youtube.com/watch?v=pbcTa-a3LBw

https://www.youtube.com/watch?v=Y7kuW1qyDng
