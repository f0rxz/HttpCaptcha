package main

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"time"

	"net/http"
	"strconv"
)


var sol int64

func captRnd5xUintFromSeed(seed int64) uint {
	rand.Seed(sol ^ seed)
	num := rand.Intn(9000) + 1000
	rand.Seed(time.Now().UTC().UnixNano())
	println(uint(num))
	return uint(num)
}

func get_sol_by_ip(ip string) int64 {
	leng := len(ip)
	loc_sol := int64(0)
	for i := 0; i < leng; i++ {
		loc_sol = (loc_sol + int64(ip[i]) << (8 * uint(i) % 64)) % 9223372036854775807;
	}
	return loc_sol
}

func CheckCaptcha(seed int64, ip string, num uint) bool {
	loc_sol := get_sol_by_ip(ip)
	minutes := time.Now().UnixNano() / 60000000000
	seed ^= loc_sol
	return captRnd5xUintFromSeed(seed ^ minutes)       == num ||
		   captRnd5xUintFromSeed(seed ^ (minutes - 1)) == num
}

var captcha_offset = 0

func GetCaptcha(seed int64, ip string) []byte {
	seed ^= get_sol_by_ip(ip) ^ time.Now().UnixNano() / 60000000000
	img := image.NewGray(image.Rect(0, 0, 23, 7))
	x := 22
	y := 1
	for X := 0; X < 23; X++ {
		for Y := 0; Y < 7; Y++ {
			clr := X * 9 + Y * 9 - rand.Intn(16) + captcha_offset
			if clr < 0 { clr += 16 }
			img.Set(X, Y, color.Gray{byte(clr)})
		}
	}
	captDrawUInt(img, x, y, captRnd5xUintFromSeed(seed))
	captcha_offset = (captcha_offset + 17) % 256
	var data bytes.Buffer
	writer := bufio.NewWriter(&data)
	//writer, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	//defer writer.Close()
	png.Encode(writer, img)
	writer.Flush()
	return data.Bytes()
}

func r50on50() bool {
	if rand.Intn(2) == 1 {
		return true
	}
	return false
}

/*func captUIntGetRndClr() color.Gray {
	c := byte(rand.Intn(193))
	return color.Gray{255 - c}
}*/

func captUIntSetPX(img * image.Gray, ox, oy, x, y int){
	img.Set(x, y, color.Gray{255 - byte((x - ox) * 12 - (y - oy) * 3 - captcha_offset)})
}

func captDrawUInt(img * image.Gray, x, y int, num uint){
	lx := 0; ly := 0
	x -= 3
	for true {
		if num == 0 { break }
		lx = x
		ly = y
		ox := rand.Intn(3) - 1
		oy := rand.Intn(3) - 1
		x += ox
		y += oy
		switch(num % 10){
			case 0:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				captUIntSetPX(img, ox, oy, x, y + 2)
				captUIntSetPX(img, ox, oy, x, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 1, y)
			case 1:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 1)
				captUIntSetPX(img, ox, oy, x + 2, y)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
			case 2:
				captUIntSetPX(img, ox, oy, x, y)
				captUIntSetPX(img, ox, oy, x + 1, y)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 2)
				}
				captUIntSetPX(img, ox, oy, x, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				captUIntSetPX(img, ox, oy, x + 2, y + 4)
			case 3:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x + 1, y)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
			case 4:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				captUIntSetPX(img, ox, oy, x, y + 2)
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
			case 5:
				captUIntSetPX(img, ox, oy, x + 2, y)
				captUIntSetPX(img, ox, oy, x + 1, y)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				captUIntSetPX(img, ox, oy, x, y + 4)
			case 6:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 1, y)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				captUIntSetPX(img, ox, oy, x, y + 2)
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
				captUIntSetPX(img, ox, oy, x, y + 3)
			case 7:
				captUIntSetPX(img, ox, oy, x, y)
				captUIntSetPX(img, ox, oy, x + 1, y)
				captUIntSetPX(img, ox, oy, x + 2, y)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 2)
					captUIntSetPX(img, ox, oy, x + 2, y + 3)
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				} else {
					captUIntSetPX(img, ox, oy, x + 1, y + 2)
					if r50on50(){
						captUIntSetPX(img, ox, oy, x, y + 3)
					} else {
						captUIntSetPX(img, ox, oy, x + 1, y + 3)
					}
					captUIntSetPX(img, ox, oy, x + 0, y + 4)
				}
			case 8:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				captUIntSetPX(img, ox, oy, x, y + 2)
				captUIntSetPX(img, ox, oy, x, y + 3)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 1, y)
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
			case 9:
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 4)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y + 4)
				}
				captUIntSetPX(img, ox, oy, x + 2, y + 3)
				captUIntSetPX(img, ox, oy, x + 2, y + 2)
				captUIntSetPX(img, ox, oy, x + 2, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x + 2, y)
				}
				captUIntSetPX(img, ox, oy, x + 1, y)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y)
				}
				captUIntSetPX(img, ox, oy, x, y + 1)
				if r50on50(){
					captUIntSetPX(img, ox, oy, x, y + 2)
				}
				captUIntSetPX(img, ox, oy, x + 1, y + 2)
		}
		x = lx
		y = ly
		x -= 6
		num /= 10
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	sol = rand.Int63n(9223372036854775807)
	http.HandleFunc("/", handler)
	panic(http.ListenAndServe("127.0.0.1:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(`<a href="/"><img src="/capt?s=` + strconv.FormatInt(rand.Int63n(9223372036854775807), 10) + `" style="border:1px solid black;height:48px;image-rendering:-moz-crisp-edges;image-rendering:-o-crisp-edges;image-rendering:-webkit-optimize-contrast;image-rendering:crisp-edges;-ms-interpolation-mode:nearest-neighbor;"></img></a>`))
		case "/capt":
			seed, err := strconv.ParseInt(r.FormValue("s"), 10, 64)
			if err == nil {
				w.Write(GetCaptcha(seed, r.Header.Get("Cf-Connecting-Ip")))
			}
		case "/capt_chk":
			seed, err := strconv.ParseInt(r.FormValue("s"), 10, 64)
			if err == nil {
				num, err := strconv.Atoi(r.FormValue("n"))
				if err == nil {
					if CheckCaptcha(seed, r.Header.Get("Cf-Connecting-Ip"), uint(num)) {
						w.Write([]byte{49})
					} else {
						w.Write([]byte{48})
					}
				}
			}
	}
}