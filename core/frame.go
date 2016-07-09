package core

// RoutineLoadDir used to
func RoutineLoadDir(r Routine, dirPath string) {

	d := new(DirInfo)
	if d.InitDir(dirPath, r.Filter()) != nil {
		return
	}

	ci := make(chan RoutineChan)
	cnt := 0
	for {

		name := d.NextFile()
		if name == nil {
			break
		}

		cnt++
		go func() {

			// file := tool.NewTFile(*name)
			// fmt.Printf(">>>>>>%v\n", *name)
			ch := r.Run(*name)
			ci <- ch
		}()
		// file := tool.NewTFile(*name)
		// if file != nil {
		//     obj.load(file)
		// }
	}

L:
	for {
		select {
		case ch := <-ci:
			if ch != nil {
				// loadCSV(dict, t)
				r.End(ch)
			}

			cnt--
			if cnt < 1 {
				break L
			}
		}
	}

}
