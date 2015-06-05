# GoWebMock

This is an easy to set up, mock server for web api's. I am a desktop programmer and have to interact with web api's from our customers quite frequently. But most of the time, the api is either unfinished, unstable or not existent at all. (Early in development)

So I needed a way to quickly mock any api to start programming. The result is this small mock server written in go.

It's nothing fancy, it's not optimized and the documentation is minimal. This will get better in the next days, but it's fairly simple to set up.

There is also no guarantee, that the server can handle more than a few connections. I haven't tested it yet, but for small tests it should defenitly be enough.

A demo config is included and it shouldn't be too hard to figure out how to use the server.

If you have any questions, write me an e-mail through github.

I hope someone has good use for this server.


#Usage:
Everything is configured through a config json. If no config file is specified (-cfg) then the server tries to load autoexec.json in the same directory.

There are two types of responses. Static and dynamic.
- Static responses serve either a string or a file from a given path
- Dynamic responses execute javascript to build the response.

Inside the javascript, there are the following predefined variables:
- request: Holds the current request object.
- config:  Holds the config, which can be defined in the autoexec.json
- storage: A simple javascript storage, which persists data between calls. (the storage gets cleared when the server is restarted)
- header:  Here you can define the headers for the response
- httpStatusCode: the response code, which should be returned

Everything located in the static folder, will be served like a webserver would.

To persist data, you can use the GlobalStorage function.
- This function takes a callback. Inside this callback you can access a global object called "data", in which the everything is stored.
- The data object is persisted to disc when the process terminates and loaded when the process starts.
- The data object is only accessible through the GlobalStorage function, because of the mutex, to prevent race conditions.
- The GlobalStorage functions returns the return value of the callback.

e.g.:
```
var globalCount = GlobalStorage(function() {
	data.count = data.count | 0; // Initialize variable, if it doesn't exist
	data.count++;
	return data.count;
});

console.log("globalCount: " + globalCount);
```
