package main
import("fmt"; "regexp")

func main() {
		var text = "Jeruk Silang anggur1"
		var cari = "Sila"
		var regex, err = regexp.Compile(`(?i:`+cari+`)`) //polanya alphabet kecil, \d harus ada angkanya, \D tidak ada angka ny

		if err != nil{
			fmt.Println(err.Error())
		}
		var hasil = regex.FindAllString(text, 2) //[jeruk silang] hanya mengambil dua string saja, -1 untuk mengambil semua
		fmt.Println(hasil)

		var cocok = regex.MatchString(text) //true. string yang cocok, bolean
		fmt.Println(cocok)

		var index1 = regex.FindStringIndex(text) //[0 5] mencari posisi string berdasarkann pola
		fmt.Println(index1)

		var stringbaru = regex.ReplaceAllString(text, "DURIAN") //DURIAN DURIAN DURIAN1DURIAN
		fmt.Println(stringbaru)
}
