package main

import (
	"os"

	"github.com/anthonydenecheau/csvtools/cli"
)

func main() {

	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

	// 1. parse arguments
	// . type fichier TC, CSAU ...
	// . repértoire à scanner

	// 2. lecture du repertoire
	// . Chargement des fichiers dans un struct
	// . loop

	// 3. Concatenation des reuultat dans un fichier
	// . génération du fichier csv pour maj LOF

	// 4. le fichier sera chargé dans un 2nd temps dans une table pour traitement
	// . génération du fichier csv pour les anos
	// . valider que le résultat pour le chien a été integré (LOF)
	// . reaffectation du résultat au nouvel evenement
	// . supprimer les evenements orphelins

	// 4. Compte rendu
}
