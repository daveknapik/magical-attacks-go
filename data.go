package main

import (
	"fmt"
	"os"
)

func getDataByFandom(f string) []string {
	var r []string
	data := getData()

	if f == "all" {
		// loop over values, concatenate into one big slice of strings
		values := make([]string, 0)

		for _, v := range data {
			values = append(values, v...)
		}

		r = unique(values)
	} else if _, ok := data[f]; ok {
		r = unique(data[f])
	} else {
		fmt.Printf("Error: Fandom '%s' does not exist\n", f)
		os.Exit(1)
	}

	return r
}

func unique(d []string) []string {
	trackDuplicates := make(map[string]bool)
	deduped := []string{}

	for _, item := range d {
		if _, found := trackDuplicates[item]; !found {
			trackDuplicates[item] = true
			deduped = append(deduped, item)
		}
	}

	return deduped
}

func getData() map[string][]string {
	precure := []string{
		"A La Mode!",
		"Aqua",
		"Arpeggio",
		"Arrow",
		"Ascension",
		"Attack",
		"Barrier",
		"Beauty",
		"Blast",
		"Blaze",
		"Blizzard",
		"Blue",
		"Burst",
		"Chain",
		"Cheerful",
		"Circle",
		"Coral",
		"Cosmo",
		"Cross",
		"Crystal",
		"Dark",
		"Decoration",
		"Diamond",
		"Diffusion",
		"Double",
		"Dream",
		"Dynamite",
		"Echo",
		"Emerald",
		"Eternal",
		"Everyone",
		"Explosion",
		"Extreme",
		"Fantastic",
		"Feather",
		"Final",
		"Fire",
		"Fire",
		"Flamingo",
		"Flash",
		"Floral",
		"Force",
		"Forever",
		"Fresh",
		"Grand",
		"Happiness",
		"Happy",
		"Harmony",
		"Healing",
		"Heart",
		"Holy",
		"Honey",
		"Hurricane",
		"Illusion",
		"Imagination",
		"Innocent",
		"Kiss",
		"Love",
		"Lovely",
		"Lucky",
		"Lumiere",
		"Marble",
		"Mint",
		"Miracle",
		"Mirror",
		"Motion",
		"Oasis",
		"Passion",
		"Passionale",
		"Peace",
		"Phoenix",
		"Pink",
		"Power",
		"Powerful",
		"Prayer",
		"Pretty",
		"Prism",
		"Punch",
		"Purification",
		"Rainbow",
		"Rainbow",
		"Reflection",
		"Reincarnation",
		"Ribbon",
		"Ripple",
		"Rising",
		"Rose",
		"Rouge",
		"Sapphire",
		"Saucer",
		"Shake",
		"Shining",
		"Shock",
		"Shoot",
		"Shooting",
		"Shot",
		"Shower",
		"Silver",
		"Slash",
		"Smash",
		"Solution",
		"Sonic",
		"Spark",
		"Sparkle",
		"Spiral",
		"Splash",
		"Star",
		"Star",
		"Stardust",
		"Starlight",
		"Storm",
		"Stream",
		"Strike",
		"Summer",
		"Sunny",
		"Sunshine",
		"Sword",
		"Temptation",
		"Tender",
		"Therapy",
		"Thunder",
		"Tomorrow",
		"Tornado",
		"Trinity",
		"True",
		"Twinkle",
		"Wall",
		"Wave",
		"Wonderful",
	}

	sailorMoon := []string{
		"Action",
		"Aerial",
		"Aqua",
		"Attack",
		"Barrage",
		"Beam",
		"Beauty",
		"Bird",
		"Break",
		"Break",
		"Burning",
		"Cannon",
		"Chain",
		"Chaos",
		"Charm",
		"Chocolate",
		"Chronos",
		"Cosmos",
		"Crash",
		"Crescent",
		"Crisis",
		"Crunch",
		"Crystal",
		"Cutter",
		"Cyclone",
		"Dark",
		"Dead",
		"Death",
		"Deep",
		"Delicious",
		"Destiny",
		"Destructive",
		"Dimension",
		"Dragon",
		"Drive",
		"Drop",
		"Escalation",
		"Evil",
		"Evolution",
		"Explosive",
		"Extinguish",
		"Fascination",
		"Fate",
		"Feather",
		"Fire",
		"Flame",
		"Flames",
		"Flare",
		"Flash",
		"Flower",
		"Freeze",
		"Freezing",
		"Fusion",
		"Galactica",
		"Gale",
		"Glaive",
		"Gorgeous",
		"Halation",
		"Hallucination",
		"Hand",
		"Healing",
		"Heart",
		"Heartache",
		"Hurricane",
		"Hyperspatial",
		"Illusion",
		"Inflation",
		"Kiss",
		"Launcher",
		"Lightning",
		"Love",
		"Magnum",
		"Mandala",
		"Meditation",
		"Mirage",
		"Mirror",
		"Misfortune",
		"Mist",
		"Moonlight",
		"Mosaic",
		"Pentagram",
		"Pink",
		"Planet",
		"Plasma",
		"Poisoning",
		"Power",
		"Power",
		"Princess",
		"Prism",
		"Punch",
		"Rainbow",
		"Rapid",
		"Reborn",
		"Reverse",
		"Revolution",
		"Rhapsody",
		"Ribbon",
		"Rise",
		"Rolling",
		"Scream",
		"Sensation",
		"Shaking",
		"Shock",
		"Slash",
		"Snake",
		"Sniper",
		"Soul",
		"Space",
		"Sparkling",
		"Spiral",
		"Spray",
		"Stardust",
		"Storm",
		"Strike",
		"Submerge",
		"Sugar",
		"Super",
		"Surprise",
		"Tempest",
		"Throw",
		"Thunder",
		"Thunderbolt",
		"Tide",
		"Tornado",
		"Tsunami",
		"Turbulence",
		"Twilight",
		"Typhoon",
		"Vicious",
		"Water",
		"World",
		"Zone",
	}

	return map[string][]string{
		"precure":    precure,
		"sailorMoon": sailorMoon,
	}
}