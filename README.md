DataDeck
===================

This program allows you to run a service on port 8000 of localhost and query the API via the url: [localhost:8000/search/songs.json](localhost:8000/search/songs.json)

**Please make sure that port 8000 is available on your machine.**

Parameters
-------------
 - All parameters are optional.
 - The order of the parameters does not matter. 
 - If there are no parameters, all songs will be returned. 
 - Before adding the parameters you must add the symbol "?".
 - If there is more than one parameter must be separated by the symbol "&".

|  Name  |  Type  |                 Description                |       Example       |
|:------:|:------:|:------------------------------------------:|:-------------------:|
| name   | string | Return songs that match the name of artist | I Gotta Feeling     |
| artist | string | Return songs that match the name of song   | The Black Eyed Peas |
| genre  | string | Return songs that match the name of genre  | Rap                 |


Examples Request
-------------

**Request**

http://localhost:8000/search/songs.json?name=Gala

**Result**

>[{"artist":"424","song":"Gala","genre":"Indie Rock","length":189}]

----------
**Request**

http://localhost:8000/search/songs.json?genre=Pop

**Result**

>[{"artist":"Chubby Checker","song":"The Twist","genre":"Pop","length":235},{"artist":"Santana","song":"Smooth","genre":"Pop","length":167},{"artist":"Los Del Rio","song":"Macarena","genre":"Pop","length":159},{"artist":"Olivia Newton-John","song":"Physical","genre":"Pop","length":195},{"artist":"Debby Boone","song":"You Light Up My Life","genre":"Pop","length":245}]

----------
**Request**

http://localhost:8000/search/songs.json?artist=LMFAO

**Result**

>[{"artist":"LMFAO","song":"Party Rock Anthem","genre":"Rap","length":189}]

----------
**Request**

http://localhost:8000/search/songs.json?artist=LMFAO&genre=Rap

**Result**

>[{"artist":"LMFAO","song":"Party Rock Anthem","genre":"Rap","length":189}]

----------

**Request**

http://localhost:8000/search/songs.json?name=Party%20Rock%20Anthem&artist=LMFAO&genre=Rap

**Result**

>[{"artist":"LMFAO","song":"Party Rock Anthem","genre":"Rap","length":189}]