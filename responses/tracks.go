package responses

import "encoding/xml"

type ResponseTrackMediaPart struct {
	XMLName xml.Name `xml:"Part"`
	//This is what we actually play
	///library/parts/77708/1574343480/file.mp3
	Key string `xml:"key,attr"`
}

type ResponseTrackMedia struct {
	XMLName xml.Name               `xml:"Media"`
	Part    ResponseTrackMediaPart `xml:"Part"`
}

type ResponseTrack struct {
	XMLName xml.Name `xml:"Track"`
	//Will always be album
	Type string `xml:"type,attr"`
	//This is the endpoint we use to play a song
	//key="/library/parts/77708/1574343480/file.mp3"
	Key string `xml:"key,attr"`
	//Album title
	Title string `xml:"title,attr"`
	//This is like the artist's ID. We can use it to reverse engineer
	//the key. This makes running commands easier
	RatingKey string `xml:"ratingKey,attr"`
	//Album Name
	ParentTitle string `xml:"parentTitle,attr"`
	//Album Name
	GrandParentTitle string             `xml:"grandparentTitle,attr"`
	Media            ResponseTrackMedia `xml:"Media"`
	//Track Number
	Index string `xml:"index,attr"`
	Image string `xml:"thumb,attr"`
}

func (r *ResponseTrack) GetPath() string {
	return r.Media.Part.Key
}

type ResponseTracksMediaContainer struct {
	XMLName xml.Name        `xml:"MediaContainer"`
	Tracks  []ResponseTrack `xml:"Track"`
}

/*
<MediaContainer size="13" allowSync="1" art="/library/metadata/2209/art/1653130965" grandparentRatingKey="2209" grandparentThumb="/library/metadata/2209/thumb/1653130965" grandparentTitle="The Dillinger Escape Plan" identifier="com.plexapp.plugins.library" key="74932" librarySectionID="1" librarySectionTitle="Adam Music" librarySectionUUID="c8aafd0e-bd89-41d3-a318-d316478227e2" mediaTagPrefix="/system/bundle/media/flags/" mediaTagVersion="1652176532" nocache="1" parentIndex="1" parentTitle="Ire Works" parentYear="2007" summary="After Miss Machine, Dillinger Escape Plan fans were divided. Many of the folks who were attached to the screaming mathematical metal of Calculating Infinity bailed on the band, disapproving of the experimental musical direction and the meathead appearance of new singer/screamer Greg Puciato. Open-minded listeners were excited about the progressive journey they were taking and many critics hailed the group as a true innovator of metalcore. Ire Works succeeds in many of the same ways that their previous album did, while branching out creatively. They continue to toy with technical metal, blistering hardcore, jazz breaks, and post-punk, but here they evolve again by adding more twists and turns with additional electronic elements. While the merging of too many styles in hardcore can make for a convoluted result (see Avenged Sevenfold&#39;s self-titled release), the added instruments and genre changeups enhance the result rather than acting as ornamental distractions. Edgy Aphex Twin-style drill&#39;n&#39;bass drum breaks and stretched and squeezed electro blips feel strangely at home next to the psychotic time-signature changes and manic riffs, especially on the tracks &#34;Sick on Sunday,&#34; &#34;Dead as History,&#34; and &#34;When Acting as a Wave.&#34; Violins, pianos, and trumpets sit nicely in the mix, and the group&#39;s willingness to take chances leads to stunning artistic endeavors rather than stale attempts at crossing genres just for the sake of being clever. Original vocalist Dimitri Minakakis makes an appearance, as does Mastodon guitarist Brent Hinds, but the most notable inclusion is drummer Gil Sharone, who proves himself an expert at picking up the slack after the departure of founding member Chris Pennie to play in Coheed and Cambria. Undoubtedly, this act added anger to fuel the fire of their heavier numbers. &#34;82588,&#34; &#34;Fix Your Face,&#34; and &#34;Party Smasher&#34; are as wicked and manic as their most difficult earlier stuff; conversely, the melodic hooks and falsetto of &#34;Black Bubblegum&#34; and the watery ambience of &#34;Mouth of Ghosts&#34; balance out the album nicely. It can be inaccessible and terrifying all at once, but in a genre overly saturated with formulaic groups, Ire Works is a true standout. If DEP aren&#39;t careful and continue down this innovative path, they could easily be labeled the Radiohead of metalcore. A visceral metal album that pushes the envelope? Who would have thunk it? ~ Jason Lymangrover" thumb="/library/metadata/74932/thumb/1653303778" title1="The Dillinger Escape Plan" title2="Ire Works" viewGroup="track" viewMode="65593">
    <Track ratingKey="74933" key="/library/metadata/74933" parentRatingKey="74932" grandparentRatingKey="2209" guid="plex://track/5d07d4bb403c64029083003c" parentGuid="plex://album/5d07c2b9403c640290902899" grandparentGuid="plex://artist/5d07bbfe403c6402904a72ab" parentStudio="Relapse Records" type="track" title="Fix Your Face" grandparentKey="/library/metadata/2209" parentKey="/library/metadata/74932" grandparentTitle="The Dillinger Escape Plan" parentTitle="Ire Works" summary="" index="1" parentIndex="1" ratingCount="93842" viewCount="4" lastViewedAt="1653419868" parentYear="2007" thumb="/library/metadata/74932/thumb/1653303778" art="/library/metadata/2209/art/1653130965" parentThumb="/library/metadata/74932/thumb/1653303778" grandparentThumb="/library/metadata/2209/thumb/1653130965" grandparentArt="/library/metadata/2209/art/1653130965" duration="161750" addedAt="1630106070" updatedAt="1653217553" musicAnalysisVersion="1">
        <Media id="77649" duration="161750" bitrate="269" audioChannels="2" audioCodec="mp3" container="mp3">
            <Part id="77708" key="/library/parts/77708/1574343480/file.mp3" duration="161750" file="D:\Adam Drew OneDrive\OneDrive\Music\The Dillinger Escape Plan\Ire Works\01 - Fix Your Face.mp3" size="5498847" container="mp3" hasThumbnail="1" />
        </Media>
    </Track>
*/
