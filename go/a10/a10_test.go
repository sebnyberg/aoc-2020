package a10_test

// func check(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func Test_day10part1(t *testing.T) {
// 	var err error
// 	var f io.Reader
// 	// Read inputs, each row is an adapter with an output joltage
// 	// Each adapter can take an input outputJoltage-3 <= jolt <= outputJoltage
// 	f, err = os.Open("input")
// 	check(err)

// 	sc := bufio.NewScanner(f)
// 	ns := make([]int, 0)
// 	var n int
// 	for sc.Scan() {
// 		n, err = strconv.Atoi(sc.Text())
// 		check(err)
// 		ns = append(ns, n)
// 	}
// 	sort.Ints(ns)

// 	prevN := ns[0]
// 	n1jolt := 0
// 	n3jolt := 1

// 	if ns[0] == 1 {
// 		n1jolt++
// 	}
// 	if ns[0] == 3 {
// 		n3jolt++
// 	}
// 	for _, n := range ns[1:] {
// 		if n == prevN {
// 			panic("prevN == n, should not happen")
// 		}
// 		// End condition
// 		if n-prevN > 3 {
// 			break
// 		}

// 		if n-prevN == 1 {
// 			n1jolt++
// 		}

// 		if n-prevN == 3 {
// 			n3jolt++
// 		}
// 		prevN = n
// 	}

// 	fmt.Println(n1jolt)
// 	fmt.Println(n3jolt)
// 	fmt.Println(n1jolt * n3jolt)
// 	t.FailNow()
// }

// func Test_day10part2(t *testing.T) {
// 	var err error
// 	var f io.Reader
// 	// Read inputs, each row is an adapter with an output joltage
// 	// Each adapter can take an input outputJoltage-3 <= jolt <= outputJoltage
// 	f, err = os.Open("input")
// 	check(err)

// 	sc := bufio.NewScanner(f)
// 	ns := make([]int, 0)
// 	var n int
// 	for sc.Scan() {
// 		n, err = strconv.Atoi(sc.Text())
// 		check(err)
// 		ns = append(ns, n)
// 	}
// 	sort.Ints(ns)

// 	t.FailNow()
// }

// // Notes:
// // * The total number of configurations is the sum of the number of configs
// // 	for each section separated by a distance of 2.

// func Test_countConfigurations(t *testing.T) {
// 	in := []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}
// 	got := countConfigurations(in)
// 	require.Equal(t, 8, got)

// 	in2 := []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31,
// 		32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52}
// 	got = countConfigurations(in2)
// 	require.Equal(t, 19208, got)
// }

// func countConfigurations(ns []int) int {
// 	var i, j int
// 	// nconfig := 1

// 	for j < len(ns) {
// 		// Tactic: increment j until ns[j] - ns[i] > 3
// 		for ns[j]-ns[i] <= 3 && j < len(ns) {
// 			j++
// 		}
// 		fmt.Println(ns[i:j])
// 		fmt.Println(1 << (j - i - 2))
// 		i = j - 1
// 		j++
// 	}
// 	return 0
// }
