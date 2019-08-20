## Go Downloader
A simple *File Download Manager* application which is able to accept a list of URLs to be downloaded through a REST endpoint and download those files locally on the system.
# Installation
```
go run main.go
```
# Design 
| Sr No.       | API                      | Comment    |
| :------      | :--------------:         | ----:      |
|    FR1       | ==GET/health==>          | return 200 OK if the application is up and running     |
|      FR2     | ==GET/downloads==>       | To download(serial/concurrent)     |
| FR3          |  ==GET/download/{id}==>  | To see the download status     |

