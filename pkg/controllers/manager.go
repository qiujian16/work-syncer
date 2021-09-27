package controllers

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"k8s.io/client-go/dynamic"
	workclient "open-cluster-management.io/api/client/work/clientset/versioned"
	workinformers "open-cluster-management.io/api/client/work/informers/externalversions"
)

// OCMManagerOptions defines the flags for ocm manager
type StatusSyncerOption struct {
	HubKubeconfigFile string
	SpokeClusterName  string
}

// NewWorkloadAgentOptions returns the flags with default value set
func NewStatusSyncerOption() *StatusSyncerOption {
	return &StatusSyncerOption{}
}

// AddFlags register and binds the default flags
func (o *StatusSyncerOption) AddFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	// This command only supports reading from config
	flags.StringVar(&o.HubKubeconfigFile, "hub-kubeconfig", o.HubKubeconfigFile, "Location of kubeconfig file to connect to hub cluster.")
	flags.StringVar(&o.SpokeClusterName, "spoke-cluster-name", o.SpokeClusterName, "Name of spoke cluster.")
}

// RunWorkloadAgent starts the controllers on agent to process work from hub.
func (o *StatusSyncerOption) RunManager(ctx context.Context, controllerContext *controllercmd.ControllerContext) error {
	// build hub client and informer
	hubRestConfig, err := clientcmd.BuildConfigFromFlags("" /* leave masterurl as empty */, o.HubKubeconfigFile)
	if err != nil {
		return err
	}

	hubWorkClient, err := workclient.NewForConfig(hubRestConfig)
	if err != nil {
		return err
	}
	// Only watch the cluster namespace on hub
	workInformerFactory := workinformers.NewSharedInformerFactoryWithOptions(hubWorkClient, 5*time.Minute, workinformers.WithNamespace(o.SpokeClusterName))

	// Build dynamic client and informer for spoke cluster
	spokeRestConfig := controllerContext.KubeConfig
	_, err = dynamic.NewForConfig(spokeRestConfig)
	if err != nil {
		return err
	}

	go workInformerFactory.Start(ctx.Done())

	<-ctx.Done()
	return nil
}
