package stringToLink

import (
	"strings"
)

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
	res := make(map[string]string)
	// word1+word2+word3
	plusString := strings.Join(strings.Split(bookName, " "), "+")
	// word1%20word2%20word3
	percString := strings.Join(strings.Split(bookName, " "), "%20")

	res[ozon] = "https://www.ozon.ru/category/knigi-16500/?category_was_predicted=true&from_global=true&text=" + plusString
	res[avito] = "https://www.avito.ru/all/knigi_i_zhurnaly?cd=1&q=" + plusString
	res[piter] = "https://www.piter.com/collection/all?q=" + plusString
	res[dmk] = "https://dmkpress.com/search/product/?search_string=" + plusString

	res[chitaiG] = "https://www.chitai-gorod.ru/search?phrase=" + percString
	res[wb] = "https://www.wildberries.ru/catalog/0/search.aspx?search=" + percString
	res[marketYa] = "https://market.yandex.ru/catalog--knigi/54510/list?text=" + percString
	res[megaMarket] = "https://megamarket.ru/catalog/?q=" + percString

	return res
}
