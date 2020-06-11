# Reddit Watcher
A basic application which listens to Reddit posts and stores them into the storage.

# How to deploy the application to your Kubernetes cluster
This application uses Kubernetes Secrets to store sensitive credentials for the Reddit Bot.

Before deploy this app you need to create secret in the k8s cluster with the name 'reddit-watcher'.

You can use this command to do that:

```shell script
kubectl create secret generic reddit-watcher \
    --from-literal=REDDIT_SUBREDDITS={subreddits to watch} \
    --from-literal=REDDIT_USER_AGENT={reddit user agent} \
    --from-literal=REDDIT_CLIENT_ID={reddit client ID} \
    --from-literal=REDDIT_SECRET={reddit secred} \
    --from-literal=REDDIT_USERNAME={reddit username} \
    --from-literal=REDDIT_PASSWORD={reddit password} \
    --from-literal=MONGODB_URI={MongoDB URI} \
    --from-literal=MONGODB_DATABASE={MongoDB database}
```
Replace placeholders to the actual values and run this command. After that check created secret with this command:

```shell script
kubectl describe secrets/reddit-watcher
```

When it is done run this command:
@TODO: Describe how to run

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