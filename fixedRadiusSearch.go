package lidario

import "math"

type fixedRadiusSearchKey struct {
	col, row int
}

type fixedRadiusSearchEntry struct {
	x, y, dist float64
	index      int
}

type fixedRadiusSearch struct {
	r, rSqr float64
	hm      map[fixedRadiusSearchKey][]fixedRadiusSearchEntry
	values  []int
	length  int
	// num_cpus: usize,
	// run_concurrently: bool,
}

func newFixedRadiusSearch(radius float64) fixedRadiusSearch {
	hm := make(map[fixedRadiusSearchKey][]fixedRadiusSearchEntry)
	values := make([]int, 0, 1000000)
	frs := fixedRadiusSearch{r: radius, rSqr: radius * radius, hm: hm, values: values, length: 0}
	return frs
}

func (frs *fixedRadiusSearch) insert(x, y float64, value int) {
	k := fixedRadiusSearchKey{col: int(math.Floor((x / frs.r))), row: int(math.Floor((y / frs.r)))}
	if val, ok := frs.hm[k]; ok {
		//do something here
		val = append(val)
	} else {

	}

	frs.values = append(frs.values, value)
	frs.length++
}

//     pub fn insert(&mut self, x: f64, y: f64, value: T) {
//         let k = FixedRadiusSearchKey { col: (x / self.r).floor() as isize, row: (y / self.r).floor() as isize };
//         let val = match self.hm.entry(k) {
//            Vacant(entry) => entry.insert(vec![]),
//            Occupied(entry) => entry.into_mut(),
//         };
//         val.push(FixedRadiusSearchEntry { x: x, y: y, index: self.length, dist: -1f64});
//         self.values.push(value);
//         self.length += 1;

//         // let mut added_value = false;
//         // if let Some(vals) = self.hm.get_mut(&k) {
//         //     self.values.push(value);
//         //     vals.push(FixedRadiusSearchEntry { x: x, y: y, index: self.length, dist: -1f64 });
//         //     added_value = true;
//         // }
//         // if !added_value {
//         //     self.values.push(value);
//         //     self.hm.insert(k, vec![FixedRadiusSearchEntry { x: x, y: y, index: self.length, dist: -1f64 }]);
//         // }
//         // self.length += 1;
//     }

//     pub fn search(&mut self, x: f64, y: f64) -> Vec<(T, f64)> {
//         let mut ret = vec![];
//         let i = (x / self.r).floor() as isize;
//         let j = (y / self.r).floor() as isize;

//         if !self.run_concurrently {
//             for m in -1..2 {
//                 for n in -1..2 {
//                     if let Some(vals) = self.hm.get(&FixedRadiusSearchKey{ col: i+m, row: j+n }) {
//                         for val in vals {
//                             // calculate the squared distance to (x,y)
//                             let dist = (x - val.x)*(x - val.x) + (y - val.y)*(y - val.y);
//                             if dist <= self.r_sqr {
//                                 ret.push((self.values[val.index], dist.sqrt()));
//                             }
//                         }
//                     }
//                 }
//             }
//             if ret.len() > 5000 && self.num_cpus > 1 {
//                 self.run_concurrently = true;
//             }
//         } else {
//             // let (tx, rx) = mpsc::channel();
//             // for m in -1..2 {
//             //     for n in -1..2 {
//             //         let tx = tx.clone();
//             //         if let Some(vals) = self.hm.get_mut(&FixedRadiusSearchKey{ col: i+m, row: j+n }) {
//             //             let vals = vals.clone();
//             //             let x = x.clone();
//             //             let y = y.clone();
//             //             let r_sqr = self.r_sqr.clone();
//             //             let tx = tx.clone();
//             //             thread::spawn(move || {
//             //                 for val in vals {
//             //                     // calculate the squared distance to (x,y)
//             //                     let dist = (x - val.x)*(x - val.x) + (y - val.y)*(y - val.y);
//             //                     if dist <= r_sqr {
//             //                         less_than_threshold.push((val.index, dist.sqrt()));
//             //                     } else {
//             //                         less_than_threshold.push((val.index, -1f64));
//             //                     }
//             //                 }
//             //                 tx.send(less_than_threshold).unwrap();
//             //             });
//             //         }
//             //     }
//             // }
//             //
//             // for _ in 0..10 {
//             //     let data = rx.recv().unwrap();
//             //     for d in data {
//             //         if d.1 >= 0f64 {
//             //             ret.push((self.values[d.0], d.1));
//             //         }
//             //     }
//             // }

//             // let mut points = vec![];
//             for m in -1..2 {
//                 for n in -1..2 {
//                     if let Some(vals) = self.hm.get_mut(&FixedRadiusSearchKey{ col: i+m, row: j+n }) {
//                         // points.extend_from_slice(&vals[..]);
//                         // points.extend(vals.iter().cloned());
//                         if vals.len() >= 5000 {
//                             calc_dist(&mut vals[..], &self.r_sqr, &x, &y);
//                             for val in vals {
//                                 if val.dist >= 0f64 {
//                                     ret.push((self.values[val.index], val.dist));
//                                 }
//                             }
//                         } else {
//                             for val in vals {
//                                 // calculate the squared distance to (x,y)
//                                 let dist = (x - val.x)*(x - val.x) + (y - val.y)*(y - val.y);
//                                 if dist <= self.r_sqr {
//                                     ret.push((self.values[val.index], dist.sqrt()));
//                                 }
//                             }
//                         }
//                     }
//                 }
//             }
//             // if points.len() >= 5000 {
//             //     calc_dist(&mut points[..], &self.r_sqr, &x, &y);
//             //     for val in points {
//             //         if val.dist >= 0f64 {
//             //             ret.push((self.values[val.index], val.dist));
//             //         }
//             //     }
//             // } else {
//             //     for val in points {
//             //         // calculate the squared distance to (x,y)
//             //         let dist = (x - val.x)*(x - val.x) + (y - val.y)*(y - val.y);
//             //         if dist <= self.r_sqr {
//             //             ret.push((self.values[val.index], dist.sqrt()));
//             //         }
//             //     }
//             // }

//             if ret.len() < 2500 {
//                 self.run_concurrently = false;
//             }
//         }

//         ret
//     }
