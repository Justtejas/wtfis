#WTFIS
- wtfis is an command line tool to query what a word/terminology means.
  
## Local Installation

- Make sure your GOPATH variables are set.
- Clone the repository ```sh git clone https://github.com/Justtejas/wtfis.git```
- Run ```sh go build -o wtfis.exe``` if on windows or else ```sh go build -o wtfis```
- Move the exe to the system32 folder to run it from anywhere ```sh move wtfis.exe C:\Windows\System32```, on linux ```sh sudo mv wtfis /usr/local/bin/```

## Usage
- If you want to just search a word use ```sh wtfis word```
- If you want to search a term with multiple words use ```sh wtfis "word word"```