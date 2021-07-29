package main

func PopulateData() (corpuses map[string][]string) {

	corpuses = make(map[string][]string)

	corpuses["pulsa"] = []string{
		"Saya mau beli pulsa dong. Jual voucher gak bang?. Mau isi pulsa dong.",
		"jual pulsa gak ya?",
		"kamu jual voucher ga?",
		"mau isi paket data bisa?",
		"mau isi pulsa bisa ga ya?",
	}

	corpuses["tiket"] = []string{
		"kamu jual tiket pesawat ga?",
		"disini jual tiket ga ya?",
		"bisa beli tiket    kereta?",
		"jual tiket apa  ya?",
		"ada tiket kereta nggak?",
		"tiket jurusan bandung ada",
		"beli satu tiket geratis satu",
		"disini jual tiket pesawat",
		"disini jual tiket bis",
		"disini jual tiket kereta",
		"disini jual tiket kapal laut",
	}

	corpuses["saldo"] = []string{
		"halo aku mau isi saldo dong",
		"eh mau topup dong bisa ga?",
		"mau nambah saldo dong bisa gak",
		"tolong bantu isi saldo dong 50 ribu",
	}

	corpuses["hotel"] = []string{
		"ada kamar kosong ga ya di tempat kamu?",
		"eh mau sewa 1 kamar dong",
		"eh  mau sewa ruang meeting dong",
		"eh fasilitas di hotel kamu apa aja ya?",
		"hotel di yogyakarta ada nggak ya?",
		"kamu nginap di hotel mana?",
		"fasilitasnya apa aja",
	}

	return
}
