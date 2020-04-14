package main;

func stringShift(s string, shift [][]int) string {

	bytes := []byte(s);

	rs := 0;
	for _, val := range shift{
		c := val[1];
		if val[0]==0{
			c = -c;
		} 

		rs+=c;
	}

	for rs < 0{

		if -rs < len(bytes){
			ending := bytes[-rs:len(bytes)];
			last := bytes[0:-rs];
			bytes = append(ending, last...);
			rs = 0;
		}else{
			rs++;

			ending := bytes[1:len(bytes)];
			last := bytes[0];
			bytes = append(ending, last);
		}
	}

	for rs > 0{
		if rs < len(bytes){
			ending := bytes[:len(bytes) -rs];
			last := bytes[len(bytes) -rs: len(bytes)];

			bytes = append(last, ending...)
			rs = 0;
		}else{
			ending := bytes[:len(bytes) -1];
			last := bytes[len(bytes) -1];
			bytes = append([]byte{last}, ending...)
	
			rs -= 1;
		}
	}

    return string(bytes);
}
