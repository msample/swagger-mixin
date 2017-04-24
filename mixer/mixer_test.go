package mixer_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/msample/swagger-mixin/mixer"
)

func TestMixin(t *testing.T) {

	primaryDoc, err := loads.Spec("../test-data/s1.yml")
	if err != nil {
		t.Fatalf("Could not load ../test-data/s1.yml: %v\n", err)
	}
	mixinDoc1, err := loads.Spec("../test-data/s2.yml")
	if err != nil {
		t.Fatalf("Could not load ../test-data/s2.yml: %v\n", err)
	}
	mixinDoc2, err := loads.Spec("../test-data/s3.yml")
	if err != nil {
		t.Fatalf("Could not load ../test-data/s3.yml: %v\n", err)
	}

	primary := primaryDoc.Spec()
	collisionCount := mixer.Mixin(primary, mixinDoc1.Spec(), mixinDoc2.Spec())
	if collisionCount != 12 {
		t.Errorf("TestMixin: Expected 10 collisions, got %v\n", collisionCount)
	}

	if len(primary.Paths.Paths) != 7 {
		t.Errorf("TestMixin: Expected 7 paths in merged, got %v\n", len(primary.Paths.Paths))
	}

	if len(primary.Definitions) != 8 {
		t.Errorf("TestMixin: Expected 8 definitions in merged, got %v\n", len(primary.Definitions))
	}

	if len(primary.Parameters) != 4 {
		t.Errorf("TestMixin: Expected 4 top level parameters in merged, got %v\n", len(primary.Parameters))
	}

	if len(primary.Responses) != 2 {
		t.Errorf("TestMixin: Expected 2 top level responses in merged, got %v\n", len(primary.Responses))
	}

}

func TestMixinFiles(t *testing.T) {
	f, err := ioutil.TempFile("", "mixerTest-")
	if err != nil {
		t.Fatal(err)
	}
	collisions, err := mixer.MixinFiles("../test-data/s1.yml", []string{"../test-data/s2.yml", "../test-data/s3.yml"}, f)
	if err != nil {
		t.Errorf("TestMixinFiles: got error: %v\n", err)
	}
	if collisions != 12 {
		t.Errorf("TestMixinFiles: expected 12 collisions, got: %v\n", collisions)
	}
	specDoc, err := loads.Spec(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	result := validate.Spec(specDoc, strfmt.Default)
	if result != nil {
		str := fmt.Sprintf("The swagger spec at %q is invalid against swagger specification %s. see errors :\n", f.Name(), specDoc.Version())
		for _, desc := range result.(*errors.CompositeError).Errors {
			str += fmt.Sprintf("- %s\n", desc)
		}
		t.Error(str)
	}
}
