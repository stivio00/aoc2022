package main

import "sort"

// recursive implementation because stack based its too hard for me
func FillSum(dir *Directory) {
	var sum int

	for _, file := range dir.Files {
		sum += file.Size
	}

	for _, directory := range dir.Directories {
		FillSum(directory)
	}

	for _, directory := range dir.Directories {
		sum += directory.Size
	}

	dir.Size = sum
}

func AllDirectoriesAtMost100000Size(dir *Directory) []*Directory {
	out := make([]*Directory, 0)

	for _, directory := range dir.Directories {
		if directory.Size <= 100000 {
			out = append(out, directory)
		}

		sub := AllDirectoriesAtMost100000Size(directory)
		out = append(out, sub...)
	}

	return out
}

func FlaDirectories(dir *Directory) []*Directory {
	flatList := make([]*Directory, 0)

	for _, directory := range dir.Directories {
		flatList = append(flatList, directory)

		sub := FlaDirectories(directory)
		flatList = append(flatList, sub...)
	}

	return flatList
}

func SumDirSizes(dirs []*Directory) int {
	var sum int = 0
	for _, dir := range dirs {
		sum += dir.Size
	}
	return sum
}

func Part1(fs *FileSystem) int {
	dirs := AllDirectoriesAtMost100000Size(fs.Root)
	sum := SumDirSizes(dirs)
	return sum
}

func Part2(fs *FileSystem) int {
	//Compute target size
	target := ComputeTarget(fs)

	allDirs := FlaDirectories(fs.Root)
	sort.Slice(allDirs, func(i, j int) bool {
		//DESC
		return allDirs[i].Size > allDirs[j].Size
	})

	dirSize := fs.Root.Size
	var i int
	for i = 0; dirSize >= target; i++ {
		//next file
		dirSize = allDirs[i+1].Size
	}

	return allDirs[i-1].Size
}

func ComputeTarget(fs *FileSystem) int {
	const total = 70000000
	const unused = 30000000

	currentUsed := fs.Root.Size
	currentUnused := total - currentUsed
	target := unused - currentUnused

	return target
}

/*
Part 1:  1391690
Part 2:  4350256
target:  4795677
*/
