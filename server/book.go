package server

import "github.com/samber/lo"

func GetRandomWisdom() string {
	wisdoms := []string{
		"Let go once in a while. You are a loose lily floating down an amber river.",
		"When in a tight spot try to imagine yourself marooned on a beautiful desert island.",
		"Add a drop of lavender to your bath and soon soak yourself calm.",
		"If you want to feel calm, eat more raw fruit and vegetables, yoghurt, milk and seeds.",
		"When you rest you are a king surveying your estate. Look at the woodland, peacocks, be king of your own calm kingdom.",
		"When you're feeling under pressure do something different. Roll up your sleeves or eat an orange.",
		"Look for things to make you laugh, If you see nothing worth laughing at pretend you do, then laugh.",
		"Have you ever seen a calm person with a loud voice? Try and speak softly once in a while.",
		"Add some lavender to milk. Leave town with an orange. Pretend you're laughing at it.",
	}

	return lo.Sample(wisdoms)
}
