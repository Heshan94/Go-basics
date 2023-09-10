# How to create a workspace?
- mkdir project_name
- cd project_name
- go work init

# Create modules inside workspace
- mkdir moduleName
- cd moduleName
- touch main.go

# Use the module in the go.work
- cd ../
- pwd 
  project_name
- go work use moduleName
https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory

# How to run ?
- cd moduleName
- go run main.go

# How to build the binary ?
- cd moduleName
- go build .

