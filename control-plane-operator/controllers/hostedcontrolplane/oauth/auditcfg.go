package oauth

import (
	configv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/openshift/hypershift/support/config"
)

const (
	auditPolicyConfigMapKey  = "policy.yaml"
	auditPolicyProfileMapKey = "profile"
)

func ReconcileAuditConfig(cm *corev1.ConfigMap, ownerRef config.OwnerRef, auditConfig configv1.Audit) error {
	ownerRef.ApplyTo(cm)
	if cm.Data == nil {
		cm.Data = map[string]string{}
	}

	cm.Data[auditPolicyConfigMapKey] = oauthPolicy
	cm.Data[auditPolicyProfileMapKey] = string(auditConfig.Profile)
	return nil
}

func AuditEnabled(auditCfgMap *corev1.ConfigMap) bool {
	return auditCfgMap.Data[auditPolicyProfileMapKey] != string(configv1.NoneAuditProfileType)
}
