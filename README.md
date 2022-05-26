![Alt text](docs/mosh_logo.png "a title")

Mosh (Music On Shell) is a Plex music player for the command line. Think of it like PlexAMP for the shell. Listen to music without breaking your flow. Supports playing via remote access (doesn't require you be on the same network as your Plex server) and uses Plex.tv's secure authorization.

## Requirements
For a compiled release you shouldn't need anything, just the executables. If you are building from source or hacking you'll need the ALSA development libs:

RHEL / Fedora:
```
$ sudo yum install alsa-lib-devel
```

Debian / Ubuntu:
```
$ sudo apt-get install libasound2-dev
```

## Setup
Mosh needs to authenticate to your Plex server and access a music library on it. The `mosh setup` command handles all of this for you. The steps are run `mosh setup`, click the link it provides you to authenticate to Plex, and then select the library you want to use from the text menu.

```
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

```
 $ mosh search artist norma
+------+------------+----------+
| ID   | TITLE      | SHORTCUT |
+------+------------+----------+
| 2956 | Norma Jean |          |
+------+------------+----------+

 $ mosh search album hail
+--------+----------+------------+----------+
| ID     | TITLE    | ARTIST     | SHORTCUT |
+--------+----------+------------+----------+
| 100358 | All Hail | Norma Jean |          |
+--------+----------+------------+----------+

```

## ID Numbers
Notice the `ID` column in the search output. These ID numbers are how you refer to specific obects like albums, artists, and songs in most Mosh commands. It is a lot easier to search, get the ID, and then use the ID than it is to type an exact album, sort, or artist name.

## Listing - the ls command
If you want to see the albums by an artist or the songs on an album you can use the `mosh ls` command to list them.

```
 $ mosh search artist pantera
+--------+---------+----------+
| ID     | TITLE   | SHORTCUT |
+--------+---------+----------+
| 100034 | Pantera |          |
+--------+---------+----------+


 $ mosh ls artist 100034
+--------+------------------------------+---------+----------+
| ID     | TITLE                        | ARTIST  | SHORTCUT |
+--------+------------------------------+---------+----------+
| 100085 | Cowboys From Hell            | Pantera |          |
| 100048 | Far Beyond Driven            | Pantera |          |
| 100035 | The Great Southern Trendkill | Pantera |          |
| 100074 | Reinventing the Steel        | Pantera |          |
| 100062 | Vulgar Display of Power      | Pantera |          |
+--------+------------------------------+---------+----------+


 $ mosh ls album 100085
+--------+-------+------------------------+-------------------+---------+
| ID     | TRACK | TITLE                  | ALBUM             | ARTIST  |
+--------+-------+------------------------+-------------------+---------+
| 100086 | 1     | Cowboys From Hell      | Cowboys From Hell | Pantera |
| 100087 | 2     | Primal Concrete Sledge | Cowboys From Hell | Pantera |
| 100088 | 3     | Psycho Holiday         | Cowboys From Hell | Pantera |
| 100089 | 4     | Heresy                 | Cowboys From Hell | Pantera |
| 100090 | 5     | Cemetery Gates         | Cowboys From Hell | Pantera |
| 100091 | 6     | Domination             | Cowboys From Hell | Pantera |
| 100092 | 7     | Shattered              | Cowboys From Hell | Pantera |
| 100093 | 8     | Clash With Reality     | Cowboys From Hell | Pantera |
| 100094 | 9     | Medicine Man           | Cowboys From Hell | Pantera |
| 100095 | 10    | Message in Blood       | Cowboys From Hell | Pantera |
| 100096 | 11    | The Sleep              | Cowboys From Hell | Pantera |
| 100097 | 12    | The Art of Shredding   | Cowboys From Hell | Pantera |
+--------+-------+------------------------+-------------------+---------+

```

## The Mosh Daemon - moshd
So far we've been looking at the mosh CLI and running queries against Plex. In order to play music we'll need to run `moshd`: The Mosh Daemon. Start `moshd` and it'll fork into the background. There's nothing else to set up or think about.

**Important: You need to have run `mosh setup` before you can start `moshd`**

```
$ ./moshd
Starting moshd...

$
```
Now that `moshd` is running we can look at playing some music!

## Play an Album
Use the `mosh play album` command to play an album:

```
 $ mosh play album 74594
Hesitation Marks by Nine Inch Nails is now playing.
```

## See what's playing
Use the `mosh get playing` command to see what's playing:

```
 $ mosh get playing
+-----------+-----------------+------------------+
| TRACK     | ARTIST          | ALBUM            |
+-----------+-----------------+------------------+
| Copy of a | Nine Inch Nails | Hesitation Marks |
+-----------+-----------------+------------------+
5:23 / 5:85 [#################---] 89 %

```

Use the `most get queue` command to see the entire play queue:

```
$ mosh get queue
+---------+-------+------------------------------+-------------------+--------+
| PLAYING | TRACK | TITLE                        | ALBUM             | ARTIST |
+---------+-------+------------------------------+-------------------+--------+
| X       | 1     | Violet                       | Live Through This | Hole   |
|         | 2     | Miss World                   | Live Through This | Hole   |
|         | 3     | Plump                        | Live Through This | Hole   |
|         | 4     | Asking for It                | Live Through This | Hole   |
|         | 5     | Jennifer’s Body              | Live Through This | Hole   |
|         | 6     | Doll Parts                   | Live Through This | Hole   |
|         | 7     | Credit in the Straight World | Live Through This | Hole   |
|         | 8     | Softer, Softest              | Live Through This | Hole   |
|         | 9     | She Walks on Me              | Live Through This | Hole   |
|         | 10    | I Think That I Would Die     | Live Through This | Hole   |
|         | 11    | Gutless                      | Live Through This | Hole   |
|         | 12    | Olympia                      | Live Through This | Hole   |
+---------+-------+------------------------------+-------------------+--------+
```

## Playback Controls
You can stop, play, and skip around the play queue with the `mosh control` commands:

```
 $ mosh control next
Went forward. Next up: Miss World by Hole from the album Live Through This

 $ mosh control next
Went forward. Next up: Plump by Hole from the album Live Through This

 $ mosh control back
Went back. Next up: Miss World by Hole from the album Live Through This

 $ mosh control stop
Playback stopped.

 $ mosh control play
Playing: Violet by Hole from the album Live Through This

```

## Shortcuts
ID numbers can be a drag when you are dealing with stuff you listen to a lot. Shortcuts to the rescue! Use shortcuts to create easy to remember and type keys for your favorite artists or albums. Then, when you want to play or explore those simply use the shortcut instead of looking up the ID.

```
 $ mosh search artist dillinger
+------+---------------------------+----------+
| ID   | TITLE                     | SHORTCUT |
+------+---------------------------+----------+
| 2209 | The Dillinger Escape Plan |          |
+------+---------------------------+----------+

 
 $ mosh shortcuts add tdep 2209
 
 $ mosh ls artist tdep
+-------+------------------+---------------------------+----------+
| ID    | TITLE            | ARTIST                    | SHORTCUT |
+-------+------------------+---------------------------+----------+
| 74932 | Ire Works        | The Dillinger Escape Plan |          |
| 74920 | Miss Machine     | The Dillinger Escape Plan |          |
| 74909 | Option Paralysis | The Dillinger Escape Plan |          |
+-------+------------------+---------------------------+----------+

 
 $ mosh shortcuts add ireworks 74932
 
 $ mosh play album ireworks
Ire Works by The Dillinger Escape Plan is now playing.

 $ mosh shortcuts list
+----------+-------+
| SHORTCUT | ID    |
+----------+-------+
| tdep     | 2209  |
| ireworks | 74932 |
| nin      | 328   |
+----------+-------+
```

*Note: You can use `mosh ls shortcuts` to view shortcuts as well as `mosh shortcuts list`.* 

If you make a mistake `shortcuts delete` has you covered:

```
 $ mosh shortcuts add oopsmistake 74932

  $ mosh shortcuts list
+-------------+-------+
| SHORTCUT    | ID    |
+-------------+-------+
| nin         | 328   |
| oopsmistake | 74932 |
| tdep        | 2209  |
| ireworks    | 74932 |
+-------------+-------+
 
 $ mosh shortcuts delete oopsmistake
 
 $ mosh shortcuts list
+----------+-------+
| SHORTCUT | ID    |
+----------+-------+
| ireworks | 74932 |
| nin      | 328   |
| tdep     | 2209  |
+----------+-------+
```