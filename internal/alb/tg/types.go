package tg

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	extensions "k8s.io/api/extensions/v1beta1"
)

// TargetGroup provides information about AWS targetGroup created.
type TargetGroup struct {
	Arn        string
	TargetType string
	Targets    []*elbv2.TargetDescription
}

// TargetGroupGroup represents an collection of targetGroups for a single ingress in AWS
type TargetGroupGroup struct {
	TGByBackend map[extensions.IngressBackend]TargetGroup
	selector    map[string]string
}

// NameGenerator provides name generation functionality for tg package.
type NameGenerator interface {
	// NameTG generates name for targetGroups.
	// Note: targetType & protocol is included here to ensure we'll create new targetGroups if one of them changed(they cannot be modified)
	NameTG(namespace string, ingressName string, serviceName, servicePort string,
		targetType string, protocol string) string
}

// TagGenerator provides tag generation functionality for tg package.
type TagGenerator interface {
	// TagTGGroup generates tags for the group of targetGroups created for a single ingress.
	TagTGGroup(namespace string, ingressName string) map[string]string

	// TagTG generates tags for a targetGroup inside a TGGroup.
	// NOTE: The final set of tags been applied to targetGroup is union of tags generated by TagTGs & TagTG.
	TagTG(serviceName string, servicePort string) map[string]string
}

// NameTagGenerator is combination of NameGenerator and TagGenerator
type NameTagGenerator interface {
	NameGenerator
	TagGenerator
}
