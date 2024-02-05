package main

import (
	"encoding/json"
	"fmt"
	"os"

	itociartifact "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs/types/ociartifact"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/accessmethods/ociartifact"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
)

const (
	SecScanLabelKey = "scan.security.kyma-project.io"
)

var labelTemplate = SecScanLabelKey + "/%s"

func main() {
	fmt.Println("ocm resource equality test for ocm version v0.6.0")
	fmt.Println("========================================")
	fmt.Println("")
	res, err := createResource("")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("initial resource:")
	toJson(res)

	res2, err := createResource("")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(2)
	}

	equiv := res.Equivalent(res2)
	fmt.Println("Resource isEquivalent to it's own copy:", equiv.IsEquivalent())
	fmt.Println("Resource isHashEqual to it's own copy:", equiv.IsHashEqual())
	fmt.Println("Resource isArtifactEqual to it's own copy:", equiv.IsArtifactEqual())
	fmt.Println("Resource isLocalHashEqual to it's own copy:", equiv.IsLocalHashEqual())

	fmt.Println("----------------------------------------")
	res3, err := createResource("-foo")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(3)
	}

	fmt.Println("another resource:")
	toJson(res3)
	equiv = res.Equivalent(res3)
	fmt.Println("Resource isLocalHashEqual to a different one:", equiv.IsLocalHashEqual())
}

func createResource(accessMod string) (*compdesc.Resource, error) {

	imageName := "template-operator"
	imageTag := "0.1.0"
	imageURL := "europe-docker.pkg.dev/kyma-project/prod/" + imageName + ":" + imageTag

	imageTypeLabel, err := generateOCMLabel("type", "third-party-image", labelTemplate)
	if err != nil {
		return nil, err
	}

	access := ociartifact.New(imageURL + accessMod)
	access.SetType(ociartifact.LegacyType)
	imageLayerResource := compdesc.Resource{
		ResourceMeta: compdesc.ResourceMeta{
			ElementMeta: compdesc.ElementMeta{
				Name:    imageName,
				Labels:  []metav1.Label{*imageTypeLabel},
				Version: imageTag,
			},
			Type:     itociartifact.LEGACY_TYPE,
			Relation: metav1.ExternalRelation,
		},
		Access: access,
	}
	return &imageLayerResource, nil
}

func generateOCMLabel(key, value, tpl string) (*metav1.Label, error) {
	return metav1.NewLabel(fmt.Sprintf(tpl, key), value, metav1.WithVersion("v1"))
}

func toJson(val any) {
	out, err := json.MarshalIndent(val, "", "    ")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(10)
	}
	fmt.Println(string(out))
}
