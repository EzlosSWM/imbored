# imbored
## A CLI tool that fetches random activities from the BoredAPI.
imbored will randomly generate activities when you're bored. It you can filter it by different types and number of participants. 

A list of accepted filter types: 
- education
- recreational
- social
- diy
- charity
- cooking
- relaxation
- music
- busywork

## Installation
```bash
git clone https://github.com/EzlosSWM/imbored.git
```
Note: You can either modify the script or use it as is, or you can use the pre-compiled script within the builds folder

## Usage
```bash
go build -o imbored
./imbored random # Prints random activity 
./imbored random --type <type of activity> # Filters the activity by defined type
./imbored random --people <int> # FIlters the activity by the number of participants
./imbored random --activities # Prints a list of accepted activities
```

## Try it with Docker 
```bash 
docker build -t imbored .
docker run imbored random 
```

## Want to contribute?
Eventually, I'd like to add make use of both "--type" & "--people" together.
Ex. 
```Bash 
$ im bored random --type=recreational --people=5
$ Have a jam session with your friends
```
