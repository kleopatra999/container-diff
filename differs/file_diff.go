package differs

import (
	"github.com/GoogleCloudPlatform/container-diff/utils"
)

type FileAnalyzer struct {
}

func (a FileAnalyzer) Name() string {
	return "FileAnalyzer"
}

// FileDiff diffs two packages and compares their contents
func (a FileAnalyzer) Diff(image1, image2 utils.Image) (utils.Result, error) {
	diff, err := diffImageFiles(image1, image2)
	return &utils.DirDiffResult{
		Image1:   image1.Source,
		Image2:   image2.Source,
		DiffType: "File",
		Diff:     diff,
	}, err
}

func (a FileAnalyzer) Analyze(image utils.Image) (utils.Result, error) {
	var result utils.FileAnalyzeResult

	imgDir, err := utils.GetDirectory(image.FSPath, true)
	if err != nil {
		return result, err
	}

	result.Image = image.Source
	result.AnalyzeType = "File"
	result.Analysis = utils.GetDirectoryEntries(imgDir)
	return &result, err
}

func diffImageFiles(image1, image2 utils.Image) (utils.DirDiff, error) {
	img1 := image1.FSPath
	img2 := image2.FSPath

	var diff utils.DirDiff

	img1Dir, err := utils.GetDirectory(img1, true)
	if err != nil {
		return diff, err
	}
	img2Dir, err := utils.GetDirectory(img2, true)
	if err != nil {
		return diff, err
	}

	diff, _ = utils.DiffDirectory(img1Dir, img2Dir)
	return diff, nil
}
