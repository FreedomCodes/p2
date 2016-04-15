package pcstore

import (
	"errors"

	"github.com/square/p2/pkg/kp/consulutil"
	"github.com/square/p2/pkg/pc/fields"
	"github.com/square/p2/pkg/types"

	klabels "github.com/square/p2/Godeps/_workspace/src/k8s.io/kubernetes/pkg/labels"
)

const podClusterTree string = "pod_clusters"

var NoPodCluster error = errors.New("No pod cluster found")

type Session interface {
	Lock(key string) (consulutil.Unlocker, error)
}

type Store interface {
	Create(
		podId types.PodID,
		availabilityZone fields.AvailabilityZone,
		clusterName fields.ClusterName,
		podSelector klabels.Selector,
		annotations fields.Annotations,
		session Session,
	) (fields.PodCluster, error)
	Get(id fields.ID) (fields.PodCluster, error)
	// Although pod clusters should always be unique for this 3-ple, this method
	// will return a slice in cases where duplicates are discovered. It is up to
	// clients to decide how to respond to such situations.
	FindWhereLabeled(
		podID types.PodID,
		availabilityZone fields.AvailabilityZone,
		clusterName fields.ClusterName,
	) ([]fields.PodCluster, error)
	Delete(id fields.ID) error
}

func IsNotExist(err error) bool {
	return err == NoPodCluster
}
