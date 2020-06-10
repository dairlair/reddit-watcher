# Reddit Watcher
A basic application which listens to Reddit posts and stores them into the storage.

# How to use it

Create the Docker image locally (without pushing it to the registry):
```shell script
make image
```

# Application structure
The application contains two parts: 

* The Watcher in `pkg/watcher` (based on the `github.com/turnage/graw` reddit client) 
* The Storage in `pkg/storage` (based on the `go.mongodb.org/mongo-driver/mongo` Mongo DB driver)

The Storage provides functionality to save posts in the Mongo DB.
The Watcher just listens Reddit and save received posts in to the storage.