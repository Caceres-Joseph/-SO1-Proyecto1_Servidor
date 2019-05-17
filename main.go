package main
import (
  "net/http" 
  "fmt" 
  "runtime"    
  "github.com/shirou/gopsutil/mem"
) 


func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))

	PrintMemUsage()
}


func main() { 
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	fmt.Println(bToMb(vmStat.Total));
	fmt.Println(bToMb(vmStat.Free));
	fmt.Println(vmStat.UsedPercent);

	return
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

func dealwithErr(err error) {
	if err != nil {
			fmt.Println(err)
			//os.Exit(-1)
	}
}
// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}