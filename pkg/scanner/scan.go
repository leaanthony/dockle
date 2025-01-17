package scanner

import (
	"context"
	"flag"
	"os"

	"github.com/goodwithtech/dockle/pkg/types"

	"github.com/goodwithtech/dockle/pkg/assessor"

	"github.com/knqyf263/fanal/analyzer"
	"github.com/knqyf263/fanal/extractor"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/xerrors"
)

func ScanImage(imageName, filePath string) (assessments []*types.Assessment, err error) {
	ctx := context.Background()
	var files extractor.FileMap

	// add required files to fanal's analyzer
	analyzer.AddRequiredFilenames(assessor.LoadRequiredFiles())
	analyzer.AddRequiredPermissions(assessor.LoadRequiredPermissions())
	if imageName != "" {
		dockerOption, err := types.GetDockerOption()
		if err != nil {
			return nil, xerrors.Errorf("failed to get docker option: %w", err)
		}
		files, err = analyzer.Analyze(ctx, imageName, dockerOption)
		if err != nil {
			return nil, xerrors.Errorf("failed to analyze image: %w", err)
		}
	} else if filePath != "" {
		rc, err := openStream(filePath)
		if err != nil {
			return nil, xerrors.Errorf("failed to open stream: %w", err)
		}

		files, err = analyzer.AnalyzeFromFile(ctx, rc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, xerrors.New("image name or image file must be specified")
	}

	assessments = assessor.GetAssessments(files)
	return assessments, nil
}

func openStream(path string) (*os.File, error) {
	if path == "-" {
		if terminal.IsTerminal(0) {
			flag.Usage()
			os.Exit(64)
		} else {
			return os.Stdin, nil
		}
	}
	return os.Open(path)
}
