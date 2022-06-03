# Dev Resources
This doc has links and info that are useful for Mosh developers.

## Plex Queries
This documentaiton is inside baseball. To get the searching and stuff working I had to reverse engineer a bunch of Plex queries. You don't need to know anything of this to use MOSH - I'm just putting these here for me or just in case some other weirdo needs them:

```
List all artists
http://<IP>:<PORT>/library/sections/<LIBRARY KEY>/all?X-Plex-Token=<TOKEN>

List all albums
http://<IP>:<PORT>/library/sections<LIBRARY KEY>albums?X-Plex-Token=<TOKEN>


The key attribute contains the link to children

<Directory allowSync="1" librarySectionID="1" 
librarySectionTitle="Adam Music" 
librarySectionUUID="c8aafd0e-bd89-41d3-a318-d316478227e2" 
ratingKey="73689" key="/library/metadata/73689/children" ... />

To get info about an item use the key

Artist info and list of their albums:
http://<IP>:<PORT>/library/metadata/2095/children/?X-Plex-Token=<TOKEN>

Album info and list of its songs:
http://<IP>:<PORT>/library/metadata/75242/children/?X-Plex-Token=<TOKEN>

In both examples we've just got the key from the parent object.

Search for an artist:
http://<IP>:<PORT>/library/sections<LIBRARY KEY>all?title=nine inch&X-Plex-Token=<TOKEN>

Search for an album:
http://<IP>:<PORT>/library/sections<LIBRARY KEY>albums?title=pony&X-Plex-Token=<TOKEN>
```

## Resources Consulted:
* [An Example of Go RPC Client and Server](https://ops.tips/gists/example-go-rpc-client-and-server/)
* [Cobra](https://cobra.dev/)
* [Cobra User Guide](https://github.com/spf13/cobra/blob/master/user_guide.md)
* [Parsing XML in Go](https://tutorialedge.net/golang/parsing-xml-with-golang/)
* [go-daemon](https://github.com/sevlyar/go-daemon)
* [pulse-simple](https://github.com/mesilliac/pulse-simple/)
* [goav](https://github.com/giorgisio/goav)
* [Authenticating with Plex](https://forums.plex.tv/t/authenticating-with-plex/609370)
* [RESTer Firefox Extension](https://addons.mozilla.org/en-US/firefox/addon/rester/)
* [Plex API Wiki](https://github.com/Arcanemagus/plex-api/wiki/Plex-Web-API-Overview)
* [Plex Media Server URL Commands](https://support.plex.tv/articles/201638786-plex-media-server-url-commands/)
* [Plex TV URLs](https://github.com/Arcanemagus/plex-api/wiki/Plex.tv#urls)
* [Beep](https://github.com/faiface/beep)
* [Beep Example](https://github.com/faiface/beep/wiki/Hello,-Beep!)
* [Go Pretty Progress](https://github.com/jedib0t/go-pretty/tree/main/progress)
* [Go Pretty Table](https://github.com/jedib0t/go-pretty/tree/main/table)
* [termpic](https://github.com/ajmalsiddiqui/termpic)