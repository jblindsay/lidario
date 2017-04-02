package lidario

import (
	"fmt"
	"os"
	"testing"
)

func TestReadLasFile(t *testing.T) {
	// fileName := "/Users/johnlindsay/Documents/data/GarveyGlenWatershed/RGB_4_529_150502_1812__0_178080_2784.las"
	fileName := "testdata/sample.las"
	var lf *LasFile
	var err error
	lf, err = NewLasFile(fileName, "r")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer lf.Close()
	fmt.Printf("%v\n", lf.Header)

	// Print the VLR data
	fmt.Println("VLRs:")
	for _, vlr := range lf.VlrData {
		fmt.Println(vlr)
	}

	j := 1000000

	x, y, z, err := lf.GetXYZ(j)
	fmt.Printf("Point %v: (%f, %f, %f) Error: %v\n", j, x, y, z, err)
	var p LasPointer

	p, err = lf.LasPoint(j)
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}
	fmt.Println("Point format:", p.Format())

	oldProgress := -1
	progress := 0
	for i := 0; i < int(lf.Header.NumberPoints); i++ {
		if p, err := lf.LasPoint(i); err != nil {
			fmt.Println(err)
			t.Fatal()
		} else {
			if i < 10 {
				pd := p.PointData()
				fmt.Printf("Point %v: (%f, %f, %f, %v, %v, %f)\n", i, pd.X, pd.Y, pd.X, pd.Intensity, pd.ClassBitField.ClassificationString(), p.GpsTimeData())
			}
			progress = int(100.0 * float64(i) / float64(lf.Header.NumberPoints))
			if progress != oldProgress {
				oldProgress = progress
				if progress%10 == 0 {
					fmt.Printf("Progress: %v\n", progress)
				}
			}
		}
	}

}

func TestWriteLasFile(t *testing.T) {
	fileName := "testdata/sample.las"
	var lf *LasFile
	var err error
	lf, err = NewLasFile(fileName, "r")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer lf.Close()

	newFileName := "testdata/newFile.las"
	newLf, err := InitializeUsingFile(newFileName, lf)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	progress := 0
	oldProgress := -1
	for i := 0; i < int(lf.Header.NumberPoints); i++ {
		if p, err := lf.LasPoint(i); err != nil {
			fmt.Println(err)
			t.Fatal()
		} else {
			if p.PointData().Z < 100.0 {
				newLf.AddLasPoint(p)
			}
		}
		progress = int(100.0 * float64(i) / float64(lf.Header.NumberPoints))
		if progress != oldProgress {
			oldProgress = progress
			if progress%10 == 0 {
				fmt.Printf("Progress: %v\n", progress)
			}
		}
	}

	newLf.Close()

	// now delete the file
	if err = os.Remove(newFileName); err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
