package main
import (
	"encoding/json"
	"net/http" 
	"fmt"  
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"math"
	"time"
) 

type JsonMemoria struct {
	Total uint64
	Libre uint64
	Porcentaje float64
}
 

/*
+----------------------------------------------
|	Funcion que retorna datos del cpu
+----------------------------------------------
*/

func totalCpu(w http.ResponseWriter, r *http.Request) {
  
	//Librería para obtener la memoria libre
	vmStat, err := cpu.Percent(0,false);
	dealwithErr(err)


	//

	//Llenando el json
	jsonMemoria := JsonMemoria{0,0,math.Floor(vmStat[0]*100)/100 }
	


	//Creando el json
	js, err := json.Marshal(jsonMemoria)
	dealwithErr(err)
  
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05")," | Retornando total de CPU")

	fmt.Println(vmStat[0]);
	//Enviando el json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	

 
	   
}

/*
+----------------------------------------------
|	Funcion que retorna datos de la ram
+----------------------------------------------
*/

func totalRam(w http.ResponseWriter, r *http.Request) {
  
	//Librería para obtener la memoria libre
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	
	//Llenando el json
	jsonMemoria := JsonMemoria{bToMb(vmStat.Total), bToMb(vmStat.Free), math.Floor(vmStat.UsedPercent*100)/100}
	

	//Creando el json
	js, err := json.Marshal(jsonMemoria)
	dealwithErr(err)
  
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05")," | Retornando total de RAM")

	//Enviando el json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}



func main() {

	fmt.Println("... Iniciando servidor .....");
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/totalRam", totalRam)
	http.HandleFunc("/totalCpu", totalCpu)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

func dealwithErr(err error) {
	if err != nil {
			fmt.Println(err)
	}
}


func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}