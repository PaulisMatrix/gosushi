# gosushi
Repo to keep track of my learning journey of golang


# Taskfile(Alternative to makefile)

One of the use case

1.  Can use its `watch` command for hot reloading of a go server so you dont have to do `go run main.go` everytime you save a file.

2.  Use it with `task <task_name> --watch`

3.  Mention files to watch under `sources`

4.  Mention which method to use for checking if the file has changes or not. Default is checksum which compares hash of
    your files under `sources` to check if they have been changed. You can also specify timestamp, then check will be
    performed using the timestamp of the file modified.

5. More [here](https://taskfile.dev/usage/#watch-tasks)

6. Or just use [air](https://github.com/cosmtrek/air)