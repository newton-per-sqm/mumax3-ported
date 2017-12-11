package engine

import (
	"github.com/mumax/3/httpfs"
	"github.com/mumax/3/util"
	"io"
)

const seperationline = `
---------------------------------------------------------------------------
`
const bibheader = `
This bibtex file is automatically generated by Mumax3.

The following list are references relevant for your simulation. If you use
the results of these simulations in any work or publication, we kindly ask
you to cite them.`

var (
	bibfile io.WriteCloser
	library map[string]*bibEntry
)

func init() {
	buildLibrary()
}

func initBib() { // inited in engine.InitIO
	if bibfile != nil {
		panic("bib already initialized")
	}
	var err error
	bibfile, err = httpfs.Create(OD() + "references.bib")
	if err != nil {
		panic(err)
	}
	util.FatalErr(err)
	fprintln(bibfile, bibheader)
	fprintln(bibfile, seperationline)
	Refer("vansteenkiste2014") // Make sure that Mumax3 is always referenced
}

type bibEntry struct {
	reason string
	bibtex string
	used   bool
}

func Refer(tag string) {
	bibentry, inLibrary := library[tag]
	if bibentry.used || !inLibrary {
		return
	}
	bibentry.used = true
	if bibfile != nil {
		fprintln(bibfile, bibentry.reason)
		fprintln(bibfile, bibentry.bibtex)
		fprintln(bibfile, seperationline)
	}
}

func buildLibrary() {

	library = make(map[string]*bibEntry)

	library["vansteenkiste2014"] = &bibEntry{
		reason: "Main paper about Mumax3",
		bibtex: `
@article{Vansteenkiste2014,
    author  = {Vansteenkiste, Arne and
               Leliaert, Jonathan and
               Dvornik, Mykola and
               Helsen, Mathias and
               Garcia-Sanchez, Felipe and
               {Van Waeyenberge}, Bartel},
    title   = {{The design and verification of Mumax3}},
    journal = {AIP Advances},
    number  = {10},
    pages   = {107133},
    volume  = {4},
    year    = {2014},
    doi     = {10.1063/1.4899186},
    url     = {http://doi.org/10.1063/1.4899186}
}`}

	library["exl2014"] = &bibEntry{
		reason: "Mumax3 uses Exl's minimizer",
		bibtex: `
@article{Exl2014,
    author  = {Exl, Lukas and
               Bance, Simon and
               Reichel, Franz and
               Schrefl, Thomas and
               {Peter Stimming}, Hans and
               Mauser, Norbert J.},
    title   = {{LaBonte's method revisited: An effective steepest
                descent method for micromagnetic energy minimization}},
    journal = {Journal of Applied Physics},
    number  = {17},
    pages   = {17D118},
    volume  = {115},
    year    = {2014},
    doi     = {10.1063/1.4862839},
    url     = {http://doi.org/10.1063/1.4862839}
}`}

	library["Lel2014"] = &bibEntry{
		reason: "Mumax3 used function ext_makegrains",
		bibtex: `
@article{Lel2014,
    author  = {Leliaert, Jonathan and
               Van de Wiele, Ben and
               Vansteenkiste, Arne and
               Laurson, Lasse and
               Durin, Gianfranco and
               Dupr{\'e}, Luc and
               Van Waeyenberge, Bartel},
    title   = {{Current-driven domain wall mobility in polycrystalline permalloy nanowires: A numerical study}},
    journal = {Journal of Applied Physics},
    number  = {23},
    pages   = {233903},
    year    = {2014},
    doi     = {10.1063/1.4883297},
    url     = {http://dx.doi.org/10.1063/1.4883297}
}`}

	library["mulkers2017"] = &bibEntry{
		reason: "Simulated system has interfacially induced DMI",
		bibtex: `
@article{Mulkers2017,
    author  = {Mulkers, Jeroen and
               Van Waeyenberge, Bartel and
	       Milo{\v{s}}evi{\'{c}}, Milorad V.},
    title   = {{Effects of spatially-engineered Dzyaloshinskii-Moriya
                interaction in ferromagnetic films}},
    journal = {Physical Review B},
    number  = {14},
    pages   = {144401},
    volume  = {95},
    year    = {2017},
    doi     = {10.1103/PhysRevB.95.144401},
    url     = {doi.org/10.1103/PhysRevB.95.144401},
}`}

	library["leliaert2017"] = &bibEntry{
		reason: "Simulated nonzero temperatures with adaptive time steps",
		bibtex: `
@article{Leliaert2017,
    author  = {Leliaert, Jonathan and
               Mulkers, Jeroen and
	       De Clercq, Jonas and
	       Coene, Annelies and
               Dvornik, Mykola and
               Van Waeyenberge, Bartel},
    title   = {{Adaptively time stepping the stochastic Landau-Lifshitz-Gilbert equation at nonzero temperature: implementation and validation in MuMax$^3$}},
    journal = {AIP Advances},
    number  = {12},
    pages   = {125010},
    volume  = {7},
    year    = {2017},
    doi     = {doi.org/10.1063/1.5003957},
    url     = {http://aip.scitation.org/doi/10.1063/1.5003957},
}`}

}
