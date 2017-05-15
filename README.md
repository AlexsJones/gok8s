```

_______________/\\\________________________________/\\\__        
 ______________\/\\\_______________________________\/\\\__       
  ______________\/\\\_______________________________\/\\\__      
   __/\\\\\\\\\\_\/\\\_____________/\\\\\\\\_________\/\\\__     
    _\/\\\//////__\/\\\\\\\\\\____/\\\/////\\\___/\\\\\\\\\__    
     _\/\\\\\\\\\\_\/\\\/////\\\__/\\\\\\\\\\\___/\\\////\\\__   
      _\////////\\\_\/\\\___\/\\\_\//\\///////___\/\\\__\/\\\__  
       __/\\\\\\\\\\_\/\\\___\/\\\__\//\\\\\\\\\\_\//\\\\\\\/\\_
        _\//////////__\///____\///____\//////////___\///////\//__
  ```

  Shell script scheduling.

  Automate repetitive tasks and save them as a `Shedfile`

  Hooya.

# Install
```
go get github.com/AlexsJones/shed
```

  ```
  clear      clear the current stack
  exit       exit the program
  help       display help
  list       list execution order
  load       Loads a local ShedFile into a schedule
  logs       logs from an execution
  push       push a k8s config-map path
  retry      retry a certain action based on index
  run        Starts running k8s config-map paths
  save       Saves out a new ShedFile
  ```

  # Example

  I want to automate a simple workflow for kubernetes deployment...

```
# In my kubernetes project directory

>>> push kubectl config view
Pushing -> kubectl
>>> push ls
Pushing -> ls
>>> list
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | kubectl          | ✓         | ✗        | ?          |
|    2 | ls               | ✓         | ✗        | ?          |
+------+------------------+-----------+----------+------------+
>>> push kubectl create -f .
Pushing -> kubectl
>>> list
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | kubectl          | ✓         | ✗        | ?          |
|    2 | ls               | ✓         | ✗        | ?          |
|    3 | kubectl          | ✓         | ✗        | ?          |
+------+------------------+-----------+----------+------------+
>>> save
Created new Shedfile...
>>> run
```
