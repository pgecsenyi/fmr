package application

import (
	"bll"
	"dal"
	"flag"
	"log"
	"os"
	"util"
)

const taskCalculate = "calculate"
const taskCompare = "compare"
const taskExport = "export"
const taskImport = "import"
const taskVerify = "verify"

// Application Contains main application logic.
type Application struct {
	config configuration
}

type configuration struct {
	task            string
	algorithm       string
	inputChecksum   string
	inputDirectory  string
	outputChecksum  string
	outputDirectory string
	outputNames     string
	basePath        string
	filter          string
}

// Initialize Initializes the application.
func (app *Application) Initialize() {

	defaultConfig := configuration{taskCalculate, dal.SHA1, "", "", "", "", "", "", ""}
	app.parseCommandLineArguments(defaultConfig)
	app.verifyConfiguration()
	app.execute()
}

func (app *Application) parseCommandLineArguments(defaultConfig configuration) {

	task := flag.String(
		"task",
		defaultConfig.task,
		"The task to execute: calculate, compare, import, export or verify. The first one calculates checksums for a"+
			" directory and stores the results in a CSV. The second compares stored checksums with the checksums of"+
			" the files in the given directory and stores filename matches. The third imports checksums from files"+
			" generated by Linux utilities or Total Commander. The fourth exports to Total Commander's formats. The"+
			" fifth verifies checksums for the files listed in the given CSV.")
	inputDirectory := flag.String(
		"indir",
		defaultConfig.inputDirectory,
		"The source directory for which the checksums will be calculated (or will be compared).")
	algorithm := flag.String("alg", defaultConfig.algorithm, "The algorithm used to calculate new checksums.")
	outputChecksum := flag.String(
		"outchk",
		defaultConfig.outputChecksum,
		"The name of the output CSV file containing checksums.")
	inputChecksum := flag.String(
		"inchk",
		defaultConfig.inputChecksum,
		"The name of the input CSV containing checksums.")
	outputNames := flag.String(
		"outnames",
		defaultConfig.outputNames,
		"The name of the output containing new file name and old filename pairs.")
	basePath := flag.String(
		"bp",
		defaultConfig.basePath,
		"The first part of the path that will not be stored in the output.")
	outputDirectory := flag.String(
		"outdir",
		defaultConfig.outputDirectory,
		"The name of the directory containing exported files.")
	filter := flag.String("filter", defaultConfig.filter, "A filter for exported filenames.")

	flag.Parse()

	app.config = configuration{
		*task, *algorithm,
		*inputChecksum, *inputDirectory,
		*outputChecksum, *outputDirectory, *outputNames,
		*basePath, *filter}
}

func (app *Application) verifyConfiguration() {

	if app.config.task != taskCalculate && app.config.task != taskCompare &&
		app.config.task != taskVerify && app.config.task != taskImport &&
		app.config.task != taskExport {
		log.Fatalln("Unknown task.")
	}
	if app.config.task == taskCalculate || app.config.task == taskCompare || app.config.task == taskImport {
		if !checkIfDirectoryExists(app.config.inputDirectory) {
			log.Fatalln("Directory " + app.config.inputDirectory + " does not exist.")
		}
	} else if app.config.task == taskExport {
		if !checkIfDirectoryExists(app.config.outputDirectory) {
			log.Fatalln("Directory " + app.config.outputDirectory + " does not exist.")
		}
	}
	if app.config.inputChecksum != "" && !util.CheckIfFileExists(app.config.inputChecksum) {
		log.Fatalln("Input file does not exist.")
	}
}

func (app *Application) execute() {

	db := dal.NewDb()
	if app.config.task == taskCalculate {
		calculator := bll.Calculator{app.config.inputDirectory, app.config.outputChecksum, app.config.basePath}
		calculator.RecordChecksumsForDirectory(&db, app.config.algorithm)
	} else if app.config.task == taskCompare {
		comparer := bll.Comparer{
			app.config.inputDirectory, app.config.inputChecksum,
			app.config.outputNames, app.config.outputChecksum,
			app.config.basePath}
		comparer.RecordNameChangesForDirectory(&db, app.config.algorithm)
	} else if app.config.task == taskExport {
		exporter := bll.NewExporter(
			app.config.inputChecksum, app.config.outputDirectory,
			app.config.filter, app.config.basePath)
		exporter.Convert(&db)
	} else if app.config.task == taskImport {
		importer := bll.NewImporter(app.config.inputDirectory, app.config.outputChecksum)
		importer.Convert(&db)
	} else if app.config.task == taskVerify {
		verifier := bll.NewVerifier(app.config.inputChecksum, app.config.basePath)
		verifier.Verify(&db)
	}
}

func checkIfDirectoryExists(path string) bool {

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}

	return false
}
