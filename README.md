# Note

```
% go mod init github.com/changyy/go-map-service-tw-covid19-rapid-antigen-tests
go: creating new go.mod: module github.com/changyy/go-map-service-tw-covid19-rapid-antigen-tests
go: to add module requirements and sums:
	go mod tidy

% go mod tidy
go: finding module for package github.com/patrickmn/go-cache
go: finding module for package github.com/gin-gonic/gin
go: found github.com/gin-gonic/gin in github.com/gin-gonic/gin v1.7.7
go: found github.com/patrickmn/go-cache in github.com/patrickmn/go-cache v2.1.0+incompatible

% go build -o bin/
# golang.org/x/sys/unix
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:27:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:40:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:28:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:43:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:59:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:75:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:90:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:105:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:121:3: //go:linkname must refer to declared function or variable
../go/pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:121:3: too many errors

% go get golang.org/x/sys/unix
go: upgraded golang.org/x/sys v0.0.0-20200116001909-b77594299b42 => v0.0.0-20220513210249-45d2b4557a2a

% go build -o bin/
% file bin/go-map-service-tw-covid19-rapid-antigen-tests 
bin/go-map-service-tw-covid19-rapid-antigen-tests: Mach-O 64-bit executable x86_64
% ./bin/go-map-service-tw-covid19-rapid-antigen-tests 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] Loaded HTML Templates (2): 
	- 
	- index.tmpl

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

# Heroku

```
% heroku git:remote -a YourProjectName
% git push heroku main
```
