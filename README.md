# GoWebMock
Simple Go server to mock web api's

The code is not the best written, but it works. optimizations will come later.

#Usage:
Everything is configured through a config json. If no config file is specified (-cfg) then the server tries to load autoexec.json in the same directory.

There are two types of responses. Static and dynamic.
Static responses serve either a string or a file from a given path
Dynamic responses execute javascript to build the response.

Inside the the javascript, there are four predefined variables:
request: Holds the current request object.
config:  Holds the config, which can be defined in the autoexec.json
storage: A simple javascript storage, which persists data between calls. (the storage gets cleared when the server is restarted)
header:  Here you can define the headers for the response
