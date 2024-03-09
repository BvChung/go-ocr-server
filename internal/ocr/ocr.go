package ocr

import (
	"context"
	"fmt"
	"io"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func DetectText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
			return err
	}

	f, err := os.Open(file)
	if err != nil {
			return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
			return err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
			return err
	}

	fmt.Fprintln(w, "Text:")
	for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
	}

	return nil
}