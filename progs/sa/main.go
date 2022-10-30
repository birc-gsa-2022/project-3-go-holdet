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
	genome := os.Args[1]
	reads := os.Args[2]

	p_genomes := shared.GeneralParser(genome, shared.Fasta)
	p_reads := shared.GeneralParser(reads, shared.Fastq)

	/*
		fo, err := os.Create("./testdata/output.txt")
		if err != nil {
			panic(err)
		}*/

	for _, gen := range p_genomes {
		sa := shared.LsdRadixSort(gen.Rec)
		for _, read := range p_reads {
			start, end := shared.BinarySearch(gen.Rec, read.Rec, sa)
			for i := start; i < end; i++ {
				shared.Sam(read.Name, gen.Name, sa[i], read.Rec)
				/*
					res := shared.SamStub(read.Name, genome.Name, sa[i], read.Rec)
					fo.Write([]byte(res))
				*/
			}
		}

	}

	fmt.Println("Search in", genome, "for the reads in", reads)
}
