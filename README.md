# gowebapp
Go - Simple Web Application

A simple program that allows you to perform CRUD operations in a store using PostgreSQL.

- File main.go:

Here lies the core of the project. Initially, the routes are initialized, followed by the opening of a local server.

- Directory templates:

Here are the main templates of the project; those whose names start with _ are used as common components for the other templates. Essentially, we need three templates that allow us to create, view, edit, and delete data.

- Directory routes:

We possess a file responsible for initializing and managing the routes in our project. Upon calling a route, we execute specific operations that are constructed within the controllers.
The controllers, for example, can load a template or do some operation like create a data.

- Directory models:

In models, the structure of the data is defined. Additionally, functions are built to operate on the database based on this structure.

- Directory DB:

Just for make DB operations like connect. The connection is made only when we wanna do operations in database, after this we close it.

- Directory controllers:

As stated in the routers, we define what we want to do when calling a route.

