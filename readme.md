Converts xml atlas maps to json

```
<TextureAtlas imagePath="sheet.png">
	<SubTexture name="chipBlackWhite.png" x="196" y="332" width="64" height="64"/>
</TextureAtlas>
```

into phaser3 compliant 

```
{
    "frames": [
        {
            "filename": "playingcards_0",
            "frame": {
                "w": 140,
                "h": 190,
                "x": 0,
                "y": 0
            },
            "anchor": {
                "x": 0.5,
                "y": 0.5
            }
        }
    ]
}
```

go get ./...

go run main.go inputfile.xml outputfile.json 'optional meta data description string'

