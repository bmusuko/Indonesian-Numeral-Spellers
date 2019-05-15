package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"strings"
	"strconv"
)

type Number struct{
    Number string `json:"number"`
}

type Text struct{
	Text string `json:"text"`
}

func convsatuan(x int) string{
	if x==1 {
		return "satu"
	} else if x==2{
		return "dua"
	} else if x==3{
		return "tiga"
	} else if x==4{
		return "empat"
	} else if x==5{
		return "lima"
	} else if x==6{
		return "enam"
	} else if x==7{
		return "tujuh"
	} else if x==8{
		return "delapan"
	} else if x==9{
		return "sembilan"
	}
	return ""
}

func convstring(x string) int {
	if(x == "satu"){
		return 1
	} else if(x == "dua"){
		return 2
	} else if(x == "tiga"){
		return 3
	} else if(x == "empat"){
		return 4
	} else if(x == "lima"){
		return 5
	} else if(x == "enam"){
		return 6
	} else if(x == "tujuh"){
		return 7
	} else if(x == "delapan"){
		return 8
	} else if(x == "sembilan"){
		return 9
	} else if (x =="nol"){
		return 0
	} else{
		return -1
	}
}

func threedigit(x,y,z int ) string{
	number := ""
	if(z == 1){
    	number += "seratus "
    } else if z>1 {
    	number += convsatuan(z)+" ratus " 
    }
    if(y == 1){
    	if(x == 0){
    		number += "sepuluh"
    	} else if (x == 1){
    		number += "sebelas"
    	} else {
    		number += convsatuan(x) +" belas "
    	}
    } else if y>1 {
    	number += convsatuan(y)+" puluh "+convsatuan(x)
    } else{
    	number += convsatuan(x)
    }
    return number
}

func threedigit2(x,y,z int ) string{	
	if(x==1 && y==0 && z==0){
		return "seribu "
	} else if !(x==0 && y==0 && z==0){
		return threedigit(x,y,z)+" ribu "
	}
	return ""
}

func threedigit3(x,y,z int ) string{	
	if !(x==0 && y==0 && z==0){
		return threedigit(x,y,z)+" juta "
	}
	return ""
}

func threedigit4(x,y,z int ) string{	
	if !(x==0 && y==0 && z==0){
		return threedigit(x,y,z)+" milyar "
	}
	return ""
}

func checkDigit(x string) bool{
	i := 0
	for i<len(x){
		if !(x[i]>='0' && x[i]<='9'){
			return false
		}
		i += 1
	}
	return true
}
func spell(w http.ResponseWriter, r *http.Request){ // menangani GET Request
	w.Header().Set("Content-Type","application/json")
	keys,_ := r.URL.Query()["number"]
	var s string = string(keys[0])
	if(!checkDigit(s)){
		angka := Text{Text: "invalid number"}
    	json.NewEncoder(w).Encode(angka)
    	return
	}
	var a [12]int
	i := 0
	for i<len(s){
		a[len(s)-i-1] = int(s[i] - '0')
		i++
	}

    number1 := threedigit(a[0],a[1],a[2])
    number2 := threedigit2(a[3],a[4],a[5])
    number3 := threedigit3(a[6],a[7],a[8])
    number4 := threedigit4(a[9],a[10],a[11])
	number := (number4+number3+number2+number1)
 
    angka := Text{Text: string(number)}
    json.NewEncoder(w).Encode(angka)
}

func read(w http.ResponseWriter, r *http.Request){ // menangani POST Request
	w.Header().Set("Content-Type","application/json")

    var lines Text
    json.NewDecoder(r.Body).Decode(&lines)
    line := lines.Text
    var kalimat []string= strings.Fields(line)
    l := len(kalimat)
    i := 0
    total := 0
    x := 0 // save ratusan
    y := 0 // save puluhan
    c := 0 // angka sebelum
    for i<l{
	    if(kalimat[i] == "sepuluh"){
	    	y = 10
	    } else if(kalimat[i] == "sebelas"){
	    	y = 11
	    } else if (kalimat[i] == "seratus"){
	    	x = 100
	    } else if(kalimat[i] == "seribu"){
	    	total += 1000
	    	x = 0
	    	y = 0
	    	c = 0
	    } else if(kalimat[i] == "ratus"){
	    	x = c*100
	    	c = 0
	    } else if(kalimat[i] == "puluh"){
	    	y = c*10
	    	c = 0
	    } else if(kalimat[i] == "belas"){
	    	y = c+10
	    	c = 0
	    } else if(kalimat[i] == "ribu"){
	    	total += (x+y+c)*1000
			x = 0
	    	y = 0
	    	c = 0
	    } else if(kalimat[i] == "juta"){
	    	total += (x+y+c)*1000000
	    	x = 0
	    	y = 0
	    	c = 0
	    } else if(kalimat[i] == "milyar"){
	    	total += (x+y+c)*1000000000
	    	x = 0
	    	y = 0
	    	c = 0
	    } else {
	    	c = convstring(kalimat[i])
	    	if(c == -1){
			    angka := Number{Number: "Invalid input"}
			    json.NewEncoder(w).Encode(angka)
			    return
	    	}
	    }
	    i += 1
    }
    total += (x+y+c)
    angka := Number{Number: strconv.Itoa(total)}
    json.NewEncoder(w).Encode(angka)
}

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/spell",spell).Methods("GET")
    r.HandleFunc("/read",read).Methods("POST")
    fmt.Print("server is running in 8081\n")
    log.Fatal(http.ListenAndServe(":8081",r))
}