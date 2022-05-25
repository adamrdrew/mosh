![Alt text](docs/mosh_logo.png "a title")

Mosh (Music On Shell) is a Plex music player for the command line. Think of it like PlexAMP for the shell. 

## Setup
Mosh needs to authenticate to your Plex server and access a music library on it. The `mosh setup` command handles all of this for you. The steps are run `mosh setup`, click the link it provides you to authenticate to Plex, and then select the library you want to use from the text menu.

```shell
 $ mosh setup
Welcome to MOSH! 🎧💿🐧
🔑 Checking Plex authorization status...
    ❌ Authorization required.
🔒 Obtaining token...

Authorize Mosh with a web browser. It should open automatically. If it doesnt then open it manually with the following URL:
https://app.plex.tv/auth#?clientID=f79d7735-864b-4ed7-a6dc-a3971824843b&code=3rieur1uptnon8tsabld80jjo&context%5Bdevice%5D%5Bproduct%5D=Mosh

💻 Waiting for authorization...
.....
Select a music library to use.
Type the number to the left of the name of the library you want.
    0) Adam Music
    1) Roddy Music
0
Library set.
We are ready to ROCK!!!!
```

## Searching
Before you can play something you need to find it. You can use the `mosh search` command along with the `artist` and `album` subcommands to search for albums or artists:

```shell
 $ mosh search artist dillinger
Welcome to MOSH! 🎧💿🐧
+------+---------------------------+
| ID   | TITLE                     |
+------+---------------------------+
| 2209 | The Dillinger Escape Plan |
+------+---------------------------+

 $ mosh search album hail
Welcome to MOSH! 🎧💿🐧
album command
+--------+----------+------------+
| ID     | TITLE    | ARTIST     |
+--------+----------+------------+
| 100358 | All Hail | Norma Jean |
+--------+----------+------------+

```

## ID Numbers
Notice the `ID` column in the search output. These ID numbers are how you refer to specific obects like albums, artists, and songs in most Mosh commands. It is a lot easier to search, get the ID, and then use the ID than it is to type an exact album, sort, or artist name.

## Listing - the ls command
If you want to see the albums by an artist or the songs on an album you can use the `mosh ls` command to list them.

```shell
 $ mosh search artist pantera
Welcome to MOSH! 🎧💿🐧
+--------+---------+
| ID     | TITLE   |
+--------+---------+
| 100034 | Pantera |
+--------+---------+

 $ mosh ls artist 100034
Welcome to MOSH! 🎧💿🐧
+--------+------------------------------+---------+
| ID     | TITLE                        | ARTIST  |
+--------+------------------------------+---------+
| 100085 | Cowboys From Hell            | Pantera |
| 100048 | Far Beyond Driven            | Pantera |
| 100035 | The Great Southern Trendkill | Pantera |
| 100074 | Reinventing the Steel        | Pantera |
| 100062 | Vulgar Display of Power      | Pantera |
+--------+------------------------------+---------+

 $ mosh ls album 100062
Welcome to MOSH! 🎧💿🐧
+--------+-------+------------------------------+-------------------------+---------+
| ID     | TRACK | TITLE                        | ALBUM                   | ARTIST  |
+--------+-------+------------------------------+-------------------------+---------+
| 100063 | 1     | Mouth for War                | Vulgar Display of Power | Pantera |
| 100064 | 2     | A New Level                  | Vulgar Display of Power | Pantera |
| 100065 | 3     | Walk                         | Vulgar Display of Power | Pantera |
| 100066 | 4     | Fucking Hostile              | Vulgar Display of Power | Pantera |
| 100067 | 5     | This Love                    | Vulgar Display of Power | Pantera |
| 100068 | 6     | Rise                         | Vulgar Display of Power | Pantera |
| 100069 | 7     | No Good (Attack the Radical) | Vulgar Display of Power | Pantera |
| 100070 | 8     | Live in a Hole               | Vulgar Display of Power | Pantera |
| 100071 | 9     | Regular People (Conceit)     | Vulgar Display of Power | Pantera |
| 100072 | 10    | By Demons Be Driven          | Vulgar Display of Power | Pantera |
| 100073 | 11    | Hollow                       | Vulgar Display of Power | Pantera |
+--------+-------+------------------------------+-------------------------+---------+
```




