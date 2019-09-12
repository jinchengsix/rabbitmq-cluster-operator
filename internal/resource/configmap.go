package resource

import (
	"strings"

	rabbitmqv1beta1 "github.com/pivotal/rabbitmq-for-kubernetes/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	serverConfigMapName = "server-conf"
)

func GenerateServerConfigMap(instance rabbitmqv1beta1.RabbitmqCluster) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ChildResourceName(serverConfigMapName),
			Namespace: instance.Namespace,
			Labels: map[string]string{
				"app":             "pivotal-rabbitmq",
				"RabbitmqCluster": instance.Name,
			},
		},
		Data: map[string]string{
			"enabled_plugins": "[" +
				"rabbitmq_management," +
				"rabbitmq_peer_discovery_k8s," +
				"rabbitmq_federation," +
				"rabbitmq_federation_management," +
				"rabbitmq_shovel," +
				"rabbitmq_shovel_management," +
				"rabbitmq_prometheus].",

			"rabbitmq.conf": strings.Join([]string{
				"cluster_formation.peer_discovery_backend = rabbit_peer_discovery_k8s",
				"cluster_formation.k8s.host = kubernetes.default.svc.cluster.local",
				"cluster_formation.k8s.address_type = hostname",
				"cluster_formation.node_cleanup.interval = 30",
				"cluster_formation.node_cleanup.only_log_warning = true",
				"cluster_partition_handling = autoheal",
				"queue_master_locator = min-masters",
			}, "\n"),
		},
	}
}
