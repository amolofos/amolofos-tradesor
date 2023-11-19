package unloader

import (
	"net/url"
)

type Unloader struct {
}

func (u *Unloader) Init() {

}

func (u *Unloader) Unload(doc string, outputTo string) (err error) {
	remoteUrl, err := url.ParseRequestURI(outputTo)
	if err != nil {
		u.unloadToWordpress(remoteUrl)
		return
	}

	u.unloadToLocalDir(outputTo)
	return
}

func (u *Unloader) unloadToLocalDir(outputTo string) {
	//	for i, v := range csvMap {
	//		fmt.Println(i)
	//
	//		fo, errFileOut := os.Create("data/" + fileName + "." + i + ".csv")
	//		if errFileOut != nil {
	//			log.Fatalln("Error opening file:", errFileOut)
	//		}
	//
	//		csvWrite := csv.NewWriter(fo)
	//		csvWrite.WriteAll(v)
	//		fo.Close()
	//
	//		if errCsvWrite := csvWrite.Error(); errCsvWrite != nil {
	//			log.Fatalln("Error writing csv:", errCsvWrite)
	//		}
	//	}
}

func (u *Unloader) unloadToWordpress(remoteUrl *url.URL) {
}
