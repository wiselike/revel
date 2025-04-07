# CHANGELOG


## v1.1.0

[[revel/revel](https://github.com/wiselike/revel)]

* bc0e27f Merge pull request #1552 from revel/1542_recursive_call
* d202b93 Merge pull request #1448 from golddranks/master
* 859e5b4 Merge pull request #1511 from jiro4989/patch-1
* 90489b1 Merge pull request #1523 from KoichiWada/handle-sigterm
* 347610c Merge branch 'develop' into handle-sigterm
* 942bd2e Merge pull request #1525 from mikyk10/feature/adding-mime-compress
* 2cb950f Merge pull request #1543 from dhiemaz/fixing-typos
* aed0d1e Merge pull request #1550 from revel/bugfix/readme-install
* bc89379 Fixed log recursive call There was a recusive loop in the logger, this fixes it closes #1542
* aa8a94d Merge pull request #1546 from braineet/master
* 741d2c8 Corrects install command for new version of go
* 65db3c0 Merge pull request #1549 from revel/bugfix/session-uuid
* 624f341 Merge branch 'develop' into bugfix/session-uuid
* 5e99db8 Refactors session uuid to use google's package
* 30f3424 Correction redis import
* f934412 fixing typos
* f261091 Added a MIME type for compressableMimes
* 9907376 Handle SIGTERM for graceful shutdown.
* 413dda3 Merge pull request #1518 from ptman/lint
* 45a4413 More lint fixes and dead code removal
* 3799c55 Merge pull request #1514 from ptman/lint
* 655927a Removed go 1.12 added go 1.15
* 060e640 Fix misspellings and some lint errors
* 921e1b4 Fix URL (http -> https)
* e12cd0e develop v1.1.0-dev
* a29f37c Fix a bug when binding to pointer

[[revel/cmd](https://github.com/wiselike/revel-cmd)]

* 86b4670 Update README.md
* bc376fb Merge pull request #211 from lujiacn/develop
* 5c8ac53 Merge branch 'develop' into develop
* c674084 Merge pull request #201 from ptman/lint
* cfe1d97 Merge branch 'develop' into lint
* 7f9f658 Merge pull request #209 from glaslos/patch-1
* 126d20c Merge pull request #210 from revel/build_process_update
* 111264c updated go.mod
* 4087c49 updated golang.org/x/tools, to avoid internal error: package xxx without types was imported from ...
* 3602eb4 Merge branch 'develop' into patch-1
* 192fc66 Merge pull request #200 from julidau/develop
* 5689f86 Merge pull request #204 from shinypb/master
* 6dba0c3 Fix bad error syntax An wrapped error message in the cmd module was referencing the wrong parameter value to be built closes revel/revel#1532
* bb926f3 Added additional pattern to test against Another different missing pacakge error thrown that can be detected and added This error occurs because a package may have been stripped down when originally loaded
* 3cd5ebb Updated launch scripts
* 25dc05b Updated Launch code Added output to error stack, so terminal errors are displayed Ficed c.Verbose, it was changed to an array which causes issues launching Removed . notation from doing anything special. This was already replaced with the -p CLI option Added documentaiton on adding the package name Started watcher with force refresh.
* 0a40a20 Merge pull request #208 from notzippy/build_process_update
* fcc1319 Fixing type
* ea5acb7 Updated shared build environments Updated check for errors. Updated go.mod Added .vscode launch
* 25d6352 Get rid of redundant space in the output of `revel new -a`
* ddec572 More linting
* 7a91d0c interrupt process on windows as well
* b562bd2 Merge pull request #199 from ptman/lint
* bf17a71 Merge branch 'develop' into lint
* 3d924a0 Lint fixes
* d64c7f1 develop v1.1.0-dev

[[revel/config](https://github.com/wiselike/revel-config)]

* no changes

[[revel/modules](https://github.com/wiselike/revel-modules)]

* 852ea71 Merge branch 'master' into develop
* 789324e Merge pull request #105 from ptman/lint
* 0a9a7f4 Update .travis.yml
* 8dbd171 Lint fixes
* 464e072 develop v1.1.0-dev

[[revel/cron](https://github.com/revel/cron)]

* 031e64e Merge pull request #5 from ptman/lint
* 7cfc261 Lint fixes

[[revel/examples](https://github.com/revel/examples)]

* e51ed5a Merge pull request #54 from realbucksavage/master
* c555714 Merge pull request #58 from teitei-tk/update_gorp_link
* 24ed869 Merge pull request #62 from ptman/lint
* 5ec13f7 More lint fixes
* b14432a Merge branch 'develop' into lint
* a33af07 Merge pull request #59 from obsti8383/master
* 46ae3a6 Removed go 1.12, added 1.15
* ced52ae Lint and revel 1.0.0 fixes
* b245628 revert to old go.sum
* b6ea1ec upgrade references to revel framework to 1.0.0 in go.mod; fixes compile errors
* 8353712 update gorp link
* be755f5 Indented the code.

[[revel/revel.github.io](https://github.com/wiselike/revel.github.io)]

* 32d1cd9 Merge branch 'master' into develop
* e72f206 Merge pull request #197 from revel/update_go_version
* 1ae75e8 Merge branch 'develop' into update_go_version
* 75172ea develop v1.1.0-dev

[[revel/heroku-buildpack-go-revel](https://github.com/revel/heroku-buildpack-go-revel)]

* no changes


## v1.0.0

[[revel/revel](https://github.com/wiselike/revel)]

* 3d1b0c3 Merge pull request #1497 from lujiacn/master
* ff2da7e Merge pull request #1498 from aacapella/feature/same-site-cookies
* c6c4c35 SameSite cookie support
* bfad570 Update server_adapter_go.go
* ff43c73 Merge pull request #1491 from notzippy/go-mod
* dbe9fee Update .travis.yml
* 38b0687 Fixed paths for test cases
* 39523bf Enhanced logging
* 59b8375 Changes to Revel for go.mod support Modified module lookup to handle lookups using the app.conf (before relied on source file) Added extra logging for routes error handling
* 1053f49 Merge pull request #1443 from lujiacn/develop
* dcafb9e Merge pull request #1488 from notzippy/go-mod
* e30c8da Merge pull request #1483 from goevexx/feature/fix-issue-1482
* 50e70f9 Updated revel to receive paths passed in Updated watcher to use master branch
* d581f71 change import to fix issue 1482
* fdc724a Merge pull request #1462 from torden/feature/fix_puretextstrict
* fe4861c Fix (#1458) the undetected self-closing tags in isPureTextStrict Fix (#1458) the always uses STRICT mode in PureText.IsSatisFied
* ae3895a added wasm mime-type
* 45ec814 Merge pull request #1439 from mukeshjeengar/hotfix/log-rotation-fixed
* d3a76ed log rotation fixed
* 2eb9067 Merge pull request #1413 from nevkontakte/patch-1
* ccf085e Merge pull request #1434 from dmjones/fix-1433
* 34e886a Don't invoke action when Before returns value
* 5b70626 Merge pull request #1427 from SYM01/hotfix/avoid-dos
* d160ecb fix issue #1424
* db7db5b remove unneccsary code assignment to nil
* 8bff5bb Update controller.go
* 16f5fef Remove a stray println call.
* 60c3d7a develop v1.0.0-dev

[[revel/cmd](https://github.com/wiselike/revel-cmd)]

* d8117a3 Merge pull request #186 from notzippy/go-mod
* 6371373 Removed version update Version control is maintained through go.mod file Modified harness to only kill the application if not responded after 60 seconds in windows
* 28ac65f Merge pull request #185 from notzippy/go-mod
* 5070fb8 Fixed issue with new and run flag Updated tests to run final test in non gopath, with new name
* 904cfa2 Added some informational messages while download
* 223bd3b Added manual scan on packages in app folder This allows for source code generation. Packages in <application>/app folder are scanned manually as opposed to the `packages.Load` scan which will fast fail on compile error, and leave you with go files with no syntax.
* 4987ee8 Added verbose logging to building / testing a no-vendor app Removed section which raises an error when examining packages, we dont need to check for errors on foreign packages since we are importing only a slice of the data
* 4bab440 Updated Revel command Added a check to see if harness had already started, saves a recompile on load Added check to source info for local import renames Removed the go/build check for path and just check existence of the path Formatting updates
* 741f492 Updated scanner Removed scanning all the import statements, this is not needed Added recursive scan for revel import path to pick up testunits
* 60b88a4 Merge pull request #180 from notzippy/go-mod
* 49eef29 Build and Historic build updates Modified GOPATH to not modify build with go.mod Updated go.mod to version 1.12 Updated harness to setup listener before killing process Updated notvendored flag to --no-vendor Updated command_config to ensure no-vendor can be build Added additional checks in source path lookup
* 9d3a554 Updates Updated NotVendored flag Updated travis matrix Updated build log
* 36bd6b9 Corrected flags
* 1d9df25 Moved test cases to run last
* ad694c0 Debug travis
* fb4b565 Debug travis Added verbose flag so we can see what is occurring, Removed checkout for revel, not needed anymore
* 20d5766 Added gomod-flags Added a gomod-flags parameter which allows you to run go mod commands on the go.mod file before the build is performed. This allows for development environments.
* 0920905 Updated to build go 1.12 and up Modified to use fsnotify directlyUpdated travis to not use go deps
* 31cb64e Check-in of command_test, remaps the go mod command to use the develop branch.
* 33abc47 Fixed remaining test
* 86736d6 Updated formating Ran through testing individually for vendored Revel applications
* 07d6784 Restructured command config Removed go/build reference in clean
* c1aee24 Corrected version detection, so that equal versions match
* f2b54f5 Updated sourceinfo Added packagepathmap to the SourceInfo, this in turn allows the RevelCLI app command to pass the source paths directly to Revel directly Added default to build to be "target" of the current folder Renamed source processor
* 3f54665 Added processor to read the functions in the imported files, and populate the SourceInfo object the same as before
* 548cbc1 Upatede Error type to SourceError Added processor object to code Verified compile errors appearing Signed-off-by: notzippy@gmail.com
* 9a9511d Updated so revel new works and revel run starts parsing the source.
* acb8fb6 Initial commit to go mod
* d201463 Merge pull request #176 from xXLokerXx/fix_windows_path
* 773f688 Merge branch 'develop' into fix_windows_path
* ca4cfa5 Merge pull request #165 from kumakichi/fixed_import_C
* 4368690 Merge pull request #179 from Laur1nMartins/Laur1nMartins/fix-linkerFlags
* cf2e617 Merge branch 'develop' into Laur1nMartins/fix-linkerFlags
* 424474a Fix linker flags inclusion in build comamnd
* 6d8fcd9 Fix sintax error
* aa459c1 Fix sintaxis error
* 0b23b3e Fix complexity
* 3f65e1e acept slash and inverted slash in src path validation
* 7dce3d8 fixed import "C"
* 5c8d5bc develop v1.0.0-dev

[[revel/config](https://github.com/wiselike/revel-config)]

* no changes

[[revel/modules](https://github.com/wiselike/revel-modules)]

* e1fdc01 Merge pull request #103 from revel/master
* 80d53e2 Merge pull request #102 from notzippy/go-mod
* 2048fce Updated build processor
* 19728d3 Added gomod removed vendor specific imports
* 515369e develop v1.0.0-dev

[[revel/cron](https://github.com/revel/cron)]

* no changes

[[revel/examples](https://github.com/revel/examples)]

* 2d2968c Merge pull request #57 from notzippy/go-mod
* dc75997 Updated examples Updated booking app to go.mod Updated chat, facebook, others app to add in go file in the root Updated travis to run tests in windows Updated travis to exclude testing fasthttp on windows
* 5b25a51 Removed persona from project, this function no longer exists in browsers

[[revel/revel.github.io](https://github.com/wiselike/revel.github.io)]

* 6cd3647 Merge pull request #196 from aacapella/feature/same-site-cookies
* 9f8f537 Merge pull request #191 from dmjones/session-value-not-found-returns-error
* d79c912 Merge pull request #194 from DGKSK8LIFE/patch-1
* 3911471 Merge pull request #195 from notzippy/develop
* 67b088f Same site cookie setting
* f5c5cb0 Corrected issues
* bba502d Update for gomod docs
* 9765ef0 Merge remote-tracking branch 'revel/master' into develop
* eedc235 fixed spelling error
* 9b9270a Explain return value when session value not found
* 24abe9a Merge pull request #184 from manfordbenjamin/master
* 4969200 Change logo and apply blue theme style to all pages
* 844fe5d Revamp homepage
* ef54af7 develop v1.0.0-dev

[[revel/heroku-buildpack-go-revel](https://github.com/revel/heroku-buildpack-go-revel)]

* no changes


## v0.21.0

### New items
* **Session Engine support** You can now define your own session engine. By default the cookie engine is the one used for the session, but with this change you can implement your own. Also controller.Session is now a `map[string]interface{}` which means you can place any object that can be serialized to a string. 
* **Added http mux support** you can now integrate Revel with packages that have their own HTTP muxers. This allows you to integrate with Hugo, Swagger or any suite of software that is out there.
* `revel.controller.reuse` app.conf option to turn on / off reuse of the controllers. Defaults to true

### Breaking changes
controller.Session is now a map[string]interface{} (previously was map[string]string) this means existing code needs to cast any values they pull from the session 
```
	if username, ok := c.Session["user"]; ok { // username is a string type
		return c.getUser(username)
	}
```
changes to
```
	if username, ok := c.Session["user"]; ok { // username is an interface{} type
		return c.getUser(username.(string))
	}
```

Deprecated log points removed
revel.ERROR, revel.TRACE, revel.DEBUG, revel.WARN have been removed

Function name change `revel.OnAppStop` Replaced revel.OnAppShutdown 

revel.SimpleStack was moved to github.com/wiselike/revel/utils.SimpleStack

### New packages required
#### Revel Framework
* github.com/twinj/uuid (revel/revel session ID generation)

#### Revel Cmd
* github.com/kr/pty (revel/cmd capture output of dep tool)

### Package changes
#### Revel Framework 
* Added stack to errors Added stack information to router when forumlating error
* Fix spelling errors from go report
* Removed deprecated loggers
* Updated travis , made windows success optional
* Exposed StopServer function to public Changed session filter to use empty call
* 577ae8b Enhancement pack for next release Added session engine support, and the session cookie engine breaking change revel.Session was map[string]string now is map[string]interface{}
* Updated shutdown to support windows environment 
* Patched shutdown support to make it work through the event engine 
* Added ENGINE_SHUTDOWN_REQUEST to events, raising this event will cause the server to shutdown 
* Assigned Server engines to receive revel.Events Added revel.OnAppStop hooks - 
* Normalized startup / shutdown hooks into their own package

#### Revel Cmd
* Modified run command to translate symlinks to absolute paths
* Tool updates Updated tool to give more meaningful errors 
* Added file system as an option to fetch the skeleton from
* Allow windows to fail on travis 
* run Command will choose CWD if no additional arguments are supplied 
* Added Revel version check, compatible lists are in model/version

#### Revel Modules
* Updated CSRF test cases 
* Added travis test for modified test engine
* Updated server-engine modules to support OnAppStop functionality.
* Reorganization, readme updates Moved auth example into its own folder 
* Updated root readme 
* Updated CSRF

#### Revel Examples
* Fixed issue with error checking closes websocket in chat
* Updated booking module to work with changed session 
* Updated to remove any references to old revel.log variables


## v0.20.0

### New items
* **Updated minimum Go requirements to Go 1.8**

#### Revel Cmd changes
See [manual](http://revel.github.io/manual/tool.html) for more information on the flags and new features
* Rewrote revel/cmd package so it has no dependencies on revel/revel - future releases for revel/cmd will not be on the same schedule as revel/revel
* Added flag support to revel/cmd ,
* Added automatic vendor creation flag, when enabled a vendor folder will be added and used to the project
* Added [DEFAULT] section to message skeleton so multiple lines work
* If port specified is 0 then proxy can will run on a random free port
* Split main file into two files so it may be invoked using other methods
* Added ability to only monitor and update build files (without running a proxy)
* Auto download revel/revel, revel/cmd - if you point your GOPATH to an empty directory the `revel` tool will still be able to create a new project
* Made application path smarter. Now supports absolute paths and relatives paths.
* modified `revel run` to pass in json code to the `revel.Init` function in place of the `mode` field. This allows for dynamic information to be passed to the Revel Server. This can be disabled by using  this flag ` --historic-run-mode`
* modified `revel new` added `-V` to auto create the vendor folder inside the application along with the Gopkg.toml file.
* modified `revel new` added `-r` to run the application after creating it.

### revel/module
* Enhanced gorp module to make database schema available
* Added server-engine  server-engine/gohttptest and a testsuite file that implements all the methods from revel/testsuite. This test engine as an alpha implementation to be able to run `go test` from the command line. Meaning your tests can now live beside the controller. An example testsuite is [here](https://github.com/revel/examples/blob/develop/booking/app/controllers/app_test.go)
* Updated static module, moved a structure into a model package

### revel/revel
* Added startup fail REVEL_FAILURE event
* Added HTTP_REQUEST_CONTEXT fetch
* Added websockets Message.Send / Message.Receive functions.
* Enhanced required validator
* Graceful shutdown added. By doing a `revel.RaiseEvent()`

#### Breaking Changes
* `Request.Context() map[string]interface()` renamed to
`Request.Args() map[string]interface()`
new function added which returns the request.Context for the server-engine that supports it
`Request.Context() context.Context`
`revel.EventHandler` signature was `func(typeOf int, value interface{}) (responseOf int)`
now it is func(typeOf Event, value interface{}) (responseOf EventResponse)


## v0.19.1
Fix import to point to the fsnotify/fsnotify.v1


## v0.19.0

# Maintenance Release

This release is focused on improving the security and resolving some issues. 

**There are no breaking changes from version 0.18**

[[revel/cmd](https://github.com/wiselike/revel-cmd)]
* Improved vendor folder detection revel/cmd#117
* Added ordering of controllers so order remains consistent in main.go revel/cmd#112
* Generate same value of `AppVersion` regardless of where Revel is run revel/cmd#108
* Added referrer policy security header revel/cmd#114

[[revel/modules](https://github.com/wiselike/revel-modules)]
* Added directory representation to static module revel/modules#46
* Gorp enhancements (added abstraction layer for transactions and database connection so both can be used), Added security fix for CSRF module revel/modules#68
* Added authorization configuration options to job page revel/modules#44

[[revel/examples](https://github.com/revel/examples)]
* General improvements and examples added revel/examples#39  revel/examples#40

## v0.18

## Upgrade path
The main breaking change is the removal of `http.Request` from the `revel.Request` object.
Everything else should just work....

## New items
* Server Engine revel/revel#998
The server engine implementation is described in the [docs](http://revel.github.io/manual/server-engine.html)
* Allow binding to a structured map. revel/revel#998 
Have a structure inside a map object which will be realized properly from params
* Gorm module revel/modules/#51
Added transaction controller
* Gorp module revel/modules/#52
* Autorun on startup in develop mode revel/cmd#95
Start the application without doing a request first using revel run ....
* Logger update revel/revel#1213
Configurable logger and added context logging on controller via controller.Log
* Before after finally panic controller method detection revel/revel#1211 
Controller methods will be automatically detected and called - similar to interceptors but without the extra code
* Float validation revel/revel#1209
Added validation for floats
* Timeago template function revel/revel#1207
Added timeago function to Revel template functions
* Authorization to jobs module revel/module#44
Added ability to specify authorization to access the jobs module routes
* Add MessageKey, ErrorKey methods to ValidationResult object revel/revel#1215
This allows the message translator to translate the keys added. So model objects can send out validation codes
* Vendor friendlier - Revel recognizes and uses `deps` (to checkout go libraries) if a vendor folder exists in the project root. 
* Updated examples to use Gorp modules and new loggers


### Breaking Changes

* `http.Request` is no longer contained in `revel.Request` revel.Request remains functionally the same but 
you cannot extract the `http.Request` from it. You can get the `http.Request` from `revel.Controller.Request.In.GetRaw().(*http.Request)`
* `http.Response.Out` Is not the http.Response and is deprecated, you can get the output writer by doing `http.Response.GetWriter()`. You can get the `http.Response` from revel.Controller.Response.Out.Server.GetRaw().(*http.Response)`

* `Websocket` changes. `revel.ServerWebsocket` is the new type of object you need to declare for controllers 
which should need to attach to websockets. Implementation of these objects have been simplified

Old
```

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()
```
New
```
func (c WebSocket) RoomSocket(user string, ws revel.ServerWebSocket) revel.Result {
	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if ws.MessageSendJSON(&event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()
```
* GORM module has been refactored into modules/orm/gorm 


### Deprecated methods
* `revel.Request.FormValue()` Is deprecated, you should use methods in the controller.Params to access this data
* `revel.Request.PostFormValue()` Is deprecated, you should use methods in the controller.Params.Form to access this data
* `revel.Request.ParseForm()` Is deprecated - not needed
* `revel.Request.ParseMultipartForm()` Is deprecated - not needed
* `revel.Request.Form` Is deprecated, you should use the controller.Params.Form to access this data
* `revel.Request.MultipartForm` Is deprecated, you should use the controller.Params.Form to access this data
* `revel.TRACE`, `revel.INFO` `revel.WARN` `revel.ERROR` are deprecated. Use new application logger `revel.AppLog` and the controller logger `controller.Log`. See [logging](http://revel.github.io/manual/logging.html) for more details.

### Features

* Pluggable server engine support. You can now implement **your own server engine**. This means if you need to listen to more then 1 IP address or port you can implement a custom server engine to do this. By default Revel uses GO http server, but also available is fasthttp server in the revel/modules repository. See the docs for more information on how to implement your own engine.

### Enhancements
* Controller instances are cached for reuse. This speeds up the request response time and prevents unnecessary garbage collection cycles.  

### Bug fixes




## v0.17

[[revel/revel](https://github.com/wiselike/revel)]

* add-validation
* i18-lang-by-param
* Added namespace to routes, controllers
* Added go 1.6 to testing
* Adds the ability to set the language by a url parameter. The route file will need to specify the parameter so that it will be picked up
* Changed url validation logic to regex
* Added new validation mehtods (IPAddr,MacAddr,Domain,URL,PureText)

[[revel/cmd](https://github.com/wiselike/revel-cmd)]

* no changes

[[revel/config](https://github.com/wiselike/revel-config)]

* no changes

[[revel/modules](https://github.com/wiselike/revel-modules)]

* Added Gorm module

[[revel/cron](https://github.com/revel/cron)]

* Updated cron task manager
* Added ability to run a specific job, reschedules job if cron is running.

[[revel/examples](https://github.com/revel/examples)]

* Gorm module (Example)

# v0.16.0

Deprecating support for golang versions prior to 1.6
### Breaking Changes

* `CurrentLocaleRenderArg` to `CurrentLocaleViewArg` for consistency
* JSON requests are now parsed by Revel, if the content type is `text/json` or `application/json`. The raw data is available in `Revel.Controller.Params.JSON`. But you can also use the automatic controller operation to load the data like you would any structure or map. See [here](http://revel.github.io/manual/parameters.html) for more details

### Features

* Modular Template Engine #1170 
* Pongo2 engine driver added revel/modules#39
* Ace engine driver added revel/modules#40
* Added i18n template support #746 

### Enhancements

* JSON request binding #1161 
* revel.SetSecretKey function added #1127 
* ResolveFormat now looks at the extension as well (this sets the content type) #936 
* Updated command to run tests using the configuration revel/cmd#61

### Bug fixes

* Updated documentation typos revel/modules#37
* Updated order of parameter map assignment #1155 
* Updated cookie lifetime for firefox #1174 
* Added test path for modules, so modules will run tests as well #1162 
* Fixed go profiler module revel/modules#20


# v0.15.0
@shawncatz released this on 2017-05-11

Deprecating support for golang versions prior to 1.7

### Breaking Changes

* None

### Features

* None

### Enhancements

* Update and improve docs revel/examples#17 revel/cmd#85

### Bug fixes

* Prevent XSS revel/revel#1153
* Improve error checking for go version detection revel/cmd#86

# v0.14.0
@notzippy released this on 2017-03-24

## Changes since v0.13.0

#### Breaking Changes
- `revel/revel`:
  - change RenderArgs to ViewArgs PR #1135
  - change RenderJson to RenderJSON PR #1057
  - change RenderHtml to RenderHTML PR #1057
  - change RenderXml to RenderXML PR #1057

#### Features
- `revel/revel`:

#### Enhancements
- `revel/revel`:


#### Bug Fixes
- `revel/revel`:


# v0.13.1
@jeevatkm released this on 2016-06-07

**Bug fix:**
- Windows path fix #1064


# v0.13.0
@jeevatkm released this on 2016-06-06

## Changes since v0.12.0

#### Breaking Changes
- `revel/revel`:
  - Application Config name changed from `watcher.*` to `watch.*`  PR #992, PR #991

#### Features
- `revel/revel`:
  - Request access log PR #1059, PR #913, #1055
  - Messages loaded from modules too PR #828
- `revel/cmd`:
  - Added `revel version` command emits the revel version and go version revel/cmd#19

#### Enhancements
- `revel/revel`:
  - Creates log directory if missing PR #1039
  - Added `application/javascript` to accepted headers PR #1022
  - You can change `Server.Addr` value via hook function PR #999
  - Improved deflate/gzip compressor PR #995
  - Consistent config name `watch.*` PR #992, PR #991
  - Defaults to HttpOnly and always secure cookies for non-dev mode #942, PR #943
  - Configurable server Read and Write Timeout via app config #936, PR #940
  - `OnAppStart` hook now supports order param too PR #935
  - Added `PutForm` and `PutFormCustom` helper method in `testing.TestSuite` #898
  - Validator supports UTF-8 string too PR #891, #841
  - Added `InitServer` method that returns `http.HandlerFunc` PR #879
  - Symlink aware processing Views, Messages and Watch mode PR #867, #673
  - Added i18n settings support unknown format PR #852
  - i18n: Make Message Translation pluggable PR #768
  - jQuery `min-2.2.4` & Bootstrap `min-3.3.6` version updated in `skeleton/public` #1063
- `revel/cmd`:
  - Revel identifies current `GOPATH` and performs `new` command; relative to directory revel/revel#1004
  - Installs package dependencies during a build PR revel/cmd#43
  - Non-200 response of test case request will correctly result into error PR revel/cmd#38
  - Websockets SSL support in `dev` mode PR revel/cmd#32
  - Won't yell about non-existent directory while cleaning PR revel/cmd#31, #908
    - [x] non-fatal errors when building #908
  - Improved warnings about route generation PR revel/cmd#25
  - Command is Symlink aware PR revel/cmd#20
  - `revel package` & `revel build` now supports environment mode PR revel/cmd#14
  - `revel clean` now cleans generated routes too PR revel/cmd#6
- `revel/config`:
  - Upstream `robfig/config` refresh and import path updated from `github.com/wiselike/revel/config` to `github.com/wiselike/revel-config`, PR #868
  - Config loading order and external configuration to override application configuration revel/config#4 [commit](https://github.com/wiselike/revel/commit/f3a422c228994978ae0a5dd837afa97248b26b41)
  - Application config error will produce insight on error PR revel/config#3 [commit](https://github.com/wiselike/revel-config/commit/85a123061070899a82f59c5ef6187e8fb4457f64)
- `revel/modules`:
  - Testrunner enhancements
    - Minor improvement on testrunner module PR #820, #895
    - Add Test Runner panels per test group PR revel/modules#12
- `revel/revel.github.io`:
  - Update `index.md` and homepage (change how samples repo is installed) PR [#85](https://github.com/wiselike/revel.github.io/pull/85)
  - Couple of UI improvements PR [#93](https://github.com/wiselike/revel.github.io/pull/93)
  - Updated techempower benchmarks Round 11 [URL](http://www.techempower.com/benchmarks/#section=data-r11)
  - Docs updated for v0.13 release
- Cross-Platform Support
  - Slashes should be normalized in paths #260, PR #1028, PR #928

#### Bug Fixes
- `revel/revel`:
  - Binder: Multipart `io.Reader` parameters needs to be closed #756
  - Default Date & Time Format correct in skeleton PR #1062, #878
  - Addressed with alternative for `json: unsupported type: <-chan struct {}` on Go 1.6 revel/revel#1037
  - Addressed one edge case, invalid Accept-Encoding header causes panic revel/revel#914


# v0.11.3
@brendensoares released this on 2015-01-04

This is a minor release to address a critical bug (#824) in v0.11.2.

Everybody is strongly encouraged to rebuild their projects with the latest version of Revel. To do it, execute the commands:

``` sh
$ go get -u github.com/wiselike/revel-cmd/revel

$ revel build github.com/myusername/myproject /path/to/destination/folder
```


# v0.11.2
on 2014-11-23

This is a minor release to address a critical bug in v0.11.0.

Everybody is strongly encouraged to rebuild their projects with the latest version of Revel. To do it, execute the commands:

``` sh
$ go get -u github.com/wiselike/revel-cmd/revel

$ revel build github.com/myusername/myproject /path/to/destination/folder
```


# v0.11.1
@pushrax released this on 2014-10-27

This is a minor release to address a compilation error in v0.11.0.


# v0.12.0
@brendensoares released this on 2015-03-25

Changes since v0.11.3:

## Breaking Changes
1. Add import path to new `testing` sub-package for all Revel tests. For example:

``` go
package tests

import "github.com/wiselike/revel/testing"

type AppTest struct {
    testing.TestSuite
}
```
1. We've relocated modules to a dedicated repo. Make sure you update your `conf/app.conf`. For example, change:

``` ini
module.static=github.com/wiselike/revel/modules/static
module.testrunner = github.com/wiselike/revel/modules/testrunner
```

to the new paths:

``` ini
module.static=github.com/wiselike/revel-modules/static
module.testrunner = github.com/wiselike/revel-modules/testrunner
```

## [ROADMAP] Focus: Improve Internal Organization

The majority of our effort here is increasing the modularity of the code within Revel so that further development can be done more productively while keeping documentation up to date.
- `revel/revel.github.io`
  - [x] Improve docs #[43](https://github.com/wiselike/revel.github.io/pull/43)
- `revel/revel`:
  - [x] Move the `revel/revel/harness` to the `revel/cmd` repo since it's only used during build time. #[714](https://github.com/wiselike/revel/issues/714)
  - [x] Move `revel/revel/modules` to the `revel/modules` repo #[785](https://github.com/wiselike/revel/issues/785)
  - [x] Move `revel/revel/samples` to the `revel/samples` repo #[784](https://github.com/wiselike/revel/issues/784)
  - [x] `testing` TestSuite #[737](https://github.com/wiselike/revel/issues/737) #[810](https://github.com/wiselike/revel/issues/810)
  - [x] Feature/sane http timeout defaults #[837](https://github.com/wiselike/revel/issues/837) PR#[843](https://github.com/wiselike/revel/issues/843) Bug Fix PR#[860](https://github.com/wiselike/revel/issues/860)
  - [x] Eagerly load templates in dev mode #[353](https://github.com/wiselike/revel/issues/353) PR#[844](https://github.com/wiselike/revel/pull/844)
  - [x] Add an option to trim whitespace from rendered HTML #[800](https://github.com/wiselike/revel/issues/800)
  - [x] Remove built-in mailer in favor of 3rd party package #[783](https://github.com/wiselike/revel/issues/783)
  - [x] Allow local reverse proxy access to jobs module status page for IPv4/6 #[481](https://github.com/wiselike/revel/issues/481) PR#[6](https://github.com/wiselike/revel-modules/pull/6) PR#[7](https://github.com/wiselike/revel-modules/pull/7)
  - [x] Add default http.Status code for render methods. #[728](https://github.com/wiselike/revel/issues/728)
  - [x] add domain for cookie #[770](https://github.com/wiselike/revel/issues/770) PR#[882](https://github.com/wiselike/revel/pull/882)
  - [x] production mode panic bug #[831](https://github.com/wiselike/revel/issues/831) PR#[881](https://github.com/wiselike/revel/pull/881)
  - [x] Fixes template loading order whether watcher is enabled or not #[844](https://github.com/wiselike/revel/issues/844)
  - [x] Fixes reverse routing wildcard bug PR#[886](https://github.com/wiselike/revel/pull/886) #[869](https://github.com/wiselike/revel/issues/869)
  - [x] Fixes router app start bug without routes. PR #[855](https://github.com/wiselike/revel/pull/855)
  - [x] Friendly URL template errors; Fixes template `url` func "index out of range" when param is `undefined` #[811](https://github.com/wiselike/revel/issues/811) PR#[880](https://github.com/wiselike/revel/pull/880)
  - [x] Make result compression conditional PR#[888](https://github.com/wiselike/revel/pull/888)
  - [x] ensure routes are loaded before returning from OnAppStart callback PR#[884](https://github.com/wiselike/revel/pull/884)
  - [x] Use "302 Found" HTTP code for redirect PR#[900](https://github.com/wiselike/revel/pull/900)
  - [x] Fix broken fake app tests PR#[899](https://github.com/wiselike/revel/pull/899)
  - [x] Optimize search of template names PR#[885](https://github.com/wiselike/revel/pull/885)
- `revel/cmd`:
  - [x] track current Revel version #[418](https://github.com/wiselike/revel/issues/418) PR#[858](https://github.com/wiselike/revel/pull/858)
  - [x] log path error After revel build? #[763](https://github.com/wiselike/revel/issues/763)
  - [x] Use a separate directory for revel project binaries #[17](https://github.com/wiselike/revel-cmd/pull/17) #[819](https://github.com/wiselike/revel/issues/819)
  - [x] Overwrite generated app files instead of deleting directory #[551](https://github.com/wiselike/revel/issues/551) PR#[23](https://github.com/wiselike/revel-cmd/pull/23)
- `revel/modules`:
  - [x] Adds runtime pprof/trace support #[9](https://github.com/wiselike/revel-modules/pull/9)
- Community Goals:
  - [x] Issue labels #[545](https://github.com/wiselike/revel/issues/545)
    - [x] Sync up labels/milestones in other repos #[721](https://github.com/wiselike/revel/issues/721)
  - [x] Update the Revel Manual to reflect current features
    - [x] [revel/revel.github.io/32](https://github.com/wiselike/revel.github.io/issues/32)
    - [x] [revel/revel.github.io/39](https://github.com/wiselike/revel.github.io/issues/39)
    - [x] Docs are obsolete, inaccessible TestRequest.testSuite #[791](https://github.com/wiselike/revel/issues/791)
    - [x] Some questions about revel & go docs #[793](https://github.com/wiselike/revel/issues/793)
  - [x] RFCs to organize features #[827](https://github.com/wiselike/revel/issues/827)

[Full list of commits](https://github.com/wiselike/revel/compare/v0.11.3...v0.12.0)


# v0.11.0
@brendensoares released this on 2014-10-26

Note, Revel 0.11 requires Go 1.3 or higher.

Changes since v0.10:

[BUG]   #729    Adding define inside the template results in an error (Changes how template file name case insensitivity is handled)

[ENH]   #769    Add swap files to gitignore
[ENH]   #766    Added passing in build flags to the go build command
[ENH]   #761    Fixing cross-compiling issue #456 setting windows path from linux
[ENH]   #759    Include upload sample's tests in travis
[ENH]   #755    Changes c.Action to be the action method name's letter casing per #635
[ENH]   #754    Adds call stack display to runtime panic in browser to match console
[ENH]   #740    Redis Cache: Add timeouts.
[ENH]   #734    watcher: treat fsnotify Op as a bitmask
[ENH]   #731    Second struct in type revel fails to find the controller
[ENH]   #725    Testrunner: show response info
[ENH]   #723    Improved compilation errors and open file from error page
[ENH]   #720    Get testrunner path from config file
[ENH]   #707    Add log.colorize option to enable/disable colorize
[ENH]   #696    Revel file upload testing
[ENH]   #694    Install dependencies at build time
[ENH]   #693    Prefer extension over Accept header
[ENH]   #692    Update fsnotify to v1 API
[ENH]   #690    Support zero downtime restarts
[ENH]   #687    Tests: request override
[ENH]   #685    Persona sample tests and bugfix
[ENH]   #598    Added README file to Revel skeleton
[ENH]   #591    Realtime rebuild
[ENH]   #573    Add AppRoot to allow changing the root path of an application

[FTR]   #606    CSRF Support

[Full list of commits](https://github.com/wiselike/revel/compare/v0.10.0...v0.11.0)


# v0.10.0
@brendensoares released this on 2014-08-10

Changes since v0.9.1:
- [FTR] #641 - Add "X-HTTP-Method-Override" to router
- [FTR] #583 - Added HttpMethodOverride filter to routes
- [FTR] #540 - watcher flag for refresh on app start
- [BUG] #681 - Case insensitive comparison for websocket upgrades (Fixes IE Websockets ...
- [BUG] #668 - Compression: Properly close gzip/deflate
- [BUG] #667 - Fix redis GetMulti and improve test coverage
- [BUG] #664 - Is compression working correct?
- [BUG] #657 - Redis Cache: panic when testing Ge
- [BUG] #637 - RedisCache: fix Get/GetMulti error return
- [BUG] #621 - Bugfix/router csv error
- [BUG] #618 - Router throws exception when parsing line with multiple default string arguments
- [BUG] #604 - Compression: Properly close gzip/deflate.
- [BUG] #567 - Fixed regex pattern to properly require message files to have a dot in filename
- [BUG] #566 - Compression fails ("unexpected EOF" in tests)
- [BUG] #287 - Don't remove the parent folders containing generated code.
- [BUG] #556 - fix for #534, also added url path to not found message
- [BUG] #534 - Websocket route not found
- [BUG] #343 - validation.Required(funtionCall).Key(...) - reflect.go:715: Failed to generate name for field.
- [ENH] #643 - Documentation Fix in Skeleton for OnAppStart
- [ENH] #674 - Removes custom `eq` template function
- [ENH] #669 - Develop compress closenotifier
- [ENH] #663 - fix for static content type not being set and defaulting to OS
- [ENH] #658 - Minor: fix niggle with import statement
- [ENH] #652 - Update the contributing guidelines
- [ENH] #651 - Use upstream gomemcache again
- [ENH] #650 - Go back to upstream memcached library
- [ENH] #612 - Fix CI package error
- [ENH] #611 - Fix "go vet" problems
- [ENH] #610 - Added MakeMultipartRequest() to the TestSuite
- [ENH] #608 - Develop compress closenotifier
- [ENH] #596 - Expose redis cache options to config
- [ENH] #581 - Make the option template tag type agnostic.
- [ENH] #576 - Defer session instantiation to first set
- [ENH] #565 - Fix #563 -- Some custom template funcs cannot be used in JavaScript cont...
- [ENH] #563 - TemplateFuncs cannot be used in JavaScript context
- [ENH] #561 - Fix missing extension from message file causing panic
- [ENH] #560 - enhancement / templateFunc `firstof`
- [ENH] #555 - adding symlink handling to the template loader and watcher processes
- [ENH] #531 - Update app.conf.template
- [ENH] #520 - Respect controller's Response.Status when action returns nil
- [ENH] #519 - Link to issues
- [ENH] #486 - Support for json compress
- [ENH] #480 - Eq implementation in template.go still necessary ?
- [ENH] #461 - Cron jobs not started until I pull a page
- [ENH] #323 - disable session/set-cookie for `Static.Serve()`

[Full list of commits](https://github.com/wiselike/revel/compare/v0.9.1...v0.10.0)


# v0.9.1
@pushrax released this on 2014-03-02

Minor patch release to address a couple bugs.

Changes since v0.9.0:
- [BUG] #529 - Wrong path was used to determine existence of `.git`
- [BUG] #532 - Fix typo for new type `ValidEmail`

The full list of commits can be found [here](https://github.com/wiselike/revel/compare/v0.9.0...v0.9.1).


# v0.9.0
@pushrax released this on 2014-02-26

## Revel GitHub Organization

We've moved development of the framework to the @revel GitHub organization, to help manage the project as Revel grows. The old import path is still valid, but will not be updated in the future.

You'll need to manually update your apps to work with the new import path. This can be done by replacing all instances of `github.com/robfig/revel` with `github.com/wiselike/revel` in your app, and running:

```
$ cd your_app_folder
$ go get -u github.com/howeyc/fsnotify  # needs updating
$ go get github.com/wiselike/revel
$ go get github.com/wiselike/revel-cmd/revel     # command line tools have moved
```

**Note:** if you have references to `github.com/robfig/revel/revel` in any files, you need to replace them with `github.com/wiselike/revel-cmd/revel` _before_ replacing `github.com/robfig/revel`! (note the prefix collision)

If you have any trouble upgrading or notice something we missed, feel free to hop in the IRC channel (#revel on Freenode) or send the mailing list a message.

Also note, the documentation is now at [revel.github.io](http://revel.github.io)!

Changes since v0.8:
- [BUG] #522 - `revel new` bug
- [BUG] - Booking sample error
- [BUG] #504 - File access via URL security issue
- [BUG] #489 - Email validator bug
- [BUG] #475 - File watcher infinite loop
- [BUG] #333 - Extensions in routes break parameters
- [FTR] #472 - Support for 3rd part app skeletons
- [ENH] #512 - Per session expiration methods
- [ENH] #496 - Type check renderArgs[CurrentLocalRenderArg]
- [ENH] #490 - App.conf manual typo
- [ENH] #487 - Make files executable on `revel build`
- [ENH] #482 - Retain input values after form valdiation
- [ENH] #473 - OnAppStart documentation
- [ENH] #466 - JSON error template quoting fix
- [ENH] #464 - Remove unneeded trace statement
- [ENH] #457 - Remove unneeded trace
- [ENH] #508 - Support arbitrary network types
- [ENH] #516 - Add Date and Message-Id mail headers

The full list of commits can be found [here](https://github.com/wiselike/revel/compare/v0.8...v0.9.0).


# v0.8
@pushrax released this on 2014-01-06

Changes since v0.7:
- [BUG] #379 - HTTP 500 error for not found public path files
- [FTR] #424 - HTTP pprof support
- [FTR] #346 - Redis Cache support
- [FTR] #292 - SMTP Mailer
- [ENH] #443 - Validator constructors to improve `v.Check()` usage
- [ENH] #439 - Basic terminal output coloring
- [ENH] #428 - Improve error message for missing `RenderArg`
- [ENH] #422 - Route embedding for modules
- [ENH] #413 - App version variable
- [ENH] #153 - $GOPATH-wide file watching aka hot loading


# v0.6
@robfig released this on 2013-09-16



