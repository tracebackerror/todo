# Goal:

Create REST API for a TO-DO application (just backend, not the UI) in Golang, which will have
the following features:
* Create a task list
* Add task into a list
* Delete task from a list
* Update an existing task
* Delete a list
Expose necessary endpoints for the above features and describe how to use the APIs with the
help of certain examples (use readme file in GitHub repository for this).



## To Start the app configure MySql Connection on Line 71 of rest.go

##How to test?
* Once connected Hit host/rest/ping -> Message from server will be received.

## Endpoints Implemented:
GET:  Fetch the data from the below todo
		   *    SUPPORTED PATH:
		   **       /rest/task
		   **       /rest/task/
		   **       /rest/task/id
		   **       /rest/task/id/
		   * Returns:
		   **       Status Code:200

		   POST: Create a todo in database
		   *    SUPPORTED PATH
		   **       /rest/task/
		   **       /rest/task
		   * Parameter:
		   **       Body:    "JSON Message of format" - {"task": "one", "status": "incomplete"}
		   * Returns:
		   **       Status Code:200

		   PUT: Update a todo in database
		   * SUPPORTED PATH
		   **       /rest/task/id/
		   **       /rest/task/id
		   * Parameter:
		   **       Body:    "JSON Message of both field" - {"task": "one", "status": "incomplete"}
		   * Returns:
		   **       Status Code:200

		   DELETE: Deletes a todo in database
		   * SUPPORTED PATH
		   **       /rest/task/id/
		   **       /rest/task/id
		   * Parameter:
		   **       Body:    None
		   * Returns:
		   **       Status Code:200
