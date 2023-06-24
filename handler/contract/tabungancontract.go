package contract

type RegisterRequest struct {
	Nama string `json:"nama"`
	NIK  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type RegisterResponse struct {
	ResponseCode int        `json:"response_code"`
	ResponseMsg  string     `json:"response_msg"`
	ResponseData DaftarData `json:"data"`
}

type DaftarData struct {
	NomorRekening string `json:"nomor_rekening"`
	Remark        string `json:"remark"`
}

