## IMDBot 

[![License-IMDBot](https://img.shields.io/hexpm/l/plug)](https://github.com/Darkhound-org/IMDBot/blob/bots_exes/LICENSE.txt)
[![Build](https://img.shields.io/appveyor/build/gruntjs/grunt)](https://github.com/Darkhound-org/IMDBot/releases/tag/20220802)
[![Cross Platform](https://img.shields.io/powershellgallery/p/DNS.1.1.1.1)](https://github.com/Darkhound-org/IMDBot#download)

A command line bot built to search details about your favourite movies/series

* [DOWNLOAD](https://github.com/Darkhound-org/IMDBot#download)
* [INSTALLATION]()
* [USAGE]()
* [DEVELOPING]()
* [API]()
* [LICENSE]()

### Download
Download the latest release from Github. IMDBot is a cross platform application and is compatible with 64 bit systems only.

### Installation
Unzip the application in any folder. IMDBot is a portable application. 
Make sure all files are properly organized.
```
   IMDBot/
      |
      |----IMDBot.exe
      |----key.exe
      |----LICENSE.txt
      |----Version
      |----data [created]
      |----movie_log.txt [created]
      └── README.md
      
Note : File extensions can vary according to your OS      
```
In this directory structure, data and movie_log.txt are created by the application on various instances.

### Usage
Note: Usage on Windows. Similar usage for other operating systems. 
Get an OMDb API from https://www.omdbapi.com/ . Refer [Get_api_key_from_omdbapi.com.md](https://github.com/Darkhound-org/IMDBot/blob/bots_exes/Get_api_key_from_omdbapi.com.md) for detailed info.
Enter the api key by running key.exe file. A data file will be created.
Open cmd.exe from the IMDBot folder and type `IMDBot.exe --help` .This will display all the commands and flags the app has along with the short description.
```
IMDBot can search information about your favourite movies and tv series right from the terminal.

Usage:
  IMDBot [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  search      Searches your favourite movie or TV series
  test        Test run

Flags:
  -h, --help      help for IMDBot
  -t, --toggle    Help message for toggle
  -v, --version   version for IMDBot

Use "IMDBot [command] --help" for more information about a command.
```
Let's run the `test` command to make sure everything is fine. The output would be this if it works without any bug.
```
Response url (http://www.omdbapi.com/?t=Interstellar&apikey="YOUR_API_KEY")
2022/08/03 12:22:46 !!..Open response url for detailed description of the movie or series..!!
2022/08/03 12:22:46 !!..No bugs found. Successfully connected to IMDB api..!!
2022/08/03 12:22:46 [ !!..IMDBot is okay..!! ]

```
Now let's search a movie. For that run the following command `IMDBot.exe search` . This will provide you with 3 choices
```
+---+-------------------+
| # | COMMANDS          |
+---+-------------------+
| 1 | Search by name    |
| 2 | Search by IMDB id |
| 3 | Exit              |
+---+-------------------+
Enter your choice [1,2 or 3] : 
```
As the commands suggest by entering choice as 1 , you can search your favourite movie by name and by entering 2, by IMDB id. An example is shown below.
```
+---+-------------------+
| # | COMMANDS          |
+---+-------------------+
| 1 | Search by name    |
| 2 | Search by IMDB id |
| 3 | Exit              |
+---+-------------------+
Enter your choice [1,2 or 3] : 1
Name of movie: Titanic
Response url (http://www.omdbapi.com/?t=Titanic&apikey="YOUR_API_KEY")
2022/08/03 12:27:53 !!..Open response url for detailed description of the movie or series..!!
2022/08/03 12:27:53 Output saved to movie_log.txt..!!
```
```
+---+-------------------+
| # | COMMANDS          |
+---+-------------------+
| 1 | Search by name    |
| 2 | Search by IMDB id |
| 3 | Exit              |
+---+-------------------+
Enter your choice [1,2 or 3] : 2
Enter IMDB id: tt0060196
Response url (http://www.omdbapi.com/?i=tt0060196&apikey="YOUR_API_KEY")
2022/08/03 12:28:48 !!..Open response url for detailed description of the movie or series..!!
2022/08/03 12:28:48 Output saved to movie_log.txt..!!
```
The response url you get is saved to movie_log.txt and is overwritten each time. In some terminals the url will be hyperlinked in the text output `Response url`. 
In the above example you have to copy paste the url into your browser to view the full output/description of the movie.
To get poster of the movie copy-paste the poster url into your browser.

### Developing
IMDBot is built in Golang with the help of Cobra-cli and OMDB api. So, first install Golang and set all go environment variables. Clone this repository `https://github.com/Darkhound-org/IMDBot.git` and make a bin folder in it. cd to IMDBot folder, Set GOBIN as the bin folder and GOPATH as IMDBot folder.
Run the following commands
```
go build main.go
```
cd to key folder and run
```
go build key.go
```
Now delete key.go and main.go files, move key.exe and main.exe [rename as IMDBot.exe] from bin folder to IMDBot and delete bin and key folders.

### Api
IMDBot uses IMDB api indirectly through the OMDBapi. Refer https://www.omdbapi.com/ for more details. For detailed info of getting an api key refer [Get_api_key_from_omdbapi.com.md](https://github.com/Darkhound-org/IMDBot/blob/bots_exes/Get_api_key_from_omdbapi.com.md) . 

### License
IMDBot is licensed under the [Apache License 2.0](https://github.com/Darkhound-org/IMDBot/blob/bots_exes/LICENSE.txt) 






