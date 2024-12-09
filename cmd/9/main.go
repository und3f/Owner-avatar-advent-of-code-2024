package main

import (
	"slices"

	"github.com/und3f/aoc/2024/fwk"
)

type Entry struct {
	size   int
	id     int
	isFree bool
}

func main() {
	input := fwk.ReadInputRunesLines()[0]

	part1(input)
	part2(input)
}

func part1(input []rune) {
	disk := buildDisk(input)
	sortDisk(disk)

	fwk.Solution(1, calcChecksum(disk))
}

func part2(input []rune) {
	disk := sortDiskFiles(input)

	fwk.Solution(2, calcChecksum(disk))
}

func buildDisk(input []rune) []rune {
	var disk []rune

	id := rune(0)
	isSpace := false
	for _, v := range input {
		sym := '.'
		if !isSpace {
			sym = id + '0'
		}
		l := int(v - '0')
		sector := make([]rune, l)
		for i := 0; i < l; i++ {
			sector[i] = sym
		}
		disk = append(disk, sector...)

		if !isSpace {
			id++
		}
		isSpace = !isSpace
	}

	return disk
}

func sortDisk(disk []rune) {
	i := 0
	for ; disk[i] != '.'; i++ {
	}

	j := len(disk) - 1
	for ; disk[j] == '.'; j-- {
	}

	for i < j {
		disk[i] = disk[j]
		disk[j] = '.'

		for ; disk[i] != '.'; i++ {
		}

		for ; disk[j] == '.'; j-- {
		}
	}
}

func sortDiskFiles(input []rune) []rune {
	var disk []Entry
	id := 0
	freeSpace := false
	for i := 0; i < len(input); i++ {
		size := int(input[i] - '0')

		var e Entry
		if !freeSpace {
			e = Entry{
				size: size,
				id:   id,
			}
			id++
		} else {
			e = Entry{
				size:   size,
				id:     -1,
				isFree: true,
			}
		}
		freeSpace = !freeSpace

		if size > 0 {
			disk = append(disk, e)
		}
	}

	for j := len(disk) - 1; j > 0; j-- {
		e := disk[j]
		if e.isFree {
			continue
		}

		for i := 0; i < j; i++ {
			space := disk[i]
			if !(space.isFree && space.size >= e.size) {
				continue
			}

			d := slices.Clone(disk[:i])
			d = append(d, e)
			space.size -= e.size
			if space.size > 0 {
				d = append(d, space)
			}

			if i+1 < j {
				d = append(d, disk[i+1:j-1]...)
			}

			freeSpace := Entry{
				size:   e.size,
				isFree: true,
				id:     -1,
			}

			// Nailed it

			if i < j-1 {
				p := disk[j-1]
				if p.isFree {
					freeSpace.size += p.size

				} else {
					d = append(d, p)
				}
			}

			if j+1 < len(disk) {
				p := disk[j+1]
				if p.isFree {
					freeSpace.size += p.size
					d = append(d, freeSpace)
				} else {
					d = append(d, freeSpace, p)
				}

				if j+2 < len(disk) {
					d = append(d, disk[j+2:]...)
				}
			} else {
				d = append(d, freeSpace)
			}

			if space.size > 0 {
				j++
			}
			disk = d
			break
		}

	}

	return stringifyDiskEntries(disk)
}

func stringifyDiskEntries(disk []Entry) []rune {
	total := 0
	for i := 0; i < len(disk); i++ {
		total += disk[i].size
	}
	drive := make([]rune, total)

	i := 0
	for _, e := range disk {
		for j := 0; j < e.size; j++ {
			sym := '.'
			if !e.isFree {
				sym = rune(e.id + '0')
			}
			drive[i+j] = sym
		}
		i += e.size
	}
	return drive
}

func calcChecksum(disk []rune) uint64 {
	var sum uint64
	for i := 0; i < len(disk); i++ {
		if disk[i] != '.' {
			sum += uint64(i) * uint64(disk[i]-'0')
		}
	}

	return sum
}
