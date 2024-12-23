package gobookprices

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	//go:embed testdata/list1.html
	mockHtmlList1 string
	//go:embed testdata/list2.html
	mockHtmlList2 string
	//go:embed testdata/list3.html
	mockHtmlList3 string
	//go:embed testdata/list4.html
	mockHtmlList4 string
	//go:embed testdata/editions1.html
	mockHtmlEditions1 string
	//go:embed testdata/editions2.html
	mockHtmlEditions2 string
	//go:embed testdata/editions3.html
	mockHtmlEditions3 string
	//go:embed testdata/work.html
	mockHtmlWork string
)

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Serve different HTML responses based on the path
	switch {
	case strings.Contains(r.URL.Path, "/review/list/"):
		// Serve different HTML responses based on the query
		switch {
		case strings.Contains(r.URL.Path, "test"):
			if strings.Contains(r.URL.RawQuery, "page=1&per_page") {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlList4))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlList3))
			}
		case strings.Contains(r.URL.RawQuery, "page=1&per_page"):
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList1))
		case strings.Contains(r.URL.RawQuery, "page=2&per_page"):
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList2))
		default:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlList3))
		}
	case strings.Contains(r.URL.Path, "/work/editions/"):
		switch {
		case strings.Contains(r.URL.RawQuery, "page=1&per_page"):
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions1))
		case strings.Contains(r.URL.RawQuery, "page=2&per_page"):
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions2))
		default:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlEditions3))
		}
	case strings.Contains(r.URL.Path, "/book/show/"):
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHtmlWork))
	}
}))

func TestGetBooksFromPage(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		want := []Book{
			{Title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", Author: "Klune, T.J.", Isbn: "", WorkUrl: "/book/show/45047384"},
			{Title: "The Fox Wife", Author: "Choo, Yangsze", Isbn: "1250266017", WorkUrl: "/book/show/127278666"},
			{Title: "The Book of Doors", Author: "Brown, Gareth", Isbn: "0063323982", WorkUrl: "/book/show/156009464"},
			{Title: "Argylle", Author: "Conway, Elly", Isbn: "0593600010", WorkUrl: "/book/show/195608705"},
			{Title: "Steelheart (The Reckoners, #1)", Author: "Sanderson, Brandon", Isbn: "0385743564", WorkUrl: "/book/show/17182126"},
			{Title: "The Lusty Argonian Maid Vol 1", Author: "Curio, Crassius", Isbn: "", WorkUrl: "/book/show/28493290"},
			{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658"},
			{Title: "Wind and Truth (The Stormlight Archive, #5)", Author: "Sanderson, Brandon", Isbn: "1250319188", WorkUrl: "/book/show/203578847"},
			{Title: "Edgedancer (The Stormlight Archive, #2.5)", Author: "Sanderson, Brandon", Isbn: "1250166543", WorkUrl: "/book/show/34703445"},
			{Title: "Red Rising (Red Rising Saga, #1)", Author: "Brown, Pierce", Isbn: "", WorkUrl: "/book/show/15839976"},
			{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196"},
			{Title: "Master and Apprentice", Author: "Gray, Claudia", Isbn: "0525619372", WorkUrl: "/book/show/40917496"},
			{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298"},
			{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137"},
			{Title: "Dawn (Legend of the Galactic Heroes, #1)", Author: "Tanaka, Yoshiki", Isbn: "1421584948", WorkUrl: "/book/show/25986983"},
			{Title: "Think and Grow Rich", Author: "Hill, Napoleon", Isbn: "", WorkUrl: "/book/show/30186948"},
			{Title: "Children of Time (Children of Time, #1)", Author: "Tchaikovsky, Adrian", Isbn: "1447273281", WorkUrl: "/book/show/25499718"},
			{Title: "The Red Knight (The Traitor Son Cycle, #1)", Author: "Cameron, Miles", Isbn: "0575113294", WorkUrl: "/book/show/13616278"},
			{Title: "The Girl and the Stars (Book of the Ice, #1)", Author: "Lawrence, Mark", Isbn: "1984805991", WorkUrl: "/book/show/51277288"},
			{Title: "The Pariah (Covenant of Steel, #1)", Author: "Ryan, Anthony", Isbn: "0316430773", WorkUrl: "/book/show/56229688"},
		}
		url := fmt.Sprint(server.URL, "/review/list/68156753", "?page=1&per_page=20&shelf=to-read")
		got, err := getBooksFromPage(url)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			// for _, b := range got {
			// 	fmt.Printf("{Title: \"%s\", Author: \"%s\", Isbn: \"%s\", WorkUrl: \"%s\"},\n", b.Title, b.Author, b.Isbn, b.WorkUrl)
			// }
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
	t.Run("test with zero results", func(t *testing.T) {
		want := []Book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
	t.Run("test with http error", func(t *testing.T) {
		// Create a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(""))
		}))
		defer server.Close()

		want := []Book{}
		got, err := getBooksFromPage(server.URL)
		switch {
		case err != nil:
			t.Errorf("error getting books:\nWant:\n%+v\nGot:\n%+v\n", want, got)
		case !reflect.DeepEqual(want, got):
			t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		}
	})
}

func TestGetBooks(t *testing.T) {
	want := []Book{
		{Title: "The House in the Cerulean Sea (Cerulean Chronicles, #1)", Author: "Klune, T.J.", Isbn: "", WorkUrl: "/book/show/45047384"},
		{Title: "The Fox Wife", Author: "Choo, Yangsze", Isbn: "1250266017", WorkUrl: "/book/show/127278666"},
		{Title: "The Book of Doors", Author: "Brown, Gareth", Isbn: "0063323982", WorkUrl: "/book/show/156009464"},
		{Title: "Argylle", Author: "Conway, Elly", Isbn: "0593600010", WorkUrl: "/book/show/195608705"},
		{Title: "Steelheart (The Reckoners, #1)", Author: "Sanderson, Brandon", Isbn: "0385743564", WorkUrl: "/book/show/17182126"},
		{Title: "The Lusty Argonian Maid Vol 1", Author: "Curio, Crassius", Isbn: "", WorkUrl: "/book/show/28493290"},
		{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658"},
		{Title: "Wind and Truth (The Stormlight Archive, #5)", Author: "Sanderson, Brandon", Isbn: "1250319188", WorkUrl: "/book/show/203578847"},
		{Title: "Edgedancer (The Stormlight Archive, #2.5)", Author: "Sanderson, Brandon", Isbn: "1250166543", WorkUrl: "/book/show/34703445"},
		{Title: "Red Rising (Red Rising Saga, #1)", Author: "Brown, Pierce", Isbn: "", WorkUrl: "/book/show/15839976"},
		{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196"},
		{Title: "Master and Apprentice", Author: "Gray, Claudia", Isbn: "0525619372", WorkUrl: "/book/show/40917496"},
		{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298"},
		{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137"},
		{Title: "Dawn (Legend of the Galactic Heroes, #1)", Author: "Tanaka, Yoshiki", Isbn: "1421584948", WorkUrl: "/book/show/25986983"},
		{Title: "Think and Grow Rich", Author: "Hill, Napoleon", Isbn: "", WorkUrl: "/book/show/30186948"},
		{Title: "Children of Time (Children of Time, #1)", Author: "Tchaikovsky, Adrian", Isbn: "1447273281", WorkUrl: "/book/show/25499718"},
		{Title: "The Red Knight (The Traitor Son Cycle, #1)", Author: "Cameron, Miles", Isbn: "0575113294", WorkUrl: "/book/show/13616278"},
		{Title: "The Girl and the Stars (Book of the Ice, #1)", Author: "Lawrence, Mark", Isbn: "1984805991", WorkUrl: "/book/show/51277288"},
		{Title: "The Pariah (Covenant of Steel, #1)", Author: "Ryan, Anthony", Isbn: "0316430773", WorkUrl: "/book/show/56229688"},
		{Title: "Ready Player Two (Ready Player One, #2)", Author: "Cline, Ernest", Isbn: "1524761338", WorkUrl: "/book/show/26082916"},
		{Title: "Red Sister (Book of the Ancestor, #1)", Author: "Lawrence, Mark", Isbn: "1101988851", WorkUrl: "/book/show/25895524"},
		{Title: "De hondenmoorden (Oxen, #1)", Author: "Jensen, Jens Henrik", Isbn: "9400505787", WorkUrl: "/book/show/25351052"},
		{Title: "Postkantoor", Author: "Bukowski, Charles", Isbn: "9023417429", WorkUrl: "/book/show/11487773"},
		{Title: "The Book That Wouldn’t Burn (The Library Trilogy, #1)", Author: "Lawrence, Mark", Isbn: "0593437918", WorkUrl: "/book/show/61612864"},
		{Title: "Bloedmaan", Author: "Nesbø, Jo", Isbn: "9403126426", WorkUrl: "/book/show/123192410"},
		{Title: "Grand Hotel Europa", Author: "Pfeijffer, Ilja Leonard", Isbn: "902952622X", WorkUrl: "/book/show/41456850"},
		{Title: "The Atlas Six (The Atlas, #1)", Author: "Blake, Olivie", Isbn: "1529095255", WorkUrl: "/book/show/61259384"},
		{Title: "Mistborn: The Wax and Wayne Series: The Alloy of Law, Shadows of Self, The Bands of Mourning", Author: "Sanderson, Brandon", Isbn: "1250293499", WorkUrl: "/book/show/37436740"},
		{Title: "The Hand of the Sun King (Pact and Pattern, #1)", Author: "Greathouse, J.T.", Isbn: "1473232902", WorkUrl: "/book/show/57596188"},
		{Title: "The Hypnotist (Joona Linna, #1)", Author: "Kepler, Lars", Isbn: "0374173958", WorkUrl: "/book/show/9835731"},
		{Title: "A Master of Djinn (Dead Djinn Universe, #1)", Author: "Clark, P. Djèlí", Isbn: "1250267676", WorkUrl: "/book/show/52504334"},
		{Title: "The Stardust Thief (The Sandsea Trilogy, #1)", Author: "Abdullah, Chelsea", Isbn: "0316368768", WorkUrl: "/book/show/58950705"},
		{Title: "Black Sun (Between Earth and Sky, #1)", Author: "Roanhorse, Rebecca", Isbn: "1534437673", WorkUrl: "/book/show/50892360"},
		{Title: "Futuristic Violence and Fancy Suits (Zoey Ashe #1)", Author: "Wong, David", Isbn: "1783291842", WorkUrl: "/book/show/25859666"},
		{Title: "The Alloy of Law (Mistborn, #4)", Author: "Sanderson, Brandon", Isbn: "0765330423", WorkUrl: "/book/show/10803121"},
		{Title: "A Darker Shade of Magic (Shades of Magic, #1)", Author: "Schwab, Victoria", Isbn: "0765376458", WorkUrl: "/book/show/22055262"},
		{Title: "The Lies of Locke Lamora (Gentleman Bastard, #1)", Author: "Lynch, Scott", Isbn: "", WorkUrl: "/book/show/29588376"},
		{Title: "Jade City (The Green Bone Saga, #1)", Author: "Lee, Fonda", Isbn: "0316440884", WorkUrl: "/book/show/43587154"},
		{Title: "Ship of Magic (The Liveship Traders, #1)", Author: "Hobb, Robin", Isbn: "B0DLT8H3NY", WorkUrl: "/book/show/6662349"},
	}
	got, err := GetBooks(server.URL, "/review/list/68156753", "to-read")
	switch {
	case err != nil:
		t.Errorf("error getting books: \nWant:\n'%+v', Got:\n'%+v'", want, got)
	case !reflect.DeepEqual(want, got):
		t.Errorf("Want:\n'%+v'\nGot:\n'%+v'", want, got)
	}
}

func TestGetEditionsFromPage(t *testing.T) {
	want := []Edition{
		{Isbn: "9788408252856", Format: "Hardcover", Language: "Spanish"},
		{Isbn: "", Format: "Kindle Edition", Language: "Spanish"},
		{Isbn: "9789044934120", Format: "ebook", Language: "Dutch"},
		{Isbn: "9788408269649", Format: "Mass Market Paperback", Language: "Spanish"},
		{Isbn: "9786190110934", Format: "Paperback", Language: "Bulgarian"},
		{Isbn: "9786180149012", Format: "Paperback", Language: "Greek, Modern (1453-)"},
		{Isbn: "9789400515277", Format: "Paperback", Language: "Dutch"},
		{Isbn: "", Format: "", Language: "Spanish"},
		{Isbn: "9788466429306", Format: "Kindle Edition", Language: "Catalan; Valencian"},
		{Isbn: "9788408255178", Format: "Kindle Edition", Language: "Spanish"},
	}
	url := fmt.Sprint(server.URL, "/work/editions/94024291", "?page=1&per_page=10")
	got, err := getEditionsFromPage(url)
	switch {
	case err != nil:
		t.Errorf("error getting editions:\nWant:\n%+v\nGot:\n%+v\n", want, got)
	case !reflect.DeepEqual(want, got):
		// for _, e := range got {
		// 	fmt.Printf("{Isbn: \"%s\", Format: \"%s\", Language: \"%s\"},\n", e.Isbn, e.Format, e.Language)
		// }
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}
}

func TestGetEditions(t *testing.T) {
	want := []Edition{
		{Isbn: "9789044934120", Format: "ebook", Language: "Dutch"},
		{Isbn: "9788408255178", Format: "Kindle Edition", Language: "Spanish"},
	}
	got, err := GetEditions(server.URL, "/work/editions/94024291", EbookFormats, []string{"Dutch", "Spanish"})
	switch {
	case err != nil:
		t.Errorf("error getting editions:\nWant:\n%+v\nGot\n%+v\n", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}
}

func TestGetEditionUrl(t *testing.T) {
	want := "/work/editions/62945242"
	got, err := getEditionUrl(server.URL, "/book/show/45047384")
	switch {
	case err != nil:
		t.Errorf("error getting edition url:\nWant:\t'%s'\nGot:\t'%s'", want, got)
	case want != got:
		t.Fatalf("\nWant:\t'%s'\nGot:\t'%+v'\n", want, got)
	}
}

func TestNewLibrary(t *testing.T) {
	want := Library{
		hostUrl:   server.URL,
		ListUrl:   "/review/list/test",
		Shelf:     "to-read",
		Formats:   EbookFormats,
		Languages: []string{"Dutch", "Spanish"},
	}
	got := NewLibrary(want.ListUrl, want.Shelf, want.Formats, want.Languages)
	got.hostUrl = server.URL

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}

	want.Books = []Book{
		{
			Title:       "The Poppy War (The Poppy War, #1)",
			Author:      "Kuang, R.F.",
			Isbn:        "0062662597",
			WorkUrl:     "/book/show/35068705",
			EditionsUrl: "/work/editions/62945242",
			Editions: []Edition{
				{Isbn: "9788408258216", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9786070791734", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9789021462691", Format: "ebook", Language: "Dutch"},
			},
		},
		{
			Title:       "The House in the Cerulean Sea (Cerulean Chronicles, #1)",
			Author:      "Klune, T.J.",
			Isbn:        "B07QPHT8CB",
			WorkUrl:     "/book/show/45047384",
			EditionsUrl: "/work/editions/62945242",
			Editions: []Edition{
				{Isbn: "9788408258216", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9786070791734", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9789021462691", Format: "ebook", Language: "Dutch"},
			},
		},
	}
	err := got.Update()

	switch {
	case err != nil:
		t.Errorf("error updating library:\nWant:%+v\nGot:\n%+v", want, got)
	case !reflect.DeepEqual(want, got):
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}
}

func TestCompareAndUpdate(t *testing.T) {
	booksOld := []Book{
			{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658", EditionsUrl: "/work/editions/62945242",
				Editions: []Edition{
					{Isbn: "9788408258216", Format: "Kindle Edition", Language: "Spanish"},
					{Isbn: "9786070791734", Format: "Kindle Edition", Language: "Spanish"},
					{Isbn: "9789021462691", Format: "ebook", Language: "Dutch"},
				}},
			{Title: "Fool's Assassin (The Fitz and the Fool, #1)", Author: "Hobb, Robin", Isbn: "0553392433", WorkUrl: "/book/show/41021196"},
			{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298", EditionsUrl: "/work/editions/62945242"},
			{Title: "Enterprise Architecture As Strategy: Creating a Foundation for Business Execution", Author: "Ross, Jeanne W.", Isbn: "1591398398", WorkUrl: "/book/show/70137"},
		}
	booksNew :=  []Book{
			{Title: "Ready Player Two (Ready Player One, #2)", Author: "Cline, Ernest", Isbn: "1524761338", WorkUrl: "/book/show/26082916"},
			{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658"},
			{Title: "Red Sister (Book of the Ancestor, #1)", Author: "Lawrence, Mark", Isbn: "1101988851", WorkUrl: "/book/show/25895524"},
			{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298"},
		}
	want := []Book{
		{Title: "Ready Player Two (Ready Player One, #2)", Author: "Cline, Ernest", Isbn: "1524761338", WorkUrl: "/book/show/26082916"},
		{Title: "Het parfum", Author: "Süskind, Patrick", Isbn: "9057134101", WorkUrl: "/book/show/3187658", EditionsUrl: "/work/editions/62945242",
			Editions: []Edition{
				{Isbn: "9788408258216", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9786070791734", Format: "Kindle Edition", Language: "Spanish"},
				{Isbn: "9789021462691", Format: "ebook", Language: "Dutch"},
			}},
		{Title: "Red Sister (Book of the Ancestor, #1)", Author: "Lawrence, Mark", Isbn: "1101988851", WorkUrl: "/book/show/25895524"},
		{Title: "Dark Disciple", Author: "Golden, Christie", Isbn: "B01EKIFQ7Y", WorkUrl: "/book/show/23277298", EditionsUrl: "/work/editions/62945242"},
	}
	got := compareAndUpdate(booksOld, booksNew)

	if !reflect.DeepEqual(want, got) {
		// t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
		t.Fatalf("\nWant:\n%+v\nGot:\n%+v\n", want, got)
	}
}
