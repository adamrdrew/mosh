![Alt text](docs/mosh_logo.png "a title")

Mosh (Music On Shell) is a Plex music player for the command line. Think of it like PlexAMP for the shell. Listen to music without breaking your flow. Supports playing via remote access (doesn't require you be on the same network as your Plex server) and uses Plex.tv's secure authorization.

## Requirements
Before you can build mosh you'll need the alsa dev package

RHEL / Fedora:
```
$ sudo yum install alsa-lib-devel
```

Debian / Ubuntu:
```
$ sudo apt-get install libasound2-dev
```

## Install

1. Clone the repo
2. Run `make build`
2. Run `sudo make install`

*Note: If you already have mosh installed and are upgrading make sure you run `mosh daemon stop` before running `make install`.*

## Setup
Mosh needs to authenticate to your Plex server and access a music library on it. The `mosh setup` command handles all of this for you. The steps are run `mosh setup`, click the link it provides you to authenticate to Plex, and then select the library you want to use from the text menu.

```
 $ mosh setup
Welcome to MOSH! π§πΏπ§
π Checking Plex authorization status...
    β Authorization required.
π Obtaining token...

Authorize Mosh with a web browser. It should open automatically. If it doesnt then open it manually with the following URL:
https://app.plex.tv/auth#?clientID=f79d7735-864b-4ed7-a6dc-a3971824843b&code=3rieur1uptnon8tsabld80jjo&context%5Bdevice%5D%5Bproduct%5D=Mosh

π» Waiting for authorization...
.....
Select a music library to use.
Type the number to the left of the name of the library you want.
    0) Adam Music
    1) Roddy Music
0
Library set.

Starting mosh daemon...
Daemon status OK - PID: 500198

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
|         | 5     | Jenniferβs Body              | Live Through This | Hole   |
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

## Config

### Directories and Ports
Mosh has a bunch of files, directories, and ports it cares about. There are sane defaults for everything, but you can also configure them all by hand with environment variables. If you decide to change any of these I assume you know enough about what you are doing to deal with any side effects.

| Description | Default | Environment Variable |
| --------------- | --------------- | --------------- |
| Config file dir | `/etc/mosh` | `MOSH_CONFIG_DIR` |
| Log dir | `/var/log/mosh` | `MOSH_LOG_DIR` |
| PID dir | `/tmp` | `MOSH_PID_DIR` |
| Daemon Port | `9666` | `MOSH_PORT` |
| Cache dir | `/tmp/mosh` | `MOSH_CACHE_DIR` |

### Config Command
You can use the `mosh config` command to change various runtime options. The options are as follows:

| Command | Valid Values | Description |
| -- | -- | -- |
| `mosh config cache-max-days` | 1 through 30 | Files with access times older than this value will be deleted on next cache prune |
| `mosh config cache-max-size` | 1024 through 16384 | If cache size exceeds this amount in MB files will be deleted from cache |
| `mosh config show-art` | true or false | Will display album art in `get playing` and `ls album` output | 

## Cache
Music you listen to is cached locally. When you play music that you've played previously mosh plays from cache. The cache system will periodically free up space by deleting old items. If the cache gets full the oldest items will be deleted until cache utilization is at 50% of max.

## Album Art
Mosh can display an ascii art representation of album art. Use the `mosh config show-art` command to enable or disable the feaure.

![Alt text](docs/art-example.png "a title")

## Debugging and Errors
The most common errors you are likely to hit are daemon not running or setup not found. The following sections should help you debug them. If you are still having trouble feel free to [opn an issue](https://github.com/adamrdrew/mosh/issues/new) along with as much info as you can (command output, daemon log file, etc.)
### Daemon Not Running
If you run a command and get this error then the daemon isn't running:
```
$ mosh play album 324
PID file not found. Daemon not running?
```
You can use the daemon commands to check status and start:
```
 $ mosh daemon status
PID file not found. Daemon not running?

 $ mosh daemon start
Starting mosh daemon...
```
If for some reason you suspect the daemon is misbehaving you can run `mosh daemon stop ; mosh daemon start` to restart it. If you are still having problems check out the log file at `/var/log/mosh/moshd.log` or wherever you set the log file if you didn't use defaults.
### Token Missing / Setup Incomplete
If you try to use mosh either without having run setup or with the config otherwise missing or damaged you'll get this error:
```
 $ mosh search artist cave
Plex token not found. Please run setup.
```
If you haven't run setup run it with `mosh setup`. If you did run setup this means that for some reason mosh can't find the config file. The default location is `/etc/mosh/config.yaml` - though it could be in an alternate location if you configured it as such. Running `mosh setup` again and reauthorizing may help.

## Development
You can use the `mosh-dev.sh` script to run a dev copy of mosh right from source. The dev instance uses a different config than prod so you can run dev and prod side by side without worrying about your config getting messed up. All of the dev instance's config files, log files, etc will show up in `mosh_tmp` which will be created on the fly. The `mosh-dev.sh` script sets up the config and then sends all remaining args to mosh, so use it just like you do the mosh comamnd:


```
$ ./mosh-dev.sh get playing
+---------------------------+-----------------+------------+
| TRACK                     | ARTIST          | ALBUM      |
+---------------------------+-----------------+------------+
| All the Love in the World | Nine Inch Nails | With Teeth |
+---------------------------+-----------------+------------+
2:48 / 5:25 [#########-----------] 47 %

 $ ./mosh-dev.sh get queue
+---------+-------+------------------------------+-----------+-----------------+
| PLAYING | TRACK | TITLE                        | ALBUM     | ARTIST          |
+---------+-------+------------------------------+-----------+-----------------+
|         | 1     | HYPERPOWER!                  | Year Zero | Nine Inch Nails |
| X       | 2     | The Beginning of the End     | Year Zero | Nine Inch Nails |
|         | 3     | Survivalism                  | Year Zero | Nine Inch Nails |
|         | 4     | The Good Soldier             | Year Zero | Nine Inch Nails |
|         | 5     | Vessel                       | Year Zero | Nine Inch Nails |
|         | 6     | Me, Iβm Not                  | Year Zero | Nine Inch Nails |
|         | 7     | Capital G                    | Year Zero | Nine Inch Nails |
|         | 8     | My Violent Heart             | Year Zero | Nine Inch Nails |
|         | 9     | The Warning                  | Year Zero | Nine Inch Nails |
|         | 10    | God Given                    | Year Zero | Nine Inch Nails |
|         | 11    | Meet Your Master             | Year Zero | Nine Inch Nails |
|         | 12    | The Greater Good             | Year Zero | Nine Inch Nails |
|         | 13    | The Great Destroyer          | Year Zero | Nine Inch Nails |
|         | 14    | Another Version of the Truth | Year Zero | Nine Inch Nails |
|         | 15    | In This Twilight             | Year Zero | Nine Inch Nails |
|         | 16    | ZeroβSum                     | Year Zero | Nine Inch Nails |
+---------+-------+------------------------------+-----------+-----------------+

```
## Tests
We've got partial test coverage for public methods. You can run the test suite from the makefile:

```
 $ make test
ok      github.com/adamrdrew/mosh/auth  0.002s 61.762s
ok      github.com/adamrdrew/mosh/config        0.002s
ok      github.com/adamrdrew/mosh/filehandler   0.584s
?       github.com/adamrdrew/mosh/ipc   [no test files]
ok      github.com/adamrdrew/mosh/library_manager       2.688s
?       github.com/adamrdrew/mosh/moshd [no test files]
?       github.com/adamrdrew/mosh/player        [no test files]
ok      github.com/adamrdrew/mosh/plex_urls     0.002s
?       github.com/adamrdrew/mosh/printer       [no test files]
?       github.com/adamrdrew/mosh/server        [no test files]
ok      github.com/adamrdrew/mosh/shortcuts     0.004s
```

Our test coverage is squarely in the "better than nothing but not that great" territory. Patches and collab welcome!