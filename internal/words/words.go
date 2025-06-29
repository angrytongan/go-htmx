package words

import "strings"

var (
	dictionary = []string{
		"cedary", "Confucianism", "bated", "folia", "knackebrod",
		"unsmelted", "parthenogeny", "reactivation", "myxopoiesis", "axman",
		"unflashing", "preresponsible", "unawakenedness", "cholecystonephrostomy", "unscenic",
		"viand", "digitalein", "honeymoonlight", "unrestrained", "enfiled",
		"femorocaudal", "musimon", "sanity", "husbandage", "soiling",
		"ahungry", "titled", "Strepsiptera", "mosker", "unforethought",
		"sideway", "cobblerfish", "peptical", "bathysmal", "phoronomics",
		"Maranta", "idiosome", "unfruitful", "Physalis", "Fabianist",
		"unartfully", "Boshas", "spaid", "syndicator", "cairned",
		"tenebrously", "stridlins", "astigmatoscope", "pantanencephalic", "foreknower",
	}
)

func Search(partial string) ([]string, error) {
	out := []string{}

	if partial == "" {
		return out, nil
	}

	for _, word := range dictionary {
		if strings.Contains(word, partial) {
			out = append(out, word)
		}
	}

	return out, nil
}
