package stringToLink

import (
	"strings"
	"sync"
	"time"
)

var (
	cacheMutex sync.Mutex
	cache      map[string]cacheEntry
)

type cacheEntry struct {
	links     map[string]string
	timestamp time.Time
}

func init() {
	cache = make(map[string]cacheEntry)
}

var (
	ozon       = "Ozon"
	avito      = "Авито"
	piter      = "Издательство Питер"
	dmk        = "Издательство ДМК"
	chitaiG    = "Читай город"
	wb         = "WildBerries"
	megaMarket = "МегаМаркет"
	marketYa   = "Яндекс Маркет"
)

func StringToLink(bookName string) map[string]string {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if result, found := cache[bookName]; found {
		if time.Since(result.timestamp).Minutes() < 60 {
			return result.links
		}
	}

	links := make(map[string]string)
	// word1+word2+word3
	plusString := strings.Join(strings.Split(bookName, " "), "+")
	// word1%20word2%20word3
	percString := strings.Join(strings.Split(bookName, " "), "%20")

	links[ozon] = "https://www.ozon.ru/category/knigi-16500/?category_was_predicted=true&from_global=true&text=" + plusString
	links[avito] = "https://www.avito.ru/all/knigi_i_zhurnaly?cd=1&q=" + plusString
	links[piter] = "https://www.piter.com/collection/all?q=" + plusString
	links[dmk] = "https://dmkpress.com/search/product/?search_string=" + plusString

	links[chitaiG] = "https://www.chitai-gorod.ru/search?phrase=" + percString
	links[wb] = "https://www.wildberries.ru/catalog/0/search.aspx?search=" + percString
	links[marketYa] = "https://market.yandex.ru/catalog--knigi/54510/list?text=" + percString
	links[megaMarket] = "https://megamarket.ru/catalog/?q=" + percString

	return links
}
