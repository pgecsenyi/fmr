package bll

import (
	"bll/testutil"
	"dal"
	"testing"
)

func TestCalculator(t *testing.T) {

	setupCalculatorTests()

	t.Run("Calculate_All", testCalculatorAll)
	t.Run("Calculate_MissingOnly", testCalculatorMissingOnly)

	tearDownCalculatorTests()
}

func setupCalculatorTests() {

	testHelper.CreateTestRootDirectory()

	testHelper.CreateTestDirectory("dir1")
	testHelper.CreateTestFileWithContent("test.txt", "Hello World!")
	testHelper.CreateTestFileWithContent("dir1/test.txt", "Lorem ipsum, dolor sit amet.")
}

func testCalculatorAll(t *testing.T) {

	// Arrange.
	expectedFingerprints := testutil.GetExpectedFingerprintsForBasicCalculation()
	memoryDatabase := dal.NewMemoryDatabase()
	testPath := testHelper.GetTestRootDirectory()
	calculator := NewCalculator(memoryDatabase, testPath, "crc32", testPath)

	// Act.
	calculator.Calculate(false)

	// Assert.
	if memoryDatabase.Fingerprints.Len() != 2 {
		t.Errorf("Wrong number of items in result set: %d.", memoryDatabase.Fingerprints.Len())
	}
	testutil.AssertContainsFingerprints(t, memoryDatabase.Fingerprints, expectedFingerprints)
}

func testCalculatorMissingOnly(t *testing.T) {

	// Arrange.
	expectedFingerprints := testutil.GetExpectedFingerprintsForBasicCalculation()
	fp1 := testutil.CreateFingerprint("test.txt", "1c291ca3", "crc32")
	memoryDatabase := dal.NewMemoryDatabase()
	memoryDatabase.AddFingerprint(fp1)
	testPath := testHelper.GetTestRootDirectory()
	calculator := NewCalculator(memoryDatabase, testPath, "crc32", testPath)

	// Act.
	calculator.Calculate(true)

	// Assert.
	if memoryDatabase.Fingerprints.Len() != 1 {
		t.Errorf("Wrong number of items in result set: %d.", memoryDatabase.Fingerprints.Len())
	}
	testutil.AssertContainsFingerprints(t, memoryDatabase.Fingerprints, expectedFingerprints)
}

func tearDownCalculatorTests() {

	testHelper.CleanUp()
}