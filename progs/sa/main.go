package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa/shared"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genomes := os.Args[1]
	reads := os.Args[2]

	p_genomes := shared.GeneralParser(genomes, shared.Fasta)
	p_reads := shared.GeneralParser(reads, shared.Fastq)

	/*fo, err := os.Create("./testdata/output.txt")
	if err != nil {
		panic(err)
	}*/

	for _, genome := range p_genomes {
		sa := shared.LsdRadixSort(genome.Rec)
		for _, read := range p_reads {
			start, end := shared.BinarySearch(genome.Rec, read.Rec, sa)
			for i := start; i < end; i++ {
				shared.Sam(read.Name, genome.Name, i, read.Rec)
				/*
					res := shared.SamStub(read.Name, genome.Name, i, read.Rec)
					fo.Write([]byte(res))
				*/
			}
		}

	}

	fmt.Println("Search in", genomes, "for the reads in", reads)
}
